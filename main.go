package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/browser"
)

type Patient struct {
	First     string `json:"first"`
	Last      string `json:"last"`
	Dob       string `json:"dob"`
	Condition string `json:"condition"`
	Allergies string `json:"allergies"`
	Emergency string `json:"emergency"`
}

var patient Patient

func updatePatient(first, last, dob, condition, allergies, emergency string) {
	patient.First = first
	patient.Last = last
	patient.Dob = dob
	patient.Condition = condition
	patient.Allergies = allergies
	patient.Emergency = emergency
}

func patientHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(patient)
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		values := r.PostForm
		first := values.Get("first")
		last := values.Get("last")
		dob := values.Get("dob")
		condition := values.Get("condition")
		allergies := values.Get("allergies")
		emergency := values.Get("emergency")
		if first == "" || last == "" || dob == "" || condition == "" || allergies == "" || emergency == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		updatePatient(first, last, dob, condition, allergies, emergency)
	default:
		w.Header().Set("allow", fmt.Sprintf("%s,%s", http.MethodGet, http.MethodPost))
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}
func main() {
	http.HandleFunc("/api/user", patientHandler)

	webDir := http.Dir("web")
	http.Handle("/", http.FileServer(webDir))

	browser.OpenURL("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
