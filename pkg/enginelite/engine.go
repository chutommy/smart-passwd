package enginelite

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// ErrInvalidRequirements is returned if request is invalid.
var ErrInvalidRequirements = errors.New("invalid requirements")

// Engine represents a password generator engine.
type Engine struct {
	c *Constructor
	s *Swapper
}

// Init initializes a new Engine.
func Init(constructor *Constructor, swapper *Swapper) *Engine {
	return &Engine{
		c: constructor,
		s: swapper,
	}
}

// Generate generates a new password using a ERequest.
func (e *Engine) Generate(_ context.Context, req *Request) (*Response, error) {
	req.helper = strings.TrimSpace(req.helper)

	switch {
	case req.helper != "" && req.length != 0:
		return nil, fmt.Errorf("%w: both length and helper defined", ErrInvalidRequirements)
	case req.helper == "" && req.length < e.c.min:
		return nil, fmt.Errorf("%w: length too small", ErrInvalidRequirements)
	case req.extraSec < 0:
		return nil, fmt.Errorf("%w: negative extra security value", ErrInvalidRequirements)
	}

	h, err := e.helper(req)
	if err != nil {
		return nil, err
	}

	p := e.swap(h)
	p = e.s.ExtraSec(p, req.extraSec)

	hs := strings.Join(h, " ")
	if req.helper != "" {
		hs = req.helper
	}

	return NewResponse(p, hs), nil
}

// swap generates a password from the helper and randomly swaps each
// letters with an upper-case or similarly looking letter.
func (e *Engine) swap(helper []string) string {
	p := []rune(strings.Join(helper, ""))
	for i, r := range p {
		switch e.s.rand.Intn(4) {
		case 0:
			continue
		case 1:
			p[i] = unicode.ToUpper(r)
		default:
			p[i] = e.s.Swap(r)
		}
	}

	return string(p)
}

// helper retrieves a helper from the request.
// If helper is not provided, it generates a new one.
func (e *Engine) helper(req *Request) ([]string, error) {
	hs := strings.ToLower(req.helper)
	h := strings.Split(hs, " ")

	return h, nil
}
