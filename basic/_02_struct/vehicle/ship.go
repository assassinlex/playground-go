package vehicle

type Ship struct {
}

// Speed 时速[km/h]
func (s *Ship) Speed() int {
	return 30
}

// unexported 测试私有方法
func (c *Ship) unexported() {
	// do nothing
}
