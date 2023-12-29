package vehicle

type MotoBike struct {
}

// Speed 时速[km/h]
func (m *MotoBike) Speed() int {
	return 80
}

// unexported 测试私有方法
func (c *MotoBike) unexported() {
	// do nothing
}
