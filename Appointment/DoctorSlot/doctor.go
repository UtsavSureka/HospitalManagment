package DoctorSlot

type DoctorDB struct {
	DocID      int    `json:"DocID"`
	DocName    string `json:"DocName"`
	DocInTime  int    `json:"DocInTime"`
	DocOutTime int    `json:"DocOutTime"`
}
