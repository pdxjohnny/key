package gen

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"log"
	"os"
)

func Gen(outFile string, length int) {
	// Generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, length)
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
	privatekey.Precompute()

	// http://golang.org/pkg/crypto/rsa/#PrivateKey.Validate
	err = privatekey.Validate()
	if err != nil {
		log.Println("ERROR: key.Gen validating private:", err)
		return
	}

	var publickey *rsa.PublicKey
	publickey = &privatekey.PublicKey

	// save private and public key separately
	privatekeyfile, err := os.Create(outFile)
	if err != nil {
		log.Println(err)
		return
	}
	privatekeyencoder := gob.NewEncoder(privatekeyfile)
	privatekeyencoder.Encode(privatekey)
	privatekeyfile.Close()

	publickeyfile, err := os.Create(outFile+".pub")
	if err != nil {
		log.Println(err)
		return
	}

	publickeyencoder := gob.NewEncoder(publickeyfile)
	publickeyencoder.Encode(publickey)
	publickeyfile.Close()
}
