package providers

import (
	"goravel/app/console"

	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type ConsoleServiceProvider struct {
}

func (receiver *ConsoleServiceProvider) Register(app foundation.Application) {
	kernel := console.Kernel{}
	facades.Schedule().Register(kernel.Schedule())
	facades.Artisan().Register(kernel.Commands())
}

func (receiver *ConsoleServiceProvider) Boot(app foundation.Application) {

}
