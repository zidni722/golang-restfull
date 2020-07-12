package controllers

import (
	"github.com/jinzhu/gorm"
	golog "github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	user_request "github.com/zidni722/golang-restfull/app/dto/request/crud"
	"github.com/zidni722/golang-restfull/app/models"
	_interface "github.com/zidni722/golang-restfull/app/repositories/interface"
	"github.com/zidni722/golang-restfull/app/web/response"
)

type UserController struct {
	Db             *gorm.DB
	UserRepository _interface.IUserRepository
}

func NewUserController(db *gorm.DB, userRepository _interface.IUserRepository) *UserController {
	return &UserController{
		Db:             db,
		UserRepository: userRepository,
	}
}

func (c *UserController) CreateUserHandler(ctx iris.Context) {
	tx := c.Db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			response.InternalServerErrorResponse(ctx, r)
			return
		}
	}()

	formRequest := user_request.NewUserRequest(ctx, c.Db, c.UserRepository)

	if err := ctx.ReadJSON(&formRequest.Form); err != nil {
		response.InternalServerErrorResponse(ctx, err)
	}

	if !formRequest.Validate() {
		return
	}

	var user models.User

	golog.Info(formRequest)

	user.Name = formRequest.Form.Name
	user.Address = formRequest.Form.Address
	user.Gender = formRequest.Form.Gender

	if err := c.UserRepository.Create(c.Db, &user); err != nil {
		tx.Rollback()
		response.InternalServerErrorResponse(ctx, err)
		return
	}

	tx.Commit()
	response.SuccessResponse(ctx, response.OK, response.SUCCESS_SAVE_USER, nil)
}
