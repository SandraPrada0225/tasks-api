package dto

//DTO representa los datos que llegan a la api
import (
	"tasks-api/internal/utils"
)

type CreateTaskDTO struct {
	Title string `json:"title"`
}

func (d *CreateTaskDTO) Validate() error {

	ve := &utils.ValidationError{}
	if d.Title == "" {
		ve.Add("title", "el titulo es obligatorio")
	}

	if len(d.Title) < 3 {
		ve.Add("title", "minimo 3 caracteres")
	}

	if len(d.Title) > 100 {
		ve.Add("title", "maximo 100 caracteres")
	}

	if ve.HasErrors() {
		return ve
	}

	return nil
}
