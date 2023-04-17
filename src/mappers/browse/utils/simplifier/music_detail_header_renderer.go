package simplifier

import (
	"fmt"
	"katze/src/mappers/browse/models"
	"katze/src/mappers/utils/columns"
	"katze/src/mappers/utils/subtitle"
	"katze/src/models/external"
)

func MusicDetailHeaderRenderer(header external.MusicDetailHeaderRenderer) (
	models.MusicDetailHeaderRenderer, error,
) {

	titleRuns := header.Title.Runs
	if len(titleRuns) == 0 {
		return models.MusicDetailHeaderRenderer{}, fmt.Errorf(
			"title runs is empty %v", titleRuns,
		)
	}

	subtitleFlex, err := subtitle.ToFlexColumns(header.Subtitle)
	if err != nil {
		return models.MusicDetailHeaderRenderer{}, err
	}
	secondSubtitleFlex, err := subtitle.ToFlexColumns(header.SecondSubtitle)
	if err != nil {
		return models.MusicDetailHeaderRenderer{}, err
	}

	duration := ""
	if secondSubLen := len(secondSubtitleFlex); secondSubLen != 0 {
		duration = secondSubtitleFlex[secondSubLen-1].Text
	}

	artists, err := columns.GetArtists(subtitleFlex)
	if err != nil {
		return models.MusicDetailHeaderRenderer{}, err
	}
	year, err := columns.GetYear(subtitleFlex)
	if err != nil {
		return models.MusicDetailHeaderRenderer{}, err
	}
	title := titleRuns[0].Text

	return models.MusicDetailHeaderRenderer{
		Thumbnails: header.Thumbnail.CroppedSquareThumbnailRenderer.Thumbnail.Thumbnails,
 
		Title:    title,
		Artists:  artists,
		Year:     year,
		Duration: duration,
	}, nil

}
