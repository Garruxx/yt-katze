package models

import "katze/src/models"

type FlexColumnRender struct {
	Top    []models.FlexColumn `json:"top"`
	Bottom []models.FlexColumn `json:"bottom"`
}
