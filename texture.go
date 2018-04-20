package oglabstr

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"image"
	"image/draw"
	_ "image/png"
	"os"
	"unsafe"
)

type Texture struct {
	rendererID         uint32
	filepath           string
	localBuffer        unsafe.Pointer
	Width, Height, Bpp int32
	Slot               uint32
}

func NewTexture(path string, slot uint32) *Texture {
	t := Texture{filepath: path, Slot: slot}

	gl.GenTextures(1, &t.rendererID)
	gl.ActiveTexture(gl.TEXTURE0 + slot)
	gl.BindTexture(gl.TEXTURE_2D, t.rendererID)

	imgFile, err := os.Open(path)
	defer imgFile.Close()
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)

	//internalFmt := int32(gl.SRGB_ALPHA)
	internalFmt := int32(gl.RGBA)
	format := uint32(gl.RGBA)
	pixType := uint32(gl.UNSIGNED_BYTE)
	dataPtr := gl.Ptr(rgba.Pix)

	t.Width = int32(rgba.Rect.Size().X)
	t.Height = int32(rgba.Rect.Size().Y)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(gl.TEXTURE_2D, 0, internalFmt, t.Width, t.Height, 0, format, pixType, dataPtr)

	gl.GenerateMipmap(gl.TEXTURE_2D)

	return &t
}

func (t *Texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0 + t.Slot)
	gl.BindTexture(gl.TEXTURE_2D, t.rendererID)
}

func (t *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
