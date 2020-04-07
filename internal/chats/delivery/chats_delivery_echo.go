package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"warpin/configs"
	"warpin/internal/chats/domain"
	"warpin/internal/chats/usecase"
	"warpin/tools"
)

type chatsDeliveryEcho struct {
	usecase usecase.ChatUsecase
}

func NewChatsDeliveryEcho(usecase usecase.ChatUsecase) *chatsDeliveryEcho {
	return &chatsDeliveryEcho{
		usecase: usecase,
	}
}

func (c *chatsDeliveryEcho) Mount(group *echo.Group) {
	// create endpoint group
	group.GET("", c.GetAll)
	group.POST("", c.Insert)
	group.GET("/room", c.Room)
	group.GET("/websocket", c.Websocket)
}

func (c *chatsDeliveryEcho) Insert(ec echo.Context) error {
	ctx := "ChatsDelivery-Insert"

	var (
		data = &domain.Chats{}
	)

	// bind data to struct
	err := ec.Bind(data)
	if err != nil {
		tools.Log(tools.ErrorLevel, err.Error(), ctx, "bind_data")
		return ec.JSON(http.StatusBadRequest, tools.ErrorDeliveryResult(http.StatusBadRequest, err.Error()))
	}

	// validate data
	err = ec.Validate(data)
	if err != nil {
		tools.Log(tools.ErrorLevel, err.Error(), ctx, "validate_data")
		return ec.JSON(http.StatusBadRequest, tools.ErrorDeliveryValidatorResult(http.StatusBadRequest, err.Error()))
	}

	// set usecase
	response, err := c.usecase.Insert(ec.Request().Context(), data)
	if err != nil {
		tools.Log(tools.ErrorLevel, err.Error(), ctx, "insert_usecase")
		return ec.JSON(http.StatusBadRequest, tools.ErrorDeliveryResult(http.StatusBadRequest, err.Error()))
	}

	// set response
	return ec.JSON(http.StatusOK, tools.SuccessDeliveryResult(http.StatusOK, response, "successfully insert chats"))
}

func (c *chatsDeliveryEcho) GetAll(ec echo.Context) error {
	ctx := "ChatsDelivery-GetAll"

	var (
		filter = &domain.Filter{}
	)

	// set parse query params
	queryFilter := tools.ParseQueryParamsToFilter(ec.QueryParams())

	// set filter
	filter.Filter = queryFilter

	// set usecase
	response, err := c.usecase.FindAll(ec.Request().Context(), filter)
	if err != nil {
		tools.Log(tools.ErrorLevel, err.Error(), ctx, "findall_usecase")
		return ec.JSON(http.StatusBadRequest, tools.ErrorDeliveryResult(http.StatusBadRequest, err.Error()))
	}

	// set response
	return ec.JSON(http.StatusOK, tools.SuccessDeliveryResult(http.StatusOK, response, "successfully get chats"))
}

func (c *chatsDeliveryEcho) Room(ec echo.Context) error {
	return ec.File("website/chats.html")
}

func (c *chatsDeliveryEcho) Websocket(ec echo.Context) error {
	ctx := "ChatsDelivery-Websocket"

	// set our name from query param
	name := ec.QueryParam("name")

	// add new websocket
	err := configs.AddConnections(ec, name)
	if err != nil {
		tools.Log(tools.ErrorLevel, err.Error(), ctx, "add_websocket")
		return ec.JSON(http.StatusBadRequest, tools.ErrorDeliveryResult(http.StatusBadRequest, err.Error()))
	}

	// set usecase
	c.usecase.Websocket(ec.Request().Context(), name)

	return nil
}
