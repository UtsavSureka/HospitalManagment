package PatientDetail

import "hospital/Patient/DatabaseConn"

func GetPatientDetail(doctor DocID) ([]PatientDB, error) {
	results, err := DatabaseConn.Dbconn.Query(`SELECT patient_name , patient_phone ,patient_email FROM doc_appointment WHERE doctor_id=?`, doctor.DoctorID)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	patients := make([]PatientDB, 0)
	for results.Next() {
		var patient PatientDB
		results.Scan(&patient.Patient_Name, &patient.Patient_Number, &patient.Patient_Email)
		patients = append(patients, patient)
	}
	return patients, nil
}
