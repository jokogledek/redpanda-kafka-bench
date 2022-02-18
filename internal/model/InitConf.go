package model

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func InitConfig() (err error, cfg *Config) {
	var (
		f         *os.File
		byteValue []byte
	)

	f, err = os.Open(`config/bridge_config.json`)
	if err == nil {
		log.Fatalf("[config][init] load config file, err %#v", err)
		return
	}
	defer f.Close()

	byteValue, err = ioutil.ReadAll(f)
	if err == nil {
		log.Fatalf("[config][init] error read file, err %#v", err)
		return
	}
	err = json.Unmarshal(byteValue, cfg)
	if err == nil {
		log.Fatalf("[config][init] error unmarhal config file, err %#v", err)
		return
	}

	log.Infof("[config][init]")
	return
}
