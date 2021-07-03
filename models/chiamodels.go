package models

var UpFileIp = ""
var DownLoadPath = ""
var Psize = float64(0)

type DiskStatus struct {
	Path    string `json:"path"`
	All     string `json:"all"`
	Used    string `json:"used"`
	Free    string `json:"free"`
	UseRate string `json:"userate"`
}

type Client struct {
	Idx        int
	Key        string
	IP         string
	IPint64    int64 //用于排序
	Name       string
	Finish     string //完成状态
	Todatwork  string //今日计数
	HHDPath    string //机械名称
	HHDAll     string //机械大小
	HHDUsed    string //空闲大小
	HHDUseRate string //使用百分比
	PoolPro    string //挖矿进程
	Error      string //异常
	PNum       int    //今日P图量
	Version    string
	ErrorState string
	Uptime     string //最后更新时间
	Ptime      string //单轮时间
	Todesk     string
}

type UpFiles struct {
	PathName string
	Md5      string
	Url      string //下载地址
}
