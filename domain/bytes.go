package domain

import "sync"

var bytesPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4)
	},
}

func IDtoBytes(id int) []byte {
	b := bytesPool.Get().([]byte)
	b[0] = byte(id >> 24)
	b[1] = byte(id >> 16)
	b[2] = byte(id >> 8)
	b[3] = byte(id)

	return b
}
