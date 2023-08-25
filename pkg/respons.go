package pkg

import (
	"ninja1cak/coffeshop-be/config"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code        int         `json:"-"`
	Status      int         `json:"status"`
	Description interface{} `json:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func (r *Response) Send(ctx *gin.Context) {
	ctx.JSON(r.Code, r)
	ctx.Abort()
	return
}

func NewResponse(code int, data *config.Result) *Response {
	respon := Response{
		Code:        code,
		Status:      code,
		Description: getStatus(code),
	}

	if respon.Code >= 400 {
		respon.Description = data.Data
	} else if data.Message != nil {
		respon.Description = data.Message
	} else {
		respon.Data = data.Data
	}

	if data.Meta != nil {
		respon.Meta = data.Meta
	}

	return &respon
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
		break
	case 201:
		desc = "Created"
		break
	case 400:
		desc = "Bad Request"
		break
	case 401:
		desc = "Unauthorized"
		break
	case 403:
		desc = "Forbidden"
		break
	case 404:
		desc = "Not Found"
		break
	case 500:
		desc = "Internal Server Error"
		break
	case 501:
		desc = "Bad Gateway"
		break
	case 304:
		desc = "Not Modified"
		break
	default:
		desc = ""
	}

	return desc
}
