package models

// PasswordReq defines the password's request.
type PasswordReq struct {
	Length        int     `json:"length"`
	ExtraSecurity int     `json:"extra_security"`
	Helper        *string `json:"helper"`
}

// PasswordResp defines the password's' response.
type PasswordResp struct {
	Passwd string `json:"passwd"`
	Helper string `json:"helper"`
}
