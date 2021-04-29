package engine

// Request represents a request for the Engine.
type Request struct {
	length   int16
	extraSec int16
	helper   string
}

// Response represents a response for the Engine.
type Response struct {
	password string
	helper   string
}

// NewRequest is a constructor for the Request.
func NewRequest(length, extraSec int16, helper string) *Request {
	return &Request{
		length:   length,
		extraSec: extraSec,
		helper:   helper,
	}
}

// Length returns a length of the Request r.
func (r *Request) Length() int16 {
	return r.length
}

// ExtraSec returns a value of the level of the extra security level of the Request r.
func (r *Request) ExtraSec() int16 {
	return r.extraSec
}

// Helper returns a helper message of the Request r.
func (r *Request) Helper() string {
	return r.helper
}

// NewResponse is a constructor for the Response.
func NewResponse(password, helper string) *Response {
	return &Response{
		password: password,
		helper:   helper,
	}
}

// Password returns a password value of the Response r.
func (r *Response) Password() string {
	return r.password
}

// Helper returns a helper value of the Response r.
func (r *Response) Helper() string {
	return r.helper
}
