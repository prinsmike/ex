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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Must be called before writing response
	t := `<html>
				<head>
					<title>
						{{.Title}}
					</title>
				</head>
				<body>
					<div>
						<form action="/" method="POST">
							<label for="text">Enter some text:</label><br />
							<input type="text" name="text" size="30"><br />
							<input type="submit" value="Submit">
						</form>
					</div>
					<div>
						{{.Body}}
					</div>
				</body>
			</html>`
	content := new(Content)
	content.Title = "Form Parse"
	if text, found := req.Form["text"]; found && text[0] != "" {
		content.Body = fmt.Sprint(req.Form)
	} else {
		content.Body = "Submit some text."
	}
	tmpl, err := template.New("test").Parse(t)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, content)
	if err != nil {
		panic(err)
	}
}
