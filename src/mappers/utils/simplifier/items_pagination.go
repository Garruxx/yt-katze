package simplifier

import (
	"fmt"
	"katze/src/mappers/models"
	"katze/src/models/external"
)

// ItemsPagination simplifies an external.ItemsPagination to a models.ItemList
func ItemsPagination(itemsPagination external.ItemsPagination) (
	itemsList models.ItemList, err error,
) {

	if itemsPagination.Error != nil {
		return itemsList, fmt.Errorf("error: %s", itemsPagination.Error.Message)
	}

	contents := itemsPagination.ContinuationContents
	if contents == nil {
		return itemsList, fmt.Errorf("error: itemsPagination.Contents is empty")
	}

	responseContext := itemsPagination.ResponseContext
	if responseContext == nil {
		return itemsList, fmt.Errorf("error: ResponseContext is empty")
	}

	shelfContinuation := contents.MusicShelfContinuation
	shelfContents := shelfContinuation.Contents
	if len(shelfContents) == 0 {
		return itemsList, nil
	}

	continuations := shelfContinuation.Continuations
	continuationID := ""
	if len(continuations) >= 1 {
		continuationID = continuations[0].NextContinuationData.Continuation
	}

	itemsList = models.ItemList{
		Contents:  shelfContents,
		VisitorID: responseContext.VisitorData,
		NextPage: continuationID,
	}

	return itemsList, nil
}
