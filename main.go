package main

import (
	"fmt"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"warpin/configs"
	chatDelivery "warpin/internal/chats/delivery"
	chatRepository "warpin/internal/chats/repository"
	chatUsecase "warpin/internal/chats/usecase"
	"warpin/tools"
)

func main() {
	ctx := "Main"

	// init environment for the firstime
	configs.InitEnv()

	// init config for the firstime
	configs.InitConfig()

	// init log for the firstime
	tools.InitLogs()

	// set repository
	repository := chatRepository.NewChatsRepository(configs.GetConfig().Sqlite.Database)

	// set usecase
	usecase := chatUsecase.NewChatsUsecase(repository)

	// set delivery
	delivery := chatDelivery.NewChatsDeliveryEcho(usecase)

	// set webserver
	e := echo.New()

	// set delivery
	chatsGroup := e.Group("/chats")
	delivery.Mount(chatsGroup)

	// to prevent panic
	//e.Use(middleware.Recover())

	// register translator
	id := id.New()
	uni := ut.New(id, id)

	// set translator language
	trans, _ := uni.GetTranslator("id")

	// set validator
	validator := validator.New()
	id_translations.RegisterDefaultTranslations(validator, trans)

	// set validator
	e.Validator = &tools.Validator{
		Validator:  validator,
		Translator: trans,
	}

	// start app
	app := fmt.Sprintf(":%s", configs.GetEnv().Port)

	// start
	go func() {
		if err := e.Start(app); err != nil {
			tools.Log(tools.PanicLevel, err.Error(), ctx, "start_app")
		}
	}()

	// check for signal if ended
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
