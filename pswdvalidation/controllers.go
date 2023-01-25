package pswdvalidation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PasswordValidation returns a response to whether the password is valid or not
func PasswordValidation(c *gin.Context) {
	var p PasswordModel

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Bad Request" + err.Error(),
		})
		return
	}

	res := rulesCases(p)

	c.JSON(http.StatusOK, res)
}

// rulesCases returns a ResponseModel based on the rules contained in PasswordModel.
// Each rule is processed to verify if the password is valid or not.
func rulesCases(p PasswordModel) ResponseModel {

	response := make([]string, 0)

	for _, r := range p.Rules {
		switch r.Rule {

		case MINSIZE:
			if !MinSize(p.Password, r.Value) {
				response = append(response, string(MINSIZE))
			}
		case MINDIGIT:
			if !MinDigit(p.Password, r.Value) {
				response = append(response, string(MINDIGIT))
			}
		case MINSPECIALCHARS:
			if !MinSpecialChars(p.Password, r.Value) {
				response = append(response, string(MINSPECIALCHARS))
			}
		case MINUPPERCASE:
			if !MinUpperCase(p.Password, r.Value) {
				response = append(response, string(MINUPPERCASE))
			}
		case MINLOWERCASE:
			if !MinLowerCase(p.Password, r.Value) {
				response = append(response, string(MINLOWERCASE))
			}
		case NOREPETED:
			if HasRepeated(p.Password) {
				response = append(response, string(NOREPETED))
			}
		}
	}

	return ResponseModel{
		Verify:  len(response) <= 1,
		NoMatch: response,
	}
}
