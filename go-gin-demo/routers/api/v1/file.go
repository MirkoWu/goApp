package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mirkowu/go-gin-demo/pkg/e"
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"github.com/mirkowu/go-gin-demo/pkg/util"
	"log"
	"net/http"
)

const UPLOAD_FILE_PATH = "static/upload_file/"

/**上传方法**/
func UploadFile(c *gin.Context) (code int, data []string, context *gin.Context) {
	//得到上传的文件
	file, header, err := c.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	code = e.ERROR_UPLOAD_FILE
	if err != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": code,
		//	"msg":  e.GetMsg(code),
		//	"data": nil,
		//})
		return code, data, c
	}
	//文件的名称
	filename := header.Filename
	fmt.Println(file, err, filename)

	//放到static/upload_file/ 文件夹下
	util.OpenFile(UPLOAD_FILE_PATH, filename)
	filePath := UPLOAD_FILE_PATH + filename
	fmt.Println("filePath=" + filePath)

	//var data []string
	if err := c.SaveUploadedFile(header, filePath); err != nil {
		logging.Fatal(err)
	} else {
		code = e.SUCCESS
		data = append(data, filePath)
	}

	return code, data, c

	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": data,
	//})
}

/**上传方法**/
func UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	code := e.ERROR_UPLOAD_FILE
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": nil,
		})
		return
	}

	files := form.File["file"]
	var data []string
	for _, file := range files {
		log.Println(file.Filename)
		filename := file.Filename
		//放到static/upload_file/ 文件夹下
		util.OpenFile(UPLOAD_FILE_PATH, filename)
		filePath := UPLOAD_FILE_PATH + filename
		fmt.Println("filePath=" + filePath)

		c.SaveUploadedFile(file, filePath)
		data = append(data, filePath)
	}
	if len(files) != 0 && len(data) == len(files) {
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
