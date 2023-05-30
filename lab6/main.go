package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"runtime"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/vbsw/glut"
)

const (
	window_width  = 900
	window_height = 900
)

var (
	angle1 = float32(10)
	angle2 = float32(10)
	angle3 = float32(10)

	v0      = float32(0)
	g       = float32(0.00051)
	surface = float32(-2)
	h0      = float32(0)

	light  = true
	moving = false
	it     = false

	// texture = 0

	light_position     = []float32{0.0, 0.0, 20.0, 1.0}
	material_diffusion = []float32{1, 1, 1, 1.0}
	light_diffusion    = []float32{1, 1, 1}

	Kl = float32(0.001)
	Kq = float32(0.002)
	Kc = float32(0.001)
)

func light_on() {
	gl.Materialfv(gl.FRONT_AND_BACK, gl.DIFFUSE, &material_diffusion[0])

	light_spot_direction := []float32{0.0, 0.0, -1.0}
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.LIGHTING)
	gl.LightModelf(gl.LIGHT_MODEL_TWO_SIDE, gl.TRUE)
	gl.Enable(gl.NORMALIZE)
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &light_diffusion[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &light_position[0])
	gl.Lightf(gl.LIGHT0, gl.SPOT_CUTOFF, 10)
	gl.Lightfv(gl.LIGHT0, gl.SPOT_DIRECTION, &light_spot_direction[0])
	gl.Lightf(gl.LIGHT0, gl.LINEAR_ATTENUATION, Kl)
	gl.Lightf(gl.LIGHT0, gl.QUADRATIC_ATTENUATION, Kq)
	gl.Lightf(gl.LIGHT0, gl.CONSTANT_ATTENUATION, Kc)
}

func drawFigure(y, s float32) {
	gl.Scalef(s, s, s)
	gl.Translatef(0, y, 0)
	gl.Rotatef(angle1, 1, 0, 0)
	gl.Rotatef(angle2, 0, 1, 0)
	gl.Rotatef(angle3, 0, 0, 1)

	glut.SolidCube(1)
}

func display(w *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	if light {
		light_on()
	}
	gl.MatrixMode(gl.MODELVIEW)

	gl.LoadIdentity()
	if !moving {
		gl.PushMatrix()
		drawFigure(0, 0.3)
		gl.PopMatrix()
	}
	if moving {
		gl.PushMatrix()
		drawFigure(h0, 0.3)
		h0 += v0
		v0 -= g
		fmt.Println(h0, v0*100)
		if h0 <= surface || h0 > 0 {
			v0 = -1 * v0
		}
		gl.PopMatrix()
	}

}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		switch key {
		case glfw.KeyEscape:
			w.SetShouldClose(true)
		case glfw.Key1:
			light_position[2] -= 0.65
		case glfw.Key2:
			light_position[2] += 0.65
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
		case glfw.KeyUp:
			if it {
				gl.Disable(gl.TEXTURE_2D)
			} else {
				gl.Enable(gl.TEXTURE_2D)
			}
			it = !it

		}
	}

}

func main() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatal("failed to initialize glfw:", err)
	}

	window, err := glfw.CreateWindow(window_width, window_height, "Lab 6", nil, nil)
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	window.SetKeyCallback(keyCallback)

	if err != nil {
		glfw.Terminate()
		log.Fatal("failed to create glfw window:", err)
	}

	if err := gl.Init(); err != nil {
		log.Fatal("failed to initialize gl:", err)
	}
	glut.Init()

	img, err := imaging.Open("1.bmp")
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

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	// Использование текстуры
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	for !window.ShouldClose() {
		display(window)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	window.Destroy()
	glfw.Terminate()

}
