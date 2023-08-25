package handlers

import (
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/internal/repositories"
	"ninja1cak/coffeshop-be/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email_user    string `db:"email_user" form:"email_user"`
	Password_user string `db:"password_user" form:"password_user"`
}

type HandlerAuth struct {
	*repositories.RepoUser
}

func NewAuth(r *repositories.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	userFromDB, err := h.GetAuthData(user.Email_user)
	if err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if *userFromDB.Status == "pending" {
		pkg.NewResponse(401, &config.Result{
			Data: "Your account is not verify",
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(userFromDB.Password_user, user.Password_user); err != nil {
		pkg.NewResponse(401, &config.Result{
			Data: "wrong password",
		}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(userFromDB.Id_user, user.Email_user, userFromDB.Role)
	token, err := jwtt.Generate()
	pkg.NewResponse(200, &config.Result{
		Data: token,
	}).Send(ctx)
	return

}

func (h *HandlerAuth) VerifyAccount(ctx *gin.Context) {
	token := ctx.Param("token")
	check, err := pkg.VerifyToken(token)
	if err != nil {
		pkg.NewResponse(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
	}
	data, err := h.UpdateStatusUser(check.Email)
	pkg.NewResponse(200, &config.Result{
		Data: data,
	}).Send(ctx)
	return

}
