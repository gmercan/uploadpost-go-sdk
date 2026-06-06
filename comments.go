package uploadpost

import "context"

// GetPostComments returns comments on an Instagram post.
// Provide either postID (numeric media ID) or postURL (full Instagram URL).
func (c *Client) GetPostComments(ctx context.Context, user, postID, postURL string) (*CommentsResponse, error) {
	params := map[string]string{
		"platform": "instagram",
		"user":     user,
	}
	if postID != "" {
		params["post_id"] = postID
	}
	if postURL != "" {
		params["post_url"] = postURL
	}

	var resp CommentsResponse
	if err := c.getJSON(ctx, "/uploadposts/comments", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplyToComment sends a private DM reply to the author of an Instagram comment.
func (c *Client) ReplyToComment(ctx context.Context, user, commentID, message string) (*ReplyResponse, error) {
	var resp ReplyResponse
	err := c.postJSON(ctx, "/uploadposts/comments/reply", map[string]string{
		"platform":   "instagram",
		"user":       user,
		"comment_id": commentID,
		"message":    message,
	}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PublicReplyToComment posts a public reply to an Instagram comment, visible
// under the original comment.
func (c *Client) PublicReplyToComment(ctx context.Context, user, commentID, message string) (*ReplyResponse, error) {
	var resp ReplyResponse
	err := c.postJSON(ctx, "/uploadposts/comments/public-reply", map[string]string{
		"platform":   "instagram",
		"user":       user,
		"comment_id": commentID,
		"message":    message,
	}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
