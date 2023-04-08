package data

import "katze/src/models"

var FLEX_COLUMN_VALID_YEAR = []models.FlexColumn{
	{
		Text: "example text",
	},
	{
		Text: "example text",
	},
	{
		Text: "2022",
	},
}

var FLEX_COLUMN_INVALID_YEAR = []models.FlexColumn{
	{
		Text: "example text",
	},
	{
		Text: "example text",
	},
	{
		Text: "202",
	},
}

var FLEX_COLUMN_INVALID_YEAR_2 = []models.FlexColumn{
	{
		Text: "example text",
	},
	{
		Text: "example text",
	},
	{
		Text: "20233",
	},
}

var FLEX_COLUMN_EMPTY = []models.FlexColumn{}
