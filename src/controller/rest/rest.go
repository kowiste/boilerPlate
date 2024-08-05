package rest

import (
	"time"

	conf "github.com/kowiste/boilerplate/src/config"
	assetapi "github.com/kowiste/boilerplate/src/controller/rest/asset"
	userapi "github.com/kowiste/boilerplate/src/controller/rest/user"
	"go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"github.com/kowiste/config"
	"go.opentelemetry.io/otel/attribute"
)

type API struct {
	router *gin.Engine
	tracer trace.Tracer
}

// Option represents a configuration option.
type Option func(*API) error

func New(opts ...Option) (api *API) {
	a := new(API)
	a.applyOptions(opts...)
	return a
}

// WithTracer sets the tracer for the API instance.
func WithTracer(tracer *trace.Tracer) Option {
	return func(a *API) error {
		a.tracer = *tracer
		return nil
	}
}

func (g *API) applyOptions(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(g); err != nil {
			return err
		}
	}
	return nil
}
func (a *API) Init() (err error) {
	c, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		return
	}
	//Adding telemetry if is active
	a.router = gin.Default()
	if a.tracer != nil {
		a.router.Use(a.telemetry)
	}

	user, err := userapi.New()
	if err != nil {
		return
	}
	user.Routes(a.router)
	asset, err := assetapi.New()
	if err != nil {
		return
	}
	asset.Routes(a.router)

	go func() {
		err = a.router.Run(":" + c.ServicePort) // listen and serve on port 8080
		if err != nil {
			panic(err)
		}
	}()
	return
}
func (a *API) telemetry(c *gin.Context) {
	ctx, span := a.tracer.Start(c.Request.Context(), c.FullPath())
	defer span.End()

	// Record the start time
	startTime := time.Now()
	c.Request = c.Request.WithContext(ctx)

	c.Next()

	// Record the end time
	endTime := time.Now()
	latency := endTime.Sub(startTime)
	span.SetAttributes(
		attribute.String("http.method", c.Request.Method),
		attribute.String("http.url", c.Request.URL.Path),
		attribute.String("http.client_ip", c.ClientIP()),
		attribute.Int("http.status_code", c.Writer.Status()),
		attribute.String("http.user_agent", c.Request.UserAgent()),
		attribute.Float64("http.latency", latency.Seconds()),
	)

}
