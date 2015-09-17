package gen

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"log"
	"os"
)

func Save(outFile string, save interface{}) {
	keyFile, err := os.Create(outFile)
	defer keyFile.Close()
	if err != nil {
		log.Println(err)
		return
	}
	keyEncoder := gob.NewEncoder(keyFile)
	keyEncoder.Encode(save)
}

func Gen(outFile string, length int) {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, length)
	if err != nil {
		log.Println("ERROR: key.Gen generating private:", err)
		return
	}

	// Note : Only used for 3rd and subsequent primes
	//log.Printf("CRTValues : Exp[%s]\n Coeff[%s]\n R[%s]\n", CRTVal[2].Exp.String(), CRTVal[2].Coeff.String(), CRTVal[2].R.String())

	// Note : if you want to have multi primes,
	//        use rsa.GenerateMultiPrimeKey() function instead of
	//        rsa.GenerateKey() function
	//        see http://golang.org/pkg/crypto/rsa/#GenerateMultiPrimeKey

	// http://golang.org/pkg/crypto/rsa/#PrivateKey.Precompute
	privateKey.Precompute()

	// http://golang.org/pkg/crypto/rsa/#PrivateKey.Validate
	err = privateKey.Validate()
	if err != nil {
		log.Println("ERROR: key.Gen validating private:", err)
		return
	}
	Save(outFile, privateKey)

	var publicKey *rsa.PublicKey
	publicKey = &privateKey.PublicKey
	Save(outFile+".pub", publicKey)
}
