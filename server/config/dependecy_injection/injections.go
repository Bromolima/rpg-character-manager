package dependecyInjection

import (
	"github.com/Bromolima/rpg-character-manager/handlers"
	"github.com/Bromolima/rpg-character-manager/repository"
	"github.com/Bromolima/rpg-character-manager/service"
	"github.com/samber/do/v2"
)

func Injections(i do.Injector) {
	do.Provide(i, handlers.NewUserHandler)
	do.Provide(i, service.NewUserService)
	do.Provide(i, repository.NewUserRepository)
}
