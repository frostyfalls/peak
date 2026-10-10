// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	_ "embed"
	"flag"
	"net/http"
	"text/template"
)

//go:embed home.min.html
var homeHtml string

func main() {
	var address string

	flag.StringVar(&address, "address", ":7035", "listen address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("tmpl").Parse(homeHtml))
		tmpl.Execute(w, nil)
	})
	http.ListenAndServe(address, nil)
}
