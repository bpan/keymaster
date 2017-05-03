/*
  Package certgen id set of utilities used to generate ssh certificates
*/
package certgen

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/ssh"
	"math/big"
	//"os"
	"os/exec"
	"time"

	"fmt"
)

const numValidHours = 24

// GetUserPubKeyFromSSSD user authorized keys content based on the running sssd configuration
func GetUserPubKeyFromSSSD(username string) (string, error) {
	cmd := exec.Command("/usr/bin/sss_ssh_authorizedkeys", username)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func goCertToFileString(c ssh.Certificate, username string) (string, error) {
	certBytes := c.Marshal()
	encoded := base64.StdEncoding.EncodeToString(certBytes)
	fileComment := "/tmp/" + username + "-cert.pub"
	return "ssh-rsa-cert-v01@openssh.com " + encoded + " " + fileComment, nil
}

// gen_user_cert a username and key, returns a short lived cert for that user
func GenSSHCertFileString(username string, userPubKey string, signer ssh.Signer, host_identity string) (string, error) {
	userKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(userPubKey))
	if err != nil {
		return "", err
	}
	keyIdentity := host_identity + "_" + username

	currentEpoch := uint64(time.Now().Unix())
	expireEpoch := currentEpoch + (3600 * numValidHours)

	// The values of the permissions are taken from the default values used
	// by ssh-keygen
	cert := ssh.Certificate{
		Key:             userKey,
		CertType:        ssh.UserCert,
		SignatureKey:    signer.PublicKey(),
		ValidPrincipals: []string{username},
		KeyId:           keyIdentity,
		ValidAfter:      currentEpoch,
		ValidBefore:     expireEpoch,
		Permissions: ssh.Permissions{Extensions: map[string]string{
			"permit-X11-forwarding":   "",
			"permit-agent-forwarding": "",
			"permit-port-forwarding":  "",
			"permit-pty":              "",
			"permit-user-rc":          ""}}}

	err = cert.SignCert(bytes.NewReader(cert.Marshal()), signer)
	if err != nil {
		return "", err
	}
	certString, err := goCertToFileString(cert, username)
	if err != nil {
		return "", err
	}
	return certString, nil
}

func GenSSHCertFileStringFromSSSDPublicKey(userName string, signer ssh.Signer, hostIdentity string) (string, error) {

	userPubKey, err := GetUserPubKeyFromSSSD(userName)
	if err != nil {
		return "", err
	}
	cert, err := GenSSHCertFileString(userName, userPubKey, signer, hostIdentity)
	if err != nil {
		return "", err
	}
	return cert, err
}

/// X509 section
func getPubKeyFromPem(pubkey string) (pub interface{}, err error) {
	block, rest := pem.Decode([]byte(pubkey))
	if block == nil || block.Type != "PUBLIC KEY" {
		err := errors.New(fmt.Sprintf("Cannot decode user public Key '%s' rest='%s'", pubkey, string(rest)))
		if block != nil {
			err = errors.New(fmt.Sprintf("public key bad type %s", block.Type))
		}
		return "", err
	}
	return x509.ParsePKIXPublicKey(block.Bytes)
}

func getPrivateKeyFromPem(privateKey string) (priv interface{}, err error) {
	//TODO handle ecdsa and other non-rsa keys
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err := errors.New("Cannot decode Private Key")
		return "", err
	}
	switch block.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	default:
		err := errors.New("Cannot process that key")
		return priv, err
	}
}

//copied from https://golang.org/src/crypto/tls/generate_cert.go
func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	//case *ecdsa.PrivateKey:
	//	return &k.PublicKey
	default:
		return nil
	}
}

func derBytesCertToCertAndPem(derBytes []byte) (*x509.Certificate, string, error) {
	cert, err := x509.ParseCertificate(derBytes)
	if err != nil {
		return nil, "", err
	}
	pemCert := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes}))
	//fmt.Printf("pem:\n%s\n", pemCert)
	return cert, pemCert, nil
}

// return both an internal representation an the pem representation of the string
// As long as the issuer value matches THEN the serial number can be different every time
func GenSelfSignedCACert(commonName string, organization string, caPriv interface{}) (*x509.Certificate, string, error) {
	//// Now do the actual work...
	notBefore := time.Now()
	notAfter := notBefore.Add(24 * 365 * 8 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, "", err
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{organization},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,
		KeyUsage:  x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign,
		//ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(caPriv), caPriv)
	if err != nil {
		return nil, "", err
	}
	return derBytesCertToCertAndPem(derBytes)

}

// From RFC 4120 section 5.2.2 (https://tools.ietf.org/html/rfc4120)
type KerberosPrincipal struct {
	Len       int      `asn1:"explicit,tag:0"`
	Principal []string `asn1:"explicit,tag:1"`
}

// From RFC 4556 section 3.2.2 (https://tools.ietf.org/html/rfc4556.html)
type KRB5PrincipalName struct {
	Realm     string            `asn1:"explicit,tag:0"`
	Principal KerberosPrincipal `asn1:"explicit,tag:1"`
}

type PKInitSANAnotherName struct {
	Id    asn1.ObjectIdentifier
	Value KRB5PrincipalName `asn1:"explicit,tag:0"`
}

// Since currently asn1 cannot mashal into GeneralString (https://github.com/golang/go/issues/18832)
// We make this hack since we know the positions of the items we want to change
func changePrintableStringToGeneralString(kerberosRealm string, inString []byte) []byte {
	position := 16
	inString[position] = 27

	position = position + 1 + len(kerberosRealm) + 14
	inString[position] = 27

	return inString
}

func genSANExtension(userName string, kerberosRealm *string) (*pkix.Extension, error) {
	if kerberosRealm == nil {
		return nil, nil
	}
	krbRealm := *kerberosRealm

	//1.3.6.1.5.2.2
	krbSanAnotherName := PKInitSANAnotherName{
		Id: []int{1, 3, 6, 1, 5, 2, 2},
		Value: KRB5PrincipalName{
			Realm:     krbRealm,
			Principal: KerberosPrincipal{Len: 1, Principal: []string{userName}},
		},
	}
	krbSanAnotherNameDer, err := asn1.Marshal(krbSanAnotherName)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("ext: %+x\n", krbSanAnotherNameDer)
	krbSanAnotherNameDer = changePrintableStringToGeneralString(krbRealm, krbSanAnotherNameDer)
	//fmt.Printf("ext: %+x\n", krbSanAnotherNameDer)

	// inspired by marshalSANs in x509.go
	var rawValues []asn1.RawValue
	rawValues = append(rawValues, asn1.RawValue{Tag: 0, Class: 2, IsCompound: true, Bytes: krbSanAnotherNameDer})

	rawSan, err := asn1.Marshal(rawValues)
	if err != nil {
		return nil, err
	}
	sanExtension := pkix.Extension{
		Id:    []int{2, 5, 29, 17},
		Value: rawSan,
	}

	return &sanExtension, nil
}

// returns an x509 cert that is with the username in the common name
func GenUserX509Cert(userName string, userPub interface{}, caCert *x509.Certificate, caPriv interface{}, kerberosRealm *string) (*x509.Certificate, string, error) {
	//// Now do the actual work...
	notBefore := time.Now()
	notAfter := notBefore.Add(time.Duration(numValidHours) * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, "", err
	}

	sanExtension, err := genSANExtension(userName, kerberosRealm)
	if err != nil {
		return nil, "", err
	}

	// need to add the extended key usage... that is special for kerberos
	//and also the client key usage
	kerberosClientExtKeyUsage := []int{1, 3, 6, 1, 5, 2, 3, 4}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   userName,
			Organization: []string{"Keymaster"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageKeyAgreement,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		UnknownExtKeyUsage:    []asn1.ObjectIdentifier{kerberosClientExtKeyUsage},
		BasicConstraintsValid: true,
		IsCA: false,
	}
	if sanExtension != nil {
		template.ExtraExtensions = []pkix.Extension{*sanExtension}
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, caCert, userPub, caPriv)
	if err != nil {
		return nil, "", err
	}
	return derBytesCertToCertAndPem(derBytes)
}
