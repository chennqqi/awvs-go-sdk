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
// TargetConfigLoginCredentials struct for TargetConfigLoginCredentials
type TargetConfigLoginCredentials struct {
	Enable bool `json:"enable,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}