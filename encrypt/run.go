package encrypt

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/key/load"
)

func Run() {
	publicKey, err := load.LoadPublic(viper.GetString("key"))
	if err != nil {
		log.Println("ERROR: encrypt.Run error loading public key: ", err)
		return
	}
	message, err := Encrypt(
		publicKey,
		[]byte(viper.GetString("message")),
	)
	if err != nil {
		log.Println("ERROR: encrypt.Run error encrypting: ", err)
		return
	}
	fmt.Printf("%s", string(message))
	return
}
