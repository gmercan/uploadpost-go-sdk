package uploadpost

import (
	"context"
	"fmt"
)

// ListRedditDetailedPosts returns detailed posts (with media info) from a
// connected Reddit account. The API may return up to 2000 posts.
func (c *Client) ListRedditDetailedPosts(ctx context.Context, profileUsername string) (*RedditDetailedPostsResponse, error) {
	if profileUsername == "" {
		return nil, fmt.Errorf("uploadpost: profileUsername is required")
	}

	var resp RedditDetailedPostsResponse
	if err := c.getJSON(ctx, "/uploadposts/reddit/detailed-posts/", map[string]string{
		"profile_username": profileUsername,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
