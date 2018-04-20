package oglabstr

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Renderer struct {
	DeltaTime, lastFrame float32
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Draw(va *VertexArray, ib *IndexBuffer, shader *Shader) {
	shader.Bind()
	va.Bind()
	ib.Bind()

	gl.DrawElements(gl.TRIANGLES, ib.Count, gl.UNSIGNED_INT, nil)
}

func (r *Renderer) DrawArrays(va *VertexArray, first, count int32, shader *Shader) {
	shader.Bind()
	va.Bind()

	gl.DrawArrays(gl.TRIANGLES, first, count)
}

func (r *Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (r *Renderer) ClearColor(rgb uint32, alpha uint8) {
	const red = 0xff0000
	const green = 0x00ff00
	const blue = 0x0000ff

	gl.ClearColor(
		float32(rgb&red>>(4*4))/255.0,
		float32(rgb&green>>(4*2))/255.0,
		float32(rgb&blue>>(4*0))/255.0,
		float32(alpha)/255.0)
}

func (r *Renderer) UpdateDeltaTime() {
	currentFrame := float32(glfw.GetTime())
	r.DeltaTime = currentFrame - r.lastFrame
	r.lastFrame = currentFrame
}
