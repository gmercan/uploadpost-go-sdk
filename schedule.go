package uploadpost

import (
	"context"
	"fmt"
)

// ListScheduled returns all scheduled posts.
func (c *Client) ListScheduled(ctx context.Context) (*ScheduledResponse, error) {
	var resp ScheduledResponse
	if err := c.getJSON(ctx, "/uploadposts/schedule", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelScheduled cancels a scheduled post by job ID.
func (c *Client) CancelScheduled(ctx context.Context, jobID string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.deleteJSON(ctx, "/uploadposts/schedule/"+jobID, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EditScheduled changes the scheduled date and/or timezone of a queued post.
func (c *Client) EditScheduled(ctx context.Context, jobID string, opts EditScheduledOptions) (*SimpleResponse, error) {
	if opts.ScheduledDate == "" && opts.Timezone == "" {
		return nil, fmt.Errorf("uploadpost: EditScheduledOptions requires at least one of ScheduledDate or Timezone")
	}

	payload := map[string]string{}
	if opts.ScheduledDate != "" {
		payload["scheduled_date"] = opts.ScheduledDate
	}
	if opts.Timezone != "" {
		payload["timezone"] = opts.Timezone
	}

	var resp SimpleResponse
	if err := c.patchJSON(ctx, "/uploadposts/schedule/"+jobID, payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}