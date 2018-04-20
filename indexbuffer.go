package oglabstr

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"unsafe"
)

type IndexBuffer struct {
	rendererID uint32
	Count      int32
}

func NewIndexBuffer(data []uint32) *IndexBuffer {
	ib := IndexBuffer{Count: int32(len(data))}
	gl.GenBuffers(1, &ib.rendererID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.rendererID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data) * int(unsafe.Sizeof(uint32(0))), gl.Ptr(data), gl.STATIC_DRAW)
	return &ib
}

func (ib *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.rendererID)
}

func (ib *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (ib *IndexBuffer) Delete() {
	gl.DeleteBuffers(1, &ib.rendererID)
}
