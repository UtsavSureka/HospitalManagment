package main

import (
	"encoding/json"
	"hospital/Patient/DatabaseConn"
	"hospital/Patient/DoctorDetail"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetDoctor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doctordetail, err := DoctorDetail.GetDoctor()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		doctorJSON, err := json.Marshal(doctordetail)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(doctorJSON)

	}
}

func main() {
	http.HandleFunc("/doctordetail", GetDoctor)
	DatabaseConn.SetUpDBConn()
	http.ListenAndServe(":8080", nil)
}
