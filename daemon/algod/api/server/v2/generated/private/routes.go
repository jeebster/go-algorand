// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/register-participation-keys/{address})
	RegisterParticipationKeys(ctx echo.Context, address string, params RegisterParticipationKeysParams) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// RegisterParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":           true,
		"fee":              true,
		"key-dilution":     true,
		"round-last-valid": true,
		"no-wait":          true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParticipationKeysParams
	// ------------- Optional query parameter "fee" -------------
	if paramValue := ctx.QueryParam("fee"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "fee", ctx.QueryParams(), &params.Fee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fee: %s", err))
	}

	// ------------- Optional query parameter "key-dilution" -------------
	if paramValue := ctx.QueryParam("key-dilution"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "key-dilution", ctx.QueryParams(), &params.KeyDilution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key-dilution: %s", err))
	}

	// ------------- Optional query parameter "round-last-valid" -------------
	if paramValue := ctx.QueryParam("round-last-valid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "round-last-valid", ctx.QueryParams(), &params.RoundLastValid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round-last-valid: %s", err))
	}

	// ------------- Optional query parameter "no-wait" -------------
	if paramValue := ctx.QueryParam("no-wait"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "no-wait", ctx.QueryParams(), &params.NoWait)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter no-wait: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterParticipationKeys(ctx, address, params)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST("/v2/register-participation-keys/:address", wrapper.RegisterParticipationKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
