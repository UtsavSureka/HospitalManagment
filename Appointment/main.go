package main

import (
	"encoding/json"
	"fmt"
	"hospital/Appointment/BookAppointment"
	"hospital/Appointment/DatabaseConn"
	"hospital/Appointment/DoctorSlot"
	"hospital/Appointment/UpdateAppointment"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetDoctorSlot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var doctordetail DoctorSlot.DoctorDB
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Line 23", err)
			return
		}
		err = json.Unmarshal(bytedata, &doctordetail)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Line 29", err)
			return
		}
		Doctordata, err := DoctorSlot.GetDoctorByTimeSlot(doctordetail)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Line 35", err)
			return
		}
		DoctorJSON, err := json.Marshal(Doctordata)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Line 41", err)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(DoctorJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func InsertAppoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var appoint BookAppointment.Appointment
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bytedata, &appoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = BookAppointment.BookAppointment(appoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func UpdateAppoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var UpdateAppoint UpdateAppointment.UpdateAppoint
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bytedata, &UpdateAppoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = UpdateAppointment.UpdateAppointment(UpdateAppoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/checkavailability", GetDoctorSlot)
	http.HandleFunc("/bookappoint", InsertAppoint)
	http.HandleFunc("/updateappoint", UpdateAppoint)
	DatabaseConn.SetUpDBConn()
	http.ListenAndServe(":8000", nil)
}
