package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	IP     string
	TapiIP string
}

var TSearchConfig Config

func InitSearchConfig(strConfigFile *string) error {

	// 读取文件
	file, err := os.Open(*strConfigFile)
	if err != nil {
		fmt.Printf("open file %s error.\n", *strConfigFile)
		return err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read file error:%v\n", err)
		return err
	}

	// 解析数据
	err = json.Unmarshal([]byte(data), &TSearchConfig)
	if err != nil {
		fmt.Println("Unmarshal data:", []byte(data))
		fmt.Println("Unmarshal error:", err)
		return err
	}

	fmt.Println(TSearchConfig)

	return nil

}
