package oglabstr

import (
	"fmt"
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"os"
	"strings"
)

type Shader struct {
	rendererID    uint32
	locationCache map[string]int32
}

func NewShaderProgram(vertexShaderSource, fragmentShaderSource string) (*Shader, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Shader{program, make(map[string]int32)}, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	if source[len(source)-1] != 0x00 {
		source += "\x00"
	}

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var length int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &length)

		log := make([]byte, length+1)
		gl.GetShaderInfoLog(shader, length, &length, &log[0])

		return 0, fmt.Errorf(
			"failed to compile shader\n"+
				"reason: %v\n"+
				"shader source: %s\n",
			string(log),
			source)
	}

	return shader, nil
}

func (s *Shader) Bind() {
	gl.UseProgram(s.rendererID)
}

func (s *Shader) Unbind() {
	gl.UseProgram(0)
}

func (s *Shader) getUniformLocation(name string) int32 {
	location, seen := s.locationCache[name]
	if seen {
		return location
	}

	var nameCstr *uint8
	if !strings.HasSuffix(name, "\x00") {
		nameCstr = gl.Str(name + "\x00")
	} else {
		nameCstr = gl.Str(name)
	}

	location = gl.GetUniformLocation(s.rendererID, nameCstr)
	if location == -1 {
		fmt.Fprintf(os.Stderr, "[Warning] Uniform %s doesn't exitst\n", name)
	}

	s.locationCache[name] = location

	return location
}

func (s *Shader) SetUniform3fv(name string, vec3 mgl32.Vec3) {
	s.Bind()
	gl.Uniform3fv(s.getUniformLocation(name), 1, &vec3[0])
}

func (s *Shader) SetUniform4fv(name string, vec4 mgl32.Vec4) {
	s.Bind()
	gl.Uniform4fv(s.getUniformLocation(name), 1, &vec4[0])
}

func (s *Shader) SetUniform1i(name string, value int32) {
	s.Bind()
	gl.Uniform1i(s.getUniformLocation(name), value)
}

func (s *Shader) SetUniformMat4fv(name string, mat mgl32.Mat4, transpose bool) {
	s.Bind()
	gl.UniformMatrix4fv(s.getUniformLocation(name), 1, transpose, &mat[0])
}

func (s *Shader) SetUniformMat3fv(name string, mat mgl32.Mat3, transpose bool) {
	s.Bind()
	gl.UniformMatrix3fv(s.getUniformLocation(name), 1, transpose, &mat[0])
}

func (s *Shader) SetUniform1f(name string, v0 float32) {
	s.Bind()
	gl.Uniform1f(s.getUniformLocation(name), v0)
}

func (s *Shader) SetUniform3f(name string, v0, v1, v2 float32) {
	s.Bind()
	gl.Uniform3f(s.getUniformLocation(name), v0, v1, v2)
}
