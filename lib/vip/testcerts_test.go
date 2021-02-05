package vip

const rootCAPem = `-----BEGIN CERTIFICATE-----
MIIE1jCCAr4CAQowDQYJKoZIhvcNAQELBQAwMTELMAkGA1UEBhMCVVMxEDAOBgNV
BAoMB1Rlc3RPcmcxEDAOBgNVBAsMB1Rlc3QgQ0EwHhcNMjEwMjA1MDA1NDUwWhcN
NDEwMTMxMDA1NDUwWjAxMQswCQYDVQQGEwJVUzEQMA4GA1UECgwHVGVzdE9yZzEQ
MA4GA1UECwwHVGVzdCBDQTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
AMepGyjzJkXS3gCBStImA3I/OY6us1MqzqQGeRJIL/13n9PHqPfXIjadJ7z0krFB
xXGT2WyVRhoc+tP1K8rfwRjNtMDjxtzphDLaUSvjhB5FvOhR0Ch6ByrFlP507FVa
MeQo+8Bex9q0SaHXtGLM5hWuYQpG6dpOlXl6DKQwqYjptzL/RfMwJKukaoc0J5e+
3xNR3XhC+ejMHp99rEORC4aLhfWhwfOhtIslIVI/HXKu72/lxIMpkhaoObBgqdQC
3e5KMg+/liZ6y/HOhN49pKuNq1SfKy6tOpI4fwZnWvKviNx8E16P4JlqPdLPAULb
aGZH7FaAwtd3bBeKXRTsasZLfEA/K2AOIucquehEH8Nkprjn6XlDbZf/1zfAqID/
lk/QTo3uhyYjj4O3G8VSZ9PCQF/n918FDOY6DtnjmY3z0ZjAZv95sf8JBbFZVHSF
ovmsWpWbRXmgorswY8yeL/jWCJ5Zp0SQhT2Bj2OmwjmNgDNnbp9BxjGHpmfJ/pxo
lJ1vDfITYYOWWejBQE6A+fsk4CQ/0IwE4jNA9MI/03yXv8qoT2eEOTsovIYMpn/9
0UjgufiiBfYHXXlAVsdHcHDfL3RxpJpFJC6M4xVLNQ0nF4gfOP4Ek2JhWEntCbTZ
je9xcvebgb675DuaAOim0fkcbPz5VRMXuw6FjaX/x6j1AgMBAAEwDQYJKoZIhvcN
AQELBQADggIBADng/A4RceeZKarKnqv3+TVBmxUrnbzUdwlrRwGFtEQZ1cVG53zP
hlfnP/zyMsyG77utJbcaKaAqOrf0XHiiDNKe/ScfCd7K0pdoS52jhenvCD8E7hko
9Ut4lCec4SFXTpeXaP0E39plmaUfB+gStZqCVxb+kOfV1C7Q7fYnUI9WzJojm2G4
IsRIEIlYn2qW7NTWUdVVprHGZgEE18MxB8aDmJV0VW/WbiiJmlNDDjHYznk5q33T
SDIRzQUEul6Zo9QUEJYy9kYHYZwhI6BRuxEJzqt2fFiGBpJ5YSwD11GDOu8SMumV
J9WwfjRpMbVjRKw0l7vQA0PrnsXt4oYzzuRUv9IQovA6CGzopCModOddrY3B1me6
e5hGdIexZUFQthu0LFpdp1fNJPoGCtKsRKXgaqTDkt0Di4lOwwzZ0CrvEqZNKvjU
keeeyuornLpaC218Qtz1AHWGLKqCfyZ0VuuNe7lrH/nGkWFzE91zr8bxiQJNGDOl
V1RUB/s4hQwCRVgiYUdKsTsD01BBcsPVBFR20Xawij5LlsFi3Dx4fIYq/Y+aa9KV
EXQO2FLxSHa0+BeKbQYjlyTcLAz0Q1qiHnUBU2Id7UQuAM25z1F4a6roHKiADxA5
AZ07E2gWGFQWsuEwC1+IUrv3QZH2Z9fNgg3M4OLTfqPr4uZNkqwQ9FN2
-----END CERTIFICATE-----`

const localhostCertPem = `-----BEGIN CERTIFICATE-----
MIID/DCCAeSgAwIBAgIFAt3gJswwDQYJKoZIhvcNAQELBQAwMTELMAkGA1UEBhMC
VVMxEDAOBgNVBAoMB1Rlc3RPcmcxEDAOBgNVBAsMB1Rlc3QgQ0EwHhcNMjEwMjA1
MDA1NDUwWhcNNDEwMTMxMDA1NDUwWjAUMRIwEAYDVQQDDAlsb2NhbGhvc3QwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDjSMe7vDHmJPbPkZEaS6glrhhS
32jG3k+Prr0ikY3fmYLFkugfM/8LZ5W9Q03EMSxyvxKH8mo1lpM6VJqWhXmyL04K
zCNQvihSF4pQDJwAB8/4BC0N8DB/s2hQyfTrMAa4z4BnJN2cF0hcAXrMJPpi+0GE
YfTSNn0BaEZMgs7cmrMwB+1m9UL02pOrwVQE6C7ddCsNis+zHP7RZ+qsD7ruv3Rr
eyzlOJJpPpB58lcIMslttUu9+yMHKZ3Zzheiw0GeMWYFthmVWij+Qazjf7MR7for
DJrBHwEhuw/USBdCC3hSLnFKG630znA0+mhEt+p1j+HGNqjSG7IDJtEhieZtAgMB
AAGjODA2MAkGA1UdEwQCMAAwFAYDVR0RBA0wC4IJbG9jYWxob3N0MBMGA1UdJQQM
MAoGCCsGAQUFBwMBMA0GCSqGSIb3DQEBCwUAA4ICAQC9ZAS+Do5OFQjpMAzWRATv
togcVhgwdSHJWYPzO+NyMX1Zsr3aFcuoFuYW5DduGb8dRzXC1IOof9J8rMbubqY0
+wZhpZ4fuxd2n08izhpk+033W+385cCjloqYZXzo7L9vHJK4M5/n68RIdmFjIxD6
uvJ0LrnBCEWL6VNjwiFoa7VxXLhAR797l5xEOXP2AVFA3lwnOxbKTgMU10VYk3XC
raNZS4CByDsNJU67IeKz/4nlPjltttKJATfEsL6zmre4GWSidMVEv5y75fA5Hogp
gKP+eXoAFNk69PXpULtvzpEqCzZbZfcWMFULczPv4RvGD/q18PDDJ7MCsflTrADf
OMY16QGhcsb1tjZtk9JYU7CIYHKeljIEW1vmfA6xOjyXTavAR9XHQRNpXTpdd1C9
GQJ8Y70jYoEE/J0GjEZEgufCOdqBEC1mQV68cfNiYV1LiyqZrC83YSuA0UdhvTEV
O15LJQZcVK2ycpKu85FaVLhwlrJUgWuN1227VFS42ApSJSCLOIgNhN20wwJFx1Ax
gkgoHV/70q0CtR0ZIDWvLwA/fXZQNboAurPR0/YMtn63PDYaGThnrnxSTaKfYBhm
lGJaD3a4/wRoZtWYkVLKa8DO3k28P9WZn9gAidGvkHgmSLYa/TftPr6SnkP4QE8C
fdyL8KJZCtk2TBQRU7U1FQ==
-----END CERTIFICATE-----`

const localhostKeyPem = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDjSMe7vDHmJPbP
kZEaS6glrhhS32jG3k+Prr0ikY3fmYLFkugfM/8LZ5W9Q03EMSxyvxKH8mo1lpM6
VJqWhXmyL04KzCNQvihSF4pQDJwAB8/4BC0N8DB/s2hQyfTrMAa4z4BnJN2cF0hc
AXrMJPpi+0GEYfTSNn0BaEZMgs7cmrMwB+1m9UL02pOrwVQE6C7ddCsNis+zHP7R
Z+qsD7ruv3RreyzlOJJpPpB58lcIMslttUu9+yMHKZ3Zzheiw0GeMWYFthmVWij+
Qazjf7MR7forDJrBHwEhuw/USBdCC3hSLnFKG630znA0+mhEt+p1j+HGNqjSG7ID
JtEhieZtAgMBAAECggEANPpF9DCNwQNktEVN+T8hVocaDFce6Rvwl6E3XNNoqnHx
1XPEv2EzVckzUgZaHh0IK0b2XtHWh98tOi92SRebojj8i+/BW5S3eUsbZkVKJ4eq
pjjy94jBBUMgUZSBjkCHkJRtbZ1eOoAC0c9WfDcyVdgTEXRzyHwC700K9dKXvTux
gJAvI7KA9kCy86793tbooKUNFWIh4sVTA8Bg78/6XTrXSgOnn1u5433HX+j8eQBD
azlA038hsJwMgDL5EYIL8lLp7PqnFzExQ+KToK3+MpRhqLUUs8dYRrxxWWJ86GTn
NeVA+wLVmNN8CQYNdoiPNCRLzxycLUteGpEGyRkiUQKBgQD4ewiNqpX+GQbt+3g/
0JvDLTUKMJt6T5VF9m1V/BjzDn27gBnWZ2AhmokXtHd2Lln0R98f4vIr/pNvsh5e
i+/YEesqAUB7fGeyGP8fyWv4JmJ2LW9U3Q4gSHveX5brsr37aclFQwCCLgSF4Sps
b035mZ7LRCGtbMhK67ZjEIgKOwKBgQDqKYpFmEfdAvJC4AwvypLy1cHQcwkWKUqf
99A0SlOEsD2f5tb0K5tJEeG7k0QbEc60G+eINwFezivp+mF5VR1TWo4F2zlrqFOW
hGZJvvOj2d88cqzvZ1D2EDO6pGwnjnfK4znR/2aBlbPHqaWk4bbMg+963ekWT66C
z5gXSjAfdwKBgQC9BOrGyaN8PhGVa5xX/xreaD9W1w8PBgcpx4H4zOM10io3PgVp
og3wwhvTmobdsfjf6OSE3yEV8ny5lSehCJaVKVxZAcbJqSNAyd8ZkSG/DYQNHhb4
2Yfravg/ezvVZI08+YYQWB3E6birF1QsVKdeXkv7mIJ96KicxZ5i0aqLDQKBgQCm
NwZnkEiyoTzGK7N62RimNwV0qbNxf63xDz9DXt8Z+OBFeN5sQ+feXksVktttO562
Snh3QFRAr6iMVaaLMde3KzhU/+AqgzmSRjk1b84SKL1Fc0E+TuAxSh7uBXigO4jd
A54valOg+Fq9B7cE6xOO6Lg9RvgFVlpRkHotTgQK0QKBgFHI5lMXMYydeAT59zb5
mIFoFLOJ+JZ4nvoNu1uFQCVKH0bARK3t8oWL7ZCRtc8vANiFS8I/JTQVO1HhW3NE
fjl8EwIhuijeojU23fzVlFdXjGGzXqHO5Bm0nDuRV6XirMF5+Lh6w+y8UYLketwc
4ru0SF3ayA6bVHGXEeS5TOkv
-----END PRIVATE KEY-----`


