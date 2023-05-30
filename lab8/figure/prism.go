package figure

import (
	"math"

	"github.com/go-gl/gl/v2.1/gl"
)

var (
	RADIUS  float64 = math.Sqrt(0.5) * 0.5
	CORNERS int     = 6
	SIZE            = 600
	HEIGHT          = 0.5
)

func drawBase(vertexes [][2]float64, normals [][3]float64, z float64, program uint32) {
	offset := 0
	if z < 0 {
		offset = len(vertexes)
	}
	gl.Color3d(z, 1, 1)
	vertexArray := make([][3]float64, len(vertexes))
	normalArray := make([][3]float64, len(vertexes))
	for i, vertex := range vertexes {
		normalArray[i] = [3]float64{normals[offset+i][0], normals[offset+i][1], normals[offset+i][2]}
		vertexArray[i] = [3]float64{vertex[0], vertex[1], z}
	}

	name := "isTexture\000"
	isTexture := gl.GetUniformLocation(program, gl.Str(name))
	gl.Uniform1f(isTexture, 0)

	gl.EnableClientState(gl.NORMAL_ARRAY)
	gl.EnableClientState(gl.VERTEX_ARRAY)

	gl.VertexPointer(3, gl.DOUBLE, 0, gl.Ptr(&vertexArray[0][0]))
	gl.NormalPointer(gl.DOUBLE, 0, gl.Ptr(&normalArray[0][0]))

	gl.DrawArrays(gl.POLYGON, 0, int32(len(vertexArray)))

	gl.DisableClientState(gl.NORMAL_ARRAY)
	gl.DisableClientState(gl.VERTEX_ARRAY)
}

func drawSideFaces(vertexes [][2]float64, normals [][3]float64, height float64, program uint32, loadedTexture uint32) {
	name := "isTexture\000"
	isTexture := gl.GetUniformLocation(program, gl.Str(name))
	gl.Uniform1f(isTexture, 0)

	gl.BindTexture(gl.TEXTURE_2D, loadedTexture)
	gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
	vertexArray := [][3]float64{}
	normalArray := [][3]float64{}
	textureArray := [][2]float64{}
	gl.Color4d(1, 1, 1, 1)
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.EnableClientState(gl.NORMAL_ARRAY)
	gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)
	gl.Uniform1f(isTexture, 1)

	for i := 0; i < len(vertexes); i++ {

		normalArray = append(normalArray,
			[3]float64{normals[i+len(vertexes)][0], normals[i+len(vertexes)][1], normals[i+len(vertexes)][2]})
		vertexArray = append(vertexArray, [3]float64{vertexes[i][0], vertexes[i][1], height / -2})

		normalArray = append(normalArray, [3]float64{normals[i][0], normals[i][1], normals[i][2]})
		vertexArray = append(vertexArray, [3]float64{vertexes[i][0], vertexes[i][1], height / 2})

		normalArray = append(normalArray, [3]float64{normals[(i+1)%len(vertexes)][0], normals[(i+1)%len(vertexes)][1], normals[(i+1)%len(vertexes)][2]})
		vertexArray = append(vertexArray, [3]float64{vertexes[(i+1)%len(vertexes)][0], vertexes[(i+1)%len(vertexes)][1], height / 2})

		normalArray = append(normalArray, [3]float64{normals[(i+1)%len(vertexes)+len(vertexes)][0],
			normals[(i+1)%len(vertexes)+len(vertexes)][1],
			normals[(i+1)%len(vertexes)+len(vertexes)][2]})
		vertexArray = append(vertexArray, [3]float64{vertexes[(i+1)%len(vertexes)][0], vertexes[(i+1)%len(vertexes)][1], height / -2})

		textureArray = append(textureArray, [2]float64{0, 0}, [2]float64{1, 0}, [2]float64{1, 1}, [2]float64{0, 1})
	}

	texture := gl.GetUniformLocation(program, gl.Str("texture\000"))
	gl.Uniform1i(texture, 0)

	gl.TexCoordPointer(2, gl.DOUBLE, 0, gl.Ptr(&textureArray[0][0]))
	gl.VertexPointer(3, gl.DOUBLE, 0, gl.Ptr(&vertexArray[0][0]))
	gl.NormalPointer(gl.DOUBLE, 0, gl.Ptr(&normalArray[0][0]))

	gl.DrawArrays(gl.QUADS, 0, int32(len(vertexArray)))

	gl.DisableClientState(gl.NORMAL_ARRAY)
	gl.DisableClientState(gl.VERTEX_ARRAY)
	gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)

	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.Uniform1f(isTexture, 0)

}

func DrawPrism(n int, program uint32, loadedTexture uint32) {
	vertexes := [][2]float64{}
	normals := make([][3]float64, n*2)
	for i := 0; i < n; i++ {
		vertexes = append(vertexes, [2]float64{
			RADIUS * math.Cos(2*math.Pi*float64(i)/float64(n)+math.Pi*45/180),
			RADIUS * math.Sin(2*math.Pi*float64(i)/float64(n)+math.Pi*45/180),
		})
	}
	for i := 0; i < n; i++ {
		normals[i][0] = -1 * (vertexes[(n+i-1)%n][0] + vertexes[(n+i+1)%n][0] - 2*vertexes[i][0])
		normals[i][1] = -1 * (vertexes[(n+i-1)%n][1] + vertexes[(n+i+1)%n][1] - 2*vertexes[i][1])
		normals[i][2] = HEIGHT / 2
		normals[n+i][0] = normals[i][0]
		normals[n+i][1] = normals[i][1]
		normals[n+i][2] = HEIGHT / -2
	}
	drawBase(vertexes, normals, HEIGHT/-2, program)
	drawBase(vertexes, normals, HEIGHT/2, program)
	drawSideFaces(vertexes, normals, HEIGHT, program, loadedTexture)
}
