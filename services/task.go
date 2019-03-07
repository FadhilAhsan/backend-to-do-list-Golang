package services

import (
	"github.com/FadhilAhsan/backend-to-do-list-Golang/models"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
	// "strconv"
)

func CreateTask(task *models.Task) error {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	// dbORM.Begin()
	task.CreatedAt = time.Now()

	id, err := dbORM.Insert(task)
	if err != nil {
		log.Println("[Error] TaskServices.CreateTask : ", err)
	    return  err
	}
	// if err == nil {
	//     dbORM.Commit()
	// } else {
	//     dbORM.Rollback()
	// }

	task.Id = id

	return  nil
}

func GetAllTask(limit int, offset int) ([]models.Task, error) {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	// user := models.Account{Email: email}

	tasks := []models.Task{}
	_, err := dbORM.QueryTable("task").Limit(limit).Offset(offset).RelatedSel().All(&tasks)
	if err != nil {
		log.Println("[Error] TaskServices.GetAllTask : ", err)
	    return  tasks, err
	}

	return tasks, nil
}

func GetTaskById(id int64) (models.Task, error) {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	// user := models.Account{Email: email}

	task := models.Task{Id:id}
	err := dbORM.Read(&task)
	if err != nil {
		log.Println("[Error] TaskServices.GetTaskById : ", err)
	    return  task, err
	}

	return task, nil
}

func UpdateTaskById(task models.Task) error {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	// user := models.Account{Email: email}

	_, err := dbORM.Update(&task)
	if err != nil {
		log.Println("[Error] TaskServices.UpdateTaskById : ", err)
	    return  err
	}

	return nil
}

func DeleteTaskById(id int64) error {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	// user := models.Account{Email: email}

	task := models.Task{Id:id}
	_,err := dbORM.Delete(&task)
	if err != nil {
		log.Println("[Error] TaskServices.DeleteTaskById : ", err)
	    return  err
	}

	return  nil
}