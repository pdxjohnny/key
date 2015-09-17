package commands

var ConfigOptions = map[string]interface{}{
	"gen": map[string]interface{}{
		"key": map[string]interface{}{
			"value": "keys/id_rsa",
			"help":  "Key file to write to",
		},
		"len": map[string]interface{}{
			"value": 2048,
			"help":  "Length of key to be generated",
		},
	},
	"encrypt": map[string]interface{}{
		"key": map[string]interface{}{
			"value": "keys/id_rsa.pub",
			"help":  "Key to used for encryption",
		},
		"message": map[string]interface{}{
			"value": "",
			"help":  "File to encrypt blank for stdin or text",
		},
	},
	"decrypt": map[string]interface{}{
		"key": map[string]interface{}{
			"value": "keys/id_rsa",
			"help":  "Key to used for decryption",
		},
		"message": map[string]interface{}{
			"value": "",
			"help":  "File to decrypt blank for stdin or text",
		},
	},
}
