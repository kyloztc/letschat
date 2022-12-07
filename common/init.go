package common

import (
	"io/ioutil"
	"letschat/util/log"

	"gopkg.in/yaml.v2"
)

var PropertiesConf *Properties
var LogUtil log.Log

func Init() {
	LogUtil = new(log.LocalLog)
	initProperties("conf/properties.yaml")
}

func initProperties(filePath string) error {
	PropertiesConf = new(Properties)
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		LogUtil.Errorf("read config file error: %v", err)
		return err
	}
	err = yaml.Unmarshal(configFile, &PropertiesConf)
	if err != nil {
		LogUtil.Errorf("yaml unmarshal error: %v", err)
		return err
	}
	return nil
}
