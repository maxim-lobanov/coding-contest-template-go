package runner

import (
	"fmt"
	"strings"
)

type OutputBuilder struct {
	buffer strings.Builder
}

func (b *OutputBuilder) WriteString(s string) {
	b.buffer.WriteString(s)
}

func (b *OutputBuilder) WriteInt(v int) {
	b.buffer.WriteString(fmt.Sprint(v))
}

func (b *OutputBuilder) WriteUInt(v uint) {
	b.buffer.WriteString(fmt.Sprint(v))
}

func (b *OutputBuilder) WriteInt64(v int64) {
	b.buffer.WriteString(fmt.Sprint(v))
}

func (b *OutputBuilder) WriteUInt64(v uint64) {
	b.buffer.WriteString(fmt.Sprint(v))
}

func (b *OutputBuilder) WriteFloat(v float64) {
	b.buffer.WriteString(fmt.Sprint(v))
}

func (o *OutputBuilder) String() string {
	return o.buffer.String()
}
