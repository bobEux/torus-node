package auth

import (
	"fmt"
	"strings"

	"github.com/torusresearch/bijson"
)

type TestVerifier struct {
	CorrectID string
}

type TestVerifierParams struct {
	IDToken string `json:"id_token"`
	ID      string `json:"id"`
}

// GetIdentifier - return identifier string for verifier
func (v *TestVerifier) GetIdentifier() string {
	return "test"
}

// CleanToken - ensure that incoming token conforms to strict format to prevent replay attacks
func (v *TestVerifier) CleanToken(token string) string {
	return strings.Trim(token, " ")
}

// VerifyRequestIdentity - verifies identity of user based on their token
func (v *TestVerifier) VerifyRequestIdentity(jsonToken *bijson.RawMessage) (bool, string, error) {
	var p TestVerifierParams
	if err := bijson.Unmarshal(*jsonToken, &p); err != nil {
		return false, "", err
	}

	p.IDToken = v.CleanToken(p.IDToken)

	if p.IDToken != v.CorrectID {
		return false, "", fmt.Errorf("Token is not blublu")
	}

	return true, p.ID, nil
}

// NewTestVerifier - Constructor for the default test verifier
func NewTestVerifier(correctID string) *TestVerifier {
	return &TestVerifier{
		CorrectID: correctID,
	}
}