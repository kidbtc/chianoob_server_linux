package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"net/http"
	"os"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	DwonLoadFile(`127.0.0.1:8180/chianoob`, "D:\\chianoob_linux")
	c.TplName = "index.tpl"
}

type files struct {
	filename     string
	filemd5      string
	downloadpath string
}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Read(p)
	r.Current += int64(n)
	fmt.Println("\r进度 %.2f%%", float64(r.Current*10000/r.Total*100)/100)
	return
}

func DwonLoadFile(url, filename string) {
	if strings.Index(url, " http://") == -1 {
		url = "http://" + url
	}
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	n, err := io.Copy(f, r.Body)
	fmt.Println(n, err)
}
