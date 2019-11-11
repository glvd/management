package management

import (
	"github.com/godcong/go-trait"
	"github.com/goextension/log"
)

func init() {
	logger := trait.NewZapSugar()
	log.Register(logger)
}
