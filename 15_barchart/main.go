package main

import (
	"net/http"
	"fmt"
	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically, the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
	*/

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true, //enables / displays the y-axis
			},
		},
		Series: []chart.Series{
			
			// first chart
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(6).WithAlpha(255),
					FillColor:   chart.GetDefaultColor(6).WithAlpha(100),
					StrokeWidth: 4,
				},
				XValues: []float64{1.0, 1.5, 2.0, 3.0, 4.0, 5.0, 6.0},
				YValues: []float64{2.0, 3.0, 4.0, 3.5, 4.5, 5.5, 2.5},
			},

			// second chart
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(255),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(100),
					StrokeWidth: 2,
				},
				XValues: []float64{1.0, 1.5, 2.0, 3.0, 4.0, 5.0, 6.0},
				YValues: []float64{1.0, 0.5, 5.0, 2.0, 4.0, 5.0, 2.0},
			},

			// end of charts
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func drawPie(res http.ResponseWriter, req *http.Request) {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 6, Label: "Sales"},
			{Value: 1, Label: "!!"},
			{Value: 7, Label: "Counter"},
		},
	}

	// PNG option
	//res.Header().Set("Content-Type", "image/png")
	//err := pie.Render(chart.PNG, res)

	// SVG option
	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := pie.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func drawBar(res http.ResponseWriter, req *http.Request) {
	sbc := chart.BarChart{
		//Height:   312,
		Width:    400,
		BarWidth: 30,
		//BarSpacing:15,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: []chart.Value{
			{Value: 5.25, Label: "Blue"},
			{Value: 4.88, Label: "Green"},
			{Value: 4.74, Label: "Gray"},
			{Value: 3.22, Label: "Orange"},
			{Value: 3, Label: "Test"},
			{Value: 2.27, Label: "??"},
			{Value: 1, Label: "!!"},
			{Value: 2.48, Label: "Sale"},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	err := sbc.Render(chart.PNG, res)
	if err != nil {
		fmt.Printf("Error rendering chart: %v\n", err)
	}
}

func main() {
	fmt.Printf("Listening on port :8080")
	http.HandleFunc("/", drawChart)
	http.HandleFunc("/pie", drawPie)
	http.HandleFunc("/bar", drawBar)
	http.ListenAndServe(":8080", nil)
}
