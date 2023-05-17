package main

import (
	"fmt"
	"math"
	"sort"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"go.uber.org/zap"
)

type point struct {
	x float64
	y float64
}

var (
	window_width  = 1280
	window_height = 720
	is_displayed  = false
	smoothing     = true
)

var logger *zap.Logger

var (
	points []point
	pixels []uint8
	edges  [][2]point
	list   map[int][]int
)

func isExtrema(y, y1, y2 float64) bool {
	return (y > y1 && y > y2) || (y < y1 && y < y2)
}

func vertexCountTwice(i, j int) bool {
	l := len(edges)
	return isExtrema(
		edges[i][j].y,
		edges[i][(j+1)%2].y,
		edges[(i-1+l)%l][j].y)
}

func addToList(x, y float64) {
	list[int(math.Floor(y))] = append(list[int(math.Floor(y))], int(math.Floor(x)))
}

func drawLine(y, x1, x2 int) {
	for i := x1 + 1; i <= x2; i++ {
		pos := (i + (window_height-y)*window_width)
		pixels[pos] = 255
	}
}

func OrderedEdgesList() {
	for i, edge := range edges {
		if vertexCountTwice(i, 0) {
			addToList(edge[0].x, edge[0].y)
		}
		addToList(edge[1].x, edge[1].y)

		dy := edge[1].y - edge[0].y
		dx := edge[1].x - edge[0].x

		if dy == 0 {
			continue
		}

		count := int(math.Ceil(math.Abs(dy)))
		dy = dy / float64(count)
		dx = dx / float64(count)

		checkEndFunc := func(i float64) bool {
			if dy > 0 {
				return edge[0].y+i*dy < edge[1].y
			} else {
				return edge[0].y+i*dy > edge[1].y
			}
		}

		for i := float64(1); checkEndFunc(i); i++ {
			addToList(edge[0].x+i*dx, edge[0].y+i*dy)
		}
	}
}

func ProcessLines() {
	for y := range list {
		sort.Ints(list[y])
		for i := 0; i < len(list[y]); i += 2 {
			drawLine(y, list[y][i], list[y][i+1])
		}
	}
}

func brezenhem0(p1, p2 point) {
	I := 255

	dx := p2.x - p1.x
	dy := p2.y - p1.y

	x := p2.x
	y := p2.y

	swap := 0

	m := dy / dx

	sx := -1
	if dx < 0 {
		sx = 1
		dx *= -1
	}

	sy := -1
	if dy < 0 {
		sy = 1
		dy *= -1
	}

	if m > 1 {
		dx, dy = dy, dx

		m = 1 / m
		swap = 1
	}

	e := float64(I) / float64(2)

	m = m * float64(I)
	w := float64(I) - m

	for i := float64(1); i <= dx+1; i++ {
		if e <= w {
			if swap == 0 {
				x += float64(sx)
			} else {
				y += float64(sy)
			}
			e += m
		} else {
			y += float64(sy)
			x += float64(sx)
			e -= w
			color := 255 - e
			// fmt.Println(color, x, y)
			pos := (int(x) + (window_height-int(y))*window_width)
			pixels[pos] = uint8(math.Floor(color))
		}
	}

}

func getPos(x, y float64) int {
	return int(x) + (window_height-int(y))*window_width
}

func brezenhem(p1, p2 point) {
	I := float64(255)
	x := p1.x
	y := p1.y

	dx := p2.x - p1.x
	dy := p2.y - p1.y

	m := I * (dy / dx)

	w := I - m

	e := m / 2

	pos := getPos(x, y)

	pixels[pos] = uint8(m / 2)

	for x < p2.x {
		x++
		if e >= w {
			y++
			e -= w
		} else {
			e += m
		}
		pos := getPos(x, y)
		pixels[pos] = uint8(e)
	}
}

func filtrate() {
	for _, edge := range edges {
		fmt.Println(edge)
		brezenhem0(edge[0], edge[1])
	}
}

func calcEdges() {
	edges = make([][2]point, len(points))
	for i, p := range points {
		nextP := points[(i+1)%len(points)]
		edges[i] = [2]point{p, nextP}
	}
}

func calcPixels() {
	pixels = make([]uint8, window_width*window_height)
	list = make(map[int][]int)

	calcEdges()
	OrderedEdgesList()
	ProcessLines()
	if smoothing {
		filtrate()
	}
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		switch key {
		case glfw.KeyEscape:
			w.SetShouldClose(true)
		case glfw.KeyD:
			if len(points) > 2 {
				is_displayed = true

				calcPixels()

				logger.Info("drawing pixels...")
			} else {
				logger.Error("can't draw pixels, array is empty")
			}
		case glfw.KeyC:
			is_displayed = false
			points = []point{}
			pixels = []uint8{}
			edges = [][2]point{}
			list = nil

			logger.Info("cleared pixels buffer")
		case glfw.KeyS:
			smoothing = !smoothing

			logger.Info("toggled smooting", zap.Bool("new value", smoothing))
		}
		if key == glfw.KeyEscape {
			w.SetShouldClose(true)
		}
	}
}

func mouseCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action == glfw.Press && button == glfw.MouseButtonLeft {
		x, y := w.GetCursorPos()
		points = append(points, point{x, y})

		logger.Info("Mouse click: ", zap.Float64("x", x), zap.Float64("y", y))
	}
}

func sizeCallback(w *glfw.Window, width, height int) {
	window_width, window_height = width, height
	gl.Viewport(0, 0, int32(width), int32(height))

	points = []point{}
	pixels = []uint8{}
	is_displayed = false
}

func display(w *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	// gl.ClearColor(0.0, 0.9, 0.9, 1.0)

	if is_displayed {
		gl.DrawPixels(int32(window_width), int32(window_height),
			gl.RED, gl.UNSIGNED_BYTE,
			unsafe.Pointer(&pixels[0]),
		)
	}
}

func main() {
	logger, _ = zap.NewDevelopment()

	// points = append(points, point{87, 104}, point{190, 624}, point{1148, 441})

	if err := glfw.Init(); err != nil {
		logger.Fatal("failed to initialize glfw:", zap.Error(err))
	}

	window, err := glfw.CreateWindow(window_width, window_height, "Lab 4", nil, nil)

	if err != nil {
		glfw.Terminate()
		logger.Fatal("failed to create glfw window:", zap.Error(err))
	}

	if err := gl.Init(); err != nil {
		logger.Fatal("failed to initialize gl:", zap.Error(err))
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	window.SetKeyCallback(keyCallback)
	window.SetMouseButtonCallback(mouseCallback)
	window.SetFramebufferSizeCallback(sizeCallback)

	w, h := window.GetFramebufferSize()
	sizeCallback(window, w, h)

	for !window.ShouldClose() {
		display(window)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	window.Destroy()
	glfw.Terminate()
}
