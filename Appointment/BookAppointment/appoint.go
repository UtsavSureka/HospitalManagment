package BookAppointment

type Appointment struct {
	Patient_Name   string `json:"PatientName"`
	Patient_Mobile int    `json:"PatientMobile"`
	Patient_email  string `json:"PatientEmail"`
	Doctor_ID      int    `json:"DoctorID"`
}
