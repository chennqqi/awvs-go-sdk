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

// Pagination struct for Pagination
type Pagination struct {
	PreviousCursor int64 `json:"previous_cursor,omitempty"`
	NextCursor int64 `json:"next_cursor,omitempty"`
}
