package serializer

import "go_memo/model"

type Task struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"吃饭"`
	Content   string `json:"content" example:"睡觉"`
	View      uint64 `json:"view" example:"32"`
	Status    int    `json:"status" example:"0"`
	CreatedAt int64  `json:"createdAt"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreatedAt: item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
