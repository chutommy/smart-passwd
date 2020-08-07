package models

// PasswordReq defines the password's request.
type PasswordReq struct {
	Length        int    `json:"length" binding:"required,min=5,max=32"`
	ExtraSecurity int    `json:"extra_security" binding:"required,min=0,max=10"`
	Helper        string `json:"helper" binding:"omitempty"`
}

// PasswordResp defines the password's' response.
type PasswordResp struct {
	Passwd string `json:"passwd"`
	Helper string `json:"helper"`
}
