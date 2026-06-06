package uploadpost

// ─────────────────────────────────────────────────────────────────────────────
// Option types
// ─────────────────────────────────────────────────────────────────────────────

// YouTubeSubtitle describes a single subtitle/caption track to attach to a
// YouTube video upload.
type YouTubeSubtitle struct {
	// Language is a BCP-47 language code, e.g. "en", "es".
	Language string
	// Name is the human-readable track name, e.g. "English".
	Name string
	// File is a local file path (SRT, VTT, SBV, SUB, ASS, SSA, TTML).
	// Mutually exclusive with URL.
	File string
	// URL is a remote URL pointing to the subtitle file.
	// Mutually exclusive with File.
	URL string
}

// VideoOptions contains all options for UploadVideo.
type VideoOptions struct {
	// ── Required ──────────────────────────────────────────────────────────────
	User      string   // profile username
	Platforms []string // e.g. []string{"tiktok", "instagram", "youtube"}

	// ── Common optional ───────────────────────────────────────────────────────
	Title             string
	Description       string
	FirstComment      string
	AltText           string
	ScheduledDate     string // ISO 8601, e.g. "2024-12-25T10:00:00Z"
	Timezone          string // e.g. "Europe/Madrid"
	AddToQueue        *bool
	MaxPostsPerSlot   *int
	AsyncUpload       *bool
	FirstCommentMedia []string // local file paths or URLs

	// Per-platform title overrides
	BlueskyTitle   string
	InstagramTitle string
	FacebookTitle  string
	TiktokTitle    string
	LinkedinTitle  string
	XTitle         string
	YoutubeTitle   string
	PinterestTitle string
	ThreadsTitle   string

	// Per-platform description overrides
	LinkedinDescription  string
	YoutubeDescription   string
	FacebookDescription  string
	TiktokDescription    string
	PinterestDescription string

	// Per-platform first-comment overrides
	InstagramFirstComment string
	FacebookFirstComment  string
	XFirstComment         string
	ThreadsFirstComment   string
	YoutubeFirstComment   string
	RedditFirstComment    string
	BlueskyFirstComment   string
	LinkedinFirstComment  string

	// ── TikTok ────────────────────────────────────────────────────────────────
	TiktokPrivacyLevel   string // PUBLIC_TO_EVERYONE | MUTUAL_FOLLOW_FRIENDS | FOLLOWER_OF_CREATOR | SELF_ONLY
	TiktokDisableDuet    *bool
	TiktokDisableComment *bool
	TiktokDisableStitch  *bool
	TiktokCoverTimestamp *int64 // milliseconds
	TiktokIsAIGC         *bool
	TiktokPostMode       string // DIRECT_POST | MEDIA_UPLOAD
	BrandContentToggle   *bool
	BrandOrganicToggle   *bool

	// ── Instagram ─────────────────────────────────────────────────────────────
	InstagramMediaType     string // REELS | STORIES
	InstagramShareToFeed   *bool
	InstagramCollaborators string
	InstagramCoverURL      string // URL string or local file path
	InstagramAudioName     string
	InstagramUserTags      string
	InstagramLocationID    string
	InstagramThumbOffset   string

	// ── YouTube ───────────────────────────────────────────────────────────────
	YouTubeTags                    []string
	YouTubeCategoryID              string
	YouTubePrivacyStatus           string // public | unlisted | private
	YouTubeEmbeddable              *bool
	YouTubeLicense                 string // youtube | creativeCommon
	YouTubePublicStatsViewable     *bool
	YouTubeThumbnailURL            string
	YouTubeSelfDeclaredMadeForKids *bool
	YouTubeContainsSyntheticMedia  *bool
	YouTubeDefaultLanguage         string
	YouTubeDefaultAudioLanguage    string
	YouTubeAllowedCountries        string
	YouTubeBlockedCountries        string
	YouTubeHasPaidProductPlacement *bool
	YouTubeRecordingDate           string
	YouTubeSubtitles               []YouTubeSubtitle

	// ── LinkedIn ──────────────────────────────────────────────────────────────
	LinkedinVisibility   string // PUBLIC | CONNECTIONS | LOGGED_IN | CONTAINER
	TargetLinkedinPageID string

	// ── Facebook ──────────────────────────────────────────────────────────────
	FacebookPageID     string
	FacebookVideoState string // PUBLISHED | DRAFT
	FacebookMediaType  string // REELS | STORIES | VIDEO
	ThumbnailURL       string // only for FacebookMediaType == "VIDEO"

	// ── Pinterest ─────────────────────────────────────────────────────────────
	PinterestBoardID                string
	PinterestAltText                string
	PinterestLink                   string
	PinterestCoverImageURL          string
	PinterestCoverImageContentType  string
	PinterestCoverImageData         string
	PinterestCoverImageKeyFrameTime *int

	// ── X (Twitter) ───────────────────────────────────────────────────────────
	XReplySettings         string // everyone | following | mentionedUsers | subscribers | verified
	XNullcast              *bool
	XTaggedUserIDs         []string
	XPlaceID               string
	XGeoPlaceID            string
	XForSuperFollowersOnly *bool
	XCommunityID           string
	XShareWithFollowers    *bool
	XDirectMessageDeepLink string
	XLongTextAsPost        *bool
	XThreadImageLayout     string // e.g. "4,4" or "2,3,1"

	// ── Threads ───────────────────────────────────────────────────────────────
	ThreadsLongTextAsPost    *bool
	ThreadsThreadMediaLayout string
	ThreadsTopicTag          string
}

// PhotosOptions contains all options for UploadPhotos.
type PhotosOptions struct {
	// ── Required ──────────────────────────────────────────────────────────────
	User      string
	Platforms []string

	// ── Common optional ───────────────────────────────────────────────────────
	Title             string
	Description       string
	FirstComment      string
	AltText           string
	ScheduledDate     string
	Timezone          string
	AddToQueue        *bool
	MaxPostsPerSlot   *int
	AsyncUpload       *bool
	FirstCommentMedia []string

	// Per-platform title overrides
	BlueskyTitle   string
	InstagramTitle string
	FacebookTitle  string
	TiktokTitle    string
	LinkedinTitle  string
	XTitle         string
	PinterestTitle string
	ThreadsTitle   string

	// Per-platform description overrides
	LinkedinDescription  string
	FacebookDescription  string
	TiktokDescription    string
	PinterestDescription string

	// Per-platform first-comment overrides
	InstagramFirstComment string
	FacebookFirstComment  string
	XFirstComment         string
	ThreadsFirstComment   string
	RedditFirstComment    string
	BlueskyFirstComment   string
	LinkedinFirstComment  string

	// ── TikTok ────────────────────────────────────────────────────────────────
	TiktokAutoAddMusic    *bool
	TiktokDisableComment  *bool
	TiktokPhotoCoverIndex *int
	BrandContentToggle    *bool
	BrandOrganicToggle    *bool

	// ── Instagram ─────────────────────────────────────────────────────────────
	InstagramMediaType     string // IMAGE | STORIES
	InstagramCollaborators string
	InstagramUserTags      string
	InstagramLocationID    string

	// ── LinkedIn ──────────────────────────────────────────────────────────────
	LinkedinVisibility   string
	TargetLinkedinPageID string

	// ── Facebook ──────────────────────────────────────────────────────────────
	FacebookPageID string

	// ── Pinterest ─────────────────────────────────────────────────────────────
	PinterestBoardID string
	PinterestAltText string
	PinterestLink    string

	// ── X (Twitter) ───────────────────────────────────────────────────────────
	XReplySettings         string
	XNullcast              *bool
	XTaggedUserIDs         []string
	XGeoPlaceID            string
	XForSuperFollowersOnly *bool
	XCommunityID           string
	XShareWithFollowers    *bool
	XDirectMessageDeepLink string
	XLongTextAsPost        *bool
	XThreadImageLayout     string

	// ── Threads ───────────────────────────────────────────────────────────────
	ThreadsLongTextAsPost    *bool
	ThreadsThreadMediaLayout string
	ThreadsTopicTag          string

	// ── Reddit ────────────────────────────────────────────────────────────────
	RedditSubreddit string
	RedditFlairID   string
}

// TextOptions contains all options for UploadText.
type TextOptions struct {
	// ── Required ──────────────────────────────────────────────────────────────
	User      string
	Platforms []string
	Title     string // the text content of the post

	// ── Common optional ───────────────────────────────────────────────────────
	FirstComment      string
	ScheduledDate     string
	Timezone          string
	AddToQueue        *bool
	MaxPostsPerSlot   *int
	AsyncUpload       *bool
	FirstCommentMedia []string

	// Per-platform title overrides
	BlueskyTitle  string
	FacebookTitle string
	TiktokTitle   string
	LinkedinTitle string
	XTitle        string
	ThreadsTitle  string

	// Per-platform first-comment overrides
	FacebookFirstComment string
	XFirstComment        string
	ThreadsFirstComment  string
	RedditFirstComment   string
	BlueskyFirstComment  string
	LinkedinFirstComment string

	// ── Link preview (generic) ────────────────────────────────────────────────
	// LinkURL applies to LinkedIn, Bluesky, Facebook (platform-specific overrides take priority).
	LinkURL         string
	LinkedinLinkURL string
	BlueskyLinkURL  string

	// ── LinkedIn ──────────────────────────────────────────────────────────────
	LinkedinVisibility   string
	TargetLinkedinPageID string
	LinkedinDescription  string

	// ── Facebook ──────────────────────────────────────────────────────────────
	FacebookPageID  string
	FacebookLinkURL string

	// ── X (Twitter) ───────────────────────────────────────────────────────────
	XReplySettings         string
	XNullcast              *bool
	XPostURL               string
	XQuoteTweetID          string
	XPollOptions           []string
	XPollDuration          *int // minutes (5–10080)
	XPollReplySettings     string
	XCardURI               string
	XGeoPlaceID            string
	XForSuperFollowersOnly *bool
	XCommunityID           string
	XShareWithFollowers    *bool
	XDirectMessageDeepLink string
	XLongTextAsPost        *bool

	// ── Threads ───────────────────────────────────────────────────────────────
	ThreadsLongTextAsPost    *bool
	ThreadsThreadMediaLayout string
	ThreadsTopicTag          string

	// ── Reddit ────────────────────────────────────────────────────────────────
	RedditSubreddit string
	RedditFlairID   string
	RedditLinkURL   string
}

// DocumentOptions contains all options for UploadDocument (LinkedIn only).
type DocumentOptions struct {
	// ── Required ──────────────────────────────────────────────────────────────
	User  string
	Title string

	// ── Optional ──────────────────────────────────────────────────────────────
	Description          string
	LinkedinVisibility   string // PUBLIC | CONNECTIONS | LOGGED_IN | CONTAINER
	TargetLinkedinPageID string
	ScheduledDate        string
	Timezone             string
	AddToQueue           *bool
	MaxPostsPerSlot      *int
	AsyncUpload          *bool
}

// HistoryOptions contains query parameters for GetHistory.
type HistoryOptions struct {
	Page  int // default 1
	Limit int // 20 | 50 | 100
}

// AnalyticsOptions contains query parameters for GetAnalytics.
type AnalyticsOptions struct {
	Platforms []string
	PageID    string // Facebook Page ID
	PageURN   string // LinkedIn page URN
}

// ImpressionsOptions contains query parameters for GetTotalImpressions.
type ImpressionsOptions struct {
	// Period shortcut: last_day, last_week, last_month, last_3months, last_year
	Period    string
	StartDate string // YYYY-MM-DD
	EndDate   string // YYYY-MM-DD
	Date      string // YYYY-MM-DD
	Platforms []string
	Breakdown bool
	Metrics   []string
}

// JWTOptions contains options for GenerateJWT.
type JWTOptions struct {
	RedirectURL        string
	LogoImage          string
	RedirectButtonText string
	Platforms          []string
	ShowCalendar       *bool
	ReadonlyCalendar   *bool
	ConnectTitle       string
	ConnectDescription string
	Language           string // en | es | de | fr | pt
}

// UserPreferencesOptions contains options for UpdateUserPreferences.
type UserPreferencesOptions struct {
	WeekStartDay *int // 0=Sunday, 1=Monday
}

// NotificationConfigOptions contains options for UpdateNotificationConfig.
type NotificationConfigOptions struct {
	WebhookEvents []string // upload_completed | social_account_connected | social_account_disconnected | social_account_reauth_required
	WebhookURL    string
}

// EditScheduledOptions contains options for EditScheduled.
type EditScheduledOptions struct {
	ScheduledDate string // ISO 8601
	Timezone      string
}

// AutoDMOptions contains options for StartAutoDM.
type AutoDMOptions struct {
	PostURL            string   // Required: Instagram post URL
	ReplyMessage       string   // Required: DM message to send
	ProfileUsername    string   // Required: profile username
	MonitoringInterval *int     // minutes, default 15, minimum 15
	TriggerKeywords    []string // optional keyword filter
}

// ─────────────────────────────────────────────────────────────────────────────
// Response types
// ─────────────────────────────────────────────────────────────────────────────

// UploadResponse is returned by upload methods.
type UploadResponse struct {
	Success   bool   `json:"success"`
	RequestID string `json:"request_id,omitempty"`
	JobID     string `json:"job_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

// StatusResponse is returned by GetStatus and GetJobStatus.
type StatusResponse struct {
	Success   bool        `json:"success"`
	Status    string      `json:"status,omitempty"`
	RequestID string      `json:"request_id,omitempty"`
	JobID     string      `json:"job_id,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

// HistoryItem represents a single upload in the history list.
type HistoryItem struct {
	RequestID string      `json:"request_id,omitempty"`
	Status    string      `json:"status,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	Platforms interface{} `json:"platforms,omitempty"`
}

// HistoryResponse is returned by GetHistory.
type HistoryResponse struct {
	Success bool          `json:"success"`
	Uploads []HistoryItem `json:"uploads,omitempty"`
	Page    int           `json:"page,omitempty"`
	Total   int           `json:"total,omitempty"`
}

// AnalyticsResponse is returned by GetAnalytics.
type AnalyticsResponse struct {
	Success   bool        `json:"success"`
	Analytics interface{} `json:"analytics,omitempty"`
}

// ImpressionsResponse is returned by GetTotalImpressions.
type ImpressionsResponse struct {
	Success          bool                   `json:"success"`
	ProfileUsername  string                 `json:"profile_username,omitempty"`
	StartDate        string                 `json:"start_date,omitempty"`
	EndDate          string                 `json:"end_date,omitempty"`
	TotalImpressions *int64                 `json:"total_impressions,omitempty"`
	Metrics          map[string]interface{} `json:"metrics,omitempty"`
	PerPlatform      map[string]interface{} `json:"per_platform,omitempty"`
	PerDay           map[string]interface{} `json:"per_day,omitempty"`
	PlatformsFilter  []string               `json:"platforms_filter,omitempty"`
}

// PostAnalyticsPlatform holds per-platform analytics for a single post.
type PostAnalyticsPlatform struct {
	Success               bool                   `json:"success"`
	PlatformPostID        string                 `json:"platform_post_id,omitempty"`
	PostURL               string                 `json:"post_url,omitempty"`
	PostMetrics           map[string]interface{} `json:"post_metrics,omitempty"`
	PostMetricsSource     string                 `json:"post_metrics_source,omitempty"`
	PostMetricsError      string                 `json:"post_metrics_error,omitempty"`
	ProfileSnapshotAtPost map[string]interface{} `json:"profile_snapshot_at_post_date,omitempty"`
	ProfileSnapshotLatest map[string]interface{} `json:"profile_snapshot_latest,omitempty"`
	ProfileSnapshotDate   string                 `json:"profile_snapshot_latest_date,omitempty"`
	ErrorMessage          string                 `json:"error_message,omitempty"`
}

// PostAnalyticsPost holds metadata about the uploaded post.
type PostAnalyticsPost struct {
	RequestID       string `json:"request_id,omitempty"`
	ProfileUsername string `json:"profile_username,omitempty"`
	PostTitle       string `json:"post_title,omitempty"`
	PostCaption     string `json:"post_caption,omitempty"`
	MediaType       string `json:"media_type,omitempty"`
	UploadTimestamp string `json:"upload_timestamp,omitempty"`
}

// PostAnalyticsResponse is returned by GetPostAnalytics.
type PostAnalyticsResponse struct {
	Success   bool                             `json:"success"`
	Post      PostAnalyticsPost                `json:"post,omitempty"`
	Platforms map[string]PostAnalyticsPlatform `json:"platforms,omitempty"`
}

// PlatformMetric describes available metrics for a single platform.
type PlatformMetric struct {
	PrimaryImpressionsField string            `json:"primary_impressions_field"`
	AvailableMetrics        []string          `json:"available_metrics"`
	MetricLabels            map[string]string `json:"metric_labels"`
}

// ScheduledItem represents a single scheduled post.
type ScheduledItem struct {
	JobID         string      `json:"job_id,omitempty"`
	Status        string      `json:"status,omitempty"`
	ScheduledDate string      `json:"scheduled_date,omitempty"`
	Timezone      string      `json:"timezone,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}

// ScheduledResponse is returned by ListScheduled.
type ScheduledResponse struct {
	Success   bool            `json:"success"`
	Scheduled []ScheduledItem `json:"scheduled,omitempty"`
}

// UserProfile represents a user/profile in the system.
type UserProfile struct {
	Username       string                 `json:"username"`
	SocialAccounts map[string]interface{} `json:"social_accounts,omitempty"`
	CreatedAt      string                 `json:"created_at,omitempty"`
}

// UsersResponse is returned by ListUsers.
type UsersResponse struct {
	Success  bool          `json:"success"`
	Profiles []UserProfile `json:"profiles,omitempty"`
}

// JWTResponse is returned by GenerateJWT.
type JWTResponse struct {
	Success       bool   `json:"success"`
	JWT           string `json:"jwt,omitempty"`
	ConnectionURL string `json:"connection_url,omitempty"`
}

// UserPreferences holds calendar and other preferences.
type UserPreferences struct {
	WeekStartDay *int                   `json:"week_start_day,omitempty"`
	Extra        map[string]interface{} `json:"-"`
}

// NotificationConfig holds webhook and notification settings.
type NotificationConfig struct {
	WebhookEvents []string `json:"webhook_events,omitempty"`
	WebhookURL    string   `json:"webhook_url,omitempty"`
}

// MediaItem represents a single post from a connected social account.
type MediaItem struct {
	ID           string `json:"id"`
	Caption      string `json:"caption,omitempty"`
	MediaType    string `json:"media_type,omitempty"` // IMAGE | VIDEO | CAROUSEL_ALBUM | TEXT
	MediaURL     string `json:"media_url,omitempty"`
	Permalink    string `json:"permalink,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
}

// MediaResponse is returned by GetMedia.
type MediaResponse struct {
	Success bool        `json:"success"`
	Media   []MediaItem `json:"media,omitempty"`
}

// Page represents a Facebook or LinkedIn page.
type Page struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// Board represents a Pinterest board.
type Board struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// PagesResponse is returned by GetFacebookPages and GetLinkedinPages.
type PagesResponse struct {
	Success bool   `json:"success"`
	Pages   []Page `json:"pages,omitempty"`
}

// BoardsResponse is returned by GetPinterestBoards.
type BoardsResponse struct {
	Success bool    `json:"success"`
	Boards  []Board `json:"boards,omitempty"`
}

// GoogleBusinessLocation represents a GBP location.
type GoogleBusinessLocation struct {
	LocationID  string `json:"location_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

// GoogleBusinessLocationsResponse is returned by GetGoogleBusinessLocations.
type GoogleBusinessLocationsResponse struct {
	Success   bool                     `json:"success"`
	Locations []GoogleBusinessLocation `json:"locations,omitempty"`
}

// Comment represents an Instagram comment.
type Comment struct {
	ID        string      `json:"id"`
	Text      string      `json:"text,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
	User      CommentUser `json:"user,omitempty"`
}

// CommentUser holds the author info for a comment.
type CommentUser struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// CommentsResponse is returned by GetPostComments.
type CommentsResponse struct {
	Success  bool      `json:"success"`
	Comments []Comment `json:"comments,omitempty"`
}

// ReplyResponse is returned by ReplyToComment and PublicReplyToComment.
type ReplyResponse struct {
	Success     bool   `json:"success"`
	ID          string `json:"id,omitempty"`
	RecipientID string `json:"recipient_id,omitempty"`
	MessageID   string `json:"message_id,omitempty"`
	Message     string `json:"message,omitempty"`
	Error       string `json:"error,omitempty"`
}

// AutoDMMonitor represents a single AutoDM monitor.
type AutoDMMonitor struct {
	MonitorID          string   `json:"monitor_id,omitempty"`
	PostURL            string   `json:"post_url,omitempty"`
	Status             string   `json:"status,omitempty"` // running | paused | resuming | stopped | expired
	ReplyMessage       string   `json:"reply_message,omitempty"`
	ProfileUsername    string   `json:"profile_username,omitempty"`
	MonitoringInterval int      `json:"monitoring_interval,omitempty"`
	TriggerKeywords    []string `json:"trigger_keywords,omitempty"`
	CreatedAt          string   `json:"created_at,omitempty"`
	StoppedAt          string   `json:"stopped_at,omitempty"`
	StopReason         string   `json:"stop_reason,omitempty"`
	IsActive           bool     `json:"is_active"`
}

// AutoDMStatusResponse is returned by GetAutoDMStatus.
type AutoDMStatusResponse struct {
	Success     bool            `json:"success"`
	Monitors    []AutoDMMonitor `json:"monitors,omitempty"`
	TotalActive int             `json:"total_active,omitempty"`
	Total       int             `json:"total,omitempty"`
}

// AutoDMLog represents a single log entry for an AutoDM monitor.
type AutoDMLog struct {
	Timestamp string `json:"timestamp,omitempty"`
	Event     string `json:"event,omitempty"`
	Details   string `json:"details,omitempty"`
}

// AutoDMLogsResponse is returned by GetAutoDMLogs.
type AutoDMLogsResponse struct {
	Success bool        `json:"success"`
	Logs    []AutoDMLog `json:"logs,omitempty"`
}

// SimpleResponse is returned by simple success/failure endpoints.
type SimpleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// CurrentUser is returned by GetCurrentUser (GET /uploadposts/me).
type CurrentUser struct {
	Success     bool                   `json:"success"`
	Message     string                 `json:"message,omitempty"`
	Email       string                 `json:"email,omitempty"`
	Plan        string                 `json:"plan,omitempty"`
	Preferences map[string]interface{} `json:"preferences,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Queue types
// ─────────────────────────────────────────────────────────────────────────────

// QueueSlot defines a single posting slot (hour + minute within a day).
type QueueSlot struct {
	Hour   int `json:"hour"`   // 0-23
	Minute int `json:"minute"` // 0-59
}

// QueueSettings represents the queue configuration for a profile.
type QueueSettings struct {
	Timezone        string      `json:"timezone,omitempty"`
	Slots           []QueueSlot `json:"slots,omitempty"`
	DaysOfWeek      []int       `json:"days_of_week,omitempty"` // 0=Monday … 6=Sunday
	MaxPostsPerSlot int         `json:"max_posts_per_slot,omitempty"`
	FullSlots       []string    `json:"full_slots,omitempty"` // ISO 8601 datetimes
}

// QueueSettingsResponse is returned by GetQueueSettings.
type QueueSettingsResponse struct {
	Success  bool          `json:"success"`
	Settings QueueSettings `json:"settings,omitempty"`
}

// QueueSlotPreview represents a single upcoming queue slot in a preview.
type QueueSlotPreview struct {
	SlotDatetime    string      `json:"slot_datetime,omitempty"` // ISO 8601 UTC
	PostCount       int         `json:"post_count"`
	MaxPostsPerSlot int         `json:"max_posts_per_slot,omitempty"`
	IsFull          bool        `json:"is_full"`
	ManuallyFull    bool        `json:"manually_full"`
	ScheduledPosts  interface{} `json:"scheduled_posts,omitempty"`
}

// QueuePreviewResponse is returned by PreviewQueue.
type QueuePreviewResponse struct {
	Success bool               `json:"success"`
	Slots   []QueueSlotPreview `json:"slots,omitempty"`
}

// NextQueueSlotResponse is returned by GetNextQueueSlot.
type NextQueueSlotResponse struct {
	Success  bool    `json:"success"`
	NextSlot *string `json:"next_slot"` // null when no slot available within 30 days
}

// UpdateQueueSettingsOptions contains options for UpdateQueueSettings.
type UpdateQueueSettingsOptions struct {
	ProfileUsername string
	Timezone        string
	Slots           []QueueSlot
	DaysOfWeek      []int
	MaxPostsPerSlot *int
}

// ─────────────────────────────────────────────────────────────────────────────
// DM types
// ─────────────────────────────────────────────────────────────────────────────

// DMConversationParticipant holds info about one participant in a conversation.
type DMConversationParticipant struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// DMMessage is a recent message within a conversation.
type DMMessage struct {
	ID        string `json:"id,omitempty"`
	Text      string `json:"text,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	FromMe    bool   `json:"from_me"`
}

// DMConversation represents a single DM conversation thread.
type DMConversation struct {
	ID           string                      `json:"id,omitempty"`
	Participants []DMConversationParticipant `json:"participants,omitempty"`
	LastMessage  *DMMessage                  `json:"last_message,omitempty"`
	UpdatedAt    string                      `json:"updated_at,omitempty"`
}

// DMConversationsResponse is returned by ListDMConversations.
type DMConversationsResponse struct {
	Success       bool             `json:"success"`
	Conversations []DMConversation `json:"conversations,omitempty"`
}

// SendDMResponse is returned by SendDirectMessage.
type SendDMResponse struct {
	Success   bool   `json:"success"`
	MessageID string `json:"message_id,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     string `json:"error,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Reddit types
// ─────────────────────────────────────────────────────────────────────────────

// RedditPostMedia holds media details for a Reddit post.
type RedditPostMedia struct {
	Type     string `json:"type,omitempty"` // image | video | external_video
	URL      string `json:"url,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`
}

// RedditDetailedPost represents a single Reddit post with full details.
type RedditDetailedPost struct {
	ID          string            `json:"id,omitempty"`
	Title       string            `json:"title,omitempty"`
	URL         string            `json:"url,omitempty"`
	Subreddit   string            `json:"subreddit,omitempty"`
	Score       int               `json:"score,omitempty"`
	Impressions int               `json:"impressions,omitempty"` // view_count or score fallback
	CreatedAt   string            `json:"created_at,omitempty"`
	Media       []RedditPostMedia `json:"media,omitempty"`
}

// RedditDetailedPostsResponse is returned by ListRedditDetailedPosts.
type RedditDetailedPostsResponse struct {
	Success bool                 `json:"success"`
	Posts   []RedditDetailedPost `json:"posts,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Webhook event types
// ─────────────────────────────────────────────────────────────────────────────

// WebhookEventType represents the type of a webhook event.
type WebhookEventType string

// Webhook event type constants sent by the Upload-Post platform.
const (
	WebhookEventUploadCompleted             WebhookEventType = "upload_completed"
	WebhookEventSocialAccountConnected      WebhookEventType = "social_account_connected"
	WebhookEventSocialAccountDisconnected   WebhookEventType = "social_account_disconnected"
	WebhookEventSocialAccountReauthRequired WebhookEventType = "social_account_reauth_required"
)

// WebhookUploadResult holds the per-platform outcome of an upload_completed event.
type WebhookUploadResult struct {
	Success   bool   `json:"success"`
	URL       string `json:"url,omitempty"`
	PublishID string `json:"publish_id,omitempty"`
	Error     string `json:"error,omitempty"`
}

// WebhookUploadCompleted is the payload for the upload_completed event.
type WebhookUploadCompleted struct {
	Event           WebhookEventType    `json:"event"`
	JobID           string              `json:"job_id,omitempty"`
	UserEmail       string              `json:"user_email,omitempty"`
	ProfileUsername string              `json:"profile_username,omitempty"`
	Platform        string              `json:"platform,omitempty"`
	MediaType       string              `json:"media_type,omitempty"`
	Title           string              `json:"title,omitempty"`
	Caption         string              `json:"caption,omitempty"`
	Result          WebhookUploadResult `json:"result,omitempty"`
	CreatedAt       string              `json:"created_at,omitempty"`
}

// WebhookSocialAccountConnected is the payload for the social_account_connected event.
type WebhookSocialAccountConnected struct {
	Event           WebhookEventType `json:"event"`
	UserEmail       string           `json:"user_email,omitempty"`
	Platform        string           `json:"platform,omitempty"`
	AccountName     string           `json:"account_name,omitempty"`
	Status          string           `json:"status,omitempty"` // "connected"
	ProfileUsername string           `json:"profile_username,omitempty"`
	CreatedAt       string           `json:"created_at,omitempty"`
}

// WebhookSocialAccountDisconnected is the payload for social_account_disconnected.
type WebhookSocialAccountDisconnected struct {
	Event           WebhookEventType `json:"event"`
	UserEmail       string           `json:"user_email,omitempty"`
	Platform        string           `json:"platform,omitempty"`
	AccountName     string           `json:"account_name,omitempty"`
	Status          string           `json:"status,omitempty"`
	Reason          string           `json:"reason,omitempty"` // manual_disconnect | account_blocked | token_refresh_threshold_exceeded | max_auth_strikes
	ProfileUsername string           `json:"profile_username,omitempty"`
	CreatedAt       string           `json:"created_at,omitempty"`
}

// WebhookSocialAccountReauthRequired is the payload for social_account_reauth_required.
type WebhookSocialAccountReauthRequired struct {
	Event           WebhookEventType `json:"event"`
	UserEmail       string           `json:"user_email,omitempty"`
	Platform        string           `json:"platform,omitempty"`
	AccountName     string           `json:"account_name,omitempty"`
	Status          string           `json:"status,omitempty"` // "reauth_required"
	ProfileUsername string           `json:"profile_username,omitempty"`
	CreatedAt       string           `json:"created_at,omitempty"`
}
