package server

import (
	"noticbackend/app/services/note"
	"noticbackend/config"
)

type Options struct {
	ServiceNote note.Service
	Config      *config.Config
}
