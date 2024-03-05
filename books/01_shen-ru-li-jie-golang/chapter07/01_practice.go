package chapter05

import "fmt"

// region -------- 抽象层 --------

// Card 显卡
type Card interface {
	Display()
}

// Memory 内存
type Memory interface {
	Storage()
}

// Cpu 处理器
type Cpu interface {
	Calculate()
}

// Computer 计算机
type Computer struct {
	cpu  Cpu
	mem  Memory
	card Card
}

// NewComputer 组装电脑
func NewComputer(cpu Cpu, mem Memory, card Card) *Computer {
	return &Computer{
		cpu:  cpu,
		mem:  mem,
		card: card,
	}
}

// Work 工作
func (c *Computer) Work() {
	c.cpu.Calculate()
	c.mem.Storage()
	c.card.Display()
}

// endregion -------- 抽象层 --------

// region -------- 实现层 --------

type InterCpu interface { // 小扩展
	Cpu
}

type CoreI7 struct {
	InterCpu
}

func (i7 *CoreI7) Calculate() {
	fmt.Println("i7 14700kf | 5.2GHz | 33M L3 ｜ 牛逼")
}

type CoreI5 struct {
	InterCpu
}

func (i5 *CoreI5) Calculate() {
	fmt.Println("i5 14600kf | 平庸 ")
}

type CoreI3 struct {
	InterCpu
}

func (i3 *CoreI3) Calculate() {
	fmt.Println("i3 .... | 垃圾 ")
}

type InterMemory struct {
	Memory
}

func (m *InterMemory) Storage() {
	fmt.Println("英特尔 32G 内存")
}

type InterCard struct {
	Card
}

func (c *InterCard) Display() {
	fmt.Println("英特尔垃圾显卡")
}

type KingstonMemory struct {
	Memory
}

func (k *KingstonMemory) Storage() {
	fmt.Println("金士顿 32G 内存")
}

type NvidiaCard struct {
	Card
}

func (n NvidiaCard) Display() {
	fmt.Println("英伟达 RTX 4090 | 牛逼")
}

// endregion -------- 实现层 --------
