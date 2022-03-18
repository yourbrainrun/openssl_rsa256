package config

import (
	"os"
	"sync"
)

type RsaConf struct {
	PrivatePath string
	PublicPath  string
}

var rsaConfInstance RsaConf
var rsaConfOnce sync.Once

func GetRsaConf() RsaConf {
	rsaConfOnce.Do(func() {
		curr, _ := os.Getwd()
		rsaConfInstance.PrivatePath = curr + "/storage/pri.pem"
		rsaConfInstance.PublicPath = curr + "/storage/pub.pem"
	})

	return rsaConfInstance
}
