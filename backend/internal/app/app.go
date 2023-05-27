package app

import (
	"github.com/BoryslavGlov/logrusx"
	"volunteering/pkg/repository"
)

type Handler struct {
	logg logrusx.Logging
	rep  repository.Repository
}

func NewApp(logg logrusx.Logging, rep repository.Repository) *Handler {
	return &Handler{logg: logg, rep: rep}
}
