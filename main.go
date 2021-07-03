package main

import (
	"chianoob_server_linux/network"
	_ "chianoob_server_linux/network"
	_ "chianoob_server_linux/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	fmt.Println(network.Get_external())
	beego.SetStaticPath("download", "static/download")
	beego.Run()
}
