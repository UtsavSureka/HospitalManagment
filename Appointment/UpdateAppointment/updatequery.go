package UpdateAppointment

import (
	"fmt"
	"hospital/Appointment/DatabaseConn"
)

func UpdateAppointment(update UpdateAppoint) error {
	var appID UpdateAppoint
	results := DatabaseConn.Dbconn.QueryRow(`SELECT appointment_ID FROM doc_appointment WHERE appointment_ID=?`, update.AppointID)
	err := results.Scan(&appID.AppointID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = DatabaseConn.Dbconn.Exec(`UPDATE doc_appointment SET patient_name=?,patient_phone=?,patient_email=? WHERE appointment_id=?`, update.PatientName, update.PatientMobile, update.PatientEmail, update.AppointID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
