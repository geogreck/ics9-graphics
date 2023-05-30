package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/geogreck/ics9-graphics/lab8/figure"
	"github.com/geogreck/ics9-graphics/lab8/shaders"
	"github.com/geogreck/ics9-graphics/lab8/textures"
	"github.com/geogreck/ics9-graphics/lab8/util"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	window_width  = 900
	window_height = 900
)

const SIZE = 600
const HEIGHT = 0.5

var (
	angle1          = float32(10)
	angle2          = float32(10)
	angle3          = float32(10)
	RADIUS  float64 = math.Sqrt(0.5) * 0.5
	CORNERS int     = 6
	it              = false

	// texture = 0

	POINT1 []float64 = []float64{0, 0.6, 0}
	POINT2 []float64 = []float64{0.6, 0.6, 0}
	POINT3 []float64 = []float64{0.6, 0, 0}

	program = uint32(0)

	generatedTexture uint32 = 0
	loadedTexture    uint32 = 0
	v0                      = float32(0)
	g                       = float32(0.00051)
	surface                 = float32(-1)
	h0                      = float32(0)

	moving = false
)

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		switch key {
		case glfw.KeyEscape:
			w.SetShouldClose(true)
		case glfw.KeyUp:
			if it {
				gl.Disable(gl.TEXTURE_2D)
			} else {
				gl.Enable(gl.TEXTURE_2D)
			}
			it = !it
		case glfw.KeyA:
			angle1 -= 10
		case glfw.KeyD:
			angle1 += 10
		case glfw.KeyW:
			angle2 -= 10
		case glfw.KeyS:
			angle2 += 10
		case glfw.KeyQ:
			angle3 -= 10
		case glfw.KeyE:
			angle3 += 10
		case glfw.KeyDown:
			if moving {
				v0 = 0
				h0 = 0
				moving = false
			} else {
				moving = true
			}
		}

	}
}

func LoadTexture() {
	imgFile, err := os.Open("1.bmp")
	if err != nil {
		log.Panicln("texture not found on disk: ", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Panicln("unsupported stride", err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		log.Panicln("unsupported stride", err)

	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	gl.GenTextures(2, &loadedTexture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, loadedTexture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
}

func drawMovingPrism(y, s float32) {
	gl.PushMatrix()

	gl.Scalef(s, s, s)
	gl.Translatef(0, y, 0)
	gl.Rotatef(angle1, 1, 0, 0)
	gl.Rotatef(angle2, 0, 1, 0)
	gl.Rotatef(angle3, 0, 0, 1)
	figure.DrawPrism(6, program, loadedTexture)

	gl.PopMatrix()

}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()

	util.SetUniformVariables(program)
	util.SetLight()
	if !moving {
		gl.PushMatrix()
		drawMovingPrism(0, 1)
		gl.PopMatrix()
	}
	if moving {
		gl.PushMatrix()
		drawMovingPrism(h0, 1)
		h0 += v0
		v0 -= g
		fmt.Println(h0, v0*100)
		if h0 <= surface || h0 > 0 {
			v0 = -1 * v0
		}
		gl.PopMatrix()
	}
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatal("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)

	window, err := glfw.CreateWindow(window_width, window_height, "Lab 8", nil, nil)
	if err != nil {
		glfw.Terminate()
		log.Fatal("failed to create glfw window:", err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		log.Fatal("failed to initialize gl:", err)
	}

	window.SetKeyCallback(keyCallback)

	util.Setup()

	// Компиляция и создание программы шейдеров
	program, err := shaders.CreateProgram(util.ReadShaders("shaders/vertex.glsl", "shaders/fragment.glsl"))
	if err != nil {
		panic(err)
	}
	gl.UseProgram(program)
	if err != nil {
		log.Fatalf("Failed to create shader program: %v", err)
	}
	generatedTexture = textures.GenerateTexture()
	defer gl.DeleteTextures(1, &generatedTexture)

	LoadTexture()
	defer gl.DeleteTextures(2, &loadedTexture)

	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)

	backLight := []float32{0.3, 0.3, 0.3, 1}
	gl.LightModelfv(gl.LIGHT_MODEL_AMBIENT, &backLight[0])

	for !window.ShouldClose() {
		display()

		window.SwapBuffers()
		glfw.PollEvents()
	}

	window.Destroy()
	glfw.Terminate()
}
