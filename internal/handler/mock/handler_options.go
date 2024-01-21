package mock

import (
	"time"

	"github.com/charmbracelet/log"

	"github.com/evg4b/uncors/internal/config"
	"github.com/spf13/afero"
)

type HandlerOption = func(*Handler)

func WithLogger(logger *log.Logger) HandlerOption {
	return func(h *Handler) {
		h.logger = logger
	}
}

func WithResponse(response config.Response) HandlerOption {
	return func(h *Handler) {
		h.response = response
	}
}

func WithFileSystem(fs afero.Fs) HandlerOption {
	return func(h *Handler) {
		h.fs = fs
	}
}

func WithAfter(after func(duration time.Duration) <-chan time.Time) HandlerOption {
	return func(h *Handler) {
		h.after = after
	}
}
