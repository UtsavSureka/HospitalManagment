package DoctorDetail

type DoctorDB struct {
	DocID      int    `json:"DocID"`
	DocName    string `json:"DocName"`
	DocInTime  string `json:"DocInTime"`
	DocOutTime string `json:"DocOutTime"`
}
