package pswdvalidation

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMinSize(t *testing.T) {
	t.Run("minSize value less than password length", func(t *testing.T) {
		password := "password"
		value := 0

		want := true
		got := MinSize(password, value)
		assert.Equal(t, got, want, "Should pass when minSize is less than password")
	})

	t.Run("minSize value bigger than password length", func(t *testing.T) {
		password := "password"
		value := 10

		want := false
		got := MinSize(password, value)
		assert.Equal(t, got, want, "Should fail when value is bigger than password")
	})

	t.Run("minSize value equal to password length", func(t *testing.T) {
		password := "123"
		value := 3

		want := true
		got := MinSize(password, value)
		assert.Equal(t, got, want, "Should pass when value is equal to password")
	})
}

func TestMinUpperCase(t *testing.T) {
	t.Run("minUpperCase no minimum upper", func(t *testing.T) {
		password := "password"
		value := 0

		want := true
		got := MinUpperCase(password, value)
		assert.Equal(t, got, want, "Should pass when there is no minimum upper case")
	})

	t.Run("minUpperCase at least one uppercase", func(t *testing.T) {
		password := "Password"
		value := 1

		want := true
		got := MinUpperCase(password, value)
		assert.Equal(t, got, want, "Shoud pass when password has at least one uppercase")
	})

	t.Run("minUpperCase password has less uppercase than value", func(t *testing.T) {
		password := "Password"
		value := 2

		want := false
		got := MinUpperCase(password, value)
		assert.Equal(t, got, want, "Should pass when password has less uppercase than value")
	})
}

func TestMinLowerCase(t *testing.T) {
	t.Run("minLowerCase has no minimum lower case", func(t *testing.T) {

		password := "PASSWORD"
		value := 0

		want := true
		got := MinLowerCase(password, value)
		assert.Equal(t, got, want, "Should pass when password has no lowercase")
	})

	t.Run("minLowerCase has at least one lowercase", func(t *testing.T) {
		password := "pASSWORD"
		value := 1

		want := true
		got := MinLowerCase(password, value)
		assert.Equal(t, got, want, "Should pass when password has at least one lowercase")
	})

	t.Run("minLowerCase has less lowercase than value", func(t *testing.T) {
		password := "PASSWOrd"
		value := 3

		want := false
		got := MinLowerCase(password, value)
		assert.Equal(t, got, want, "Should pass when password has less lowercase than value")
	})
}

func TestMinDigit(t *testing.T) {
	t.Run("minDigit has no minimum digit", func(t *testing.T) {
		password := "password"
		value := 0

		want := true
		got := MinDigit(password, value)
		assert.Equal(t, got, want, "Should pass when password has no digit")
	})

	t.Run("minDigit has at least one digit", func(t *testing.T) {
		password := "password1"
		value := 1

		want := true
		got := MinDigit(password, value)
		assert.Equal(t, got, want, "Should pass when password has at least one digit")
	})

	t.Run("minDigit password has less digit than value", func(t *testing.T) {
		password := "password12"
		value := 3

		want := false
		got := MinDigit(password, value)
		assert.Equal(t, got, want, "Should pass when password has less digits than value")
	})
}

func TestMinSpecialChars(t *testing.T) {
	t.Run("minSpecialChars has no minimum special char", func(t *testing.T) {
		password := "password"
		value := 0

		want := true
		got := MinSpecialChars(password, value)
		assert.Equal(t, got, want, "Should pass when password has no special char")
	})

	t.Run("minSpecialChars has at least one special char", func(t *testing.T) {
		password := "password!"
		value := 1

		want := true
		got := MinSpecialChars(password, value)
		assert.Equal(t, got, want, "Should pass when password has at least one special char")
	})

	t.Run("minSpecialChars password has less than the minimum special chars", func(t *testing.T) {
		password := "password!@"
		value := 3

		want := false
		got := MinSpecialChars(password, value)
		assert.Equal(t, got, want, "Should pass when password has less the minimum special chars")
	})
}

func TestHasRepeated(t *testing.T) {
	t.Run("hasRepeated password with no repeat chars", func(t *testing.T) {
		password := "pasword"

		want := false
		got := HasRepeated(password)
		assert.Equal(t, got, want, "Should pass when norepeat is off")
	})

	t.Run("hasRepeated password with no consecutive repeated char", func(t *testing.T) {
		password := "abab"

		want := false
		got := HasRepeated(password)
		assert.Equal(t, got, want, "Should pass when password has no consecutive repeated char")
	})

	t.Run("hasRepeated password has consecutive repeated char", func(t *testing.T) {
		password := "password"

		want := true
		got := HasRepeated(password)
		assert.Equal(t, got, want, "Should pass when password has consecutive repeated char")
	})

}

func TestRouteVerify(t *testing.T) {

	r := gin.Default()

	v1 := r.Group("/api")
	ValidadePassword(v1)

	p := PasswordModel{
		Password: "PsworD2414@!$#",
		Rules: []RuleModel{
			{Rule: MINSIZE,
				Value: 2},
		},
	}

	jsonValue, _ := json.Marshal(p)
	req, _ := http.NewRequest("POST", "/api/verify", bytes.NewBuffer(jsonValue))

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should be status ok")

	mockResponse := ResponseModel{
		Verify:  true,
		NoMatch: []string{},
	}

	mockJson, _ := json.Marshal(mockResponse)
	responseBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, mockJson, responseBody)
}
