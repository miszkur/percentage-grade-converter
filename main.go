package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

type gradingTable struct {
	A grade
	B grade
	C grade
	D grade
	E grade
	F grade
}

type grade struct {
	Max int
	Min int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		maxPoints := req.FormValue("maxPoints")

		max, _ := strconv.Atoi(maxPoints)
		A := computeGrade(100, 100, max, nil)
		B := computeGrade(99, 88, max, &A)
		C := computeGrade(87, 73, max, &B)
		D := computeGrade(72, 58, max, &C)
		E := computeGrade(57, 45, max, &D)
		F := computeGrade(44, 0, max, &E)

		grading := gradingTable{A, B, C, D, E, F}
		tpl.ExecuteTemplate(w, "index.gohtml", grading)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func computeGrade(maxPercentage int, minPercentage int, maxPoints int, prevGrade *grade) grade {
	if prevGrade == nil {
		return grade{
			Max: maxPoints,
			Min: (minPercentage * maxPoints) / 100,
		}
	}

	return grade{
		Max: prevGrade.Min - 1,
		Min: (minPercentage * maxPoints) / 100,
	}
}
