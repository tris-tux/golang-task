package db

import (
	"github.com/tris-tux/go-task/backend/schema"
)

type Static struct{}

func (s *Static) GetAll() ([]schema.Task, error) {
	taskList := []schema.Task{
		{
			TaskId:   1,
			Title: "AAA",
			AcctionTime: 1652461200,
			CreateTime: 1652461200,
			UpdateTime: 1652461200,
			IdFinished: false,
		},
		{
			TaskId:   2,
			Title: "BBB",
			AcctionTime: 1652462100,
			CreateTime: 1652462100,
			UpdateTime: 1652462100,
			IdFinished: false,
		},
		{
			TaskId:   3,
			Title: "CCC",
			AcctionTime: 1652661200,
			CreateTime: 1652661200,
			UpdateTime: 1652661200,
			IdFinished: false,
		},
	}
	return taskList, nil
}

func (s *Static) Insert(task *schema.Task) (int, error) {
	return 0, nil
}

func (s *Static) Update(task *schema.Task) error {
	return nil
}

func (s *Static) Delete(id int) error {
	return nil
}

func (s *Static) Close() {}
