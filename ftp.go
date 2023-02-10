package common

import "github.com/micro/go-micro/v2/config"

type FtpConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Port int64  `json:"port"`
}

//获取mysql配置
func GetFtpFromConsul(cfg config.Config, path ...string) (*FtpConfig, error) {
	ftpCfg := &FtpConfig{}
	err := cfg.Get(path...).Scan(ftpCfg)
	if err != nil {
		return ftpCfg, err
	}
	return ftpCfg, nil
}
