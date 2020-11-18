package handle

import (
	openapi "github.com/birnamwood/go-open-api/openapi/gen"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
}

func (si userHandler) FindPets(ctx echo.Context, params openapi.FindUsersParams) error {
	// GET /pets の処理記載
	return nil
}

func (si userHandler) AddUser(ctx echo.Context) error {
	b := openapi.NewUser{}
	// リクエスト Body は echo の APIを利用
	ctx.Bind(&b)

	// POST /pets の処理記載
	return nil
}

func (si userHandler) DeletePet(ctx echo.Context, id int64) error {
	// DELETE /pets/{id} の処理記載
	return nil
}

func (si userHandler) FindUserById(ctx echo.Context, id int64) error {
	// GET /pets/{id} の処理記載
	return nil
}
