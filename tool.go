package common

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"math/rand"
	"strconv"
	"time"
)

//通过json tag对结构体赋值
func SwapTo(req, category interface{}) (err error) {
	dataBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataBytes, category)
}

//6位验证码生成
func Captcha() string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}
	return code
}

//初始化
type MicroOptions struct {
	ConsulCfg      *config.Config
	ConsulRegister *registry.Registry
	DB             *gorm.DB
}

func Init() (*MicroOptions, error) {
	//配置中心
	consulCfg, err := GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		return nil, err
	}
	//注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	},
	)
	//获取mysql配置
	mysqlInfo, err := GetMysqlFromConsul(consulCfg, "mysql")
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+
		"@tcp("+mysqlInfo.Host+":"+strconv.FormatInt(mysqlInfo.Port, 10)+
		")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return &MicroOptions{
		ConsulCfg:      &consulCfg,
		ConsulRegister: &consulRegister,
		DB:             db,
	}, nil
}

//string转json再转map[string]interface{}
func HttpJsonToMap(body string) (*map[string]interface{}, error) {
	var post map[string]interface{}
	err := json.Unmarshal([]byte(body), &post)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	param := post["params"].(map[string]interface{})
	return &param, nil
}
