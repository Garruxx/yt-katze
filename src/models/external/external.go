package external

type ArtistTracklistPagination struct {
	ResponseContext      *ArtistTracklistPaginationResponseContext      `json:"responseContext,omitempty"`
	ContinuationContents *ArtistTracklistPaginationContinuationContents `json:"continuationContents,omitempty"`
	TrackingParams       *string                                        `json:"trackingParams,omitempty"`
	Error                *ArtistTracklistPaginationError                `json:"error,omitempty"`
}

type SubtitleBadges []SubtitleBadge

type SubtitleBadge struct {
	MusicInlineBadgeRenderer MusicInlineBadgeRenderer `json:"musicInlineBadgeRenderer"`
}

type ArtistTracklistPaginationContinuationContents struct {
	MusicPlaylistShelfContinuation MusicPlaylistShelfContinuation `json:"musicPlaylistShelfContinuation"`
}

type MusicPlaylistShelfContinuation struct {
	Contents                []MusicPlaylistShelfContinuationContent `json:"contents"`
	ShelfDivider            ShelfDivider                            `json:"shelfDivider"`
	TrackingParams          string                                  `json:"trackingParams"`
	Continuations           []Continuation                          `json:"continuations,omitempty"`
	ContentsMultiSelectable bool                                    `json:"contentsMultiSelectable"`
}

type MusicPlaylistShelfContinuationContent struct {
	MusicResponsiveListItemRenderer PurpleMusicResponsiveListItemRenderer `json:"musicResponsiveListItemRenderer"`
}

type PurpleMusicResponsiveListItemRenderer struct {
	TrackingParams      string                 `json:"trackingParams"`
	Thumbnail           ThumbnailRendererClass `json:"thumbnail"`
	Overlay             PurpleOverlay          `json:"overlay"`
	FlexColumns         []FluffyFlexColumn     `json:"flexColumns"`
	FixedColumns        []FixedColumn          `json:"fixedColumns"`
	Menu                PurpleMenu             `json:"menu"`
	PlaylistItemData    PurplePlaylistItemData `json:"playlistItemData"`
	MultiSelectCheckbox MultiSelectCheckbox    `json:"multiSelectCheckbox"`
	ItemHeight          string                 `json:"itemHeight"`
	Index               TitleElement           `json:"index"`
	Badges              *[]Badge               `json:"badges,omitempty"`
}

type Badge struct {
	MusicInlineBadgeRenderer MusicInlineBadgeRenderer `json:"musicInlineBadgeRenderer"`
}

type MusicInlineBadgeRenderer struct {
	TrackingParams    string                      `json:"trackingParams"`
	Icon              Icon                        `json:"icon"`
	AccessibilityData SubscribeAccessibilityClass `json:"accessibilityData"`
}

type SubscribeAccessibilityClass struct {
	AccessibilityData AccessibilityAccessibility `json:"accessibilityData"`
}

type AccessibilityAccessibility struct {
	Label string `json:"label"`
}

type Icon struct {
	IconType string `json:"iconType"`
}

type FixedColumn struct {
	MusicResponsiveListItemFixedColumnRenderer FluffyMusicResponsiveListItemFlexColumnRenderer `json:"musicResponsiveListItemFixedColumnRenderer"`
}

type TitleElement struct {
	Runs []DescriptionRun `json:"runs"`
}

type DescriptionRun struct {
	Text string `json:"text"`
}

type PurpleRun struct {
	Text               string                    `json:"text"`
	NavigationEndpoint *PurpleNavigationEndpoint `json:"navigationEndpoint,omitempty"`
}

type PurpleNavigationEndpoint struct {
	ClickTrackingParams string               `json:"clickTrackingParams"`
	WatchEndpoint       *PurpleWatchEndpoint `json:"watchEndpoint,omitempty"`
	BrowseEndpoint      *OnTapBrowseEndpoint `json:"browseEndpoint,omitempty"`
}

type OnTapBrowseEndpoint struct {
	BrowseID                              string                                `json:"browseId"`
	BrowseEndpointContextSupportedConfigs BrowseEndpointContextSupportedConfigs `json:"browseEndpointContextSupportedConfigs"`
}

type BrowseEndpointContextSupportedConfigs struct {
	BrowseEndpointContextMusicConfig BrowseEndpointContextMusicConfig `json:"browseEndpointContextMusicConfig"`
}

type BrowseEndpointContextMusicConfig struct {
	PageType string `json:"pageType"`
}

type PurpleWatchEndpoint struct {
	VideoID                            string                              `json:"videoId"`
	PlaylistID                         string                              `json:"playlistId"`
	LoggingContext                     LoggingContext                      `json:"loggingContext"`
	WatchEndpointMusicSupportedConfigs *WatchEndpointMusicSupportedConfigs `json:"watchEndpointMusicSupportedConfigs,omitempty"`
	Params                             *string                             `json:"params,omitempty"`
	PlaylistSetVideoID                 *string                             `json:"playlistSetVideoId,omitempty"`
	Index                              *int64                              `json:"index,omitempty"`
}

type LoggingContext struct {
	VssLoggingContext VssLoggingContext `json:"vssLoggingContext"`
}

type VssLoggingContext struct {
	SerializedContextData string `json:"serializedContextData"`
}

type WatchEndpointMusicSupportedConfigs struct {
	WatchEndpointMusicConfig WatchEndpointMusicConfig `json:"watchEndpointMusicConfig"`
}

type WatchEndpointMusicConfig struct {
	MusicVideoType string `json:"musicVideoType"`
}

type PurpleMenu struct {
	MenuRenderer PurpleMenuRenderer `json:"menuRenderer"`
}

type PurpleMenuRenderer struct {
	Items           []PurpleItem                `json:"items"`
	TrackingParams  string                      `json:"trackingParams"`
	TopLevelButtons []PurpleTopLevelButton      `json:"topLevelButtons"`
	Accessibility   SubscribeAccessibilityClass `json:"accessibility"`
}

type PurpleItem struct {
	MenuNavigationItemRenderer *MenuItemRenderer `json:"menuNavigationItemRenderer,omitempty"`
	MenuServiceItemRenderer    *MenuItemRenderer `json:"menuServiceItemRenderer,omitempty"`
}

type MenuItemRenderer struct {
	Text               TitleElement                                  `json:"text"`
	Icon               Icon                                          `json:"icon"`
	NavigationEndpoint *MenuNavigationItemRendererNavigationEndpoint `json:"navigationEndpoint,omitempty"`
	TrackingParams     string                                        `json:"trackingParams"`
	ServiceEndpoint    *MenuNavigationItemRendererServiceEndpoint    `json:"serviceEndpoint,omitempty"`
}

type MenuNavigationItemRendererNavigationEndpoint struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	WatchEndpoint         *PurpleWatchEndpoint   `json:"watchEndpoint,omitempty"`
	ModalEndpoint         *ModalEndpoint         `json:"modalEndpoint,omitempty"`
	BrowseEndpoint        *OnTapBrowseEndpoint   `json:"browseEndpoint,omitempty"`
	ShareEntityEndpoint   *ShareEntityEndpoint   `json:"shareEntityEndpoint,omitempty"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint `json:"watchPlaylistEndpoint,omitempty"`
}

type ModalEndpoint struct {
	Modal Modal `json:"modal"`
}

type Modal struct {
	ModalWithTitleAndButtonRenderer ModalWithTitleAndButtonRenderer `json:"modalWithTitleAndButtonRenderer"`
}

type ModalWithTitleAndButtonRenderer struct {
	Title   TitleElement                          `json:"title"`
	Content TitleElement                          `json:"content"`
	Button  ModalWithTitleAndButtonRendererButton `json:"button"`
}

type ModalWithTitleAndButtonRendererButton struct {
	ButtonRenderer PurpleButtonRenderer `json:"buttonRenderer"`
}

type PurpleButtonRenderer struct {
	Style              string                   `json:"style"`
	IsDisabled         bool                     `json:"isDisabled"`
	Text               TitleElement             `json:"text"`
	NavigationEndpoint FluffyNavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams     string                   `json:"trackingParams"`
}

type FluffyNavigationEndpoint struct {
	ClickTrackingParams string         `json:"clickTrackingParams"`
	SignInEndpoint      SignInEndpoint `json:"signInEndpoint"`
}

type SignInEndpoint struct {
	Hack bool `json:"hack"`
}

type ShareEntityEndpoint struct {
	SerializedShareEntity string `json:"serializedShareEntity"`
	SharePanelType        string `json:"sharePanelType"`
}

type WatchPlaylistEndpoint struct {
	PlaylistID string  `json:"playlistId"`
	Params     *string `json:"params,omitempty"`
}

type MenuNavigationItemRendererServiceEndpoint struct {
	ClickTrackingParams string           `json:"clickTrackingParams"`
	QueueAddEndpoint    QueueAddEndpoint `json:"queueAddEndpoint"`
}

type QueueAddEndpoint struct {
	QueueTarget         QueueTarget      `json:"queueTarget"`
	QueueInsertPosition string           `json:"queueInsertPosition"`
	Commands            []CommandElement `json:"commands"`
}

type CommandElement struct {
	ClickTrackingParams string           `json:"clickTrackingParams"`
	AddToToastAction    AddToToastAction `json:"addToToastAction"`
}

type AddToToastAction struct {
	Item AddToToastActionItem `json:"item"`
}

type AddToToastActionItem struct {
	NotificationTextRenderer NotificationTextRenderer `json:"notificationTextRenderer"`
}

type NotificationTextRenderer struct {
	SuccessResponseText TitleElement `json:"successResponseText"`
	TrackingParams      string       `json:"trackingParams"`
}

type QueueTarget struct {
	VideoID    *string `json:"videoId,omitempty"`
	PlaylistID *string `json:"playlistId,omitempty"`
}

type PurpleTopLevelButton struct {
	LikeButtonRenderer LikeButtonRenderer `json:"likeButtonRenderer"`
}

type LikeButtonRenderer struct {
	Target                    PlaylistItemData               `json:"target"`
	LikeStatus                string                         `json:"likeStatus"`
	TrackingParams            string                         `json:"trackingParams"`
	LikesAllowed              bool                           `json:"likesAllowed"`
	DislikeNavigationEndpoint DefaultNavigationEndpointClass `json:"dislikeNavigationEndpoint"`
	LikeCommand               DefaultNavigationEndpointClass `json:"likeCommand"`
}

type DefaultNavigationEndpointClass struct {
	ClickTrackingParams string        `json:"clickTrackingParams"`
	ModalEndpoint       ModalEndpoint `json:"modalEndpoint"`
}

type PlaylistItemData struct {
	VideoID string `json:"videoId"`
}

type MultiSelectCheckbox struct {
	CheckboxRenderer CheckboxRenderer `json:"checkboxRenderer"`
}

type CheckboxRenderer struct {
	OnSelectionChangeCommand OnSelectionChangeCommand `json:"onSelectionChangeCommand"`
	CheckedState             string                   `json:"checkedState"`
	TrackingParams           string                   `json:"trackingParams"`
}

type OnSelectionChangeCommand struct {
	ClickTrackingParams           string                        `json:"clickTrackingParams"`
	UpdateMultiSelectStateCommand UpdateMultiSelectStateCommand `json:"updateMultiSelectStateCommand"`
}

type UpdateMultiSelectStateCommand struct {
	MultiSelectParams string `json:"multiSelectParams"`
	MultiSelectItem   string `json:"multiSelectItem"`
}

type PurpleOverlay struct {
	MusicItemThumbnailOverlayRenderer PurpleMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type PurpleMusicItemThumbnailOverlayRenderer struct {
	Background      Background    `json:"background"`
	Content         PurpleContent `json:"content"`
	ContentPosition string        `json:"contentPosition"`
	DisplayStyle    string        `json:"displayStyle"`
}

type Background struct {
	VerticalGradient VerticalGradient `json:"verticalGradient"`
}

type VerticalGradient struct {
	GradientLayerColors []string `json:"gradientLayerColors"`
}

type PurpleContent struct {
	MusicPlayButtonRenderer PurpleMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type PurpleMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint NavigationEndpoint          `json:"playNavigationEndpoint"`
	TrackingParams         string                      `json:"trackingParams"`
	PlayIcon               Icon                        `json:"playIcon"`
	PauseIcon              Icon                        `json:"pauseIcon"`
	IconColor              int64                       `json:"iconColor"`
	BackgroundColor        int64                       `json:"backgroundColor"`
	ActiveBackgroundColor  int64                       `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                       `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                        `json:"playingIcon"`
	IconLoadingColor       int64                       `json:"iconLoadingColor"`
	ActiveScaleFactor      int64                       `json:"activeScaleFactor"`
	ButtonSize             string                      `json:"buttonSize"`
	RippleTarget           string                      `json:"rippleTarget"`
	AccessibilityPlayData  SubscribeAccessibilityClass `json:"accessibilityPlayData"`
	AccessibilityPauseData SubscribeAccessibilityClass `json:"accessibilityPauseData"`
}

type NavigationEndpoint struct {
	ClickTrackingParams string              `json:"clickTrackingParams"`
	WatchEndpoint       PurpleWatchEndpoint `json:"watchEndpoint"`
}

type PurplePlaylistItemData struct {
	PlaylistSetVideoID string `json:"playlistSetVideoId"`
	VideoID            string `json:"videoId"`
}

type ThumbnailRendererClass struct {
	MusicThumbnailRenderer MusicThumbnailRenderer `json:"musicThumbnailRenderer"`
}

type MusicThumbnailRenderer struct {
	Thumbnail      MusicThumbnailRendererThumbnail `json:"thumbnail"`
	ThumbnailCrop  string                          `json:"thumbnailCrop"`
	ThumbnailScale string                          `json:"thumbnailScale"`
	TrackingParams string                          `json:"trackingParams"`
}

type MusicThumbnailRendererThumbnail struct {
	Thumbnails []ThumbnailElement `json:"thumbnails"`
}

type ThumbnailElement struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type ArtistTracklistPaginationError struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Errors  []ErrorElement `json:"errors"`
	Status  string         `json:"status"`
}

type ErrorElement struct {
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
}

type ArtistTracklistPaginationResponseContext struct {
	VisitorData           string                 `json:"visitorData"`
	ServiceTrackingParams []ServiceTrackingParam `json:"serviceTrackingParams"`
}

type ServiceTrackingParam struct {
	Service string  `json:"service"`
	Params  []Param `json:"params"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ArtistTracklist struct {
	ResponseContext *ArtistTracklistPaginationResponseContext `json:"responseContext,omitempty"`
	Contents        *ArtistTracklistContents                  `json:"contents,omitempty"`
	Header          *ArtistTracklistHeader                    `json:"header,omitempty"`
	TrackingParams  *string                                   `json:"trackingParams,omitempty"`
	Error           *ArtistTracklistPaginationError           `json:"error,omitempty"`
}

type ArtistTracklistContents struct {
	SingleColumnBrowseResultsRenderer PurpleSingleColumnBrowseResultsRenderer `json:"singleColumnBrowseResultsRenderer"`
}

type PurpleSingleColumnBrowseResultsRenderer struct {
	Tabs []PurpleTab `json:"tabs"`
}

type PurpleTab struct {
	TabRenderer PurpleTabRenderer `json:"tabRenderer"`
}

type PurpleTabRenderer struct {
	Content        FluffyContent `json:"content"`
	TrackingParams string        `json:"trackingParams"`
}

type FluffyContent struct {
	SectionListRenderer PurpleSectionListRenderer `json:"sectionListRenderer"`
}

type PurpleSectionListRenderer struct {
	Contents       []TentacledContent `json:"contents"`
	TrackingParams string             `json:"trackingParams"`
}

type TentacledContent struct {
	MusicPlaylistShelfRenderer MusicPlaylistShelfRenderer `json:"musicPlaylistShelfRenderer"`
}

type MusicPlaylistShelfRenderer struct {
	PlaylistID              string                                  `json:"playlistId"`
	Contents                []MusicPlaylistShelfContinuationContent `json:"contents"`
	CollapsedItemCount      int64                                   `json:"collapsedItemCount"`
	TrackingParams          string                                  `json:"trackingParams"`
	Continuations           []Continuation                          `json:"continuations"`
	ContentsMultiSelectable bool                                    `json:"contentsMultiSelectable"`
}

type Continuation struct {
	NextContinuationData NextContinuationData `json:"nextContinuationData"`
}

type NextContinuationData struct {
	Continuation        string `json:"continuation"`
	ClickTrackingParams string `json:"clickTrackingParams"`
}

type ArtistTracklistHeader struct {
	MusicHeaderRenderer MusicRenderer `json:"musicHeaderRenderer"`
}

type MusicRenderer struct {
	Title          TitleElement `json:"title"`
	TrackingParams string       `json:"trackingParams"`
}

type MusicPagination ItemsPagination
type ArtistPagination ItemsPagination
type AlbumsPagination ItemsPagination
type ItemsPagination struct {
	ResponseContext      *TracksPaginationResponseContext      `json:"responseContext,omitempty"`
	ContinuationContents *TracksPaginationContinuationContents `json:"continuationContents,omitempty"`
	TrackingParams       *string                               `json:"trackingParams,omitempty"`
	Header               *TracksPaginationHeader               `json:"header,omitempty"`
	Contents             *TracksPaginationContents             `json:"contents,omitempty"`
	Error                *ArtistTracklistPaginationError       `json:"error,omitempty"`
}

type TracksPaginationContents struct {
	TabbedSearchResultsRenderer PurpleTabbedSearchResultsRenderer `json:"tabbedSearchResultsRenderer"`
}

type PurpleTabbedSearchResultsRenderer struct {
	Tabs []FluffyTab `json:"tabs"`
}

type FluffyTab struct {
	TabRenderer FluffyTabRenderer `json:"tabRenderer"`
}

type FluffyTabRenderer struct {
	Title          string        `json:"title"`
	Selected       bool          `json:"selected"`
	Content        StickyContent `json:"content"`
	TabIdentifier  string        `json:"tabIdentifier"`
	TrackingParams string        `json:"trackingParams"`
}

type StickyContent struct {
	SectionListRenderer FluffySectionListRenderer `json:"sectionListRenderer"`
}

type FluffySectionListRenderer struct {
	Contents       []IndigoContent           `json:"contents"`
	TrackingParams string                    `json:"trackingParams"`
	Header         MusicHeaderRendererHeader `json:"header"`
}

type IndigoContent struct {
	MusicShelfRenderer PurpleMusicShelfRenderer `json:"musicShelfRenderer"`
}

type PurpleMusicShelfRenderer struct {
	Title          TitleElement         `json:"title"`
	Contents       []MischievousContent `json:"contents"`
	TrackingParams string               `json:"trackingParams"`
	Continuations  []Continuation       `json:"continuations"`
	ShelfDivider   ShelfDivider         `json:"shelfDivider"`
}

type FluffyFlexColumn struct {
	MusicResponsiveListItemFlexColumnRenderer FluffyMusicResponsiveListItemFlexColumnRenderer `json:"musicResponsiveListItemFlexColumnRenderer"`
}

type FluffyMusicResponsiveListItemFlexColumnRenderer struct {
	Text            TitleClass `json:"text"`
	DisplayPriority string     `json:"displayPriority"`
	Size            string     `json:"size"`
}

type TitleClass struct {
	Runs []FluffyRun `json:"runs"`
}

type FluffyRun struct {
	Text               string `json:"text"`
	NavigationEndpoint *OnTap `json:"navigationEndpoint,omitempty"`
}

type OnTap struct {
	ClickTrackingParams string               `json:"clickTrackingParams"`
	WatchEndpoint       *OnTapWatchEndpoint  `json:"watchEndpoint,omitempty"`
	BrowseEndpoint      *OnTapBrowseEndpoint `json:"browseEndpoint,omitempty"`
}

type OnTapWatchEndpoint struct {
	VideoID                            string                             `json:"videoId"`
	WatchEndpointMusicSupportedConfigs WatchEndpointMusicSupportedConfigs `json:"watchEndpointMusicSupportedConfigs"`
}

type MusicCardShelfRendererMenu struct {
	MenuRenderer FluffyMenuRenderer `json:"menuRenderer"`
}

type FluffyMenuRenderer struct {
	Items          []FluffyItem                `json:"items"`
	TrackingParams string                      `json:"trackingParams"`
	Accessibility  SubscribeAccessibilityClass `json:"accessibility"`
}

type FluffyItem struct {
	MenuNavigationItemRenderer    *MenuItemRenderer              `json:"menuNavigationItemRenderer,omitempty"`
	MenuServiceItemRenderer       *MenuItemRenderer              `json:"menuServiceItemRenderer,omitempty"`
	ToggleMenuServiceItemRenderer *ToggleMenuServiceItemRenderer `json:"toggleMenuServiceItemRenderer,omitempty"`
}

type ToggleMenuServiceItemRenderer struct {
	DefaultText            TitleElement                   `json:"defaultText"`
	DefaultIcon            Icon                           `json:"defaultIcon"`
	DefaultServiceEndpoint DefaultNavigationEndpointClass `json:"defaultServiceEndpoint"`
	ToggledText            TitleElement                   `json:"toggledText"`
	ToggledIcon            Icon                           `json:"toggledIcon"`
	TrackingParams         string                         `json:"trackingParams"`
	ToggledServiceEndpoint *ToggledServiceEndpoint        `json:"toggledServiceEndpoint,omitempty"`
}

type ToggledServiceEndpoint struct {
	ClickTrackingParams string       `json:"clickTrackingParams"`
	LikeEndpoint        LikeEndpoint `json:"likeEndpoint"`
}

type LikeEndpoint struct {
	Status string      `json:"status"`
	Target TargetClass `json:"target"`
}

type TargetClass struct {
	PlaylistID string `json:"playlistId"`
}

type ThumbnailOverlayClass struct {
	MusicItemThumbnailOverlayRenderer FluffyMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type FluffyMusicItemThumbnailOverlayRenderer struct {
	Background      Background      `json:"background"`
	Content         IndecentContent `json:"content"`
	ContentPosition string          `json:"contentPosition"`
	DisplayStyle    string          `json:"displayStyle"`
}

type IndecentContent struct {
	MusicPlayButtonRenderer FluffyMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type FluffyMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint PurplePlayNavigationEndpoint `json:"playNavigationEndpoint"`
	TrackingParams         string                       `json:"trackingParams"`
	PlayIcon               Icon                         `json:"playIcon"`
	PauseIcon              Icon                         `json:"pauseIcon"`
	IconColor              int64                        `json:"iconColor"`
	BackgroundColor        int64                        `json:"backgroundColor"`
	ActiveBackgroundColor  int64                        `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                        `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                         `json:"playingIcon"`
	IconLoadingColor       int64                        `json:"iconLoadingColor"`
	ActiveScaleFactor      int64                        `json:"activeScaleFactor"`
	ButtonSize             string                       `json:"buttonSize"`
	RippleTarget           string                       `json:"rippleTarget"`
	AccessibilityPlayData  SubscribeAccessibilityClass  `json:"accessibilityPlayData"`
	AccessibilityPauseData SubscribeAccessibilityClass  `json:"accessibilityPauseData"`
}

type PurplePlayNavigationEndpoint struct {
	ClickTrackingParams string             `json:"clickTrackingParams"`
	WatchEndpoint       OnTapWatchEndpoint `json:"watchEndpoint"`
}

type ShelfDivider struct {
	MusicShelfDividerRenderer MusicShelfDividerRenderer `json:"musicShelfDividerRenderer"`
}

type MusicShelfDividerRenderer struct {
	Hidden bool `json:"hidden"`
}

type MusicHeaderRendererHeader struct {
	ChipCloudRenderer PurpleChipCloudRenderer `json:"chipCloudRenderer"`
}

type PurpleChipCloudRenderer struct {
	Chips                []PurpleChip `json:"chips"`
	CollapsedRowCount    int64        `json:"collapsedRowCount"`
	TrackingParams       string       `json:"trackingParams"`
	HorizontalScrollable bool         `json:"horizontalScrollable"`
}

type PurpleChip struct {
	ChipCloudChipRenderer PurpleChipCloudChipRenderer `json:"chipCloudChipRenderer"`
}

type PurpleChipCloudChipRenderer struct {
	Style              Style                                   `json:"style"`
	NavigationEndpoint ChipCloudChipRendererNavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams     string                                  `json:"trackingParams"`
	Icon               *Icon                                   `json:"icon,omitempty"`
	AccessibilityData  SubscribeAccessibilityClass             `json:"accessibilityData"`
	IsSelected         bool                                    `json:"isSelected"`
	Text               *TitleElement                           `json:"text,omitempty"`
	UniqueID           *string                                 `json:"uniqueId,omitempty"`
}

type ChipCloudChipRendererNavigationEndpoint struct {
	ClickTrackingParams string         `json:"clickTrackingParams"`
	SearchEndpoint      SearchEndpoint `json:"searchEndpoint"`
}

type SearchEndpoint struct {
	Query  string  `json:"query"`
	Params *string `json:"params,omitempty"`
}

type Style struct {
	StyleType string `json:"styleType"`
}

type TracksPaginationContinuationContents struct {
	MusicShelfContinuation MusicShelf `json:"musicShelfContinuation"`
}

type MusicShelf FluffyMusicShelfRenderer

type TracksPaginationHeader struct {
	MusicHeaderRenderer MusicHeaderRenderer `json:"musicHeaderRenderer"`
}

type MusicHeaderRenderer struct {
	Header         MusicHeaderRendererHeader `json:"header"`
	TrackingParams string                    `json:"trackingParams"`
}

type TracksPaginationResponseContext struct {
	VisitorData           string                 `json:"visitorData"`
	ServiceTrackingParams []ServiceTrackingParam `json:"serviceTrackingParams"`
	MaxAgeSeconds         int64                  `json:"maxAgeSeconds"`
}

type ArtistList MusicList
type AlbumList MusicList
type MusicList ItemsList
type ItemsList struct {
	ResponseContext *TracksPaginationResponseContext `json:"responseContext,omitempty"`
	Contents        *TracklistContents               `json:"contents,omitempty"`
	TrackingParams  *string                          `json:"trackingParams,omitempty"`
	Error           *ArtistTracklistPaginationError  `json:"error,omitempty"`
}
type ArtistTcracklist struct {
	ResponseContext *ArtistTracklistPaginationResponseContext `json:"responseContext,omitempty"`
	Contents        *ArtistTracklistContents                  `json:"contents,omitempty"`
	Header          *ArtistTracklistHeader                    `json:"header,omitempty"`
	TrackingParams  *string                                   `json:"trackingParams,omitempty"`
	Error           *ArtistTracklistPaginationError           `json:"error,omitempty"`
}

type TracklistContents struct {
	TabbedSearchResultsRenderer FluffyTabbedSearchResultsRenderer `json:"tabbedSearchResultsRenderer"`
}

type FluffyTabbedSearchResultsRenderer struct {
	Tabs []TentacledTab `json:"tabs"`
}

type TentacledTab struct {
	TabRenderer TentacledTabRenderer `json:"tabRenderer"`
}

type TentacledTabRenderer struct {
	Title          string           `json:"title"`
	Selected       bool             `json:"selected"`
	Content        HilariousContent `json:"content"`
	TabIdentifier  string           `json:"tabIdentifier"`
	TrackingParams string           `json:"trackingParams"`
}

type HilariousContent struct {
	SectionListRenderer TentacledSectionListRenderer `json:"sectionListRenderer"`
}

type TentacledSectionListRenderer struct {
	Contents       []AmbitiousContent        `json:"contents"`
	TrackingParams string                    `json:"trackingParams"`
	Header         MusicHeaderRendererHeader `json:"header"`
}

type AmbitiousContent struct {
	MusicShelfRenderer MusicShelf `json:"musicShelfRenderer"`
}

type General struct {
	ResponseContext *TracksPaginationResponseContext `json:"responseContext,omitempty"`
	Contents        *SearchResultContents            `json:"contents,omitempty"`
	TrackingParams  *string                          `json:"trackingParams,omitempty"`
	Error           *ArtistTracklistPaginationError  `json:"error,omitempty"`
}

type SearchResultContents struct {
	TabbedSearchResultsRenderer TentacledTabbedSearchResultsRenderer `json:"tabbedSearchResultsRenderer"`
}

type TentacledTabbedSearchResultsRenderer struct {
	Tabs []StickyTab `json:"tabs"`
}

type StickyTab struct {
	TabRenderer StickyTabRenderer `json:"tabRenderer"`
}

type StickyTabRenderer struct {
	Title          string         `json:"title"`
	Selected       bool           `json:"selected"`
	Content        CunningContent `json:"content"`
	TabIdentifier  string         `json:"tabIdentifier"`
	TrackingParams string         `json:"trackingParams"`
}

type CunningContent struct {
	SectionListRenderer StickySectionListRenderer `json:"sectionListRenderer"`
}

type StickySectionListRenderer struct {
	Contents       []MagentaContent `json:"contents"`
	TrackingParams string           `json:"trackingParams"`
	Header         PurpleHeader     `json:"header"`
}

type MagentaContent struct {
	MusicCardShelfRenderer *MusicCardShelfRenderer   `json:"musicCardShelfRenderer,omitempty"`
	MusicShelfRenderer     *FluffyMusicShelfRenderer `json:"musicShelfRenderer,omitempty"`
}

type MusicCardShelfRenderer struct {
	TrackingParams   string                          `json:"trackingParams"`
	Thumbnail        ThumbnailRendererClass          `json:"thumbnail"`
	Title            TitleClass                      `json:"title"`
	Subtitle         Subtitle                        `json:"subtitle"`
	Buttons          []PlayButtonElement             `json:"buttons"`
	Menu             MusicCardShelfRendererMenu      `json:"menu"`
	OnTap            OnTap                           `json:"onTap"`
	Header           MusicCardShelfRendererHeader    `json:"header"`
	ThumbnailOverlay *ThumbnailOverlayClass          `json:"thumbnailOverlay,omitempty"`
	SubtitleBadges   SubtitleBadges                  `json:"subtitleBadges"`
	Contents         []MusicCardShelfRendererContent `json:"contents,omitempty"`
	EndIcon          *Icon                           `json:"endIcon,omitempty"`
}

type PlayButtonElement struct {
	ButtonRenderer PlayButtonButtonRenderer `json:"buttonRenderer"`
}

type PlayButtonButtonRenderer struct {
	Style              string                       `json:"style"`
	Size               *string                      `json:"size,omitempty"`
	IsDisabled         *bool                        `json:"isDisabled,omitempty"`
	Text               TitleElement                 `json:"text"`
	Icon               Icon                         `json:"icon"`
	Accessibility      AccessibilityAccessibility   `json:"accessibility"`
	TrackingParams     string                       `json:"trackingParams"`
	AccessibilityData  *SubscribeAccessibilityClass `json:"accessibilityData,omitempty"`
	Command            *ButtonRendererCommand       `json:"command,omitempty"`
	NavigationEndpoint *TentacledNavigationEndpoint `json:"navigationEndpoint,omitempty"`
}

type ButtonRendererCommand struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	WatchEndpoint         *CommandWatchEndpoint  `json:"watchEndpoint,omitempty"`
	ModalEndpoint         *ModalEndpoint         `json:"modalEndpoint,omitempty"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint `json:"watchPlaylistEndpoint,omitempty"`
	ShareEntityEndpoint   *ShareEntityEndpoint   `json:"shareEntityEndpoint,omitempty"`
}

type CommandWatchEndpoint struct {
	VideoID                            string                             `json:"videoId"`
	Params                             string                             `json:"params"`
	WatchEndpointMusicSupportedConfigs WatchEndpointMusicSupportedConfigs `json:"watchEndpointMusicSupportedConfigs"`
}

type TentacledNavigationEndpoint struct {
	ClickTrackingParams   string               `json:"clickTrackingParams"`
	WatchEndpoint         *PurpleWatchEndpoint `json:"watchEndpoint,omitempty"`
	WatchPlaylistEndpoint *TargetClass         `json:"watchPlaylistEndpoint,omitempty"`
}

type MusicCardShelfRendererContent struct {
	MusicResponsiveListItemRenderer TentacledMusicResponsiveListItemRenderer `json:"musicResponsiveListItemRenderer"`
}

type TentacledMusicResponsiveListItemRenderer struct {
	TrackingParams         string                                             `json:"trackingParams"`
	Thumbnail              ThumbnailRendererClass                             `json:"thumbnail"`
	Overlay                FluffyOverlay                                      `json:"overlay"`
	FlexColumns            []FluffyFlexColumn                                 `json:"flexColumns"`
	Menu                   MusicCardShelfRendererMenu                         `json:"menu"`
	PlaylistItemData       *PlaylistItemData                                  `json:"playlistItemData,omitempty"`
	FlexColumnDisplayStyle string                                             `json:"flexColumnDisplayStyle"`
	ItemHeight             string                                             `json:"itemHeight"`
	Badges                 []Badge                                            `json:"badges,omitempty"`
	NavigationEndpoint     *MusicResponsiveListItemRendererNavigationEndpoint `json:"navigationEndpoint,omitempty"`
}

type MusicResponsiveListItemRendererNavigationEndpoint struct {
	ClickTrackingParams string              `json:"clickTrackingParams"`
	BrowseEndpoint      OnTapBrowseEndpoint `json:"browseEndpoint"`
}

type FluffyOverlay struct {
	MusicItemThumbnailOverlayRenderer TentacledMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type TentacledMusicItemThumbnailOverlayRenderer struct {
	Background      Background    `json:"background"`
	Content         FriskyContent `json:"content"`
	ContentPosition string        `json:"contentPosition"`
	DisplayStyle    string        `json:"displayStyle"`
}

type FriskyContent struct {
	MusicPlayButtonRenderer TentacledMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type TentacledMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint FluffyPlayNavigationEndpoint `json:"playNavigationEndpoint"`
	TrackingParams         string                       `json:"trackingParams"`
	PlayIcon               Icon                         `json:"playIcon"`
	PauseIcon              Icon                         `json:"pauseIcon"`
	IconColor              int64                        `json:"iconColor"`
	BackgroundColor        int64                        `json:"backgroundColor"`
	ActiveBackgroundColor  int64                        `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                        `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                         `json:"playingIcon"`
	IconLoadingColor       int64                        `json:"iconLoadingColor"`
	ActiveScaleFactor      int64                        `json:"activeScaleFactor"`
	ButtonSize             string                       `json:"buttonSize"`
	RippleTarget           string                       `json:"rippleTarget"`
	AccessibilityPlayData  SubscribeAccessibilityClass  `json:"accessibilityPlayData"`
	AccessibilityPauseData SubscribeAccessibilityClass  `json:"accessibilityPauseData"`
}

type FluffyPlayNavigationEndpoint struct {
	ClickTrackingParams   string              `json:"clickTrackingParams"`
	WatchEndpoint         *OnTapWatchEndpoint `json:"watchEndpoint,omitempty"`
	WatchPlaylistEndpoint *TargetClass        `json:"watchPlaylistEndpoint,omitempty"`
}

type MusicCardShelfRendererHeader struct {
	MusicCardShelfHeaderBasicRenderer MusicRenderer `json:"musicCardShelfHeaderBasicRenderer"`
}

type Subtitle struct {
	Runs []TentacledRun `json:"runs"`
}

type TentacledRun struct {
	Text               string                                             `json:"text"`
	NavigationEndpoint *MusicResponsiveListItemRendererNavigationEndpoint `json:"navigationEndpoint,omitempty"`
}

type FluffyMusicShelfRenderer struct {
	Title               TitleElement         `json:"title"`
	Contents            []MischievousContent `json:"contents"`
	TrackingParams      string               `json:"trackingParams"`
	ShelfDivider        ShelfDivider         `json:"shelfDivider"`
	Continuations       []Continuation       `json:"continuations"`
	AutoReloadWhenEmpty *bool                `json:"autoReloadWhenEmpty,omitempty"`

	BottomText     TitleElement                            `json:"bottomText"`
	BottomEndpoint ChipCloudChipRendererNavigationEndpoint `json:"bottomEndpoint"`
}

type MischievousContent struct {
	MusicResponsiveListItemRenderer StickyMusicResponsiveListItemRenderer `json:"musicResponsiveListItemRenderer"`
}

type StickyMusicResponsiveListItemRenderer struct {
	TrackingParams         string                                             `json:"trackingParams"`
	Thumbnail              ThumbnailRendererClass                             `json:"thumbnail"`
	Overlay                *TentacledOverlay                                  `json:"overlay,omitempty"`
	FlexColumns            []FluffyFlexColumn                                 `json:"flexColumns"`
	Menu                   MusicCardShelfRendererMenu                         `json:"menu"`
	PlaylistItemData       *PlaylistItemData                                  `json:"playlistItemData,omitempty"`
	FlexColumnDisplayStyle string                                             `json:"flexColumnDisplayStyle"`
	ItemHeight             string                                             `json:"itemHeight"`
	NavigationEndpoint     *MusicResponsiveListItemRendererNavigationEndpoint `json:"navigationEndpoint,omitempty"`
	Badges                 *[]Badge                                           `json:"badges,omitempty"`
}

type TentacledOverlay struct {
	MusicItemThumbnailOverlayRenderer StickyMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type StickyMusicItemThumbnailOverlayRenderer struct {
	Background      Background           `json:"background"`
	Content         BraggadociousContent `json:"content"`
	ContentPosition string               `json:"contentPosition"`
	DisplayStyle    string               `json:"displayStyle"`
}

type BraggadociousContent struct {
	MusicPlayButtonRenderer StickyMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type StickyMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint TentacledPlayNavigationEndpoint `json:"playNavigationEndpoint"`
	TrackingParams         string                          `json:"trackingParams"`
	PlayIcon               Icon                            `json:"playIcon"`
	PauseIcon              Icon                            `json:"pauseIcon"`
	IconColor              int64                           `json:"iconColor"`
	BackgroundColor        int64                           `json:"backgroundColor"`
	ActiveBackgroundColor  int64                           `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                           `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                            `json:"playingIcon"`
	IconLoadingColor       int64                           `json:"iconLoadingColor"`
	ActiveScaleFactor      int64                           `json:"activeScaleFactor"`
	ButtonSize             string                          `json:"buttonSize"`
	RippleTarget           string                          `json:"rippleTarget"`
	AccessibilityPlayData  SubscribeAccessibilityClass     `json:"accessibilityPlayData"`
	AccessibilityPauseData SubscribeAccessibilityClass     `json:"accessibilityPauseData"`
}

type TentacledPlayNavigationEndpoint struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	WatchEndpoint         *OnTapWatchEndpoint    `json:"watchEndpoint,omitempty"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint `json:"watchPlaylistEndpoint,omitempty"`
}

type PurpleHeader struct {
	ChipCloudRenderer FluffyChipCloudRenderer `json:"chipCloudRenderer"`
}

type FluffyChipCloudRenderer struct {
	Chips                []FluffyChip `json:"chips"`
	CollapsedRowCount    int64        `json:"collapsedRowCount"`
	TrackingParams       string       `json:"trackingParams"`
	HorizontalScrollable bool         `json:"horizontalScrollable"`
}

type FluffyChip struct {
	ChipCloudChipRenderer FluffyChipCloudChipRenderer `json:"chipCloudChipRenderer"`
}

type FluffyChipCloudChipRenderer struct {
	Style              Style                                   `json:"style"`
	Text               TitleElement                            `json:"text"`
	NavigationEndpoint ChipCloudChipRendererNavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams     string                                  `json:"trackingParams"`
	AccessibilityData  SubscribeAccessibilityClass             `json:"accessibilityData"`
	IsSelected         bool                                    `json:"isSelected"`
	UniqueID           string                                  `json:"uniqueId"`
}

type Artist struct {
	ResponseContext *TracksPaginationResponseContext `json:"responseContext,omitempty"`
	Contents        *ArtistContents                  `json:"contents,omitempty"`
	Header          *ArtistHeader                    `json:"header,omitempty"`
	TrackingParams  *string                          `json:"trackingParams,omitempty"`
	Error           *ArtistTracklistPaginationError  `json:"error,omitempty"`
}

type ArtistContents struct {
	SingleColumnBrowseResultsRenderer FluffySingleColumnBrowseResultsRenderer `json:"singleColumnBrowseResultsRenderer"`
}

type FluffySingleColumnBrowseResultsRenderer struct {
	Tabs []IndigoTab `json:"tabs"`
}

type IndigoTab struct {
	TabRenderer IndigoTabRenderer `json:"tabRenderer"`
}

type IndigoTabRenderer struct {
	Content        Content1 `json:"content"`
	TrackingParams string   `json:"trackingParams"`
}

type Content1 struct {
	SectionListRenderer IndigoSectionListRenderer `json:"sectionListRenderer"`
}

type IndigoSectionListRenderer struct {
	Contents       []Content2 `json:"contents"`
	TrackingParams string     `json:"trackingParams"`
}

type Content2 struct {
	MusicShelfRenderer            *TentacledMusicShelfRenderer   `json:"musicShelfRenderer,omitempty"`
	MusicCarouselShelfRenderer    *MusicCarouselShelfRenderer    `json:"musicCarouselShelfRenderer,omitempty"`
	MusicDescriptionShelfRenderer *MusicDescriptionShelfRenderer `json:"musicDescriptionShelfRenderer,omitempty"`
}

type MusicCarouselShelfRenderer struct {
	Header         MusicCarouselShelfRendererHeader    `json:"header"`
	Contents       []MusicCarouselShelfRendererContent `json:"contents"`
	TrackingParams string                              `json:"trackingParams"`
	ItemSize       string                              `json:"itemSize"`
}

type MusicCarouselShelfRendererContent struct {
	MusicTwoRowItemRenderer ContentMusicTwoRowItemRenderer `json:"musicTwoRowItemRenderer"`
}

type ContentMusicTwoRowItemRenderer struct {
	ThumbnailRenderer  ThumbnailRendererClass                     `json:"thumbnailRenderer"`
	AspectRatio        string                                     `json:"aspectRatio"`
	Title              MusicCarouselShelfBasicHeaderRendererTitle `json:"title"`
	Subtitle           Subtitle                                   `json:"subtitle"`
	NavigationEndpoint MusicTwoRowItemRendererNavigationEndpoint  `json:"navigationEndpoint"`
	TrackingParams     string                                     `json:"trackingParams"`
	Menu               MusicCardShelfRendererMenu                 `json:"menu"`
	ThumbnailOverlay   *PurpleThumbnailOverlay                    `json:"thumbnailOverlay,omitempty"`
	SubtitleBadges     *[]Badge                                   `json:"subtitleBadges,omitempty"`
}

type MusicTwoRowItemRendererNavigationEndpoint struct {
	ClickTrackingParams string                        `json:"clickTrackingParams"`
	BrowseEndpoint      *BottomEndpointBrowseEndpoint `json:"browseEndpoint,omitempty"`
	WatchEndpoint       *PurpleWatchEndpoint          `json:"watchEndpoint,omitempty"`
}

type BottomEndpointBrowseEndpoint struct {
	BrowseID                              string                                `json:"browseId"`
	Params                                *string                               `json:"params,omitempty"`
	BrowseEndpointContextSupportedConfigs BrowseEndpointContextSupportedConfigs `json:"browseEndpointContextSupportedConfigs"`
}

type PurpleThumbnailOverlay struct {
	MusicItemThumbnailOverlayRenderer IndigoMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type IndigoMusicItemThumbnailOverlayRenderer struct {
	Background      Background `json:"background"`
	Content         Content3   `json:"content"`
	ContentPosition string     `json:"contentPosition"`
	DisplayStyle    string     `json:"displayStyle"`
}

type Content3 struct {
	MusicPlayButtonRenderer IndigoMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type IndigoMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint StickyPlayNavigationEndpoint `json:"playNavigationEndpoint"`
	TrackingParams         string                       `json:"trackingParams"`
	PlayIcon               Icon                         `json:"playIcon"`
	PauseIcon              Icon                         `json:"pauseIcon"`
	IconColor              int64                        `json:"iconColor"`
	BackgroundColor        int64                        `json:"backgroundColor"`
	ActiveBackgroundColor  int64                        `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                        `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                         `json:"playingIcon"`
	IconLoadingColor       int64                        `json:"iconLoadingColor"`
	ActiveScaleFactor      float64                      `json:"activeScaleFactor"`
	ButtonSize             string                       `json:"buttonSize"`
	RippleTarget           string                       `json:"rippleTarget"`
	AccessibilityPlayData  SubscribeAccessibilityClass  `json:"accessibilityPlayData"`
	AccessibilityPauseData SubscribeAccessibilityClass  `json:"accessibilityPauseData"`
}

type StickyPlayNavigationEndpoint struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint `json:"watchPlaylistEndpoint,omitempty"`
	WatchEndpoint         *PurpleWatchEndpoint   `json:"watchEndpoint,omitempty"`
}

type MusicCarouselShelfBasicHeaderRendererTitle struct {
	Runs []StickyRun `json:"runs"`
}

type StickyRun struct {
	Text               string             `json:"text"`
	NavigationEndpoint *RunBottomEndpoint `json:"navigationEndpoint,omitempty"`
}

type RunBottomEndpoint struct {
	ClickTrackingParams string                       `json:"clickTrackingParams"`
	BrowseEndpoint      BottomEndpointBrowseEndpoint `json:"browseEndpoint"`
}

type MusicCarouselShelfRendererHeader struct {
	MusicCarouselShelfBasicHeaderRenderer MusicCarouselShelfBasicHeaderRenderer `json:"musicCarouselShelfBasicHeaderRenderer"`
}

type MusicCarouselShelfBasicHeaderRenderer struct {
	Title             MusicCarouselShelfBasicHeaderRendererTitle `json:"title"`
	AccessibilityData SubscribeAccessibilityClass                `json:"accessibilityData"`
	HeaderStyle       string                                     `json:"headerStyle"`
	TrackingParams    string                                     `json:"trackingParams"`
	MoreContentButton *MoreContentButton                         `json:"moreContentButton,omitempty"`
}

type MoreContentButton struct {
	ButtonRenderer MoreContentButtonButtonRenderer `json:"buttonRenderer"`
}

type MoreContentButtonButtonRenderer struct {
	Style              string                      `json:"style"`
	Text               TitleElement                `json:"text"`
	NavigationEndpoint RunBottomEndpoint           `json:"navigationEndpoint"`
	TrackingParams     string                      `json:"trackingParams"`
	AccessibilityData  SubscribeAccessibilityClass `json:"accessibilityData"`
}

type MusicDescriptionShelfRenderer struct {
	Header         TitleElement `json:"header"`
	Subheader      TitleElement `json:"subheader"`
	Description    TitleElement `json:"description"`
	MoreButton     MoreButton   `json:"moreButton"`
	TrackingParams string       `json:"trackingParams"`
}

type MoreButton struct {
	ToggleButtonRenderer MoreButtonToggleButtonRenderer `json:"toggleButtonRenderer"`
}

type MoreButtonToggleButtonRenderer struct {
	IsToggled      bool         `json:"isToggled"`
	IsDisabled     bool         `json:"isDisabled"`
	DefaultIcon    Icon         `json:"defaultIcon"`
	DefaultText    TitleElement `json:"defaultText"`
	ToggledIcon    Icon         `json:"toggledIcon"`
	ToggledText    TitleElement `json:"toggledText"`
	TrackingParams string       `json:"trackingParams"`
}

type TentacledMusicShelfRenderer struct {
	Title          MusicCarouselShelfBasicHeaderRendererTitle `json:"title"`
	Contents       []MischievousContent                       `json:"contents"`
	TrackingParams string                                     `json:"trackingParams"`
	BottomText     TitleElement                               `json:"bottomText"`
	BottomEndpoint RunBottomEndpoint                          `json:"bottomEndpoint"`
	ShelfDivider   ShelfDivider                               `json:"shelfDivider"`
}

type ArtistHeader struct {
	MusicImmersiveHeaderRenderer MusicImmersiveHeaderRenderer `json:"musicImmersiveHeaderRenderer"`
}

type MusicImmersiveHeaderRenderer struct {
	Title              TitleElement                     `json:"title"`
	SubscriptionButton SubscriptionButton               `json:"subscriptionButton"`
	Description        TitleElement                     `json:"description"`
	MoreButton         MoreButton                       `json:"moreButton"`
	Menu               MusicImmersiveHeaderRendererMenu `json:"menu"`
	Thumbnail          ThumbnailRendererClass           `json:"thumbnail"`
	TrackingParams     string                           `json:"trackingParams"`
	PlayButton         PlayButtonElement                `json:"playButton"`
	StartRadioButton   StartRadioButton                 `json:"startRadioButton"`
	ShareEndpoint      ShareEndpoint                    `json:"shareEndpoint"`
}

type MusicImmersiveHeaderRendererMenu struct {
	MenuRenderer TentacledMenuRenderer `json:"menuRenderer"`
}

type TentacledMenuRenderer struct {
	Items          []TentacledItem             `json:"items"`
	TrackingParams string                      `json:"trackingParams"`
	Accessibility  SubscribeAccessibilityClass `json:"accessibility"`
}

type TentacledItem struct {
	MenuNavigationItemRenderer MenuItemRenderer `json:"menuNavigationItemRenderer"`
}

type ShareEndpoint struct {
	ClickTrackingParams string              `json:"clickTrackingParams"`
	ShareEntityEndpoint ShareEntityEndpoint `json:"shareEntityEndpoint"`
}

type StartRadioButton struct {
	ButtonRenderer StartRadioButtonButtonRenderer `json:"buttonRenderer"`
}

type StartRadioButtonButtonRenderer struct {
	Text               TitleElement               `json:"text"`
	Icon               Icon                       `json:"icon"`
	NavigationEndpoint NavigationEndpoint         `json:"navigationEndpoint"`
	Accessibility      AccessibilityAccessibility `json:"accessibility"`
	TrackingParams     string                     `json:"trackingParams"`
}

type SubscriptionButton struct {
	SubscribeButtonRenderer SubscribeButtonRenderer `json:"subscribeButtonRenderer"`
}

type SubscribeButtonRenderer struct {
	SubscriberCountText              TitleElement                   `json:"subscriberCountText"`
	Subscribed                       bool                           `json:"subscribed"`
	Enabled                          bool                           `json:"enabled"`
	Type                             string                         `json:"type"`
	ChannelID                        string                         `json:"channelId"`
	ShowPreferences                  bool                           `json:"showPreferences"`
	SubscriberCountWithSubscribeText TitleElement                   `json:"subscriberCountWithSubscribeText"`
	SubscribedButtonText             TitleElement                   `json:"subscribedButtonText"`
	UnsubscribedButtonText           TitleElement                   `json:"unsubscribedButtonText"`
	TrackingParams                   string                         `json:"trackingParams"`
	UnsubscribeButtonText            TitleElement                   `json:"unsubscribeButtonText"`
	ServiceEndpoints                 []ServiceEndpointElement       `json:"serviceEndpoints"`
	LongSubscriberCountText          Text                           `json:"longSubscriberCountText"`
	ShortSubscriberCountText         TitleElement                   `json:"shortSubscriberCountText"`
	SubscribeAccessibility           SubscribeAccessibilityClass    `json:"subscribeAccessibility"`
	UnsubscribeAccessibility         SubscribeAccessibilityClass    `json:"unsubscribeAccessibility"`
	SignInEndpoint                   DefaultNavigationEndpointClass `json:"signInEndpoint"`
}

type Text struct {
	Runs          []DescriptionRun            `json:"runs"`
	Accessibility SubscribeAccessibilityClass `json:"accessibility"`
}

type ServiceEndpointElement struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	SubscribeEndpoint     *SubscribeEndpoint     `json:"subscribeEndpoint,omitempty"`
	SignalServiceEndpoint *SignalServiceEndpoint `json:"signalServiceEndpoint,omitempty"`
}

type SignalServiceEndpoint struct {
	Signal  string   `json:"signal"`
	Actions []Action `json:"actions"`
}

type Action struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	OpenPopupAction     OpenPopupAction `json:"openPopupAction"`
}

type OpenPopupAction struct {
	Popup     Popup  `json:"popup"`
	PopupType string `json:"popupType"`
}

type Popup struct {
	ConfirmDialogRenderer ConfirmDialogRenderer `json:"confirmDialogRenderer"`
}

type ConfirmDialogRenderer struct {
	TrackingParams string         `json:"trackingParams"`
	DialogMessages []TitleElement `json:"dialogMessages"`
	ConfirmButton  Button         `json:"confirmButton"`
	CancelButton   Button         `json:"cancelButton"`
}

type Button struct {
	ButtonRenderer CancelButtonButtonRenderer `json:"buttonRenderer"`
}

type CancelButtonButtonRenderer struct {
	Style           string                         `json:"style"`
	Size            string                         `json:"size"`
	Text            TitleElement                   `json:"text"`
	TrackingParams  string                         `json:"trackingParams"`
	ServiceEndpoint *ButtonRendererServiceEndpoint `json:"serviceEndpoint,omitempty"`
}

type ButtonRendererServiceEndpoint struct {
	ClickTrackingParams string              `json:"clickTrackingParams"`
	UnsubscribeEndpoint UnsubscribeEndpoint `json:"unsubscribeEndpoint"`
}

type UnsubscribeEndpoint struct {
	ChannelIDS []string `json:"channelIds"`
}

type SubscribeEndpoint struct {
	ChannelIDS []string `json:"channelIds"`
	Params     string   `json:"params"`
}

type Tracklist struct {
	ResponseContext *ArtistTracklistPaginationResponseContext `json:"responseContext,omitempty"`
	Contents        *AlbumContents                            `json:"contents,omitempty"`
	Header          *AlbumHeader                              `json:"header,omitempty"`
	TrackingParams  *string                                   `json:"trackingParams,omitempty"`
	Microformat     *Microformat                              `json:"microformat,omitempty"`
	Error           *ArtistTracklistPaginationError           `json:"error,omitempty"`
}

type AlbumContents struct {
	SingleColumnBrowseResultsRenderer TentacledSingleColumnBrowseResultsRenderer `json:"singleColumnBrowseResultsRenderer"`
}

type TentacledSingleColumnBrowseResultsRenderer struct {
	Tabs []IndecentTab `json:"tabs"`
}

type IndecentTab struct {
	TabRenderer IndecentTabRenderer `json:"tabRenderer"`
}

type IndecentTabRenderer struct {
	Content        Content5 `json:"content"`
	TrackingParams string   `json:"trackingParams"`
}

type Content5 struct {
	SectionListRenderer IndecentSectionListRenderer `json:"sectionListRenderer"`
}

type IndecentSectionListRenderer struct {
	Contents       []Content6 `json:"contents"`
	TrackingParams string     `json:"trackingParams"`
}

type Content6 struct {
	MusicShelfRenderer StickyMusicShelfRenderer `json:"musicShelfRenderer"`
}

type StickyMusicShelfRenderer struct {
	Contents                []MusicPlaylistShelfContinuationContent `json:"contents"`
	TrackingParams          string                                  `json:"trackingParams"`
	ShelfDivider            ShelfDivider                            `json:"shelfDivider"`
	ContentsMultiSelectable bool                                    `json:"contentsMultiSelectable"`
}

// type Content7 struct {
// 	MusicResponsiveListItemRenderer IndecentMusicResponsiveListItemRenderer `json:"musicResponsiveListItemRenderer"`
// }

// type IndecentMusicResponsiveListItemRenderer struct {
// 	TrackingParams      string                 `json:"trackingParams"`
// 	Overlay             PurpleOverlay          `json:"overlay"`
// 	FlexColumns         []TentacledFlexColumn  `json:"flexColumns"`
// 	FixedColumns        []FixedColumn          `json:"fixedColumns"`
// 	Menu                PurpleMenu             `json:"menu"`
// 	PlaylistItemData    PurplePlaylistItemData `json:"playlistItemData"`
// 	ItemHeight          string                 `json:"itemHeight"`
// 	Index               TitleElement           `json:"index"`
// 	MultiSelectCheckbox MultiSelectCheckbox    `json:"multiSelectCheckbox"`
// }

// type TentacledFlexColumn struct {
// 	MusicResponsiveListItemFlexColumnRenderer TentacledMusicResponsiveListItemFlexColumnRenderer `json:"musicResponsiveListItemFlexColumnRenderer"`
// }

// type TentacledMusicResponsiveListItemFlexColumnRenderer struct {
// 	Text            FluffyText `json:"text"`
// 	DisplayPriority string     `json:"displayPriority"`
// }

type FluffyText struct {
	Runs []IndigoRun `json:"runs,omitempty"`
}

type IndigoRun struct {
	Text               string             `json:"text"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
}

type AlbumHeader struct {
	MusicDetailHeaderRenderer MusicDetailHeaderRenderer `json:"musicDetailHeaderRenderer"`
}

type MusicDetailHeaderRenderer struct {
	Title          TitleElement                       `json:"title"`
	Subtitle       Subtitle                           `json:"subtitle"`
	Menu           MusicDetailHeaderRendererMenu      `json:"menu"`
	Thumbnail      MusicDetailHeaderRendererThumbnail `json:"thumbnail"`
	TrackingParams string                             `json:"trackingParams"`
	MoreButton     MoreButton                         `json:"moreButton"`
	SecondSubtitle Subtitle                           `json:"secondSubtitle"`
}

type MusicDetailHeaderRendererMenu struct {
	MenuRenderer StickyMenuRenderer `json:"menuRenderer"`
}

type StickyMenuRenderer struct {
	Items           []PurpleItem                `json:"items"`
	TrackingParams  string                      `json:"trackingParams"`
	TopLevelButtons []FluffyTopLevelButton      `json:"topLevelButtons"`
	Accessibility   SubscribeAccessibilityClass `json:"accessibility"`
}

type FluffyTopLevelButton struct {
	ButtonRenderer       *PlayButtonButtonRenderer           `json:"buttonRenderer,omitempty"`
	ToggleButtonRenderer *TopLevelButtonToggleButtonRenderer `json:"toggleButtonRenderer,omitempty"`
}

type TopLevelButtonToggleButtonRenderer struct {
	IsToggled                 bool                           `json:"isToggled"`
	IsDisabled                bool                           `json:"isDisabled"`
	DefaultIcon               Icon                           `json:"defaultIcon"`
	DefaultText               Text                           `json:"defaultText"`
	ToggledIcon               Icon                           `json:"toggledIcon"`
	ToggledText               Text                           `json:"toggledText"`
	TrackingParams            string                         `json:"trackingParams"`
	DefaultNavigationEndpoint DefaultNavigationEndpointClass `json:"defaultNavigationEndpoint"`
}

type MusicDetailHeaderRendererThumbnail struct {
	CroppedSquareThumbnailRenderer CroppedSquareThumbnailRenderer `json:"croppedSquareThumbnailRenderer"`
}

type CroppedSquareThumbnailRenderer struct {
	Thumbnail      MusicThumbnailRendererThumbnail `json:"thumbnail"`
	TrackingParams string                          `json:"trackingParams"`
}

type Microformat struct {
	MicroformatDataRenderer MicroformatDataRenderer `json:"microformatDataRenderer"`
}

type MicroformatDataRenderer struct {
	URLCanonical string `json:"urlCanonical"`
}

type ArtistAlbums ArtistTwoRowItem
type ArtistSingles ArtistTwoRowItem
type ArtistTwoRowItem struct {
	ResponseContext *ArtistAlbumsResponseContext    `json:"responseContext,omitempty"`
	Contents        *ArtistAlbumsContents           `json:"contents,omitempty"`
	Header          *ArtistTracklistHeader          `json:"header,omitempty"`
	TrackingParams  *string                         `json:"trackingParams,omitempty"`
	Error           *ArtistTracklistPaginationError `json:"error,omitempty"`
}

type ArtistAlbumsContents struct {
	SingleColumnBrowseResultsRenderer StickySingleColumnBrowseResultsRenderer `json:"singleColumnBrowseResultsRenderer"`
}

type StickySingleColumnBrowseResultsRenderer struct {
	Tabs []HilariousTab `json:"tabs"`
}

type HilariousTab struct {
	TabRenderer HilariousTabRenderer `json:"tabRenderer"`
}

type HilariousTabRenderer struct {
	Content        Content8 `json:"content"`
	TrackingParams string   `json:"trackingParams"`
}

type Content8 struct {
	SectionListRenderer HilariousSectionListRenderer `json:"sectionListRenderer"`
}

type HilariousSectionListRenderer struct {
	Contents       []Content9 `json:"contents"`
	TrackingParams string     `json:"trackingParams"`
}

type Content9 struct {
	GridRenderer GridRenderer `json:"gridRenderer"`
}

type GridRenderer struct {
	Items          []MusicCarouselShelfRendererContent `json:"items"`
	TrackingParams string                              `json:"trackingParams"`
	Header         GridRendererHeader                  `json:"header"`
}

type GridRendererHeader struct {
	GridHeaderRenderer GridHeaderRenderer `json:"gridHeaderRenderer"`
}

type GridHeaderRenderer struct {
	Title TitleElement `json:"title"`
}

type FluffyThumbnailOverlay struct {
	MusicItemThumbnailOverlayRenderer IndecentMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type IndecentMusicItemThumbnailOverlayRenderer struct {
	Background      Background `json:"background"`
	Content         Content10  `json:"content"`
	ContentPosition string     `json:"contentPosition"`
	DisplayStyle    string     `json:"displayStyle"`
}

type Content10 struct {
	MusicPlayButtonRenderer IndecentMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type IndecentMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint IndigoPlayNavigationEndpoint `json:"playNavigationEndpoint"`
	TrackingParams         string                       `json:"trackingParams"`
	PlayIcon               Icon                         `json:"playIcon"`
	PauseIcon              Icon                         `json:"pauseIcon"`
	IconColor              int64                        `json:"iconColor"`
	BackgroundColor        int64                        `json:"backgroundColor"`
	ActiveBackgroundColor  int64                        `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                        `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                         `json:"playingIcon"`
	IconLoadingColor       int64                        `json:"iconLoadingColor"`
	ActiveScaleFactor      float64                      `json:"activeScaleFactor"`
	ButtonSize             string                       `json:"buttonSize"`
	RippleTarget           string                       `json:"rippleTarget"`
	AccessibilityPlayData  SubscribeAccessibilityClass  `json:"accessibilityPlayData"`
	AccessibilityPauseData SubscribeAccessibilityClass  `json:"accessibilityPauseData"`
}

type IndigoPlayNavigationEndpoint struct {
	ClickTrackingParams   string      `json:"clickTrackingParams"`
	WatchPlaylistEndpoint TargetClass `json:"watchPlaylistEndpoint"`
}

type ArtistAlbumsResponseContext struct {
	ServiceTrackingParams []ServiceTrackingParam `json:"serviceTrackingParams"`
	MaxAgeSeconds         int64                  `json:"maxAgeSeconds"`
	VisitorData           string                 `json:"visitorData"`
}
