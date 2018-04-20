package oglabstr

import (
	"github.com/go-gl/gl/v4.3-core/gl"
)

type VertexArray struct {
	rendererID uint32
}

func NewVertexArray() *VertexArray {
	va := VertexArray{}
	gl.GenVertexArrays(1, &va.rendererID)
	return &va
}

func (va *VertexArray) AddBuffer(vb *VertexBuffer, layout *VertexBufferLayout) {
	va.Bind()
	vb.Bind()

	offset := 0
	for i, elem := range layout.Elements {
		gl.EnableVertexAttribArray(uint32(i))
		gl.VertexAttribPointer(uint32(i), elem.count, elem.elemType, elem.normalized,
			layout.Stride, gl.PtrOffset(offset))
		offset += int(elem.count * GetSizeOfType(elem.elemType))
	}
}

func (va *VertexArray) Bind() {
	gl.BindVertexArray(va.rendererID)
}

func (va *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}
