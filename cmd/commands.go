package cmd

import "clitodo/pkg/domain"

type TaskAdded struct {
	IsSucces bool
	Item     domain.Item
}

type AddTaskTrigger bool
