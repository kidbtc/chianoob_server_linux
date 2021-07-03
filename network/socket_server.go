package network

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//保存在线用户
type Client struct {
	Conn net.Conn
	Name string
	Addr string
	Msg  string
}

var message = make(chan string)
var OlineMap = make(map[string]Client)
var quitChan = make(chan bool)
var Lock sync.Mutex

func init() {
	go func() {
		fmt.Print("服务器已启动")
		port := beego.AppConfig.String("socketport")
		listener, err := net.Listen("tcp", ":"+port) //创建socket,绑定绑定端口,实现监听
		if err != nil {
			fmt.Println(err)
		}
		defer listener.Close() //延迟关闭socket
		for {
			conn, err := listener.Accept() //创建连接
			if err != nil {
				fmt.Println("Listen.Accept err= ", err)
				continue
			}
			go HandleCoon(conn) //创建协程处理连接
		}
	}() //数据	监听
}

//广播消息
//func Manager()  {
//	OlineMap = make(map[string]Client)
//	for {
//		msg := <- message
//		for _,cli:=range  OlineMap{
//			cli.C <- msg
//		}
//	}
//}

func HandleCoon(conn net.Conn) {
	defer conn.Close()                    //延迟关闭连接
	cliAddr := conn.RemoteAddr().String() //客户端的IP和端口号
	beego.Info("客户端:" + cliAddr + "已连接")
	cli := Client{conn, cliAddr, cliAddr, cliAddr}
	Lock.Lock()
	OlineMap[cliAddr] = cli
	Lock.Unlock()
	isQuit := make(chan bool)
	hasData := make(chan bool)
	go func() {
		for {
			var buf []byte = make([]byte, 2048)
			n, err := conn.Read(buf)
			if n == 0 { //对方断开或者出异常
				isQuit <- true
				fmt.Println("conn.Read err ", err)
				fmt.Println(cliAddr, " 关闭连接")
				return
			}
			msg := string(buf[:n])
			cli.Msg = msg
			Lock.Lock()
			OlineMap[cliAddr] = cli
			Lock.Unlock()
			hasData <- true
		}
	}()

	for {
		select {
		case <-isQuit:
			Lock.Lock()
			delete(OlineMap, cliAddr) //当前用户从map移除
			Lock.Unlock()
			//此处可以写退出日志
			fmt.Println(cliAddr, " 当前用户从map移除")
			return
		case <-hasData:
			conn.Write([]byte("1"))
			beego.Info(cli.Msg)
		case <-time.After(180 * time.Second):
			Lock.Lock()
			delete(OlineMap, cliAddr) //当前用户从map移除
			Lock.Unlock()
			fmt.Println(cliAddr, " 超时关闭")
			//此处可以写退出日志
			return
		}
	}
}

//IP格式化
func IpFormat(ip string) int64 {
	//192.168.1.2
	iparry := strings.Split(ip, ".")
	if len(iparry) != 4 {
		return 999999999999
	}
	for i := 0; i < len(iparry); i++ {
		if len(iparry[i]) == 1 {
			iparry[i] = "00" + iparry[i]
		}
		if len(iparry[i]) == 2 {
			iparry[i] = "0" + iparry[i]
		}
	}
	ip64, _ := strconv.ParseInt(iparry[0]+iparry[1]+iparry[2]+iparry[3], 10, 64)
	return ip64
}

//IP格式化
func LabelFormat(name string) int64 {
	//192.168.1.2
	intlable, _ := strconv.Atoi(strings.Replace(name, "-", "", -1))
	return int64(intlable)
}

//取外网ip
func Get_external() string {
	responseClient, errClient := http.Get("http://ip.dhcp.cn/?ip") // 获取外网 IP
	if errClient != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		return ""
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer responseClient.Body.Close()

	body, _ := ioutil.ReadAll(responseClient.Body)
	clientIP := fmt.Sprintf("%s", string(body))
	return clientIP
}
