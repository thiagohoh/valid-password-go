package pswdvalidation

type RuleType string

// Rules for password validation
const (
	MINSIZE         RuleType = "minSize"
	MINSPECIALCHARS RuleType = "minSpecialChars"
	NOREPETED       RuleType = "noRepeted"
	MINDIGIT        RuleType = "minDigit"
	MINLOWERCASE    RuleType = "minLowerCase"
	MINUPPERCASE    RuleType = "minUpperCase"
)

// RuleModel represents the rules for password validation.
type RuleModel struct {
	Rule  RuleType `json:"rule"`
	Value int      `json:"value"`
}

// PasswordModel represents how data must be shaped
type PasswordModel struct {
	Password string      `json:"password"`
	Rules    []RuleModel `json:"rules"`
}

// ResponseModel represents the respone for password validation
type ResponseModel struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
