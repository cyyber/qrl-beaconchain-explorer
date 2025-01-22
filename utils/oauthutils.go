package utils

import (
	"encoding/json"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

const InvalidRequest = "invalid_request"
const InvalidClient = "invalid_client"
const InvalidGrant = "invalid_grant"
const UnauthorizedClient = "unauthorized_client"
const UnsupportedGrantType = "unsupported_grant_type"
const InvalidScope = "invalid_scope"
const AccessDenied = "access_denied"
const ServerError = "server_error"
const TemporarilyUnavailable = "temporarily_unavailable"

// UserIDKey Key for context access to get the validated userID
const ClaimsContextKey = "ClaimsKey"

const JsonBodyKey = "JsonBodyKey"
const JsonBodyNakedKey = "JsonBodyNakedKey"

var logger = logrus.New().WithField("module", "oauth")
var signingMethod = jwt.SigningMethodHS256

// CustomClaims Structure of JWT body, contains standard JWT claims and userID as a custom claim
type CustomClaims struct {
	UserID   uint64 `json:"userID"`
	AppID    uint64 `json:"appID"`
	DeviceID uint64 `json:"deviceID"`
	Package  string `json:"package"`
	Theme    string `json:"theme"`
	jwt.StandardClaims
}

// OAuthResponse Structure of an successful OAuth response
type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// OAuthErrorResponse Structure of an OAuth error response
type OAuthErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

func stripOffBearerFromToken(tokenString string) (string, error) {
	if len(tokenString) > 6 && strings.ToUpper(tokenString[0:6]) == "BEARER" {
		return tokenString[7:], nil
	}
	return tokenString, nil //"", errors.New("Only bearer tokens are supported, got: " + tokenString)
}

// SendOAuthResponse creates and sends a OAuth response according to RFC6749
func SendOAuthResponse(j *json.Encoder, route, accessToken, refreshToken string, expiresIn int) {
	response := OAuthResponse{
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		TokenType:    "bearer",
		RefreshToken: refreshToken,
	}
	err := j.Encode(response)

	if err != nil {
		logger.Errorf("error serializing json error for API %v route: %v", route, err)
	}
}

// SendOAuthErrorResponse creates and sends a OAuth error response according to RFC6749
func SendOAuthErrorResponse(j *json.Encoder, route, errString, description string) {
	response := OAuthErrorResponse{
		Error:       errString,
		Description: description,
	}
	err := j.Encode(response)

	if err != nil {
		logger.Errorf("error serializing json error for API %v route: %v", route, err)
	}
}
