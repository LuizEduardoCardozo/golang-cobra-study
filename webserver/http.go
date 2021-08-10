package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LuizEduardoCardozo/CobraStudy/app"
)

type Server struct {
	Port string
}

func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.ParseFloat(r.URL.Query().Get("A"), 64)
	b, _ := strconv.ParseFloat(r.URL.Query().Get("B"), 64)

	calc := app.NewCalc()

	calc.A = a
	calc.B = b

	w.Write([]byte(fmt.Sprintf("%.3f + %.3f = %.3f", calc.A, calc.B, calc.Sum())))

}

func (s Server) Serve() {
	http.HandleFunc("/", s.SumHandler)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
