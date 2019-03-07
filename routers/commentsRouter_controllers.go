package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:id`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/backend-to-do-list-Golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
