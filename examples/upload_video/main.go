package main

import (
	"context"
	"fmt"
	"log"

	uploadpost "github.com/gmercan/uploadpost-go-sdk"
)

func main() {
	client := uploadpost.New("YOUR_API_KEY")
	ctx := context.Background()

	// ── Upload a video ────────────────────────────────────────────────────────
	resp, err := client.UploadVideo(ctx, "./video.mp4", uploadpost.VideoOptions{
		User:      "my-profile",
		Platforms: []string{"tiktok", "instagram", "youtube"},
		Title:     "Check out this awesome video!",

		// Schedule for later (optional)
		// ScheduledDate: "2024-12-25T10:00:00Z",
		// Timezone:      "Europe/Madrid",

		// First comment (optional)
		FirstComment: "Thanks for watching!",

		// TikTok-specific
		TiktokPrivacyLevel: "PUBLIC_TO_EVERYONE",

		// Instagram-specific
		InstagramMediaType: "REELS",

		// YouTube-specific
		YouTubePrivacyStatus: "public",
		YouTubeTags:          []string{"tutorial", "coding"},
	})
	if err != nil {
		log.Fatalf("upload failed: %v", err)
	}
	fmt.Printf("Upload response: %+v\n", resp)

	// ── Poll async status ─────────────────────────────────────────────────────
	if resp.RequestID != "" {
		status, err := client.GetStatus(ctx, resp.RequestID)
		if err != nil {
			log.Fatalf("status check failed: %v", err)
		}
		fmt.Printf("Status: %s\n", status.Status)
	}

	// ── Upload photos ─────────────────────────────────────────────────────────
	photoResp, err := client.UploadPhotos(ctx,
		[]string{"./photo1.jpg", "./photo2.jpg"},
		uploadpost.PhotosOptions{
			User:               "my-profile",
			Platforms:          []string{"instagram", "facebook", "x"},
			Title:              "Check out these photos!",
			InstagramMediaType: "IMAGE",
		},
	)
	if err != nil {
		log.Fatalf("photo upload failed: %v", err)
	}
	fmt.Printf("Photos response: %+v\n", photoResp)

	// ── Upload text post ──────────────────────────────────────────────────────
	textResp, err := client.UploadText(ctx, uploadpost.TextOptions{
		User:      "my-profile",
		Platforms: []string{"x", "linkedin", "threads"},
		Title:     "Just shipped a new feature! Check it out.",

		// X poll
		XPollOptions:  []string{"Yes", "No", "Maybe"},
		XPollDuration: uploadpost.Int(1440), // 24 hours

		// LinkedIn company page (optional)
		// TargetLinkedinPageID: "company-page-id",
	})
	if err != nil {
		log.Fatalf("text upload failed: %v", err)
	}
	fmt.Printf("Text response: %+v\n", textResp)

	// ── User management ───────────────────────────────────────────────────────
	users, err := client.ListUsers(ctx)
	if err != nil {
		log.Fatalf("list users failed: %v", err)
	}
	for _, u := range users.Profiles {
		fmt.Printf("Profile: %s\n", u.Username)
	}

	// ── Generate JWT for white-label integration ───────────────────────────────
	jwtResp, err := client.GenerateJWT(ctx, "my-profile", uploadpost.JWTOptions{
		RedirectURL: "https://yourapp.com/callback",
		Platforms:   []string{"tiktok", "instagram"},
		Language:    "en",
	})
	if err != nil {
		log.Fatalf("generate JWT failed: %v", err)
	}
	fmt.Printf("Connection URL: %s\n", jwtResp.ConnectionURL)
}