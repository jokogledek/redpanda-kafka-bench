package model

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func InitConfig() (cfg *Config, err error) {
	var (
		f         *os.File
		byteValue []byte
	)

	f, err = os.Open(`config/bridge_config.json`)
	if err != nil {
		log.Errorf("[config][init] load config file, err %#v", err)
		return
	}
	defer f.Close()

	byteValue, err = ioutil.ReadAll(f)
	if err != nil {
		log.Errorf("[config][init] error read file, err %#v", err)
		return
	}
	err = json.Unmarshal(byteValue, &cfg)
	if err != nil {
		log.Errorf("[config][init] error unmarhal config file, err %#v", err)
		return
	}
	log.Info("[config][init] config loaded")
	return
}
