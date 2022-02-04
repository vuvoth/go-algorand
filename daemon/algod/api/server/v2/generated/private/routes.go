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
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

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

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

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

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

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

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

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
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

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
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

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
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
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
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
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
