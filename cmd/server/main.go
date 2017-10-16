package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"server/pkg/calculator"
)

type apiHandler struct{}

func (*apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Direct api calls not implemented yet.")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
		<!doctype html>
		
		<html lang="en">
		<head>
		<meta charset="utf-8">
		
		<title>Calculator</title>
		<meta name="description" content="Calculator">
		<meta name="author" content="Milos Jovanov">
		
		<link rel="stylesheet" href="static/css/style.css">
		
		<!--[if lt IE 9]>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
		<![endif]-->
		</head>
		
		<body>
		 	<div class="content">
				<form method="post">
					<h4>Calculator</h4>
					<p>Allowed operators are '+', '-', '*' and '/'"</p>
					<p>Syntax: number operator number</p>
					<p>Examples: "5 + 5", "5 - 5", "5 * 5" and "5 / 5" </p>
					<input name="expression" value={{ .Expr }} required>
					<input type="submit" value="Calculate">
				</form>
				{{ if .Result }}<h4>Result: {{ .Result }}</h4>{{ end }}
			</div>
		</body>
		</html>`

	// set the encoding
	w.Header().Add("Content-type", "text/html")

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result := ""
	if r.FormValue("expression") != "" {
		result = calculator.Calculate(r.FormValue("expression"))
	}

	// prepare the data
	data := struct {
		Result string
		Expr   string
	}{
		Result: result,
		Expr:   r.FormValue("expression"),
	}

	// parse the template
	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		fmt.Println("Failed to parse template;", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func startServer(address string) {
	http.Handle("/api/calculate", &apiHandler{})
	http.HandleFunc("/calculate", htmlHandler)

	//Load css
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, nil)
}

func main() {
	var addr = flag.String("addr", "", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	if *addr != "" {
		startServer(*addr)
	} else {
		startServer("0.0.0.0:8080")
	}
}
