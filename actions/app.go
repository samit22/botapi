package actions

import (
  "os"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/x/sessions"
	"github.com/samitghimire/botapi/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			SessionName:  "_botapi_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))
    app.Use(AuthorizeRequest)

		g := app.Group("/api/v1")
		// g.Use(APIAuthorizer)
		g.GET("/", HomeHandler)
		g.GET("/person" , GetPersonInfo)

	}

	return app
}

// Middleware to Authorize Token Based Request
func AuthorizeRequest(next buffalo.Handler) buffalo.Handler {
  return func(c buffalo.Context) error {
    token := c.Request().Header.Get("API-TOKEN")
    if token == os.Getenv("APP_TOKEN") {
      return next(c)
    } else {
      return c.Render(404, r.JSON(map[string]string{"message": "UnAuthorized Token"}))
    }
  }
}
