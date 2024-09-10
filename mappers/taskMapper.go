package mappers

import (
	"todo/dto"
	"todo/types"
)

func fromTaskToDto(task *types.Task) *dto.TaskDto {
	return &dto.TaskDto{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      dto.Status(task.Status),
		Deleted:     task.Deleted,
	}
}

func FromTasksToDto(tasks []*types.Task) []*dto.TaskDto {
	var taskDtos []*dto.TaskDto
	for _, task := range tasks {
		var taskDto *dto.TaskDto = fromTaskToDto(task)
		taskDtos = append(taskDtos, taskDto)
	}
	return taskDtos
}
