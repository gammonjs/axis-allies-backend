package application

import (
	"axis-allies-backend/pkg/contracts/utility"
	"axis-allies-backend/pkg/contracts/web/api"
	"axis-allies-backend/pkg/handler"
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Router        api.Router
	Configuration utility.Configuration
	Log           utility.Logger
}

func (server Server) ListenAndServe() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	httpServer := &http.Server{
		Addr:    server.Configuration.ServerUrl(),
		Handler: server.Router,
	}

	home := &handler.Home{}
	server.Router.Get("/", home.Handler)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			server.Log.Error("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	server.Log.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		server.Log.Error("Server forced to shutdown: ", err)
	}

	server.Log.Info("Server exiting")
}
