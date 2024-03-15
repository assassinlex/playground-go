package znet

import (
	"fmt"
	"playground/books/01_shen-ru-li-jie-golang/zinx/utils"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
	"strconv"
)

// MsgHandle 请求处理器
type MsgHandle struct {
	Apis           map[uint32]ziface.IRouter
	WorkerPoolSize uint32                 // worker 数量
	TaskQueue      []chan ziface.IRequest // worker 消息队列
}

// NewMsgHandle 构造器
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Printf("api msg-id: %d is nout found\n", request.GetMsgID())
	}
	// handler.PreHandle(request)
	handler.Handle(request)
	// handler.PostHandle(request)
}

// AddRouter 添加消息处理逻辑
func (m *MsgHandle) AddRouter(id uint32, router ziface.IRouter) {
	if _, ok := m.Apis[id]; ok {
		panic("repeated api: " + strconv.Itoa(int(id)))
	}
	m.Apis[id] = router
	fmt.Printf("api %d registed.\n", id)
}

func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

// SendMsgToTaskQueue 向任务队列发送请求
func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 对连接 id 取模的结果来决定将任务分配到对应的 worker(goroutine) 上
	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	m.TaskQueue[workerID] <- request
	fmt.Printf(
		"add connection id = %d, msg id = %d to woker id = %d\n",
		request.GetConnection().GetConnID(),
		request.GetMsgID(),
		workerID,
	)
}

// StartOneWorker 启动一个 worker
func (m *MsgHandle) StartOneWorker(wid int, taskQueue chan ziface.IRequest) {
	fmt.Printf("worker id = %d is started\n", wid)
	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}
