package crypto4go

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRSAEncrypt(t *testing.T) {
	var pub = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZ8lmnG4TMFqJty0FoqAHKxEIs
IcZaM2E8PXjXoQ+iszREsqi1B5Lmq7GeJ1/9N+OGDIjpDHnEfMMlHrj+5gYSTPab
bLGCvtcluPbI6R+uJz3uYGtPzqn4EKiNvC1ixANLmbhdqbb3KAkCcRltZOZYSerG
VE069nKjmSWRRlWJkQIDAQAB
-----END PUBLIC KEY-----`)

	var pri = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDZ8lmnG4TMFqJty0FoqAHKxEIsIcZaM2E8PXjXoQ+iszREsqi1
B5Lmq7GeJ1/9N+OGDIjpDHnEfMMlHrj+5gYSTPabbLGCvtcluPbI6R+uJz3uYGtP
zqn4EKiNvC1ixANLmbhdqbb3KAkCcRltZOZYSerGVE069nKjmSWRRlWJkQIDAQAB
AoGBAKGAiRbfuYRSsYKSv6GB/fH3hOGXFZj5wfAVzVpcK23xRaYyjfm35w+v4yrD
GspVg/BtkXbAm+sSWLlFDuk0IwJGhqxmIJ6TvxRFDxep4Kz5LjPlTvG1/mvSKKW6
1uUZquT2Ll0qy2hwyui8K08+CWVxAO5NMskYHaztF8QArjgBAkEA8UtE8pawG/Cj
24kj//y10f36Yt9tQ00/Nu7hfLXJe292zWn0cCEZG2ukkt6kQQtNoUBpRMTj9cQR
hd+2hPgX8QJBAOc6z0sJUWwG6m13nSVlu6j2wmZTp6W9U52WNR364L1UDfn7MI/X
7roW1SdwQwdYNVgwvt2N1MNJheTxt5/ZK6ECQQCs3Ta08J2UNq69LZ+72ejMWz7R
LK3TZHjgOv0R4g5JPw6GlNzIo/2ftls92QEllBp2ZnXEDaYewOuo1B+nXTGRAkEA
h8qCp+dN+KnLDAQ9thObdCuNmHgyMOQRca8ffH6zcpwlJRP9vcuqd4AnJ2UHCA4m
Ladav1OmihToW74T/vyTYQJAAf+PKLRD+O2CuwcZJG0taEqL0RR+kXMEd0wLp4EN
pFwKUHMh+rm4/Asgy126+rS6Hr0QuNuoJuQbAr3Q28h7PQ==
-----END RSA PRIVATE KEY-----`)

	var c, _ = RSAEncrypt([]byte("Sent when the application is about to move from active to inactive state. This can occur for certain types of temporary interruptions (such as an incoming phone call or SMS message) or when the user quits the application and it begins the transition to the background state."), pub)
	fmt.Println(hex.EncodeToString(c))

	var p, _ = RSADecryptWithPKCS1(c, pri)
	fmt.Println(string(p))
}
