package simplifier

import (
	"fmt"
	"katze/src/mappers/models"
	"katze/src/mappers/utils"
	"katze/src/mappers/utils/columns"
	"katze/src/models/external"
)

// MusicResponsiveListItemRenderer simplifies an external.MischievousContent 
// to a models.MusicResponsiveListItemRenderer
func MusicResponsiveListItemRenderer(
	musicResponsiveListItemRenderer external.MischievousContent,
) (
	models.MusicResponsiveListItemRenderer, error,
) {
	item := musicResponsiveListItemRenderer.MusicResponsiveListItemRenderer
	// Get FlexColumns
	flexColumns, err := columns.Mapper(item.FlexColumns)
	if err != nil {
		return models.MusicResponsiveListItemRenderer{}, err
	}
	// Get Id
	id := ""
	pageType := ""
	// Get id and pageType if it is a music video
	if item.PlaylistItemData != nil {
		id = item.PlaylistItemData.VideoID
		pageType = "MUSIC_TYPE"
	}

	// Get id and pageType if it is a other type
	if item.NavigationEndpoint != nil {
		browseEndpoint := item.NavigationEndpoint.BrowseEndpoint
		pageType = browseEndpoint.BrowseEndpointContextSupportedConfigs.
			BrowseEndpointContextMusicConfig.PageType
		id = browseEndpoint.BrowseID
	}

	// Valid id and pageType
	if id == "" {
		err := fmt.Errorf("error: id is empty")
		return models.MusicResponsiveListItemRenderer{}, err
	}

	if pageType == "" {
		err := fmt.Errorf("error: pageType is empty")
		return models.MusicResponsiveListItemRenderer{}, err
	}

	thumbnails := utils.GetThumbnail(item.Thumbnail)
	explicit := item.Badges != nil
	return models.MusicResponsiveListItemRenderer{
		ID:          id,
		Thumbnails:  thumbnails,
		FlexColumns: flexColumns,
		PageType:    pageType,
		Explicit:    explicit,
	}, nil
}
