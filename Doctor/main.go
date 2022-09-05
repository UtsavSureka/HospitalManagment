package main

import (
	"encoding/json"
	"hospital/Doctor/PatientDetail"
	"hospital/Patient/DatabaseConn"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func getPatientDetail(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var DocId PatientDetail.DocID
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bytedata, &DocId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		Patient, err := PatientDetail.GetPatientDetail(DocId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		patientdetailJSON, err := json.Marshal(Patient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application.json")
		w.Write(patientdetailJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func main() {
	http.HandleFunc("/showpatient", getPatientDetail)
	DatabaseConn.SetUpDBConn()
	http.ListenAndServe(":5000", nil)
}
