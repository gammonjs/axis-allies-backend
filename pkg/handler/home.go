package handler

import (
	"axis-allies-backend/pkg/contracts/web/api"
	"net/http"
)

type HomeResponse struct {
	Hello string
}

type Home struct{}

func (home Home) Handler(context api.Context) {
	content := HomeResponse{Hello: "World"}
	context.JSON(http.StatusOK, content)
}
