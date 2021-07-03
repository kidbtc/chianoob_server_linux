package controllers

import (
	"chianoob_server_linux/models"
	"chianoob_server_linux/network"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type ChiaPlotterStaticController struct {
	beego.Controller
}

func (c *ChiaPlotterStaticController) Getall() {
	models.Psize = 0
	chiaclients := []models.Client{}
	//chiaclientsort := []models.ChiaClient{}

	network.Lock.Lock()
	olineMap := network.OlineMap //全局赋值
	//名称 ip 状态 完成度 当前目标 今日量 矿池监控 设置 版本号
	//st = "chia"+"|"+clinet.Name+"|"+clinet.IP+"|"+clinet.Error + "|" + clinet.Finish +"|"+ hhddisk + "|" + clinet.PNum+"|"+clinet.PoolPro+"|"+models.Version
	//i:=0

	for k, _ := range olineMap {
		client := models.Client{}
		temp := strings.Split(olineMap[k].Msg, "|")
		if len(temp) > 7 {
			client.Name = temp[1]
			client.IP = temp[2]
			client.Key = k
			client.Error, client.ErrorState = geterror(client)
			client.Finish = temp[4]
			client.HHDPath, client.HHDUsed, client.HHDAll, client.HHDUseRate = ReadHddInfo(temp[5])
			client.PNum, _ = strconv.Atoi(temp[6])
			client.PoolPro = temp[7]
			client.Version = temp[8]
			client.Uptime = temp[9]
			client.Ptime = temp[10]
			if len(temp) == 12 {
				client.Todesk = temp[11]
			}
			models.Psize = models.Psize + float64(client.PNum)
			//client.Error,client.ErrorState=geterror(client)
			if strings.Index(client.Name, ".") != -1 {
				client.IPint64 = network.IpFormat(client.Name) //格式化IP 转int64
			}
			if strings.Index(client.Name, "-") != -1 {
				client.IPint64 = network.LabelFormat(client.Name) //格式化IP 转int64
			}
		}
		chiaclients = append(chiaclients, client)
	}
	network.Lock.Unlock()

	c.Data["Chiastatic"] = BubbleAsort(chiaclients)
	models.Psize, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", models.Psize/1024), 64)
	c.Data["Psize"] = models.Psize
	c.Data["Line"] = len(chiaclients)
	c.TplName = "Console.html"
}

//解析当前硬盘信息
func ReadHddInfo(diskinfostr string) (path, used, all, userate string) {
	temparray := strings.Split(diskinfostr, "/")
	if len(temparray) == 4 {
		path = temparray[0]
		used = temparray[1]
		all = temparray[2]
		userate = temparray[3]
	}
	return path, used, all, userate
}

//排序
func BubbleAsort(values []models.Client) []models.Client {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i].IPint64 > values[j].IPint64 {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	for i := 0; i < len(values); i++ {
		values[i].Idx = i + 1
	}
	return values
}

//异常状态
func geterror(chiaClient models.Client) (str string, state string) {
	//diskconfig, _ :=beego.AppConfig.Int("errordisk")
	//if len(strings.Split(chiaClient.HDD,"/")) < diskconfig{
	//	return "硬盘丢失","danger"
	//}
	//if len(strings.Split(chiaClient.HDD,"/")) == 7{
	//	return "硬盘丢失","danger"
	//}
	return "正常", "success"
}
