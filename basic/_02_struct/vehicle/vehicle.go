package vehicle

// Vehicle 载具
type Vehicle interface {
	Speed() int
	unexported()
}
