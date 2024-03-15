package utils

import (
	"encoding/json"
	"os"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

type GlobalObj struct {
	TcpServer        ziface.IServer // 全局 Server 对象
	Host             string         // 主机 IP
	TcpPort          int            // 监听端口号
	Name             string         // 服务名称
	Version          string         // 版本号
	MaxPacketSize    uint32         // 数据包最大值
	MaxConn          int            // 最大连接数
	WorkerPoolSize   uint32         // 当前工作池 worker 数量
	MaxWorkerTaskLen uint32         // 工作 worker 对应负责的任务队列最大任务存储数量
	ConfFilePath     string
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Host:             "0.0.0.0",
		TcpPort:          7777,
		Name:             "ZinxServerApp",
		Version:          "v0.4",
		MaxPacketSize:    4096,
		MaxConn:          12000,
		ConfFilePath:     "conf/zinx.json",
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
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
