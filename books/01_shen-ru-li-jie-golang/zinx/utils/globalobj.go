package utils

import (
	"encoding/json"
	"os"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

type GlobalObj struct {
	TcpServer     ziface.IServer // 全局 Server 对象
	Host          string         // 主机 IP
	TcpPort       int            // 监听端口号
	Name          string         // 服务名称
	Version       string         // 版本号
	MaxPacketSize uint32         // 数据包最大值
	MaxConn       int            // 最大连接数
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Host:          "0.0.0.0",
		TcpPort:       7777,
		Name:          "ZinxServerApp",
		Version:       "v0.4",
		MaxPacketSize: 4096,
		MaxConn:       12000,
	}
	GlobalObject.Reload()
}

// Reload 加载配置文件
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("../conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
