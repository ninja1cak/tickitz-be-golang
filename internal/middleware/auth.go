package middleware

import (
	"ninja1cak/coffeshop-be/config"
	"ninja1cak/coffeshop-be/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsVerify(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.NewResponse(401, &config.Result{
				Data: "Please login",
			}).Send(ctx)
			return
		}
		if !strings.Contains(header, "Bearer") {
			pkg.NewResponse(401, &config.Result{
				Data: "Invalid header",
			}).Send(ctx)
			return
		}

		token := strings.Replace(header, "Bearer ", "", -1)
		check, err := pkg.VerifyToken(token)
		if err != nil {
			pkg.NewResponse(401, &config.Result{
				Data: err.Error(),
			}).Send(ctx)
			return
		}

		for _, value := range role {
			if value == check.Role {
				valid = true
			}
		}

		if !valid {
			pkg.NewResponse(401, &config.Result{
				Data: "You not have permission to access",
			}).Send(ctx)
			return
		}
		ctx.Set("user_id", check.Id)
		ctx.Set("email", check.Email)
		ctx.Next()
	}
}
