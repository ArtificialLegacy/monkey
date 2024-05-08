package vm

import (
	"github.com/ArtificialLegacy/monkey/pkg/code"
	"github.com/ArtificialLegacy/monkey/pkg/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{cl: cl, ip: -1, basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
