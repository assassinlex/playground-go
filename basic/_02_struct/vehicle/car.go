package vehicle

type Car struct {
}

// Speed 时速[km/h]
func (c *Car) Speed() int {
	return 120
}

// unexported 测试私有方法
func (c *Car) unexported() {
	// do nothing
}

func (c *Car) NumWheels() int {
	return 4
}

func (c *Car) String() string {
	return "this is a car"
}

func (c *Car) move() {
	// do nothing
}
