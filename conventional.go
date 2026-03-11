import (
	"image/color"
	"os"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)
type DataPoint struct {
	NodeCount float64
	Throughput float64
}
func buildData() []DataPoint {
	return []DataPoint{
		
	}
}
func convertToXY(data []DataPoint) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i, d := range data {
		pts[i].X = d.NodeCount
		pts[i].Y = d.Throughput
	}
	return pts
}
func createPlot() *plot.Plot {
	p := plot.New()
	p.Title.Text = "Throughput Scalability Under Conventional Locking"
	p.X.Label.Text = "Cluster Size (Nodes)"
	p.Y.Label.Text = "Throughput (txn/sec)"
	p.BackgroundColor = color.White
	return p
}
func addLine(p *plot.Plot, pts plotter.XYs) {
	line, _ := plotter.NewLine(pts)
	line.Color = color.RGBA{R: 0, G: 0, B: 200, A: 255}
	line.Width = vg.Points(2)
	p.Add(line)
}
func addPoints(p *plot.Plot, pts plotter.XYs) {
	scatter, _ := plotter.NewScatter(pts)
	scatter.Shape = plot.PyramidGlyph{}
	scatter.Color = color.RGBA{R: 200, G: 0, B: 0, A: 255}
	scatter.Radius = vg.Points(4)
	p.Add(scatter)
}
func savePlot(p *plot.Plot, file string) {
	writer, _ := os.Create(file)
	defer writer.Close()
	canvas := vgimg.New(6*vg.Inch, 4*vg.Inch)
	dc := draw.New(canvas)
	p.Draw(dc)
	png := vgimg.PngCanvas{Canvas: canvas}
	png.WriteTo(writer)
}
func main() {
	data := buildData()
	points := convertToXY(data)
	plot := createPlot()
	addLine(plot, points)
	addPoints(plot, points)
	plot.Save(6*vg.Inch,4*vg.Inch, "throughput_conventional.png")
}
