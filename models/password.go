package models

// PasswordReq defines the password's request.
type PasswordReq struct {
	Length    int     `json:"length"`
	ExtraNums int     `json:"extra_nums"`
	ExtraSpec int     `json:"extra_spec"`
	Helper    *string `json:"helper"`
}

// PasswordResp defines the password's' response.
type PasswordResp struct {
	Passwd string `json:"passwd"`
	Helper string `json:"helper"`
}
