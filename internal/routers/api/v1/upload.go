package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lgl/blog-service/internal/service"
	"github.com/lgl/blog-service/pkg/app"
	"github.com/lgl/blog-service/pkg/convert"
	"github.com/lgl/blog-service/pkg/errcode"
	"github.com/lgl/blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.NewService(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		errResp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
