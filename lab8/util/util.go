package util

import (
	"log"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

var (
	ambientMode  int         = 0
	ambient      [][]float32 = [][]float32{{0, 0, 0, 1}, {1, 1, 1, 0.5}, {1, 1, 1, 1}, {0.5, 0.5, 0.5, 1}, {0, 1, 0, 1}}
	diffuseMode  int         = 0
	diffuse      [][]float32 = [][]float32{{1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 0.5}, {0.5, 0.5, 0.5, 1}, {0, 1, 0, 1}}
	specularMode int         = 0
	specular     [][]float32 = [][]float32{{1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 0.5}, {0.5, 0.5, 0.5, 1}, {0, 1, 0, 1}}

	lightPosition []float32 = []float32{0, 0, 1, 1}

	setInfinityDistantLight = false
)

func Setup() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.NORMALIZE)
	gl.Enable(gl.COLOR_MATERIAL)
	gl.Enable(gl.TEXTURE_2D)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.NORMALIZE)
	gl.Enable(gl.COLOR_MATERIAL)
	gl.Enable(gl.TEXTURE_2D)
}

func ReadShaders(vertexfile, fragmentfile string) (string, string) {
	fragmentShader, err := os.ReadFile(vertexfile)
	if err != nil {
		log.Fatal(err)
	}
	vertexShader, err := os.ReadFile(fragmentfile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Read shaders")

	return string(fragmentShader), string(vertexShader)
}

func SetUniformVariables(program uint32) {
	uniformLightPos := gl.GetUniformLocation(program, gl.Str("lightPos\000"))
	gl.Uniform3fv(uniformLightPos, 1, &lightPosition[0])

	uniformAmbient := gl.GetUniformLocation(program, gl.Str("ambient\000"))
	gl.Uniform4fv(uniformAmbient, 1, &ambient[ambientMode][0])
	uniformDiffuse := gl.GetUniformLocation(program, gl.Str("diffuse\000"))
	gl.Uniform4fv(uniformDiffuse, 1, &diffuse[diffuseMode][0])
}

func SetLight() {
	if setInfinityDistantLight {
		lightPosition[3] = 0
	} else {
		lightPosition[3] = 1
	}
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[ambientMode][0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[diffuseMode][0])
	gl.Lightfv(gl.LIGHT0, gl.SPECULAR, &specular[specularMode][0])

	gl.Color3d(1, 1, 1)
	gl.PointSize(10)
	gl.Normal3b(0, 0, -1)

	gl.Begin(gl.POINTS)
	gl.Vertex3fv(&lightPosition[0])
	gl.End()
}
