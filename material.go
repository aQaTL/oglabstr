package oglabstr

type Material struct {
	Name                                     string
	diffuseName, specularName, shininessName string

	Diffuse, Specular *Texture
	Shininess         float32
}

func NewMaterial(name string, diffuse, specular *Texture, shininess float32) *Material {
	return &Material{
		Name:          name,
		diffuseName:   name + ".diffuse",
		specularName:  name + ".specular",
		shininessName: name + ".shininess",

		Diffuse:   diffuse,
		Specular:  specular,
		Shininess: shininess,
	}
}

func (m *Material) Bind(shader *Shader) {
	shader.SetInt(m.diffuseName, int32(m.Diffuse.Slot))
	shader.SetInt(m.specularName, int32(m.Specular.Slot))
	shader.SetFloat(m.shininessName, m.Shininess)
}
