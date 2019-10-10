package bootstrap

import (
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
	"github.com/zidni722/golang-restfull/app/web/response"
)

type Configurator func(bootstrapper *Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
	Sessions     *sessions.Sessions
}

// New returns a new Bootstrapper
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews loads the template
func (b *Bootstrapper) SetupViews(viewsDir string) {
	b.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
}

// SetupSessions initializes the sessions, optionally
func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

// SetupWebsockets prepare the websocket server
func (b *Bootstrapper) SetupWebsockets(endpoint string, onConnection websocket.ConnectionFunc) {
	ws := websocket.New(websocket.Config{})
	ws.OnConnection(onConnection)

	b.Get(endpoint, ws.Handler())
	b.Any("/iris-ws.js", func(context iris.Context) {
		context.Write(websocket.ClientSource)
	})
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnErrorCode(response.INTERNAL_SERVER_ERROR, func(ctx context.Context) {
		response.InternalServerErrorResponse(ctx, nil)
	})
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public/"

	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Boostrapper's context
func (b *Bootstrapper) Configure(cfgs ...Configurator) {
	for _, cfg := range cfgs {
		cfg(b)
	}
}

// Bootstrap prepares our application.
// Return itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupSessions(24*time.Hour, []byte("zidni-api"), []byte("zidni-api-key"))
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr"
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
