package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.json") // 指定配置文件路径
	viper.AddConfigPath("/")             // 查找配置文件所在的路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 读取
	fmt.Println(viper.Get("name"))

	// 覆盖设置，并写入文件
	//viper.Set("name", "xiaoCai")
	//viper.WriteConfig()

	fmt.Println(viper.GetString("example.metric.host"))

}
