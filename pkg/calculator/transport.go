package calculator

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	expr := r.FormValue("expression")

	return Request{
		Expression: expr,
	}, nil
}

// writes response from endpoint to client
func encodePlainResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(Response)

	w.Header().Add("Content-type", "text/plain")

	_, err := fmt.Fprint(w, resp.Result)
	return err
}

// writes decorated response from endpoint to client
func encodeHTMLResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(Response)

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
				<input name="expression" value="{{ if .Expr }} {{ .Expr }} {{ end }}" type="text" required>
				<input type="submit" value="Calculate">
			</form>
			{{ if .Result }}<h4>Result: {{ .Result }}</h4>{{ end }}
		</div>
	</body>
	</html>`

	w.Header().Add("Content-type", "text/html")

	// prepare the data
	data := struct {
		Result float64
		Error  error
		Expr   string
	}{
		Result: resp.Result,
		Error:  resp.Error,
		Expr:   resp.Expression,
	}

	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}

// NewHTTPHandler creates greeter handlers
func NewHTTPHandler(endpoint endpoint.Endpoint) http.Handler {
	m := http.NewServeMux()

	//Load css
	m.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	m.Handle("/api/calculate", httptransport.NewServer(
		endpoint,
		decodeRequest,
		encodePlainResponse,
	))

	m.Handle("/calculate", httptransport.NewServer(
		endpoint,
		decodeRequest,
		encodeHTMLResponse,
	))
	return m
}
