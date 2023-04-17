package simplifier

import (
	"fmt"
	"katze/src/mappers/models"
	"katze/src/models/external"
)

// ItemsList simplifies an external.ItemsList to a models.ItemList
func ItemsList(itemList external.ItemsList) (
	result models.ItemList, err error,
) {

	if itemList.Error != nil {
		return result, fmt.Errorf("error: %s", itemList.Error.Message)
	}

	contents := itemList.Contents
	if contents == nil {
		return result, fmt.Errorf("error: itemList.Contents is empty")
	}

	responseContext := itemList.ResponseContext
	if responseContext == nil {
		return result, fmt.Errorf("error: ResponseContext is empty")
	}

	tabs := contents.TabbedSearchResultsRenderer.Tabs
	if len(tabs) == 0 {
		return result, fmt.Errorf("tabs is empty")
	}

	tabContents := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(tabContents) <= 0 {
		return result, fmt.Errorf("tabContents is empty")
	}

	shelfRenderer := tabContents[0].MusicShelfRenderer
	shelfContents := shelfRenderer.Contents
	if len(shelfContents) == 0 {
		return result, nil
	}

	continuation := shelfRenderer.Continuations
	continuationID := ""
	if len(continuation) >= 1 {
		continuationID = continuation[0].NextContinuationData.Continuation
	}

	result = models.ItemList{
		Contents:  shelfContents,
		VisitorID: responseContext.VisitorData,
		NextPage:  continuationID,
	}

	return result, nil
}
