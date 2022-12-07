package common

type Properties struct {
	MysqlConfig MysqlProperties `yaml:"mysqlConfig"`
}

type MysqlProperties struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"userName"`
	Pwd      string `yaml:"pwd"`
}
