package models

import "katze/src/models/external"

type ItemList struct {
	Contents  []external.MischievousContent
	NextPage  string
	VisitorID string
}