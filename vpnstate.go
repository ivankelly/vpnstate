package main

import( "fmt"
	"net/http"
	"html/template"
	"strconv"
	"time"
)

var state = false

var htmlTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <META http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <META name="viewport" content="width=device-width, initial-scale=1">
    <title>VPN state</title>
    <link href="/css/screen.css" rel="stylesheet" type="text/css" />
    <style type="text/css">
      body {
        text-align: center;
        font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
        font-size: 10em;
        height: 100%%;
        background-color: {{if .}}green{{else}}red{{end}};
        color: {{if .}}white{{else}}black{{end}};
      }
    </style>
    <script>
      window.onclick = function() {
        document.getElementById("toggler").submit()
      }
    </script>
  </head>
  <body onClick="foobar">
    <form id="toggler" action="/toggle" method="POST">
    </form>
    <div onClick="foobar">{{if .}}ON{{else}}OFF{{end}}</div>
  </body>
</html>
`

var plainTmpl = `{{if .}}ON{{else}}OFF{{end}}`

func display(w http.ResponseWriter, r *http.Request) {
	tmpl,_ := template.New("Foobar").Parse(htmlTmpl)
	tmpl.Execute(w, state)
}

func toggle(w http.ResponseWriter, r *http.Request) {
	state = !state
	display(w, r)
}

func poll(w http.ResponseWriter, r *http.Request) {
	expected := r.FormValue("expected-current")
	expected_state := false
	if expected == "ON" {
		expected_state = true
	}
	timeout,err := strconv.Atoi(r.FormValue("timeout"))
	for i := 0; err == nil && expected_state == state && i < timeout; i++ {
		time.Sleep(1 * time.Second)
	}
	tmpl,_ := template.New("Foobar").Parse(plainTmpl)
	tmpl.Execute(w, state)
}

func main() {
	fmt.Printf("hello, fuckers\n")
	http.HandleFunc("/", display)
	http.HandleFunc("/toggle", toggle)
	http.HandleFunc("/poll", poll)
	http.ListenAndServe(":3001", nil)
}
