package textures

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v2.1/gl"
)

type Texture struct {
	Width  int
	Height int
	Pixels []uint8
}

func ReadTexture(path string) Texture {
	img, err := imaging.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Преобразование изображения в формат RGBA
	rgba := imaging.Clone(img)

	// Получение размеров изображения
	bounds := rgba.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Получение пиксельных данных изображения
	pixels := rgba.Pix

	return Texture{
		Width:  width,
		Height: height,
		Pixels: pixels,
	}
}

func GenerateTexture() uint32 {
	var generatedTexture uint32

	data := [2][2][4]uint8{{{255, 0, 0, 0}, {255, 255, 0, 0}}, {{0, 255, 0, 0}, {0, 0, 255, 0}}}
	gl.GenTextures(1, &generatedTexture)
	gl.BindTexture(gl.TEXTURE_2D, generatedTexture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, 2, 2, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&data[0][0][0]))
	log.Println(generatedTexture)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return generatedTexture
}
