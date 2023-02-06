package common

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
