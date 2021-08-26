package gardens_models

import (
	completed_tasks_models "github.com/ice-cream-backend/models/v1/completed_tasks"
	rules_models "github.com/ice-cream-backend/models/v1/rules"
)

type Rules struct {
	ID             string `bson:"_id" json:"_id"`
	Name           string `bson:"name" json:"name"`
	Description    string `bson:"description" json:"description"`
	UserFireBaseId string `bson:"userFireBaseId" json:"userFireBaseId"`
}

type GardensFullyPopulated struct {
	Garden Rules `json:"garden"`
	Rules []rules_models.Rules `json:"rules"`
	CompletedTasks []completed_tasks_models.CompletedTasks `json:"completedTasks"`
}