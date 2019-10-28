package warts

var addressContainer map[uint32]*Address
var addrCounter uint32

//Init
func Init() {
	addressContainer = make(map[uint32]*Address, 0)
	addrCounter = 0
}
