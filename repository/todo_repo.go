package repository

import (
	"Permanent-Storage-Todo/model"
	"gorm.io/gorm"
	"errors"
)

var (
	ErrNotFound = errors.New("Record not found")
)

type Todo_repo interface{
	Todo_writer
	Todo_reader
}

type Todo_writer interface{
	Create_record(todo *model.Todo) error
	Update_record_by_id(new_todo *model.Todo) error
	Update_records_by_category(category string, new_status bool) error
	Delete_record_by_id(id int) error
	Delete_all_records() error
}

type Todo_reader interface{
	Get_all_records() ([]model.Todo, error)
	Get_record_by_id(id int) (*model.Todo, error)
	Get_records_by_category(category string) ([]model.Todo, error)
	Get_records_by_status(status bool) ([]model.Todo, error)
	Get_records_by_title(title string) ([]model.Todo, error)
}

type postgres_repo struct{
	db *gorm.DB
}

func Init_postgres_repo(new_db *gorm.DB) Todo_repo{
	//initialize new repo and return its address as a repo interface.
	var new_repo = postgres_repo{db: new_db}
	return &new_repo
}

func (repo *postgres_repo) Create_record(todo *model.Todo) error{
	//add the new object to the database.
	result := repo.db.Create(todo)
	//return the result.
	return result.Error
}

func (repo *postgres_repo) Update_record_by_id(new_todo *model.Todo) error{
	//save the new modified object.
	result := repo.db.Save(new_todo)
	//return the result of process.
	return result.Error
}

func (repo *postgres_repo) Update_records_by_category(category string, new_status bool) error{
	//update completed status of records in the passed category.
	result := repo.db.Model(&model.Todo{}).Where("category = ?", category).Update("completed", new_status)
	return result.Error
}

func (repo *postgres_repo) Delete_record_by_id(id int) error{
	//delete the todo that have the passed id.
	result := repo.db.Where("id = ?", id).Delete(&model.Todo{})
	//check the result
	if result.Error != nil{
		return result.Error
	}else if result.RowsAffected == 0{
		return ErrNotFound
	}else{
		return nil
	}
}

func (repo *postgres_repo) Delete_all_records() error{
	//delete all records in the database.
	result := repo.db.Exec("DELETE FROM todos")
	return result.Error
}

func (repo *postgres_repo) Get_all_records() ([]model.Todo, error){
	//initialize a slice for records.
	records := make([]model.Todo, 0)
	//get all records in the database.
	result := repo.db.Order("id asc").Find(&records)
	//if something went wrong, return empty slice and error value.
	if result.Error != nil{
		return nil, result.Error
	}
	return records, nil
}

func (repo *postgres_repo) Get_record_by_id(id int) (*model.Todo, error){
	//define an object for the target record.
	var todo model.Todo
	result := repo.db.First(&todo, id)
	//check the result of process.
	switch{
	case result.Error == nil:
		return &todo, nil
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, ErrNotFound
	default:
		return nil, result.Error
	}
}

func (repo *postgres_repo) Get_records_by_category(category string) ([]model.Todo, error){
	//define a slice for records.
	records := make([]model.Todo, 0)
	//get all records of the target category.
	result := repo.db.Where("category = ?", category).Find(&records)
	//check the result of process.
	if result.Error != nil{
		//return empty slice and error value.
		return nil, result.Error
	}
	return records, nil
}

func (repo *postgres_repo) Get_records_by_status(status bool) ([]model.Todo, error){
	//define a slice for the records.
	records := make([]model.Todo, 0)
	//get all records that has the required status.
	result := repo.db.Where("completed = ?", status).Find(&records)
	//check the result of the process.
	if result.Error != nil{
		//return empty slice and error value.
		return nil, result.Error
	}
	return records, nil
}

func (repo *postgres_repo) Get_records_by_title(title string) ([]model.Todo, error){
	//define a slice for records.
	records := make([]model.Todo, 0)
	//get all records that match the required title.
	result := repo.db.Where("title ILIKE ?", "%"+title+"%").Find(&records)
	//check the result of the process.
	if result.Error != nil{
		return nil, result.Error
	}
	return records, nil
}
