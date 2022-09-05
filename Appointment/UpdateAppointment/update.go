package UpdateAppointment

type UpdateAppoint struct {
	PatientName   string `json:"PatientName"`
	PatientMobile int    `json:"PatientMobile"`
	PatientEmail  string `json:"PatientEmail"`
	AppointID     int    `json:"AppointID"`
}
