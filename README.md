# Upload-Post SDK for Go

Official Go client for the [Upload-Post API](https://www.upload-post.com) — cross-platform social media publishing.

Upload videos, photos, text posts, and documents to **TikTok, Instagram, YouTube, LinkedIn, Facebook, Pinterest, Threads, Reddit, Bluesky, and X (Twitter)** with a single API call.

## Installation

```bash
go get github.com/Upload-Post/upload-post-go-sdk
```

Requires Go 1.21+. No external dependencies — stdlib only.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    uploadpost "github.com/Upload-Post/upload-post-go-sdk"
)

func main() {
    client := uploadpost.New("YOUR_API_KEY")
    ctx := context.Background()

    resp, err := client.UploadVideo(ctx, "./video.mp4", uploadpost.VideoOptions{
        User:      "my-profile",
        Platforms: []string{"tiktok", "instagram", "youtube"},
        Title:     "Check out this video!",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp.RequestID)
}
```

## Features

- **Video Upload** — TikTok, Instagram, YouTube, LinkedIn, Facebook, Pinterest, Threads, Bluesky, X
- **Photo Upload** — TikTok, Instagram, LinkedIn, Facebook, Pinterest, Threads, Reddit, Bluesky, X
- **Text Posts** — X, LinkedIn, Facebook, Threads, Reddit, Bluesky
- **Document Upload** — LinkedIn (PDF, PPT, PPTX, DOC, DOCX)
- **Scheduling** — Schedule posts for any future date
- **Posting Queue** — Add posts to your configured queue
- **First Comments** — Auto-post a first comment after publishing
- **Analytics** — Profile and post-level engagement metrics
- **AutoDM** — Automatically DM commenters on Instagram posts
- **User Management** — Create, list, and delete profiles; generate JWTs for white-label integration
- **No external dependencies** — pure stdlib

## API Reference

### Upload Video

```go
resp, err := client.UploadVideo(ctx, "./video.mp4", uploadpost.VideoOptions{
    User:      "my-profile",
    Platforms: []string{"tiktok", "instagram", "youtube"},
    Title:     "My awesome video",

    // Scheduling (optional)
    ScheduledDate: "2024-12-25T10:00:00Z",
    Timezone:      "Europe/Madrid",

    // First comment (optional)
    FirstComment: "Thanks for watching!",

    // TikTok
    TiktokPrivacyLevel: "PUBLIC_TO_EVERYONE",

    // Instagram
    InstagramMediaType: "REELS",

    // YouTube
    YouTubePrivacyStatus: "public",
    YouTubeTags:          []string{"tutorial", "coding"},
    YouTubeCategoryID:    "22",

    // X poll not applicable for video; see UploadText
})
```

### Upload Photos

```go
resp, err := client.UploadPhotos(ctx,
    []string{"./photo1.jpg", "https://example.com/photo2.jpg"},
    uploadpost.PhotosOptions{
        User:               "my-profile",
        Platforms:          []string{"instagram", "facebook", "x"},
        Title:              "Check out these photos!",
        AddToQueue:         uploadpost.Bool(true),
        InstagramMediaType: "IMAGE",
    },
)
```

### Upload Text Post

```go
resp, err := client.UploadText(ctx, uploadpost.TextOptions{
    User:      "my-profile",
    Platforms: []string{"x", "linkedin", "threads"},
    Title:     "Just shipped a new feature!",

    // X poll
    XPollOptions:  []string{"Option A", "Option B"},
    XPollDuration: uploadpost.Int(1440), // 24 hours in minutes

    // LinkedIn company page
    TargetLinkedinPageID: "company-page-id",
})
```

### Upload Document (LinkedIn)

```go
resp, err := client.UploadDocument(ctx, "./presentation.pdf", uploadpost.DocumentOptions{
    User:                 "my-profile",
    Title:                "Q4 2024 Report",
    Description:          "Our latest quarterly results",
    LinkedinVisibility:   "PUBLIC",
    TargetLinkedinPageID: "company-page-id",
})
```

### Check Upload Status

```go
// Async upload status by request_id
status, err := client.GetStatus(ctx, resp.RequestID)

// Scheduled/queued post status by job_id
status, err := client.GetJobStatus(ctx, resp.JobID)
```

### Upload History

```go
history, err := client.GetHistory(ctx, uploadpost.HistoryOptions{Page: 1, Limit: 20})
for _, item := range history.Uploads {
    fmt.Println(item.RequestID, item.Status)
}
```

### Scheduled Posts

```go
// List
scheduled, err := client.ListScheduled(ctx)

// Reschedule
_, err = client.EditScheduled(ctx, "job-id", uploadpost.EditScheduledOptions{
    ScheduledDate: "2024-12-26T15:00:00Z",
    Timezone:      "America/New_York",
})

// Cancel
_, err = client.CancelScheduled(ctx, "job-id")
```

### Analytics

```go
// Profile analytics
analytics, err := client.GetAnalytics(ctx, "my-profile", uploadpost.AnalyticsOptions{
    Platforms: []string{"instagram", "tiktok"},
})

// Total impressions over a period
impressions, err := client.GetTotalImpressions(ctx, "my-profile", uploadpost.ImpressionsOptions{
    Period:    "last_month",
    Breakdown: true,
})

// Per-post analytics
postAnalytics, err := client.GetPostAnalytics(ctx, "request-id")
```

### Media

```go
// Recent posts from a connected account
media, err := client.GetMedia(ctx, "instagram", "my-profile", "")

// LinkedIn organisation page
media, err := client.GetMedia(ctx, "linkedin", "my-profile", "12345")
```

### User Management

```go
// List profiles
users, err := client.ListUsers(ctx)

// Create / delete
_, err = client.CreateUser(ctx, "new-profile")
_, err = client.DeleteUser(ctx, "old-profile")

// Generate JWT for white-label integration
jwt, err := client.GenerateJWT(ctx, "my-profile", uploadpost.JWTOptions{
    RedirectURL: "https://yourapp.com/callback",
    Platforms:   []string{"tiktok", "instagram"},
    Language:    "en",
})
fmt.Println(jwt.ConnectionURL)
```

### Instagram Comments

```go
// Get comments
comments, err := client.GetPostComments(ctx, "my-profile", "", "https://www.instagram.com/p/...")

// Private DM reply
_, err = client.ReplyToComment(ctx, "my-profile", commentID, "Thanks!")

// Public reply
_, err = client.PublicReplyToComment(ctx, "my-profile", commentID, "Thanks for the comment!")
```

### AutoDM

```go
// Start a monitor — automatically DM every commenter on the post
_, err = client.StartAutoDM(ctx, uploadpost.AutoDMOptions{
    PostURL:         "https://www.instagram.com/p/...",
    ReplyMessage:    "Thanks for commenting! Here is the link: https://example.com",
    ProfileUsername: "my-profile",
    TriggerKeywords: []string{"link", "more info"},
})

// Check status
status, err := client.GetAutoDMStatus(ctx, false)

// Pause / resume / stop / delete
_, err = client.PauseAutoDM(ctx, "monitor-id")
_, err = client.ResumeAutoDM(ctx, "monitor-id")
_, err = client.StopAutoDM(ctx, "monitor-id")
_, err = client.DeleteAutoDM(ctx, "monitor-id")
```

### Helper Endpoints

```go
fbPages, err := client.GetFacebookPages(ctx, "my-profile")
liPages, err := client.GetLinkedinPages(ctx, "my-profile")
boards,  err := client.GetPinterestBoards(ctx, "my-profile")
locs,    err := client.GetGoogleBusinessLocations(ctx, "my-profile")
```

## Error Handling

```go
resp, err := client.UploadVideo(ctx, "./video.mp4", opts)
if err != nil {
    var apiErr *uploadpost.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
    } else {
        fmt.Printf("network error: %v\n", err)
    }
    return
}
```

## Helper Functions

Use `uploadpost.Bool(v)` and `uploadpost.Int(v)` to create pointers for optional
fields that have a meaningful zero value:

```go
opts := uploadpost.VideoOptions{
    AsyncUpload:        uploadpost.Bool(false),   // explicit false
    YouTubeEmbeddable:  uploadpost.Bool(true),
    MaxPostsPerSlot:    uploadpost.Int(3),
}
```

## Platform-Specific Options

### TikTok (Video)
| Field | Description |
|-------|-------------|
| `TiktokPrivacyLevel` | `PUBLIC_TO_EVERYONE`, `MUTUAL_FOLLOW_FRIENDS`, `FOLLOWER_OF_CREATOR`, `SELF_ONLY` |
| `TiktokDisableDuet` | Disable duet |
| `TiktokDisableComment` | Disable comments |
| `TiktokDisableStitch` | Disable stitch |
| `TiktokCoverTimestamp` | Cover frame timestamp (ms) |
| `TiktokIsAIGC` | AI-generated content flag |
| `TiktokPostMode` | `DIRECT_POST` or `MEDIA_UPLOAD` |

### Instagram
| Field | Description |
|-------|-------------|
| `InstagramMediaType` | `REELS`, `STORIES` (video) / `IMAGE`, `STORIES` (photo) |
| `InstagramShareToFeed` | Share Reels/Stories to feed |
| `InstagramCollaborators` | Comma-separated collaborator usernames |
| `InstagramCoverURL` | Cover image URL or local file path |
| `InstagramAudioName` | Audio track name |
| `InstagramUserTags` | Comma-separated user tags |

### YouTube
| Field | Description |
|-------|-------------|
| `YouTubePrivacyStatus` | `public`, `unlisted`, `private` |
| `YouTubeTags` | Tag slice |
| `YouTubeCategoryID` | Category ID (e.g. `"22"` for People & Blogs) |
| `YouTubeThumbnailURL` | Custom thumbnail URL |
| `YouTubeSubtitles` | `[]YouTubeSubtitle` — language, name, file/URL |

### X (Twitter)
| Field | Description |
|-------|-------------|
| `XReplySettings` | `everyone`, `following`, `mentionedUsers`, `subscribers`, `verified` |
| `XPollOptions` | Poll choices (2–4 items) |
| `XPollDuration` | Poll duration in minutes (5–10080) |
| `XQuoteTweetID` | Tweet ID to quote |
| `XThreadImageLayout` | e.g. `"4,4"` or `"2,3,1"` |

## Links

- [Upload-Post Website](https://www.upload-post.com)
- [API Documentation](https://docs.upload-post.com)
- [Dashboard](https://app.upload-post.com)

## License

MIT