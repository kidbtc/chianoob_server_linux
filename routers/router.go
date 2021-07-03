package routers

import (
	"chianoob_server_linux/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/server", &controllers.ChiaPlotterStaticController{}, "*:Getall")
	//UpFileController
	beego.Router("/batchup", &controllers.UpFileController{}, "*:Batchup")          //批量升级
	beego.Router("/batchrebootpc", &controllers.UpFileController{}, "*:Batchrebpc") //批量重启机器
	beego.Router("/batchreboothp", &controllers.UpFileController{}, "*:Batchrebhp") //批量重启挖矿
	beego.Router("/upfile", &controllers.UpFileController{}, "*:Upfile")
	beego.Router("/chianoob", &controllers.UpFileController{}, "*:Get")
	beego.Router("/filemd5", &controllers.UpFileController{}, "*:FileMd5")
	beego.Router("/reboot", &controllers.UpFileController{}, "*:Reboot")
	beego.Router("/resthpool", &controllers.UpFileController{}, "*:Resthpool")

}
