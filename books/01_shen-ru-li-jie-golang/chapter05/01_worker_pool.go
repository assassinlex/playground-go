package chapter05

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// WorkerManager 管理 goroutine
type WorkerManager struct {
	workerChan chan *worker // 监控 worker 是否已经死亡
	nWorkers   int          // 监控 worker 数量
}

// NewWorkerManager WorkerManager 构造器
func NewWorkerManager(nWorkers int) *WorkerManager {
	return &WorkerManager{
		workerChan: make(chan *worker, nWorkers),
		nWorkers:   nWorkers,
	}
}

func (wm *WorkerManager) StartWorkerPool() {
	for i := 0; i < wm.nWorkers; i++ {
		_worker := &worker{id: i}
		go _worker.Work(wm.workerChan)
	}
	wm.KeepLiveWorkers() // 启动保活监控
}

// KeepLiveWorkers worker 保活
func (wm *WorkerManager) KeepLiveWorkers() {
	for _worker := range wm.workerChan {
		if _worker.err == nil {
			fmt.Printf("worker %d succeed.\n", _worker.id)
		} else {
			fmt.Printf("worker %d stopped with err: [%v]\n", _worker.id, _worker.err)
			_worker.err = nil
		}
		go _worker.Work(wm.workerChan)
	}
}

// worker 子 goroutine
type worker struct {
	id  int   // goroutine id
	err error // 错误
}

// Work 执行任务
func (w *worker) Work(workerChan chan<- *worker) (err error) {
	defer func() { // 向 WorkerManager 发送退出通知
		var ok bool
		if r := recover(); r != nil {
			if err, ok = r.(error); ok {
				w.err = err
			} else {
				w.err = fmt.Errorf("panic happend with [ %v ]", r)
			}
		} else {
			w.err = err
		}
		workerChan <- w // 通知 WorkerManager, 当前 worker 已退出
	}()

	// 执行任务
	time.Sleep(3 * time.Second)               // 模拟任务耗时
	fmt.Printf("start worker id: %d\n", w.id) // 任务实际逻辑

	// 模拟任务
	if v := rand.Intn(10); v >= 7 { // 70% 任务成功退出
		runtime.Goexit()
	} else { // 70% 任务异常退出
		panic("worker panic")
	}
	return err
}
