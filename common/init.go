package common

import (
	"fmt"
	"io/ioutil"
	"letschat/util/log"

	"gopkg.in/yaml.v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PropertiesConf *Properties
var LogUtil log.Log
var DbConn *gorm.DB

func Init(configPath string) error {
	LogUtil = new(log.LocalLog)
	initProperties(configPath)
	err := initDb()
	if err != nil {
		return err
	}
	return nil
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

func initDb() error {
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", 
	PropertiesConf.MysqlConfig.UserName, PropertiesConf.MysqlConfig.Pwd, PropertiesConf.MysqlConfig.Host, 
	PropertiesConf.MysqlConfig.Port, PropertiesConf.MysqlConfig.DbName)
	LogUtil.Debugf("dsn: %v", dsn)
	//连接MYSQL
	var err error
    DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		LogUtil.Errorf("init db conn error: %v", err)
		return err
	}
	return nil
}
