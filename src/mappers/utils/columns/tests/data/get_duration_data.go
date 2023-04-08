package data

import "katze/src/models"

var FLEX_COLUMN_VALID_DURATION = []models.FlexColumn{
	{
		Text: "example text",
	},
	{
		Text: "example text",
	},
	{
		Text: "20:15:01",
	},
}

var FLEX_COLUMN_INVALID_DURATION = []models.FlexColumn{
	{
		Text: "example text",
	},
	{
		Text: "example text",
	},
	{
		Text: "60:10",
	},
}

var FLEX_COLUMN_INVALID_DURATION_2 = []models.FlexColumn{
	{
		Text: "example text",
	},
	{
		Text: "example text",
	},
	{
		Text: "32",
	},
}

var FLEX_COLUMN_DURATION_EMPTY = []models.FlexColumn{}
