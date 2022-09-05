package BookAppointment

import (
	"fmt"
	"hospital/Appointment/DatabaseConn"
)

func BookAppointment(appoint Appointment) error {
	var app Appointment
	results := DatabaseConn.Dbconn.QueryRow(`SELECT Doc_id FROM doctorlist WHERE Doc_id=?`, appoint.Doctor_ID)
	err := results.Scan(&app.Doctor_ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = DatabaseConn.Dbconn.Exec(`INSERT INTO doc_appointment(patient_name,patient_phone,patient_email,doctor_id) VALUES(?,?,?,?)`, appoint.Patient_Name, appoint.Patient_Mobile, appoint.Patient_email, appoint.Doctor_ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
