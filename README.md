# upload-post-go-sdk

> **Unofficial** Go SDK for the [Upload-Post API](https://www.upload-post.com).  
> Created and maintained by **[Gokhan MERCANOGLU](https://github.com/gmercan)** as an open-source project.  
> This package is not affiliated with or endorsed by Upload-Post.

[![Go Reference](https://pkg.go.dev/badge/github.com/gmercan/uploadpost-go-sdk.svg)](https://pkg.go.dev/github.com/gmercan/uploadpost-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/gmercan/uploadpost-go-sdk)](https://goreportcard.com/report/github.com/gmercan/uploadpost-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Upload videos, photos, text posts, and documents to **TikTok, Instagram, YouTube, LinkedIn, Facebook, Pinterest, Threads, Reddit, Bluesky, and X (Twitter)** with a single API call — no external dependencies, pure Go stdlib.

---

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Authentication](#authentication)
- [Upload Methods](#upload-methods)
  - [Upload Video](#upload-video)
  - [Upload Photos](#upload-photos)
  - [Upload Text Post](#upload-text-post)
  - [Upload Document (LinkedIn)](#upload-document-linkedin)
- [Status & History](#status--history)
- [Scheduling](#scheduling)
- [Queue Management](#queue-management)
- [Analytics](#analytics)
- [Media](#media)
- [User Management](#user-management)
- [Comments & DMs](#comments--dms)
- [AutoDM Monitors](#autodm-monitors)
- [Webhook Events](#webhook-events)
- [Helper Endpoints](#helper-endpoints)
- [Error Handling](#error-handling)
- [Platform-Specific Options Reference](#platform-specific-options-reference)
- [Links](#links)
- [License](#license)

---

## Features

- **Video Upload** — TikTok, Instagram, YouTube, LinkedIn, Facebook, Pinterest, Threads, Bluesky, X, Google Business
- **Photo Upload** — TikTok, Instagram, LinkedIn, Facebook, Pinterest, Threads, Reddit, Bluesky, X, Google Business
- **Text Posts** — X, LinkedIn, Facebook, Threads, Reddit, Bluesky, Google Business
- **Document Upload** — LinkedIn (PDF, PPT, PPTX, DOC, DOCX)
- **Scheduling** — Schedule posts to any future date with timezone support
- **Posting Queue** — Add posts to your configured queue slots
- **First Comments** — Auto-post a first comment right after publishing
- **Analytics** — Profile-level and per-post engagement metrics across all platforms
- **Queue Management** — Full control over slot settings, previews, and manual overrides
- **AutoDM** — Automatically DM Instagram commenters with keyword filtering
- **Instagram Comments & DMs** — Read comments, send private and public replies
- **User Management** — Create and manage profiles; generate JWTs for white-label integration
- **Webhook Parsing** — Typed parsers for all Upload-Post webhook event payloads
- **Zero dependencies** — only Go standard library

---

## Requirements

- Go 1.21 or later
- An [Upload-Post](https://app.upload-post.com) account and API key

---

## Installation

```bash
go get github.com/gmercan/uploadpost-go-sdk
```

---

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    uploadpost "github.com/gmercan/uploadpost-go-sdk"
)

func main() {
    client := uploadpost.New("YOUR_API_KEY")

    resp, err := client.UploadVideo(context.Background(), "./video.mp4", uploadpost.VideoOptions{
        User:      "my-profile",
        Platforms: []string{"tiktok", "instagram", "youtube"},
        Title:     "Check out this video!",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Request ID:", resp.RequestID)
}
```

Get your API key from the [Upload-Post dashboard](https://app.upload-post.com).

---

## Authentication

Every request is authenticated with your API key via the `Authorization: Apikey <key>` header — the SDK handles this automatically.

The only exception is `ValidateJWT`, which uses `Authorization: Bearer <jwt>` and is also handled transparently by the SDK.

```go
// Standard client
client := uploadpost.New("YOUR_API_KEY")

// With custom HTTP client or base URL override
client := uploadpost.NewWithOptions("YOUR_API_KEY", uploadpost.ClientOptions{
    HTTPClient: myHttpClient,
    BaseURL:    "https://api.upload-post.com/api", // default
})
```

---

## Upload Methods

### Upload Video

Upload a local file or a remote video URL to one or more platforms simultaneously.

```go
ctx := context.Background()

resp, err := client.UploadVideo(ctx, "./video.mp4", uploadpost.VideoOptions{
    // Required
    User:      "my-profile",
    Platforms: []string{"tiktok", "instagram", "youtube", "linkedin"},
    Title:     "My awesome video 🎬",

    // Scheduling — post at a specific time
    ScheduledDate: "2025-01-15T10:00:00Z",
    Timezone:      "Europe/Istanbul",

    // Or add to your queue instead
    // AddToQueue: uploadpost.Bool(true),

    // First comment posted right after publishing
    FirstComment: "Thanks for watching! Drop a ❤️ below.",

    // Async upload (default: true) — returns request_id immediately
    AsyncUpload: uploadpost.Bool(true),

    // ── TikTok ──────────────────────────────────────────────────────────────
    TiktokPrivacyLevel:   "PUBLIC_TO_EVERYONE",
    TiktokDisableComment: uploadpost.Bool(false),
    TiktokDisableDuet:    uploadpost.Bool(false),
    TiktokDisableStitch:  uploadpost.Bool(false),
    TiktokIsAIGC:         uploadpost.Bool(false), // AI-generated content flag

    // ── Instagram ───────────────────────────────────────────────────────────
    InstagramMediaType:   "REELS",
    InstagramShareToFeed: uploadpost.Bool(true),
    InstagramCoverURL:    "./cover.jpg", // local file or URL

    // ── YouTube ─────────────────────────────────────────────────────────────
    YouTubePrivacyStatus:  "public",
    YouTubeTags:           []string{"golang", "tutorial", "api"},
    YouTubeCategoryID:     "22", // People & Blogs
    YouTubeEmbeddable:     uploadpost.Bool(true),
    YouTubeThumbnailURL:   "https://example.com/thumb.jpg",

    // YouTube subtitles
    YouTubeSubtitles: []uploadpost.YouTubeSubtitle{
        {Language: "en", Name: "English", File: "./subtitles_en.srt"},
        {Language: "es", Name: "Español", URL: "https://example.com/subs_es.vtt"},
    },

    // ── LinkedIn ────────────────────────────────────────────────────────────
    LinkedinVisibility:   "PUBLIC",
    TargetLinkedinPageID: "company-page-id", // omit for personal profile

    // ── Per-platform title overrides ────────────────────────────────────────
    TiktokTitle:    "TikTok-specific caption #fyp",
    YoutubeTitle:   "Full YouTube Title — Much Longer Is Fine Here",
    LinkedinTitle:  "Professional LinkedIn headline",

    // ── Per-platform first comment overrides ────────────────────────────────
    YoutubeFirstComment: "Subscribe for more! 🔔",
    InstagramFirstComment: "Save this for later 📌",
})
if err != nil {
    log.Fatal(err)
}
fmt.Println("Upload queued, request_id:", resp.RequestID)
```

**Upload from URL instead of a local file:**

```go
resp, err := client.UploadVideo(ctx, "https://example.com/video.mp4", uploadpost.VideoOptions{
    User:      "my-profile",
    Platforms: []string{"tiktok"},
    Title:     "Remote video",
})
```

---

### Upload Photos

Upload one or more photos (carousel) or a single image to multiple platforms.

```go
resp, err := client.UploadPhotos(ctx,
    []string{
        "./photo1.jpg",
        "./photo2.jpg",
        "https://example.com/photo3.jpg", // mix of local and remote is fine
    },
    uploadpost.PhotosOptions{
        User:      "my-profile",
        Platforms: []string{"instagram", "facebook", "x", "linkedin", "tiktok"},
        Title:     "Behind the scenes 📸",
        AltText:   "Team working on the new project",

        // Queue instead of immediate post
        AddToQueue:      uploadpost.Bool(true),
        MaxPostsPerSlot: uploadpost.Int(2),

        // ── Instagram ───────────────────────────────────────────────────────
        InstagramMediaType:     "IMAGE", // IMAGE or STORIES
        InstagramCollaborators: "friend_username,partner_username",

        // ── TikTok ──────────────────────────────────────────────────────────
        TiktokAutoAddMusic:    uploadpost.Bool(true),
        TiktokPhotoCoverIndex: uploadpost.Int(0), // 0-based cover photo index

        // ── Facebook ────────────────────────────────────────────────────────
        FacebookPageID: "123456789",

        // ── Pinterest ───────────────────────────────────────────────────────
        PinterestBoardID: "board-id-here",
        PinterestLink:    "https://yourwebsite.com",
        PinterestAltText: "Behind the scenes photo",

        // ── Reddit ──────────────────────────────────────────────────────────
        RedditSubreddit: "golang", // without r/
        RedditFlairID:   "flair-template-id",

        // ── X (Twitter) ─────────────────────────────────────────────────────
        XReplySettings:     "everyone",
        XTaggedUserIDs:     []string{"user_id_1", "user_id_2"},
        XThreadImageLayout: "2,2", // 2 images per tweet in a thread

        // ── Threads ─────────────────────────────────────────────────────────
        ThreadsThreadMediaLayout: "5,5", // split 10 photos into 2 posts of 5
        ThreadsTopicTag:          "photography",
    },
)
```

---

### Upload Text Post

Post text-only content to X, LinkedIn, Facebook, Threads, Reddit, Bluesky, and Google Business.

```go
// Simple text post
resp, err := client.UploadText(ctx, uploadpost.TextOptions{
    User:      "my-profile",
    Platforms: []string{"x", "linkedin", "threads", "bluesky"},
    Title:     "Just shipped something cool. Check the link 👇",
    LinkURL:   "https://yourwebsite.com", // generic link preview (LinkedIn, Bluesky, Facebook)
})

// X poll
resp, err = client.UploadText(ctx, uploadpost.TextOptions{
    User:      "my-profile",
    Platforms: []string{"x"},
    Title:     "What's your favorite Go web framework?",

    XPollOptions:       []string{"Gin", "Echo", "Fiber", "Chi"},
    XPollDuration:      uploadpost.Int(1440), // 24 hours in minutes
    XPollReplySettings: "everyone",
})

// Quote tweet
resp, err = client.UploadText(ctx, uploadpost.TextOptions{
    User:          "my-profile",
    Platforms:     []string{"x"},
    Title:         "This is a great point 👆",
    XQuoteTweetID: "1234567890123456789",
})

// Reddit link post
resp, err = client.UploadText(ctx, uploadpost.TextOptions{
    User:            "my-profile",
    Platforms:       []string{"reddit"},
    Title:           "Check out this Go SDK for Upload-Post",
    RedditSubreddit: "golang",
    RedditLinkURL:   "https://github.com/gmercan/uploadpost-go-sdk",
})

// LinkedIn company page post with link preview
resp, err = client.UploadText(ctx, uploadpost.TextOptions{
    User:                 "my-profile",
    Platforms:            []string{"linkedin"},
    Title:                "Excited to announce our new product launch!",
    TargetLinkedinPageID: "company-page-id",
    LinkedinLinkURL:      "https://yourproduct.com",
    LinkedinVisibility:   "PUBLIC",
})

// Scheduled post to multiple platforms
resp, err = client.UploadText(ctx, uploadpost.TextOptions{
    User:          "my-profile",
    Platforms:     []string{"x", "linkedin", "facebook", "threads"},
    Title:         "Good morning! Here is your daily tip 🌅",
    ScheduledDate: "2025-06-10T08:00:00Z",
    Timezone:      "America/New_York",

    // Different first comments per platform
    XFirstComment:       "RT if this helped!",
    LinkedinFirstComment: "Follow for more tips.",
})
```

---

### Upload Document (LinkedIn)

Upload PDF, PPT, PPTX, DOC, or DOCX as a native LinkedIn document post.

```go
resp, err := client.UploadDocument(ctx, "./Q2_2025_Report.pdf", uploadpost.DocumentOptions{
    User:                 "my-profile",
    Title:                "Q2 2025 Company Report",
    Description:          "Highlights from our second quarter — record growth across all segments.",
    LinkedinVisibility:   "PUBLIC",
    TargetLinkedinPageID: "company-page-id", // omit for personal profile
    ScheduledDate:        "2025-07-01T09:00:00Z",
    Timezone:             "Europe/London",
})
```

---

## Status & History

### Check async upload status

```go
// After an async upload, poll with the returned request_id
status, err := client.GetStatus(ctx, resp.RequestID)
fmt.Printf("Status: %s\n", status.Status) // pending | processing | completed | failed

// For scheduled or queued posts, use the job_id
status, err = client.GetJobStatus(ctx, resp.JobID)
```

### Upload history

```go
history, err := client.GetHistory(ctx, uploadpost.HistoryOptions{
    Page:  1,
    Limit: 50, // 20 | 50 | 100
})
for _, item := range history.Uploads {
    fmt.Printf("[%s] %s — %s\n", item.CreatedAt, item.RequestID, item.Status)
}
```

---

## Scheduling

```go
// List all scheduled posts
scheduled, err := client.ListScheduled(ctx)
for _, s := range scheduled.Scheduled {
    fmt.Printf("Job %s scheduled for %s\n", s.JobID, s.ScheduledDate)
}

// Reschedule a post
_, err = client.EditScheduled(ctx, "job-id-here", uploadpost.EditScheduledOptions{
    ScheduledDate: "2025-08-20T15:00:00Z",
    Timezone:      "Asia/Tokyo",
})

// Cancel a scheduled post
_, err = client.CancelScheduled(ctx, "job-id-here")
```

---

## Queue Management

```go
// Get queue settings for a profile
settings, err := client.GetQueueSettings(ctx, "my-profile")
fmt.Printf("Timezone: %s, Slots: %d\n", settings.Settings.Timezone, len(settings.Settings.Slots))

// Update queue settings
// DaysOfWeek: 0=Monday … 6=Sunday
_, err = client.UpdateQueueSettings(ctx, uploadpost.UpdateQueueSettingsOptions{
    ProfileUsername: "my-profile",
    Timezone:        "Europe/Istanbul",
    DaysOfWeek:      []int{0, 1, 2, 3, 4}, // Monday–Friday
    Slots: []uploadpost.QueueSlot{
        {Hour: 9, Minute: 0},
        {Hour: 13, Minute: 30},
        {Hour: 18, Minute: 0},
    },
    MaxPostsPerSlot: uploadpost.Int(1),
})

// Preview upcoming queue slots
preview, err := client.PreviewQueue(ctx, "my-profile", 10)
for _, slot := range preview.Slots {
    fmt.Printf("Slot %s — %d/%d posts (full: %v)\n",
        slot.SlotDatetime, slot.PostCount, slot.MaxPostsPerSlot, slot.IsFull)
}

// Get the next available queue slot
next, err := client.GetNextQueueSlot(ctx, "my-profile")
if next.NextSlot != nil {
    fmt.Println("Next slot:", *next.NextSlot)
} else {
    fmt.Println("No available slot in the next 30 days")
}

// Manually mark a slot as full
_, err = client.MarkQueueSlotFull(ctx, "my-profile", "2025-06-10T09:00:00Z")

// Unmark it
_, err = client.UnmarkQueueSlotFull(ctx, "my-profile", "2025-06-10T09:00:00Z")
```

---

## Analytics

### Profile analytics

```go
analytics, err := client.GetAnalytics(ctx, "my-profile", uploadpost.AnalyticsOptions{
    Platforms: []string{"instagram", "tiktok", "youtube"},
    PageID:    "fb-page-id",  // required for Facebook analytics
    PageURN:   "li-page-urn", // optional, for LinkedIn org page
})
```

### Total impressions over a period

```go
// Using a period shortcut
impressions, err := client.GetTotalImpressions(ctx, "my-profile", uploadpost.ImpressionsOptions{
    Period:    "last_month", // last_day | last_week | last_month | last_3months | last_year
    Breakdown: true,         // include per-platform and per-day breakdown
})

// Using a custom date range with specific metrics
impressions, err = client.GetTotalImpressions(ctx, "my-profile", uploadpost.ImpressionsOptions{
    StartDate: "2025-01-01",
    EndDate:   "2025-03-31",
    Platforms: []string{"instagram", "tiktok"},
    Metrics:   []string{"likes", "comments", "shares"},
})
fmt.Println("Total impressions:", *impressions.TotalImpressions)
```

### Per-post analytics

```go
// By Upload-Post request_id
postAnalytics, err := client.GetPostAnalytics(ctx, "request-id-here")
for platform, data := range postAnalytics.Platforms {
    fmt.Printf("%s: %+v\n", platform, data.PostMetrics)
}

// By native platform post ID (works for organic posts too)
postAnalytics, err = client.GetPostAnalyticsByPlatformID(ctx,
    "instagram-media-id",
    "instagram",
    "my-profile",
)

// Available metrics per platform
metrics, err := client.GetPlatformMetrics(ctx)
for platform, cfg := range metrics {
    fmt.Printf("%s primary field: %s\n", platform, cfg.PrimaryImpressionsField)
}
```

---

## Media

```go
// Recent posts from any connected platform
media, err := client.GetMedia(ctx, "instagram", "my-profile", "")
for _, item := range media.Media {
    fmt.Printf("[%s] %s — %s\n", item.MediaType, item.ID, item.Permalink)
}

// LinkedIn personal profile
media, err = client.GetMedia(ctx, "linkedin", "my-profile", "me")

// LinkedIn organization page
media, err = client.GetMedia(ctx, "linkedin", "my-profile", "12345678")

// Reddit detailed posts (with media info, up to 2000 posts)
posts, err := client.ListRedditDetailedPosts(ctx, "my-profile")
for _, p := range posts.Posts {
    fmt.Printf("r/%s — %s (%d impressions)\n", p.Subreddit, p.Title, p.Impressions)
}
```

---

## User Management

```go
// List all profiles under your API key
users, err := client.ListUsers(ctx)
fmt.Printf("Plan: %s, Profiles: %d\n", users.Profiles[0].Username, len(users.Profiles))

// Get a single profile
profile, err := client.GetUserProfile(ctx, "my-profile")
fmt.Println(profile.Username, profile.CreatedAt)

// Validate your API key and get account info
me, err := client.GetCurrentUser(ctx)
fmt.Printf("Email: %s, Plan: %s\n", me.Email, me.Plan)

// Create and delete profiles
_, err = client.CreateUser(ctx, "new-profile")
_, err = client.DeleteUser(ctx, "old-profile")

// Generate a JWT link for white-label social account connection
jwt, err := client.GenerateJWT(ctx, "my-profile", uploadpost.JWTOptions{
    RedirectURL:        "https://yourapp.com/connected",
    LogoImage:          "https://yourapp.com/logo.png",
    ConnectTitle:       "Connect your social accounts",
    ConnectDescription: "Link your TikTok, Instagram and YouTube to start posting.",
    Platforms:          []string{"tiktok", "instagram", "youtube"},
    ShowCalendar:       uploadpost.Bool(true),
    Language:           "en", // en | es | de | fr | pt
})
fmt.Println("Share this URL with your user:", jwt.ConnectionURL)

// Validate a JWT (uses Bearer auth automatically)
valid, err := client.ValidateJWT(ctx, "the-jwt-token")
fmt.Println("Valid:", valid.Success)

// User preferences
prefs, err := client.GetUserPreferences(ctx)
_, err = client.UpdateUserPreferences(ctx, uploadpost.UserPreferencesOptions{
    WeekStartDay: uploadpost.Int(1), // 0=Sunday, 1=Monday
})

// Webhook / notification config
config, err := client.GetNotificationConfig(ctx)
fmt.Println("Webhook URL:", config.WebhookURL)

_, err = client.UpdateNotificationConfig(ctx, uploadpost.NotificationConfigOptions{
    WebhookURL: "https://yourapp.com/webhooks/uploadpost",
    WebhookEvents: []string{
        "upload_completed",
        "social_account_connected",
        "social_account_disconnected",
        "social_account_reauth_required",
    },
})
```

---

## Comments & DMs

### Instagram comments

```go
// Get comments by post URL
comments, err := client.GetPostComments(ctx, "my-profile", "", "https://www.instagram.com/p/ABC123/")

// Or by numeric media ID
comments, err = client.GetPostComments(ctx, "my-profile", "17854360229135492", "")

for _, c := range comments.Comments {
    fmt.Printf("@%s: %s\n", c.User.Username, c.Text)
}

// Send a private DM reply (7-day comment window applies)
_, err = client.ReplyToComment(ctx, "my-profile", comments.Comments[0].ID, "Thanks for your comment! 🙌")

// Post a public reply visible under the comment
_, err = client.PublicReplyToComment(ctx, "my-profile", comments.Comments[0].ID, "Great point! 👏")
```

### Direct messages

```go
// List DM conversations
convs, err := client.ListDMConversations(ctx, "instagram", "my-profile")
for _, conv := range convs.Conversations {
    fmt.Printf("Conversation %s with %d participants\n", conv.ID, len(conv.Participants))
}

// Send a DM (recipient_id from comments or conversations)
_, err = client.SendDirectMessage(ctx, "instagram", "my-profile", "recipient-user-id", "Hey, thanks for reaching out!")
```

---

## AutoDM Monitors

AutoDM monitors watch your Instagram posts 24/7 and automatically send private DMs to new commenters — with optional keyword filtering.

**Limits:** 2 new monitors per profile per day · monitors auto-expire after 15 days · 500 DMs/day on paid plans.

```go
// Start a monitor — DM everyone who comments
_, err := client.StartAutoDM(ctx, uploadpost.AutoDMOptions{
    PostURL:            "https://www.instagram.com/p/ABC123/",
    ReplyMessage:       "Thanks for commenting! Here is the free guide: https://example.com/guide",
    ProfileUsername:    "my-profile",
    MonitoringInterval: uploadpost.Int(15), // minutes, minimum 15
})

// Start with keyword filter — only DM comments containing these words
_, err = client.StartAutoDM(ctx, uploadpost.AutoDMOptions{
    PostURL:            "https://www.instagram.com/p/ABC123/",
    ReplyMessage:       "Here is the link you asked for: https://example.com",
    ProfileUsername:    "my-profile",
    TriggerKeywords:    []string{"link", "guide", "more info"},
})

// List active monitors
status, err := client.GetAutoDMStatus(ctx, false)
fmt.Printf("Active monitors: %d\n", status.TotalActive)

// Include stopped/expired monitors
status, err = client.GetAutoDMStatus(ctx, true)
for _, m := range status.Monitors {
    fmt.Printf("[%s] %s — %s\n", m.Status, m.MonitorID, m.PostURL)
}

// Get activity logs for a monitor
logs, err := client.GetAutoDMLogs(ctx, "monitor-id-here")
for _, l := range logs.Logs {
    fmt.Printf("%s: %s\n", l.Timestamp, l.Event)
}

// Pause / resume (data is preserved)
_, err = client.PauseAutoDM(ctx, "monitor-id-here")
_, err = client.ResumeAutoDM(ctx, "monitor-id-here")

// Stop (data preserved, visible with includeInactive=true)
_, err = client.StopAutoDM(ctx, "monitor-id-here")

// Delete permanently
_, err = client.DeleteAutoDM(ctx, "monitor-id-here")
```

---

## Webhook Events

Set up a webhook URL via `UpdateNotificationConfig` and use `ParseWebhookEvent` to parse inbound payloads with full type safety.

```go
import "net/http"

http.HandleFunc("/webhooks/uploadpost", func(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)

    evt, err := uploadpost.ParseWebhookEvent(body)
    if err != nil {
        http.Error(w, "bad payload", 400)
        return
    }

    switch e := evt.(type) {
    case *uploadpost.WebhookUploadCompleted:
        fmt.Printf("Upload done — platform: %s, success: %v, url: %s\n",
            e.Platform, e.Result.Success, e.Result.URL)

    case *uploadpost.WebhookSocialAccountConnected:
        fmt.Printf("Account connected — @%s on %s\n", e.AccountName, e.Platform)

    case *uploadpost.WebhookSocialAccountDisconnected:
        fmt.Printf("Account disconnected — @%s on %s (reason: %s)\n",
            e.AccountName, e.Platform, e.Reason)
        // Prompt re-auth: generate a JWT for the user
        // jwt, _ := client.GenerateJWT(ctx, e.ProfileUsername, uploadpost.JWTOptions{})

    case *uploadpost.WebhookSocialAccountReauthRequired:
        fmt.Printf("Reauth required — @%s needs to reconnect %s\n",
            e.ProfileUsername, e.Platform)

    default:
        // Unknown event type — forward to logging system
        fmt.Printf("Unknown webhook event: %+v\n", e)
    }

    w.WriteHeader(http.StatusOK)
})
```

---

## Helper Endpoints

```go
// Facebook pages linked to a profile (needed for facebook_page_id)
pages, err := client.GetFacebookPages(ctx, "my-profile")
for _, p := range pages.Pages {
    fmt.Printf("Page: %s (%s)\n", p.Name, p.ID)
}

// LinkedIn organization pages
liPages, err := client.GetLinkedinPages(ctx, "my-profile")

// Pinterest boards
boards, err := client.GetPinterestBoards(ctx, "my-profile")
for _, b := range boards.Boards {
    fmt.Printf("Board: %s (%s)\n", b.Name, b.ID)
}

// Google Business Profile locations
locations, err := client.GetGoogleBusinessLocations(ctx, "my-profile")
for _, l := range locations.Locations {
    fmt.Printf("Location: %s — ID: %s\n", l.DisplayName, l.LocationID)
}

// Select a GBP location for the profile
_, err = client.SelectGoogleBusinessLocation(ctx, "accounts/123/locations/456", "my-profile")
```

---

## Error Handling

```go
import "errors"

resp, err := client.UploadVideo(ctx, "./video.mp4", opts)
if err != nil {
    var apiErr *uploadpost.APIError
    if errors.As(err, &apiErr) {
        switch apiErr.StatusCode {
        case 400:
            fmt.Println("Bad request:", apiErr.Message)
        case 401:
            fmt.Println("Invalid or missing API key")
        case 403:
            fmt.Println("Forbidden:", apiErr.Message)
        case 404:
            fmt.Println("Not found:", apiErr.Message)
        case 429:
            fmt.Println("Rate limit reached — slow down or upgrade your plan")
        case 500:
            fmt.Println("Upload-Post server error, try again later")
        default:
            fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
        }
        return
    }
    // Network or serialization error
    log.Fatal("Unexpected error:", err)
}
```

---

## Platform-Specific Options Reference

### Helper functions

Use `uploadpost.Bool(v)` and `uploadpost.Int(v)` to pass optional fields that have a meaningful zero value:

```go
opts := uploadpost.VideoOptions{
    AsyncUpload:    uploadpost.Bool(false),  // explicit false, not "omit"
    YouTubeEmbeddable: uploadpost.Bool(true),
    MaxPostsPerSlot:   uploadpost.Int(3),
}
```

### TikTok (Video)

| Field | Type | Description |
|-------|------|-------------|
| `TiktokPrivacyLevel` | `string` | `PUBLIC_TO_EVERYONE` · `MUTUAL_FOLLOW_FRIENDS` · `FOLLOWER_OF_CREATOR` · `SELF_ONLY` |
| `TiktokDisableDuet` | `*bool` | Disable duet feature |
| `TiktokDisableComment` | `*bool` | Disable comments |
| `TiktokDisableStitch` | `*bool` | Disable stitch feature |
| `TiktokCoverTimestamp` | `*int64` | Cover frame timestamp in milliseconds |
| `TiktokIsAIGC` | `*bool` | Mark as AI-generated content |
| `TiktokPostMode` | `string` | `DIRECT_POST` · `MEDIA_UPLOAD` |
| `BrandContentToggle` | `*bool` | Branded content disclosure |
| `BrandOrganicToggle` | `*bool` | Brand organic disclosure |

### TikTok (Photos)

| Field | Type | Description |
|-------|------|-------------|
| `TiktokAutoAddMusic` | `*bool` | Auto-add background music |
| `TiktokPhotoCoverIndex` | `*int` | 0-based index of the cover photo |
| `TiktokDisableComment` | `*bool` | Disable comments |

### Instagram

| Field | Type | Description |
|-------|------|-------------|
| `InstagramMediaType` | `string` | Video: `REELS` · `STORIES` / Photo: `IMAGE` · `STORIES` |
| `InstagramShareToFeed` | `*bool` | Share Reels/Stories to main feed |
| `InstagramCollaborators` | `string` | Comma-separated collaborator usernames |
| `InstagramCoverURL` | `string` | Cover image: URL or local file path |
| `InstagramAudioName` | `string` | Audio track name |
| `InstagramUserTags` | `string` | Comma-separated tagged usernames |
| `InstagramLocationID` | `string` | Instagram location ID |
| `InstagramThumbOffset` | `string` | Thumbnail offset |

### YouTube

| Field | Type | Description |
|-------|------|-------------|
| `YouTubePrivacyStatus` | `string` | `public` · `unlisted` · `private` |
| `YouTubeTags` | `[]string` | Video tags |
| `YouTubeCategoryID` | `string` | Category ID (e.g. `"22"` = People & Blogs) |
| `YouTubeEmbeddable` | `*bool` | Allow embedding on external sites |
| `YouTubeLicense` | `string` | `youtube` · `creativeCommon` |
| `YouTubeThumbnailURL` | `string` | Custom thumbnail URL |
| `YouTubeSelfDeclaredMadeForKids` | `*bool` | COPPA — made for kids flag |
| `YouTubeContainsSyntheticMedia` | `*bool` | AI/synthetic content disclosure |
| `YouTubeDefaultLanguage` | `string` | BCP-47 language code for title/description |
| `YouTubeDefaultAudioLanguage` | `string` | BCP-47 audio language code |
| `YouTubeSubtitles` | `[]YouTubeSubtitle` | Subtitle tracks (file path or URL per language) |
| `YouTubeAllowedCountries` | `string` | Comma-separated allowed country codes |
| `YouTubeBlockedCountries` | `string` | Comma-separated blocked country codes |
| `YouTubeRecordingDate` | `string` | ISO 8601 recording date |
| `YouTubeHasPaidProductPlacement` | `*bool` | Paid product placement disclosure |

### LinkedIn

| Field | Type | Description |
|-------|------|-------------|
| `LinkedinVisibility` | `string` | `PUBLIC` · `CONNECTIONS` · `LOGGED_IN` · `CONTAINER` |
| `TargetLinkedinPageID` | `string` | Numeric org page ID (omit for personal profile) |

### Facebook

| Field | Type | Description |
|-------|------|-------------|
| `FacebookPageID` | `string` | Facebook Page ID (required for page posts) |
| `FacebookVideoState` | `string` | `PUBLISHED` · `DRAFT` |
| `FacebookMediaType` | `string` | `REELS` · `STORIES` · `VIDEO` |
| `ThumbnailURL` | `string` | Thumbnail URL (only for `VIDEO` type) |

### Pinterest

| Field | Type | Description |
|-------|------|-------------|
| `PinterestBoardID` | `string` | Target board ID |
| `PinterestLink` | `string` | Destination URL |
| `PinterestAltText` | `string` | Alt text for accessibility |
| `PinterestCoverImageURL` | `string` | Cover image URL (video only) |
| `PinterestCoverImageKeyFrameTime` | `*int` | Key frame time in milliseconds |

### X (Twitter)

| Field | Type | Description |
|-------|------|-------------|
| `XReplySettings` | `string` | `everyone` · `following` · `mentionedUsers` · `subscribers` · `verified` |
| `XNullcast` | `*bool` | Promoted-only post (not shown on timeline) |
| `XTaggedUserIDs` | `[]string` | User IDs to tag in media |
| `XQuoteTweetID` | `string` | Tweet ID to quote |
| `XPollOptions` | `[]string` | Poll choices (2–4 items) |
| `XPollDuration` | `*int` | Poll duration in minutes (5–10080) |
| `XCommunityID` | `string` | Post to a specific community |
| `XShareWithFollowers` | `*bool` | Also share community post with followers |
| `XLongTextAsPost` | `*bool` | Publish long text as a single post |
| `XThreadImageLayout` | `string` | Images per tweet in a thread, e.g. `"4,4"` |
| `XForSuperFollowersOnly` | `*bool` | Exclusive content for super followers |

### Threads

| Field | Type | Description |
|-------|------|-------------|
| `ThreadsLongTextAsPost` | `*bool` | Post long text as a single post (not a thread) |
| `ThreadsThreadMediaLayout` | `string` | Media split, e.g. `"5,5"` for 2 posts of 5 each |
| `ThreadsTopicTag` | `string` | Topic tag (1–50 chars, no periods or ampersands) |

### Reddit

| Field | Type | Description |
|-------|------|-------------|
| `RedditSubreddit` | `string` | Subreddit name without `r/` |
| `RedditFlairID` | `string` | Flair template ID |

---

## Links

| Resource | URL |
|----------|-----|
| Upload-Post Website | https://www.upload-post.com |
| API Documentation | https://docs.upload-post.com |
| Dashboard | https://app.upload-post.com |
| npm SDK (official) | https://github.com/Upload-Post/upload-post-npm |

---

## License

MIT © [GokhanMERCANOGLU](https://github.com/gmercan)

See [LICENSE](LICENSE) for the full text.

> **Disclaimer:** This is an unofficial, community-maintained SDK. It is not affiliated with, officially connected to, or endorsed by Upload-Post.