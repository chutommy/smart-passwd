package server

// GenRequest represents a request for the password generation.
// It holds all requirements for a new generation.
type GenRequest struct {
	Length        int    `json:"len" binding:"omitempty,min=5,max=32"`
	ExtraSecurity int    `json:"extra" binding:"omitempty,min=0,max=10"`
	Helper        string `json:"helper" binding:"omitempty,max=60"`
}

// GenResponse represents a response of a password generation.
type GenResponse struct {
	Passwd string `json:"password"`
	Helper string `json:"helper"`
}
