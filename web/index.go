package web

// IndexTmpl defines the HTML template for index page
const IndexTmpl = `<!DOCTYPE html>
<html>
	<head>
		<title>blocky</title>
	</head>
	<body>
		<h1>blocky</h1>
		<ul>
		{{range .}}
			<li><a href="{{.URL}}">{{.Title}}</a></li>
		{{end}}
		</ul>
		</body>
	</html>`
