package app

import "fmt"

func Start() {
	fmt.Println("HELLO WORLD 2")
	ConfigInit()
	fmt.Println(AppConfig.Test)
}
