/*
 * AWVS12 client api
 *
 * Awvs12 client api [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// LoginReq struct for LoginReq
type LoginReq struct {
	Email string `json:"email"`
	// sha256(password)
	Password       string `json:"password"`
	RememberMe     bool   `json:"remember_me,omitempty"`
	LogoutPrevious bool   `json:"logout_previous,omitempty"`
}
