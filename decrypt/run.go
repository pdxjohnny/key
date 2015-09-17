package decrypt

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/key/load"
)

func Run() {
	message, err := ioutil.ReadFile(viper.GetString("message"))
	if err != nil {
		log.Println("ERROR: encrypt.Run error loading message: ", err)
		return
	}
	publicKey, err := load.LoadPrivate(viper.GetString("key"))
	if err != nil {
		log.Println("ERROR: encrypt.Run error loading private key: ", err)
		return
	}
	message, err = Decrypt(
		publicKey,
		message,
	)
	if err != nil {
		log.Println("ERROR: decrypt.Run error decrypting: ", err)
		return
	}
	fmt.Println(string(message))
	return
}
