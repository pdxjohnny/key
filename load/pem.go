package load

import (
  "encoding/pem"

)

func PEM(data string) *rsa.PublicKey {
  // https://golang.org/pkg/encoding/pem/
	block, rest := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}
}
