// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/madalingrecescu/PizzaDelivery/internal/swagger/restapi/operations"
	"github.com/madalingrecescu/PizzaDelivery/internal/swagger/restapi/operations/users"
)

//go:generate swagger generate server --target ../../swagger --name Pizzeria --spec ../../../configs/swagger/users_swagger.yaml --principal interface{}

func configureFlags(api *operations.PizzeriaAPI) {
	// handlers.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PizzeriaAPI) http.Handler {
	// configure the handlers here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// handlers.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// handlers.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.UsersPostLoginHandler == nil {
		api.UsersPostLoginHandler = users.PostLoginHandlerFunc(func(params users.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostLogin has not yet been implemented")
		})
	}
	if api.UsersPostSignupHandler == nil {
		api.UsersPostSignupHandler = users.PostSignupHandlerFunc(func(params users.PostSignupParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostSignup has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
