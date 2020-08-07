package models

// PasswordReq defines the password's request.
type PasswordReq struct {
	Length    int
	ExtraNums int
	ExtraSpec int
	Helper    *string
}

// PasswordResp defines the password's' response.
type PasswordResp struct {
	Passwd string
	Helper string
}
