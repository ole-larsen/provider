// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetVkCallbackURL generates an URL for the get vk callback operation
type GetVkCallbackURL struct {
	Authuser *string
	Prompt   *string
	Scope    *string
	State    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetVkCallbackURL) WithBasePath(bp string) *GetVkCallbackURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetVkCallbackURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetVkCallbackURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/vk/callback"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var authuserQ string
	if o.Authuser != nil {
		authuserQ = *o.Authuser
	}
	if authuserQ != "" {
		qs.Set("authuser", authuserQ)
	}

	var promptQ string
	if o.Prompt != nil {
		promptQ = *o.Prompt
	}
	if promptQ != "" {
		qs.Set("prompt", promptQ)
	}

	var scopeQ string
	if o.Scope != nil {
		scopeQ = *o.Scope
	}
	if scopeQ != "" {
		qs.Set("scope", scopeQ)
	}

	var stateQ string
	if o.State != nil {
		stateQ = *o.State
	}
	if stateQ != "" {
		qs.Set("state", stateQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetVkCallbackURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetVkCallbackURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetVkCallbackURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetVkCallbackURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetVkCallbackURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetVkCallbackURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
