package models

import "katze/src/models"

type FlexaColumnRender struct {
	Top    []models.FlexColumn `json:"top"`
	Bottom []models.FlexColumn `json:"bottom"`
}
