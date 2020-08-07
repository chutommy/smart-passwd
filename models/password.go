package models

// PasswordReq defines the password's request.
type PasswordReq struct {
	Length        int    `json:"len" binding:"required,min=5,max=32"`
	ExtraSecurity int    `json:"extra" binding:"required,min=0,max=10"`
	Helper        string `json:"note" binding:"omitempty"`
}

// PasswordResp defines the password's' response.
type PasswordResp struct {
	Passwd string `json:"password"`
	Helper string `json:"note"`
}
