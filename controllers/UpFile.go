package controllers

import (
	"chianoob_server_linux/models"
	"chianoob_server_linux/myfunc"
	"chianoob_server_linux/network"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type UpFileController struct {
	beego.Controller
}

func (c *UpFileController) Get() {
	// 业务逻辑处理，例如先检测用户权限

	// 下载服务器上当前目录/data/file.zip文件， 下载后的文件名为:压缩包1.zip
	c.Ctx.Output.Download("static/download/chianoob_linux", "chianoob_linux")

}

func (c *UpFileController) Upfile() {
	key := c.GetString("key")
	network.OlineMap[key].Conn.Write([]byte("upfile"))
	fmt.Println("Send:", key, "upfile")
	c.Redirect("/server", 302)
}

func (c *UpFileController) Batchup() {
	key := c.GetString("key")
	fmt.Println("rev:", key)
	keys := strings.Split(key, "|")
	if len(keys) > 0 {
		//network.OlineMap[key].Conn.Write([]byte("upfile"))
		for _, temp := range keys {
			if temp != "" {
				network.OlineMap[temp].Conn.Write([]byte("upfile"))
				fmt.Println("Send:", temp, "upfile")
			}
		}
	}
	c.Redirect("/server", 302)
}

func (c *UpFileController) Batchrebpc() {
	key := c.GetString("key")
	fmt.Println("rev:", key)
	keys := strings.Split(key, "|")
	if len(keys) > 0 {
		for _, temp := range keys {
			if temp != "" {
				network.OlineMap[temp].Conn.Write([]byte("reboot"))
				fmt.Println("Send:", temp, "reboot")
			}
		}
	}
	c.Redirect("/server", 302)
}

func (c *UpFileController) Reboot() {
	key := c.GetString("key")
	network.OlineMap[key].Conn.Write([]byte("reboot"))
	fmt.Println("Send:", key, "reboot")
	c.Redirect("/server", 302)
}

func (c *UpFileController) Batchrebhp() {
	key := c.GetString("key")
	fmt.Println("rev:", key)
	keys := strings.Split(key, "|")
	if len(keys) > 0 {
		for _, temp := range keys {
			if temp != "" {
				network.OlineMap[temp].Conn.Write([]byte("resthpool"))
				fmt.Println("Send:", temp, "resthpool")
			}
		}
	}
	c.Redirect("/server", 302)
}

//resthpool

func (c *UpFileController) Resthpool() {
	key := c.GetString("key")
	network.OlineMap[key].Conn.Write([]byte("resthpool"))
	fmt.Println("Send:", key, "resthpool")
	c.Redirect("/server", 302)
}

func (c *UpFileController) FileMd5() {
	// 业务逻辑处理，例如先检测用户权限
	upfile := myfunc.GetFilesFromDir(`static/download/`)
	upfile = setupurl(upfile)
	upfile = setpathname(upfile, `chianoob_linux`, `/www/server/chiabee/chianoob_linux`)
	upfile = setpathname(upfile, `mount_disk.sh`, `/www/server/chiabee/mount_disk.sh`)
	fmt.Println(upfile)
	jsonStr, _ := json.Marshal(upfile) //
	c.Ctx.ResponseWriter.Write([]byte(jsonStr))

	// 下载服务器上当前目录/data/file.zip文件， 下载后的文件名为:压缩包1.zip
	//this.Ctx.ResponseWriter("")
}

func setupurl(upfile []models.UpFiles) []models.UpFiles {
	port := beego.AppConfig.String("httpport")
	if models.DownLoadPath == "" {
		models.DownLoadPath = network.Get_external() + ":" + port
	}
	for i := 0; i < len(upfile); i++ {
		upfile[i].Url = models.DownLoadPath + "/download/" + upfile[i].PathName
	}
	return upfile
}

func setpathname(upfile []models.UpFiles, name, newname string) []models.UpFiles {
	for i := 0; i < len(upfile); i++ {
		if upfile[i].PathName == name {
			upfile[i].PathName = newname
		}
	}
	return upfile
}
