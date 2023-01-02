package shortxid

import (
	"sync/atomic"
	"time"

	"github.com/jxskiss/base62"
	"golang.org/x/exp/constraints"
)

const (
	epochMillis = uint64(1672491600) * 1000 // starts 2023 and ends 2057 (40 bits)
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
	now := g.TimeFunc()
	ctr := atomic.AddUint64(&g.counter, 1)

	newIDint := ((now & 0xffffffffff) << (8 * 3)) | (ctr & 0xffff) | (g.id)
	return g.globalPrepend + prepend + string(g.encoder.FormatUint(newIDint))
}
