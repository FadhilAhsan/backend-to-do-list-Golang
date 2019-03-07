package controllers

import (
	"github.com/FadhilAhsan/backend-to-do-list-Golang/models"
	"encoding/json"
	"github.com/FadhilAhsan/backend-to-do-list-Golang/services"
	"github.com/astaxie/beego"
	"log"
)

// Operations about Users
type TaskController struct {
	beego.Controller
}

// @Title CreateTask
// @Description create tasks
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {object} models.JSONResponse
// @Failure 403 {object} models.JSONResponse
// @router / [post]
func (this *TaskController) Post() {
	task := models.Task{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &task)
	if err != nil || task.Task == "" {
		log.Println("[Error] TaskController.Post : ", err)
		log.Println("[Error] TaskController.Post Task : -->", task.Task)
		this.Ctx.Output.SetStatus(400)
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "Task is required,",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}


	err = services.CreateTask(&task)
	if err != nil {
		log.Println("[Error] TaskController.Post : ", err)
		this.Ctx.Output.SetStatus(400)
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "Task failed created",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	meta := models.JSONResponseMeta{
		Code:    200,
		Status:  "success",
		Message: "Success created task",
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: task,
	}

	this.Data["json"] = result
	this.Ctx.Output.SetStatus(200)
	this.ServeJSON()
}

// @Title GetAll
// @Description get all tasks
// @Success 200 {object} models.JSONResponse
// @Failure 403 {object} models.JSONResponse
// @router / [get]
func (this *TaskController) GetAll() {
	limit, errLimit := this.GetInt("limit")
	if errLimit != nil {
		limit = 5
	}

	offset, errOffset := this.GetInt("offset")
	if errOffset != nil {
		offset = 0
	}

	data, err := services.GetAllTask(limit,offset) 
	if err != nil {
		log.Println("[Error] TaskController.GetAll : ", err)
		this.Ctx.Output.SetStatus(403)
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "failed get all task",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	meta := models.JSONResponseMeta{
		Code:    200,
		Status:  "success",
		Message: "Success get all task",
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: data,
	}

	this.Data["json"] = result
	this.Ctx.Output.SetStatus(200)
	this.ServeJSON()
}

// @Title Get
// @Description get task by id
// @Param	id		path 	string	true		"ID Task"
// @Success 200 {object} models.JSONResponse
// @Failure 403 {object} models.JSONResponse
// @router /:id [get]
func (this *TaskController) Get() {
	id, err := this.GetInt64(":id")
	if err != nil {
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "ID must a number",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	data, err := services.GetTaskById(id) 
	if err != nil {
		log.Println("[Error] TaskController.Get : ", err)
		this.Ctx.Output.SetStatus(403)
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "failed get task by id",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	meta := models.JSONResponseMeta{
		Code:    200,
		Status:  "success",
		Message: "Success get task by id",
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: data,
	}

	this.Data["json"] = result
	this.Ctx.Output.SetStatus(200)
	this.ServeJSON()
}

// @Title Update
// @Description update tasks by id
// @Param	id		path 	string	true		"ID Task"
// @Param	body		body 	models.Task	true		"Task"
// @Success 200 {object} models.JSONResponse
// @Failure 403 {object} models.JSONResponse
// @router /:id [post]
func (this *TaskController) Update() {
	id, err := this.GetInt64(":id")
	if err != nil {
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "ID must a number",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	task := models.Task{Id:id}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &task)
	if err != nil || task.Task == "" {
		log.Println("[Error] TaskController.Update : ", err)
		this.Ctx.Output.SetStatus(400)
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "Task is required,",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	err = services.UpdateTaskById(task) 
	if err != nil {
		log.Println("[Error] TaskController.Update : ", err)
		this.Ctx.Output.SetStatus(403)
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "failed update task by id",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	meta := models.JSONResponseMeta{
		Code:    200,
		Status:  "success",
		Message: "Success update task by id",
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: task,
	}

	this.Data["json"] = result
	this.Ctx.Output.SetStatus(200)
	this.ServeJSON()
}

// @Title Delete
// @Description delete task by id
// @Param	id		path 	string	true		"ID Task"
// @Success 200 {object} models.JSONResponse
// @Failure 403 {object} models.JSONResponse
// @router /:id [delete]
func (this *TaskController) Delete() {
	id, err := this.GetInt64(":id")
	if err != nil {
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "ID must a number",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	err = services.DeleteTaskById(id) 
	if err != nil {
		log.Println("[Error] TaskController.Delete : ", err)
		this.Ctx.Output.SetStatus(403)
		meta := models.JSONResponseMeta{
			Code:    403,
			Status:  "Failed",
			Message: "failed get task by id",
		}
		this.Data["json"] = models.JSONResponse{Meta: meta}
		this.ServeJSON()
		return
	}

	meta := models.JSONResponseMeta{
		Code:    200,
		Status:  "success",
		Message: "Success deleted task by id",
	}

	this.Data["json"] = models.JSONResponse{Meta: meta}
	this.Ctx.Output.SetStatus(200)
	this.ServeJSON()
}


