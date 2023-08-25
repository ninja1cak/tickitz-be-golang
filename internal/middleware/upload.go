package middleware

import (
	"log"
	"net/http"
	"ninja1cak/coffeshop-be/pkg"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		log.Println(err)
		if err.Error() == "http: no such file" {
			emptyString := ""
			ctx.Set("image", &emptyString)
			ctx.Next()
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing file"})
		return
	}
	src, err := file.Open()
	result, err := pkg.Cloudinary(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to open file",
		})
		return
	}
	log.Println(result)
	ctx.Set("image", &result)
	ctx.Next()
}
