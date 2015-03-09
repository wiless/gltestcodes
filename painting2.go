package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/gl/2.0"
)

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
	Npoints   int
	YMax      float64
	yvals     []float32
	mypainter *qml.Painter
}

func (g *GoPlot) init() {
	// width := int32(g.Int("width"))
	// g.Npoints = g.Int("nPoints")

	g.yvals = make([]float32, g.Npoints)
	for i := 0; i < g.Npoints; i++ {
		g.yvals[i] = float32(rand.Int31n(int32(g.YMax)))
	}

	log.Printf("Init() : Seems %#v has only %d points with %v height ", g, g.Npoints, g.YMax)
	// g.yvals = vlib.Randsrc(int(g.Npoints), height).ToVectorF()

}
func (g *GoPlot) Clicked() {
	log.Printf("\nPlot was clicked  : ")
	g.Paint(g.mypainter)
	log.Println("Old sample ", g.yvals[10])
	for i := 0; i < g.Npoints; i++ {
		g.yvals[i] = float32(rand.Int31n(int32(g.YMax)))
	}
	log.Println("New sample ", g.yvals[10])
}

func (g *GoPlot) Paint(p *qml.Painter) {
	// if g.N == 0 {
	g.mypainter = p

	log.Printf("\n Painter knows :  %#v \n %#v \n Npoints = %v , Height = %v", g, g.Object)

	g.init()
	gl := GL.API(p)

	// gl.Enable(GL.BLEND)
	// gl.BlendFunc(GL.SRC_ALPHA, GL.ONE_MINUS_SRC_ALPHA)
	gl.LineWidth(2)

	gl.Begin(GL.QUADS)
	gl.Color4f(0.1, .1, 0, .4)
	width := float32(g.Float64("width"))
	height := float32(g.Float64("height"))
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(0, height)
	gl.End()

	gl.Color4f(.5, .5, 0.1, .3)
	gl.Begin(GL.LINE_STRIP)
	gl.Vertex2f(0, float32(g.yvals[0]))
	for i := 1; i < int(g.Npoints); i++ {
		gl.Vertex2f(float32(i)*10, float32(g.yvals[i]))
	}
	gl.End()

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

	component, err := engine.LoadFile("plotter.qml")
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)
	// ctx := engine.Context()
	// ctx.SetVar(name, value)
	// button := win.ObjectByName("okbutton")
	// button.On("clicked", Clicked)
	// var goplotobj GoPlot
	// goplotobj.Object = win.ObjectByName("myobject")
	// goplotobj.Box = win.ObjectByName("myrect")
	// goplotobj.plotter = win.ObjectByName("myobject")
	// goplotobj.Object.On("mousePressed", goplotobj.Clicked)
	// ctx := engine.Context()

	// log.Printf("\n ObjectRetrieved is Plotter : %#v \n Object:=%#v \n Common:=%#v", goplotobj.plotter, goplotobj.Object, goplotobj.Common().TypeName())

	// log.Println(button.Property("width"))
	win.Set("x", 0)
	win.Set("y", 0)
	// win.Set("width", 1000)
	// win.Set("height", 768)

	win.Show()
	win.Wait()

	return nil
}
