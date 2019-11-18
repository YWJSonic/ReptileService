package routineswitch

// Info ...
type Info struct {
	isColse bool
	// closeChan chan bool
}

// Close close routing
func (I *Info) Close() {
	I.isColse = true
}

// IsClose ...
func (I *Info) IsClose() bool {
	return I.isColse
}
