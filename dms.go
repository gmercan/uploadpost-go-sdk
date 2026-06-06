package uploadpost

import (
	"context"
	"fmt"
)

// SendDirectMessage sends a private Instagram DM to a recipient.
// recipientID can be obtained from GetPostComments or ListDMConversations.
func (c *Client) SendDirectMessage(ctx context.Context, platform, user, recipientID, message string) (*SendDMResponse, error) {
	if platform == "" {
		return nil, fmt.Errorf("uploadpost: platform is required")
	}
	if user == "" {
		return nil, fmt.Errorf("uploadpost: user is required")
	}
	if recipientID == "" {
		return nil, fmt.Errorf("uploadpost: recipientID is required")
	}
	if message == "" {
		return nil, fmt.Errorf("uploadpost: message is required")
	}

	var resp SendDMResponse
	if err := c.postJSON(ctx, "/uploadposts/dms/send", map[string]string{
		"platform":     platform,
		"user":         user,
		"recipient_id": recipientID,
		"message":      message,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDMConversations returns the DM conversation threads for a connected account.
func (c *Client) ListDMConversations(ctx context.Context, platform, user string) (*DMConversationsResponse, error) {
	if platform == "" {
		return nil, fmt.Errorf("uploadpost: platform is required")
	}
	if user == "" {
		return nil, fmt.Errorf("uploadpost: user is required")
	}

	var resp DMConversationsResponse
	if err := c.getJSON(ctx, "/uploadposts/dms/conversations", map[string]string{
		"platform": platform,
		"user":     user,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
