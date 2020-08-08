package models

// PasswordReq defines the password's request.
type PasswordReq struct {
	Length        int    `json:"len" binding:"omitempty,min=5,max=32"`
	ExtraSecurity int    `json:"extra" binding:"omitempty,min=0,max=10"`
	Helper        string `json:"helper" binding:"omitempty,max=60"`
}

// PasswordResp defines the password's' response.
type PasswordResp struct {
	Passwd string `json:"password"`
	Helper string `json:"helper"`
}
