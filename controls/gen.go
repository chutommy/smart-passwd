package controls

import (
	"strings"

	models "github.com/chutified/smart-passwd/models"
	"github.com/pkg/errors"
)

// Generate handles the password's request and generates the password,
// which satisfies the requirements.
func (c *Controller) Generate(preq *models.PasswordReq) (*models.PasswordResp, error) {

	var phrase string
	var helper string
	// get phrase
	if preq.Helper == "" {

		// generate words
		ws, err := c.newPhrase(preq.Length)
		if err != nil {
			return nil, errors.Wrap(err, "generating new phrase")
		}

		// handle words
		phrase, helper, err = c.composeWords(ws)
		if err != nil {
			return nil, errors.Wrap(err, "composing generated words")
		}

	} else {
		phrase = strings.Replace(preq.Helper, " ", "", -1)
		helper = preq.Helper
	}

	// transform the password
	phrase = c.transform(phrase)
	nums, specs := extraSecurityLvl(preq.ExtraSecurity)
	phrase = c.randomAdds(phrase, nums, specs)

	// assing into the response
	presp := &models.PasswordResp{
		Passwd: phrase,
		Helper: helper,
	}

	return presp, nil
}
