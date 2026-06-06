package uploadpost

import (
	"context"
	"fmt"
	"strconv"
)

// UploadVideo uploads a video file or URL to one or more social media platforms.
// videoPathOrURL can be a local file path (e.g. "./video.mp4") or an HTTPS URL.
func (c *Client) UploadVideo(ctx context.Context, videoPathOrURL string, opts VideoOptions) (*UploadResponse, error) {
	if opts.User == "" {
		return nil, fmt.Errorf("uploadpost: VideoOptions.User is required")
	}
	if len(opts.Platforms) == 0 {
		return nil, fmt.Errorf("uploadpost: VideoOptions.Platforms must not be empty")
	}

	fb := newFormBuilder()

	if err := fb.setFile("video", videoPathOrURL); err != nil {
		return nil, err
	}

	addCommonParams(fb, commonParams{
		User:           opts.User,
		Title:          opts.Title,
		Platforms:      opts.Platforms,
		FirstComment:   opts.FirstComment,
		AltText:        opts.AltText,
		ScheduledDate:  opts.ScheduledDate,
		Timezone:       opts.Timezone,
		AddToQueue:     opts.AddToQueue,
		MaxPostsPerSlot: opts.MaxPostsPerSlot,
		AsyncUpload:    opts.AsyncUpload,
		FirstCommentMedia: opts.FirstCommentMedia,

		BlueskyTitle:   opts.BlueskyTitle,
		InstagramTitle: opts.InstagramTitle,
		FacebookTitle:  opts.FacebookTitle,
		TiktokTitle:    opts.TiktokTitle,
		LinkedinTitle:  opts.LinkedinTitle,
		XTitle:         opts.XTitle,
		YoutubeTitle:   opts.YoutubeTitle,
		PinterestTitle: opts.PinterestTitle,
		ThreadsTitle:   opts.ThreadsTitle,

		Description:          opts.Description,
		LinkedinDescription:  opts.LinkedinDescription,
		YoutubeDescription:   opts.YoutubeDescription,
		FacebookDescription:  opts.FacebookDescription,
		TiktokDescription:    opts.TiktokDescription,
		PinterestDescription: opts.PinterestDescription,

		InstagramFirstComment: opts.InstagramFirstComment,
		FacebookFirstComment:  opts.FacebookFirstComment,
		XFirstComment:         opts.XFirstComment,
		ThreadsFirstComment:   opts.ThreadsFirstComment,
		YoutubeFirstComment:   opts.YoutubeFirstComment,
		RedditFirstComment:    opts.RedditFirstComment,
		BlueskyFirstComment:   opts.BlueskyFirstComment,
		LinkedinFirstComment:  opts.LinkedinFirstComment,
	})

	for _, p := range opts.Platforms {
		switch p {
		case "tiktok":
			addTiktokVideoParams(fb, opts)
		case "instagram":
			addInstagramVideoParams(fb, opts)
		case "youtube":
			if err := addYoutubeParams(fb, opts); err != nil {
				return nil, err
			}
		case "linkedin":
			fb.setIfNotEmpty("visibility", opts.LinkedinVisibility)
			fb.setIfNotEmpty("target_linkedin_page_id", opts.TargetLinkedinPageID)
		case "facebook":
			fb.setIfNotEmpty("facebook_page_id", opts.FacebookPageID)
			fb.setIfNotEmpty("video_state", opts.FacebookVideoState)
			fb.setIfNotEmpty("facebook_media_type", opts.FacebookMediaType)
			fb.setIfNotEmpty("thumbnail_url", opts.ThumbnailURL)
		case "pinterest":
			addPinterestVideoParams(fb, opts)
		case "x":
			addXMediaParams(fb, xMediaParams{
				ReplySettings:         opts.XReplySettings,
				Nullcast:              opts.XNullcast,
				TaggedUserIDs:         opts.XTaggedUserIDs,
				PlaceID:               opts.XPlaceID,
				GeoPlaceID:            opts.XGeoPlaceID,
				ForSuperFollowersOnly: opts.XForSuperFollowersOnly,
				CommunityID:           opts.XCommunityID,
				ShareWithFollowers:    opts.XShareWithFollowers,
				DirectMessageDeepLink: opts.XDirectMessageDeepLink,
				LongTextAsPost:        opts.XLongTextAsPost,
				ThreadImageLayout:     opts.XThreadImageLayout,
			})
		case "threads":
			addThreadsParams(fb, opts.ThreadsLongTextAsPost, opts.ThreadsThreadMediaLayout, opts.ThreadsTopicTag)
		}
	}

	var resp UploadResponse
	if err := c.postForm(ctx, "/upload", fb, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadPhotos uploads one or more photos (local file paths or URLs) to social
// media platforms.
func (c *Client) UploadPhotos(ctx context.Context, photosPathsOrURLs []string, opts PhotosOptions) (*UploadResponse, error) {
	if opts.User == "" {
		return nil, fmt.Errorf("uploadpost: PhotosOptions.User is required")
	}
	if len(opts.Platforms) == 0 {
		return nil, fmt.Errorf("uploadpost: PhotosOptions.Platforms must not be empty")
	}
	if len(photosPathsOrURLs) == 0 {
		return nil, fmt.Errorf("uploadpost: photosPathsOrURLs must not be empty")
	}

	fb := newFormBuilder()

	if err := fb.setFiles("photos[]", photosPathsOrURLs); err != nil {
		return nil, err
	}

	addCommonParams(fb, commonParams{
		User:           opts.User,
		Title:          opts.Title,
		Platforms:      opts.Platforms,
		FirstComment:   opts.FirstComment,
		AltText:        opts.AltText,
		ScheduledDate:  opts.ScheduledDate,
		Timezone:       opts.Timezone,
		AddToQueue:     opts.AddToQueue,
		MaxPostsPerSlot: opts.MaxPostsPerSlot,
		AsyncUpload:    opts.AsyncUpload,
		FirstCommentMedia: opts.FirstCommentMedia,

		BlueskyTitle:   opts.BlueskyTitle,
		InstagramTitle: opts.InstagramTitle,
		FacebookTitle:  opts.FacebookTitle,
		TiktokTitle:    opts.TiktokTitle,
		LinkedinTitle:  opts.LinkedinTitle,
		XTitle:         opts.XTitle,
		PinterestTitle: opts.PinterestTitle,
		ThreadsTitle:   opts.ThreadsTitle,

		Description:          opts.Description,
		LinkedinDescription:  opts.LinkedinDescription,
		FacebookDescription:  opts.FacebookDescription,
		TiktokDescription:    opts.TiktokDescription,
		PinterestDescription: opts.PinterestDescription,

		InstagramFirstComment: opts.InstagramFirstComment,
		FacebookFirstComment:  opts.FacebookFirstComment,
		XFirstComment:         opts.XFirstComment,
		ThreadsFirstComment:   opts.ThreadsFirstComment,
		RedditFirstComment:    opts.RedditFirstComment,
		BlueskyFirstComment:   opts.BlueskyFirstComment,
		LinkedinFirstComment:  opts.LinkedinFirstComment,
	})

	for _, p := range opts.Platforms {
		switch p {
		case "tiktok":
			fb.setBool("auto_add_music", opts.TiktokAutoAddMusic)
			fb.setBool("disable_comment", opts.TiktokDisableComment)
			fb.setInt("photo_cover_index", opts.TiktokPhotoCoverIndex)
			fb.setBool("brand_content_toggle", opts.BrandContentToggle)
			fb.setBool("brand_organic_toggle", opts.BrandOrganicToggle)
		case "instagram":
			fb.setIfNotEmpty("media_type", opts.InstagramMediaType)
			fb.setIfNotEmpty("collaborators", opts.InstagramCollaborators)
			fb.setIfNotEmpty("user_tags", opts.InstagramUserTags)
			fb.setIfNotEmpty("location_id", opts.InstagramLocationID)
		case "linkedin":
			fb.setIfNotEmpty("visibility", opts.LinkedinVisibility)
			fb.setIfNotEmpty("target_linkedin_page_id", opts.TargetLinkedinPageID)
		case "facebook":
			fb.setIfNotEmpty("facebook_page_id", opts.FacebookPageID)
		case "pinterest":
			fb.setIfNotEmpty("pinterest_board_id", opts.PinterestBoardID)
			fb.setIfNotEmpty("pinterest_alt_text", opts.PinterestAltText)
			fb.setIfNotEmpty("pinterest_link", opts.PinterestLink)
		case "x":
			addXMediaParams(fb, xMediaParams{
				ReplySettings:         opts.XReplySettings,
				Nullcast:              opts.XNullcast,
				TaggedUserIDs:         opts.XTaggedUserIDs,
				GeoPlaceID:            opts.XGeoPlaceID,
				ForSuperFollowersOnly: opts.XForSuperFollowersOnly,
				CommunityID:           opts.XCommunityID,
				ShareWithFollowers:    opts.XShareWithFollowers,
				DirectMessageDeepLink: opts.XDirectMessageDeepLink,
				LongTextAsPost:        opts.XLongTextAsPost,
				ThreadImageLayout:     opts.XThreadImageLayout,
			})
		case "threads":
			addThreadsParams(fb, opts.ThreadsLongTextAsPost, opts.ThreadsThreadMediaLayout, opts.ThreadsTopicTag)
		case "reddit":
			fb.setIfNotEmpty("subreddit", opts.RedditSubreddit)
			fb.setIfNotEmpty("flair_id", opts.RedditFlairID)
		}
	}

	var resp UploadResponse
	if err := c.postForm(ctx, "/upload_photos", fb, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadText uploads a text post to social media platforms.
func (c *Client) UploadText(ctx context.Context, opts TextOptions) (*UploadResponse, error) {
	if opts.User == "" {
		return nil, fmt.Errorf("uploadpost: TextOptions.User is required")
	}
	if len(opts.Platforms) == 0 {
		return nil, fmt.Errorf("uploadpost: TextOptions.Platforms must not be empty")
	}

	fb := newFormBuilder()

	addCommonParams(fb, commonParams{
		User:      opts.User,
		Title:     opts.Title,
		Platforms: opts.Platforms,

		FirstComment:  opts.FirstComment,
		ScheduledDate: opts.ScheduledDate,
		Timezone:      opts.Timezone,
		AddToQueue:    opts.AddToQueue,
		MaxPostsPerSlot: opts.MaxPostsPerSlot,
		AsyncUpload:   opts.AsyncUpload,
		FirstCommentMedia: opts.FirstCommentMedia,

		BlueskyTitle:  opts.BlueskyTitle,
		FacebookTitle: opts.FacebookTitle,
		TiktokTitle:   opts.TiktokTitle,
		LinkedinTitle: opts.LinkedinTitle,
		XTitle:        opts.XTitle,
		ThreadsTitle:  opts.ThreadsTitle,

		LinkedinDescription: opts.LinkedinDescription,

		FacebookFirstComment: opts.FacebookFirstComment,
		XFirstComment:        opts.XFirstComment,
		ThreadsFirstComment:  opts.ThreadsFirstComment,
		RedditFirstComment:   opts.RedditFirstComment,
		BlueskyFirstComment:  opts.BlueskyFirstComment,
		LinkedinFirstComment: opts.LinkedinFirstComment,
	})

	fb.setIfNotEmpty("link_url", opts.LinkURL)

	for _, p := range opts.Platforms {
		switch p {
		case "linkedin":
			fb.setIfNotEmpty("visibility", opts.LinkedinVisibility)
			fb.setIfNotEmpty("target_linkedin_page_id", opts.TargetLinkedinPageID)
			linkURL := opts.LinkedinLinkURL
			if linkURL == "" {
				linkURL = opts.LinkURL
			}
			fb.setIfNotEmpty("linkedin_link_url", linkURL)
		case "facebook":
			fb.setIfNotEmpty("facebook_page_id", opts.FacebookPageID)
			fb.setIfNotEmpty("facebook_link_url", opts.FacebookLinkURL)
		case "x":
			addXTextParams(fb, opts)
		case "threads":
			addThreadsParams(fb, opts.ThreadsLongTextAsPost, opts.ThreadsThreadMediaLayout, opts.ThreadsTopicTag)
		case "reddit":
			fb.setIfNotEmpty("subreddit", opts.RedditSubreddit)
			fb.setIfNotEmpty("flair_id", opts.RedditFlairID)
			linkURL := opts.RedditLinkURL
			if linkURL == "" {
				linkURL = opts.LinkURL
			}
			fb.setIfNotEmpty("reddit_link_url", linkURL)
		case "bluesky":
			linkURL := opts.BlueskyLinkURL
			if linkURL == "" {
				linkURL = opts.LinkURL
			}
			fb.setIfNotEmpty("bluesky_link_url", linkURL)
		}
	}

	var resp UploadResponse
	if err := c.postForm(ctx, "/upload_text", fb, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadDocument uploads a document (PDF, PPT, PPTX, DOC, DOCX) to LinkedIn.
// documentPathOrURL can be a local file path or an HTTPS URL.
func (c *Client) UploadDocument(ctx context.Context, documentPathOrURL string, opts DocumentOptions) (*UploadResponse, error) {
	if opts.User == "" {
		return nil, fmt.Errorf("uploadpost: DocumentOptions.User is required")
	}
	if opts.Title == "" {
		return nil, fmt.Errorf("uploadpost: DocumentOptions.Title is required")
	}

	fb := newFormBuilder()

	if err := fb.setFile("document", documentPathOrURL); err != nil {
		return nil, err
	}

	fb.set("user", opts.User)
	fb.set("title", opts.Title)
	fb.set("platform[]", "linkedin")

	fb.setIfNotEmpty("description", opts.Description)
	fb.setIfNotEmpty("scheduled_date", opts.ScheduledDate)
	fb.setIfNotEmpty("timezone", opts.Timezone)
	fb.setBool("add_to_queue", opts.AddToQueue)
	fb.setInt("max_posts_per_slot", opts.MaxPostsPerSlot)
	fb.setBool("async_upload", opts.AsyncUpload)
	fb.setIfNotEmpty("visibility", opts.LinkedinVisibility)
	fb.setIfNotEmpty("target_linkedin_page_id", opts.TargetLinkedinPageID)

	var resp UploadResponse
	if err := c.postForm(ctx, "/upload_document", fb, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// Internal helpers
// ─────────────────────────────────────────────────────────────────────────────

type commonParams struct {
	User      string
	Title     string
	Platforms []string

	FirstComment  string
	AltText       string
	ScheduledDate string
	Timezone      string
	AddToQueue    *bool
	MaxPostsPerSlot *int
	AsyncUpload   *bool
	FirstCommentMedia []string

	BlueskyTitle   string
	InstagramTitle string
	FacebookTitle  string
	TiktokTitle    string
	LinkedinTitle  string
	XTitle         string
	YoutubeTitle   string
	PinterestTitle string
	ThreadsTitle   string

	Description          string
	LinkedinDescription  string
	YoutubeDescription   string
	FacebookDescription  string
	TiktokDescription    string
	PinterestDescription string

	InstagramFirstComment string
	FacebookFirstComment  string
	XFirstComment         string
	ThreadsFirstComment   string
	YoutubeFirstComment   string
	RedditFirstComment    string
	BlueskyFirstComment   string
	LinkedinFirstComment  string
}

func addCommonParams(fb *formBuilder, p commonParams) {
	fb.set("user", p.User)
	fb.setIfNotEmpty("title", p.Title)
	fb.setArray("platform[]", p.Platforms)

	fb.setIfNotEmpty("first_comment", p.FirstComment)
	fb.setIfNotEmpty("alt_text", p.AltText)
	fb.setIfNotEmpty("scheduled_date", p.ScheduledDate)
	fb.setIfNotEmpty("timezone", p.Timezone)
	fb.setBool("add_to_queue", p.AddToQueue)
	fb.setInt("max_posts_per_slot", p.MaxPostsPerSlot)
	fb.setBool("async_upload", p.AsyncUpload)

	// Platform-specific title overrides
	fb.setIfNotEmpty("bluesky_title", p.BlueskyTitle)
	fb.setIfNotEmpty("instagram_title", p.InstagramTitle)
	fb.setIfNotEmpty("facebook_title", p.FacebookTitle)
	fb.setIfNotEmpty("tiktok_title", p.TiktokTitle)
	fb.setIfNotEmpty("linkedin_title", p.LinkedinTitle)
	fb.setIfNotEmpty("x_title", p.XTitle)
	fb.setIfNotEmpty("youtube_title", p.YoutubeTitle)
	fb.setIfNotEmpty("pinterest_title", p.PinterestTitle)
	fb.setIfNotEmpty("threads_title", p.ThreadsTitle)

	// Descriptions
	fb.setIfNotEmpty("description", p.Description)
	fb.setIfNotEmpty("linkedin_description", p.LinkedinDescription)
	fb.setIfNotEmpty("youtube_description", p.YoutubeDescription)
	fb.setIfNotEmpty("facebook_description", p.FacebookDescription)
	fb.setIfNotEmpty("tiktok_description", p.TiktokDescription)
	fb.setIfNotEmpty("pinterest_description", p.PinterestDescription)

	// Per-platform first comments
	fb.setIfNotEmpty("instagram_first_comment", p.InstagramFirstComment)
	fb.setIfNotEmpty("facebook_first_comment", p.FacebookFirstComment)
	fb.setIfNotEmpty("x_first_comment", p.XFirstComment)
	fb.setIfNotEmpty("threads_first_comment", p.ThreadsFirstComment)
	fb.setIfNotEmpty("youtube_first_comment", p.YoutubeFirstComment)
	fb.setIfNotEmpty("reddit_first_comment", p.RedditFirstComment)
	fb.setIfNotEmpty("bluesky_first_comment", p.BlueskyFirstComment)
	fb.setIfNotEmpty("linkedin_first_comment", p.LinkedinFirstComment)

	// First comment media
	for _, m := range p.FirstCommentMedia {
		_ = fb.setFile("first_comment_media[]", m)
	}
}

func addTiktokVideoParams(fb *formBuilder, opts VideoOptions) {
	fb.setBool("disable_comment", opts.TiktokDisableComment)
	fb.setBool("brand_content_toggle", opts.BrandContentToggle)
	fb.setBool("brand_organic_toggle", opts.BrandOrganicToggle)
	fb.setIfNotEmpty("privacy_level", opts.TiktokPrivacyLevel)
	fb.setBool("disable_duet", opts.TiktokDisableDuet)
	fb.setBool("disable_stitch", opts.TiktokDisableStitch)
	fb.setInt64("cover_timestamp", opts.TiktokCoverTimestamp)
	fb.setBool("is_aigc", opts.TiktokIsAIGC)
	fb.setIfNotEmpty("post_mode", opts.TiktokPostMode)
}

func addInstagramVideoParams(fb *formBuilder, opts VideoOptions) {
	fb.setIfNotEmpty("media_type", opts.InstagramMediaType)
	fb.setIfNotEmpty("collaborators", opts.InstagramCollaborators)
	fb.setIfNotEmpty("user_tags", opts.InstagramUserTags)
	fb.setIfNotEmpty("location_id", opts.InstagramLocationID)
	fb.setBool("share_to_feed", opts.InstagramShareToFeed)

	if opts.InstagramCoverURL != "" {
		if isURL(opts.InstagramCoverURL) {
			fb.set("cover_url", opts.InstagramCoverURL)
		} else {
			_ = fb.setFile("cover_image", opts.InstagramCoverURL)
		}
	}

	fb.setIfNotEmpty("audio_name", opts.InstagramAudioName)
	fb.setIfNotEmpty("thumb_offset", opts.InstagramThumbOffset)
}

func addYoutubeParams(fb *formBuilder, opts VideoOptions) error {
	fb.setArray("tags[]", opts.YouTubeTags)
	fb.setIfNotEmpty("categoryId", opts.YouTubeCategoryID)
	fb.setIfNotEmpty("privacyStatus", opts.YouTubePrivacyStatus)
	fb.setBool("embeddable", opts.YouTubeEmbeddable)
	fb.setIfNotEmpty("license", opts.YouTubeLicense)
	fb.setBool("publicStatsViewable", opts.YouTubePublicStatsViewable)
	fb.setIfNotEmpty("thumbnail_url", opts.YouTubeThumbnailURL)
	fb.setBool("selfDeclaredMadeForKids", opts.YouTubeSelfDeclaredMadeForKids)
	fb.setBool("containsSyntheticMedia", opts.YouTubeContainsSyntheticMedia)
	fb.setIfNotEmpty("defaultLanguage", opts.YouTubeDefaultLanguage)
	fb.setIfNotEmpty("defaultAudioLanguage", opts.YouTubeDefaultAudioLanguage)
	fb.setIfNotEmpty("allowedCountries", opts.YouTubeAllowedCountries)
	fb.setIfNotEmpty("blockedCountries", opts.YouTubeBlockedCountries)
	fb.setBool("hasPaidProductPlacement", opts.YouTubeHasPaidProductPlacement)
	fb.setIfNotEmpty("recordingDate", opts.YouTubeRecordingDate)

	for i, sub := range opts.YouTubeSubtitles {
		if sub.Language == "" {
			continue
		}
		idx := strconv.Itoa(i)
		fb.set("youtube_subtitle_language_"+idx, sub.Language)
		fb.setIfNotEmpty("youtube_subtitle_name_"+idx, sub.Name)
		if sub.File != "" {
			if err := fb.setFile("youtube_subtitle_file_"+idx, sub.File); err != nil {
				return err
			}
		} else if sub.URL != "" {
			fb.set("youtube_subtitle_file_"+idx, sub.URL)
		}
	}
	return nil
}

func addPinterestVideoParams(fb *formBuilder, opts VideoOptions) {
	fb.setIfNotEmpty("pinterest_board_id", opts.PinterestBoardID)
	fb.setIfNotEmpty("pinterest_alt_text", opts.PinterestAltText)
	fb.setIfNotEmpty("pinterest_link", opts.PinterestLink)
	fb.setIfNotEmpty("pinterest_cover_image_url", opts.PinterestCoverImageURL)
	fb.setIfNotEmpty("pinterest_cover_image_content_type", opts.PinterestCoverImageContentType)
	fb.setIfNotEmpty("pinterest_cover_image_data", opts.PinterestCoverImageData)
	fb.setInt("pinterest_cover_image_key_frame_time", opts.PinterestCoverImageKeyFrameTime)
}

type xMediaParams struct {
	ReplySettings         string
	Nullcast              *bool
	TaggedUserIDs         []string
	PlaceID               string
	GeoPlaceID            string
	ForSuperFollowersOnly *bool
	CommunityID           string
	ShareWithFollowers    *bool
	DirectMessageDeepLink string
	LongTextAsPost        *bool
	ThreadImageLayout     string
}

func addXMediaParams(fb *formBuilder, p xMediaParams) {
	if p.ReplySettings != "" && p.ReplySettings != "everyone" {
		fb.set("reply_settings", p.ReplySettings)
	}
	fb.setBool("nullcast", p.Nullcast)
	fb.setArray("tagged_user_ids[]", p.TaggedUserIDs)
	fb.setIfNotEmpty("place_id", p.PlaceID)
	fb.setIfNotEmpty("geo_place_id", p.GeoPlaceID)
	fb.setBool("for_super_followers_only", p.ForSuperFollowersOnly)
	fb.setIfNotEmpty("community_id", p.CommunityID)
	fb.setBool("share_with_followers", p.ShareWithFollowers)
	fb.setIfNotEmpty("direct_message_deep_link", p.DirectMessageDeepLink)
	fb.setBool("x_long_text_as_post", p.LongTextAsPost)
	fb.setIfNotEmpty("x_thread_image_layout", p.ThreadImageLayout)
}

func addXTextParams(fb *formBuilder, opts TextOptions) {
	if opts.XReplySettings != "" && opts.XReplySettings != "everyone" {
		fb.set("reply_settings", opts.XReplySettings)
	}
	fb.setBool("nullcast", opts.XNullcast)
	fb.setIfNotEmpty("geo_place_id", opts.XGeoPlaceID)
	fb.setBool("for_super_followers_only", opts.XForSuperFollowersOnly)
	fb.setIfNotEmpty("community_id", opts.XCommunityID)
	fb.setBool("share_with_followers", opts.XShareWithFollowers)
	fb.setIfNotEmpty("direct_message_deep_link", opts.XDirectMessageDeepLink)
	fb.setBool("x_long_text_as_post", opts.XLongTextAsPost)
	fb.setIfNotEmpty("post_url", opts.XPostURL)
	fb.setIfNotEmpty("quote_tweet_id", opts.XQuoteTweetID)
	fb.setIfNotEmpty("card_uri", opts.XCardURI)

	if len(opts.XPollOptions) > 0 {
		fb.setArray("poll_options[]", opts.XPollOptions)
		fb.setInt("poll_duration", opts.XPollDuration)
		fb.setIfNotEmpty("poll_reply_settings", opts.XPollReplySettings)
	}
}

func addThreadsParams(fb *formBuilder, longTextAsPost *bool, mediaLayout, topicTag string) {
	fb.setBool("threads_long_text_as_post", longTextAsPost)
	fb.setIfNotEmpty("threads_thread_media_layout", mediaLayout)
	fb.setIfNotEmpty("threads_topic_tag", topicTag)
}

