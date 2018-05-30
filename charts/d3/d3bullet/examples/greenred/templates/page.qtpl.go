// This file is automatically generated by qtc from "main.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line main.qtpl:1
package templates

//line main.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line main.qtpl:1
import "github.com/grokify/gocharts/charts/d3/d3bullet"

//line main.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line main.qtpl:2
func StreamCharts(qw422016 *qt422016.Writer, data ChartsData) {
	//line main.qtpl:2
	qw422016.N().S(`
<!doctype html>
<html>
	<head>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.6/d3.min.js"></script>
	</head>
	<body>

		<h1>D3 Bullet Example</h1>

<div style="clear:both;height:1em"></div>

`)
	//line main.qtpl:14
	qw422016.N().S(d3bullet.GetExampleCSSGreenRed(true))
	//line main.qtpl:14
	qw422016.N().S(`

<script>

var data = `)
	//line main.qtpl:18
	qw422016.N().Z(data.DataInt64.GetBulletDataJSON())
	//line main.qtpl:18
	qw422016.N().S(`;

`)
	//line main.qtpl:20
	qw422016.N().Z(d3bullet.GetJS())
	//line main.qtpl:20
	qw422016.N().S(`

var margin = {top: 5, right: 40, bottom: 20, left: 160},
    width = 960 - margin.left - margin.right,
    height = 50 - margin.top - margin.bottom;

var chart = d3.bullet()
    .width(width)
    .height(height);

`)
	//line main.qtpl:30
	qw422016.N().S(d3bullet.GetExampleJS())
	//line main.qtpl:30
	qw422016.N().S(`

</script>

<div style="clear:both;height:1em"></div>

	</body>
</html>
`)
//line main.qtpl:38
}

//line main.qtpl:38
func WriteCharts(qq422016 qtio422016.Writer, data ChartsData) {
	//line main.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line main.qtpl:38
	StreamCharts(qw422016, data)
	//line main.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:38
}

//line main.qtpl:38
func Charts(data ChartsData) string {
	//line main.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
	//line main.qtpl:38
	WriteCharts(qb422016, data)
	//line main.qtpl:38
	qs422016 := string(qb422016.B)
	//line main.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
	//line main.qtpl:38
	return qs422016
//line main.qtpl:38
}