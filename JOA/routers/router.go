package routers

import (
	"golang/JOA/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.FileController{}, "post:HandleUpload")
	beego.Router("/process", &controllers.FileController{}, "post:ProcessUploadedFile")
	userNs := beego.NewNamespace("/user",
		beego.NSInclude(
			&controllers.UserController{},
		),
	)
	beego.AddNamespace(userNs)
}
