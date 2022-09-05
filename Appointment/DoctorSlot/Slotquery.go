package DoctorSlot

import "hospital/Appointment/DatabaseConn"

func GetDoctorByTimeSlot(Doctor DoctorDB) ([]DoctorDB, error) {
	results, err := DatabaseConn.Dbconn.Query(`SELECT * FROM doctorlist WHERE In_time<=? AND Out_time>=?`, Doctor.DocInTime, Doctor.DocOutTime)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	doctors := make([]DoctorDB, 0)
	for results.Next() {
		var doctorSlot DoctorDB
		results.Scan(&doctorSlot.DocID, &doctorSlot.DocName, &doctorSlot.DocInTime, &doctorSlot.DocOutTime)
		doctors = append(doctors, doctorSlot)
	}
	return doctors, nil

}
