package shortxid

import (
	"sync/atomic"
	"time"

	"github.com/jxskiss/base62"
	"golang.org/x/exp/constraints"
)

const (
	epochMillis = uint64(1672491600000) // starts 2023 and ends 2040 (39 bits)
	encoding    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type Generator struct {
	id            uint64
	counter       uint64
	globalPrepend string
	encoder       *base62.Encoding
	TimeFunc      func() uint64
}

func NewGenerator[I constraints.Integer](id I, globalPrepend string) *Generator {
	return &Generator{
		id:            (uint64(id) & 0xff) << (8 * 2),
		counter:       0,
		globalPrepend: globalPrepend,
		encoder:       base62.NewEncoding(encoding),
		TimeFunc: func() uint64 {
			return uint64(time.Now().UnixMilli()) - epochMillis
		},
	}
}

func (g *Generator) NewID(prepend string) string {
	now := g.TimeFunc() & 0xffffffffff
	ctr := atomic.AddUint64(&g.counter, 1) & 0xffff

	newIDint := (now << (8 * 3)) | ctr | g.id | uint64(0x8000000000000000)
	return g.globalPrepend + prepend + string(g.encoder.FormatUint(newIDint))
}
