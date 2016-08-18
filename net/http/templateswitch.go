package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Content struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Must be called before writing response

	t1 := `<html>
				<head>
					<title>
						{{.Title}}
					</title>
				</head>
				<body style="background-color: #CCC">
					<div>
						<form action="/" method="POST">
							<label for="templates">Select a template to use:</label><br />
							<input type="radio" name="templates" value="Template 1" checked="yes" /> : Template 1<br />
							<input type="radio" name="templates" value="Template 2" /> : Template 2<br />
							<input type="submit" value="Submit">
						</form>
					</div>
					<div>
						{{.Body}}
					</div>
				</body>
			</html>`

	t2 := `<html>
				<head>
					<title>
						{{.Title}}
					</title>
				</head>
				<body style="background-color: yellow">
					<div>
						<form action="/" method="POST">
							<label for="templates">Select a template to use:</label><br />
							<input type="radio" name="templates" value="Template 1" /> Template 1<br />
							<input type="radio" name="templates" value="Template 2" checked="yes" /> Template 2<br />
							<input type="submit" value="Submit">
						</form>
					</div>
					<div>
						{{.Body}}
					</div>
				</body>
			</html>`

	content := new(Content)
	content.Title = "Template Switch"
	content.Body = fmt.Sprint(req.Form)
	tmpl := template.New("Template Switch")
	selected, found := req.Form["templates"]

	if found && selected[0] == "Template 2" {
		tmpl, err = tmpl.Parse(t2)
	} else {
		tmpl, err = tmpl.Parse(t1)
	}
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, content)
	if err != nil {
		panic(err)
	}
}
