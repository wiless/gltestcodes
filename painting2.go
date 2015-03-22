package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/gl/2.0"
)

var redraw bool
var defaultApp appInfo
var myplot GoPlot

type appInfo struct {
	Sinewaves float64
}

func (a *appInfo) UpdateScale(value float64) {
	a.Sinewaves = value
	// log.Println("Value udpated to ", a.Sinewaves)
	myplot.Call("update")

}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}

type GoPlot struct {
	qml.Object
	Name      string
	Color     color.RGBA
	Npoints   int
	YMax      float64
	yvals     []float32
	mypainter *qml.Painter
	Update    bool
}

func (g *GoPlot) init() {
	// width := int32(g.Int("width"))
	// g.Npoints = g.Int("nPoints")

	// log.Printf("\n Initializing  :  Npoints = %v , Height = %v", int(g.Object.Property("width").(float64)), float32(g.Object.Property("height").(float64)))
	g.YMax = g.Object.Property("height").(float64) / 2.0
	zeroline := float32(g.Object.Property("height").(float64)) / 2.0
	g.Npoints = int(g.Object.Property("width").(float64))
	g.yvals = make([]float32, g.Npoints)
	omega := defaultApp.Sinewaves //float64(g.Npoints) /
	// fmt.Printf("\n frequency = %v", omega)
	for i := 0; i < g.Npoints; i++ {
		theta := float64(2.0 * math.Pi * float64(i) * omega / float64(g.Npoints))
		g.yvals[i] = float32(g.YMax*math.Sin(theta)) + zeroline
		// g.yvals[i] = float32(i) + zeroline
	}

	// log.Printf("\n Init() : Seems %#v has only %d points with %v height ", g, g.Npoints, g.YMax)
	// log.Printf("\n%v*Sine Values(%d) : %v... %v", g.YMax, g.Npoints, g.yvals[0:10], g.yvals[g.Npoints-10:])
	// g.yvals = vlib.Randsrc(int(g.Npoints), height).ToVectorF()

}
func (g *GoPlot) Clicked() {
	g.Update = true
	log.Printf("\nPlot was clicked  : ")

	g.YMax = g.Object.Property("height").(float64)
	zeroline := float32(g.Object.Property("height").(float64)) / 2.0
	g.Npoints = int(g.Object.Property("width").(float64))
	g.yvals = make([]float32, g.Npoints)

	log.Println("Old sample ", g.yvals[10])
	for i := 0; i < g.Npoints; i++ {
		g.yvals[i] = zeroline + (rand.Float32()*float32(g.YMax) - float32(g.YMax)/2.0)
	}
	log.Println("New sample ", g.yvals[10])

	g.Call("update")
	// g.mypainter.Call("update")
}

func toColor4f(c color.RGBA) (r, g, b, a float32) {
	return float32(c.R) / 256.0, float32(c.G) / 256.0, float32(c.B) / 256.0, float32(c.A) / 256.0
}

func (g *GoPlot) Paint(p *qml.Painter) {
	empty := color.RGBA{}
	if g.Color == empty {
		g.Color = color.RGBA{00, 0, 0, 80}
		// log.Printf("DEFAULT SET TO  %v ", g.Color)
	} else {
		log.Println("Custom color ", g.Color)
	}

	red, green, blue, alpha := toColor4f(g.Color)
	g.mypainter = p
	if !g.Update { /// if not initialized even once
		g.init()
	}

	gl := GL.API(p)
	// gl.Enable(GL.BLEND)
	// gl.BlendFunc(GL.SRC_ALPHA, GL.ONE_MINUS_SRC_ALPHA)

	gl.Begin(GL.QUADS)
	// gl.Color4f(0.1, .1, 0, .4)
	// gl.Color4f(00, 0, 0, 0.5)
	gl.Color4f(red, green, blue, alpha)

	width := float32(g.Float64("width"))
	height := float32(g.Float64("height"))
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(0, height)
	gl.End()

	// glColor(solid_color)
	// glBegin(GL_POLYGON)
	// glVertex Commands here
	// glEnd
	// glDisable(GL_DEPTH_TEST)
	// glColor(outline_color)
	// glBegin(GL_LINE_LOOP)
	// glVertex Commands here
	// glEnd
	// glEnable(GL_DEPTH_TEST)

	gl.LineWidth(1)

	/// Draw grid lines
	gl.LineStipple(1, 0x00ff)
	gl.Enable(GL.LINE_STIPPLE)
	gl.Color4f(.8, .8, .8, .3)
	gl.Begin(GL.LINE_STRIP)
	gl.Vertex2f(0, float32(g.YMax))
	gl.Vertex2f(float32(g.Npoints), float32(g.YMax))
	gl.End()
	gl.Disable(GL.LINE_STIPPLE)

	gl.Color4f(.5, .5, 0.1, .3)
	gl.Begin(GL.LINE_STRIP)
	// gl.Begin(GL.POLYGON)
	gl.Vertex2f(0, float32(g.yvals[0]))
	for i := 1; i < len(g.yvals); i++ {
		gl.Vertex2f(float32(i), float32(g.yvals[i]))
	}

	// gl.Vertex2f(float32(len(g.yvals))-1, height/2.0)
	gl.End()

	// gl.Disable(GL.DEPTH_TEST)
	// gl.Color4f(.5, 0, 0, .2)
	// gl.Begin(GL.LINE_LOOP)
	// gl.Vertex2f(0, float32(g.yvals[0]))
	// for i := 1; i < len(g.yvals); i++ {
	// 	gl.Vertex2f(float32(i), float32(g.yvals[i]))
	// }
	// gl.End()
	// gl.Enable(GL.DEPTH_TEST)

}

func (g *GoPlot) SetYmax(height float64) {
	log.Printf("Height Set  = %v", height)
	g.YMax = height
	log.Printf("Height from obj  = %v", g.Int("YMax"))
	g.init()
}

func (g *GoPlot) SetNpoints(npoints int) {
	log.Printf("Npoints Set : %v", npoints)
	g.Npoints = npoints
}

func run() error {
	defaultApp.Sinewaves = 4.0
	objs := make([]qml.TypeSpec, 1)
	objs[0] = qml.TypeSpec{Init: func(r *GoPlot, obj qml.Object) {
		r.Object = obj
		r.Npoints = 0
		r.YMax = 111
		r.yvals = []float32{1, 1, 1}
		obj.On("mousePressed", r.Clicked)
	}}
	qml.RegisterTypes("GoExtensions", 1, 0, objs)

	engine := qml.NewEngine()
	ctx := engine.Context()
	ctx.SetVar("myapp", &defaultApp)

	component, err := engine.LoadFile("plotter.qml")
	if err != nil {
		return err
	}
	win := component.CreateWindow(nil)
	myplot.Object = win.ObjectByName("myobject")

	win.Set("x", 0)
	win.Set("y", 0)

	win.Show()
	win.Wait()

	return nil
}
