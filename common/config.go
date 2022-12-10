package common

type Properties struct {
	MysqlConfig MysqlProperties `yaml:"mysqlConfig"`
}

type MysqlProperties struct {
	Host     string `yaml:"host"`
	Port     int `yaml:"port"`
	UserName string `yaml:"userName"`
	Pwd      string `yaml:"pwd"`
	DbName string `yaml:"dbName"`
}
