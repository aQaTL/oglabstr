package oglabstr

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"unsafe"
	"fmt"
)

type VertexBuffer struct {
	rendererID uint32
}

func NewVertexBuffer(data interface{}, size int) *VertexBuffer {
	vb := VertexBuffer{}
	gl.GenBuffers(1, &vb.rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	return &vb
}

func (ib *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, ib.rendererID)
}

func (ib *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (ib *VertexBuffer) Delete() {
	gl.DeleteBuffers(1, &ib.rendererID)
}

type VertexBufferElement struct {
	elemType   uint32
	count      int32
	normalized bool
}

type VertexBufferLayout struct {
	Elements []VertexBufferElement
	Stride   int32
}

func NewVertexBufferLayout() *VertexBufferLayout {
	return &VertexBufferLayout{}
}

func (vbl *VertexBufferLayout) push(element VertexBufferElement) {
	vbl.Elements = append(vbl.Elements, element)
}

func (vbl *VertexBufferLayout) Pushf32(count int32) {
	vbl.push(VertexBufferElement{gl.FLOAT, count, false})
	vbl.Stride += count * GetSizeOfType(gl.FLOAT)
}

func (vbl *VertexBufferLayout) Pushui32(count int32) {
	vbl.push(VertexBufferElement{gl.UNSIGNED_INT, count, false})
	vbl.Stride += count * GetSizeOfType(gl.UNSIGNED_INT)
}

func (vbl *VertexBufferLayout) Pushub(count int32) {
	vbl.push(VertexBufferElement{gl.UNSIGNED_BYTE, count, true})
	vbl.Stride += count * GetSizeOfType(gl.UNSIGNED_BYTE)
}

func GetSizeOfType(glType uint32) int32 {
	switch glType {
	case gl.FLOAT:
		return int32(unsafe.Sizeof(float32(0)))
	case gl.UNSIGNED_INT:
		return int32(unsafe.Sizeof(uint32(0)))
	case gl.UNSIGNED_BYTE:
		return int32(unsafe.Sizeof(byte(0)))
	default:
		panic(fmt.Errorf("type %v not supported", glType))
	}
}
