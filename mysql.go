package common

import "github.com/micro/go-micro/v2/config"

type MySqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

//获取mysql配置
func GetMysqlFromConsul(cfg config.Config, path ...string) (*MySqlConfig, error) {
	mysqlCfg := &MySqlConfig{}
	err := cfg.Get(path...).Scan(mysqlCfg)
	if err != nil {
		return mysqlCfg, err
	}
	return mysqlCfg, nil
}
