package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewRouter(app *Resolver) http.Handler {
	r := gin.Default()

	graphqlPlayground := playground.Handler(
		"Contents admin api",
		"/query",
	)

	graphqlRoot := handler.New(NewExecutableSchema(
		Config{
			Resolvers: app,
		},
	))

	graphqlRoot.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	graphqlRoot.AddTransport(transport.Options{})
	graphqlRoot.AddTransport(transport.GET{})
	graphqlRoot.AddTransport(transport.POST{})
	graphqlRoot.AddTransport(transport.MultipartForm{
		MaxUploadSize: 400 << 20,
	})

	graphqlRoot.SetQueryCache(lru.New(1000))

	graphqlRoot.Use(extension.Introspection{})
	graphqlRoot.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	r.GET("/", gin.WrapF(graphqlPlayground))
	r.Any("/query", gin.WrapH(graphqlRoot))

	return r
}
