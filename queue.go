package uploadpost

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetQueueSettings returns the posting queue configuration for a profile.
func (c *Client) GetQueueSettings(ctx context.Context, profileUsername string) (*QueueSettingsResponse, error) {
	if profileUsername == "" {
		return nil, fmt.Errorf("uploadpost: profileUsername is required")
	}
	var resp QueueSettingsResponse
	if err := c.getJSON(ctx, "/uploadposts/queue/settings", map[string]string{
		"profile_username": profileUsername,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateQueueSettings updates the posting queue configuration for a profile.
// DaysOfWeek uses 0=Monday … 6=Sunday. Slot hours must be 0-23, minutes 0-59.
// At most 24 slots are allowed. MaxPostsPerSlot must be 1-100.
func (c *Client) UpdateQueueSettings(ctx context.Context, opts UpdateQueueSettingsOptions) (*SimpleResponse, error) {
	if opts.ProfileUsername == "" {
		return nil, fmt.Errorf("uploadpost: UpdateQueueSettingsOptions.ProfileUsername is required")
	}
	if len(opts.Slots) > 24 {
		return nil, fmt.Errorf("uploadpost: at most 24 queue slots are allowed (got %d)", len(opts.Slots))
	}
	for _, s := range opts.Slots {
		if s.Hour < 0 || s.Hour > 23 {
			return nil, fmt.Errorf("uploadpost: slot hour must be 0-23 (got %d)", s.Hour)
		}
		if s.Minute < 0 || s.Minute > 59 {
			return nil, fmt.Errorf("uploadpost: slot minute must be 0-59 (got %d)", s.Minute)
		}
	}
	for _, d := range opts.DaysOfWeek {
		if d < 0 || d > 6 {
			return nil, fmt.Errorf("uploadpost: days_of_week values must be 0-6 (got %d)", d)
		}
	}
	if opts.MaxPostsPerSlot != nil && (*opts.MaxPostsPerSlot < 1 || *opts.MaxPostsPerSlot > 100) {
		return nil, fmt.Errorf("uploadpost: max_posts_per_slot must be 1-100 (got %d)", *opts.MaxPostsPerSlot)
	}

	payload := map[string]interface{}{
		"profile_username": opts.ProfileUsername,
	}
	if opts.Timezone != "" {
		payload["timezone"] = opts.Timezone
	}
	if len(opts.Slots) > 0 {
		payload["slots"] = opts.Slots
	}
	if len(opts.DaysOfWeek) > 0 {
		payload["days_of_week"] = opts.DaysOfWeek
	}
	if opts.MaxPostsPerSlot != nil {
		payload["max_posts_per_slot"] = *opts.MaxPostsPerSlot
	}

	var resp SimpleResponse
	if err := c.postJSON(ctx, "/uploadposts/queue/settings", payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewQueue returns the upcoming queue slots for a profile.
// count defaults to 10 when zero; maximum is 50.
func (c *Client) PreviewQueue(ctx context.Context, profileUsername string, count int) (*QueuePreviewResponse, error) {
	if profileUsername == "" {
		return nil, fmt.Errorf("uploadpost: profileUsername is required")
	}
	if count == 0 {
		count = 10
	}
	if count > 50 {
		return nil, fmt.Errorf("uploadpost: count must be ≤ 50 (got %d)", count)
	}

	// Use raw JSON decode because the API may return either "scheduled_posts" or
	// "scheduled_post" (singular) in the slot object for backward compatibility.
	params := map[string]string{
		"profile_username": profileUsername,
		"count":            itoa(count),
	}
	rawData, statusCode, err := c.doRaw(ctx, "GET", buildQueryURL("/uploadposts/queue/preview", params), nil, "")
	if err != nil {
		return nil, err
	}
	if statusCode >= 400 {
		var m map[string]json.RawMessage
		if json.Unmarshal(rawData, &m) == nil {
			return nil, &APIError{StatusCode: statusCode, Message: extractErrorMessage(m)}
		}
		return nil, &APIError{StatusCode: statusCode, Message: string(rawData)}
	}
	var resp QueuePreviewResponse
	if err := json.Unmarshal(rawData, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarkQueueSlotFull manually marks a queue slot as full so no more posts are
// assigned to it. slotDatetime must be an ISO 8601 UTC string.
func (c *Client) MarkQueueSlotFull(ctx context.Context, profileUsername, slotDatetime string) (*SimpleResponse, error) {
	if profileUsername == "" {
		return nil, fmt.Errorf("uploadpost: profileUsername is required")
	}
	if slotDatetime == "" {
		return nil, fmt.Errorf("uploadpost: slotDatetime is required (ISO 8601 UTC)")
	}
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/uploadposts/queue/slot-full", map[string]string{
		"profile_username": profileUsername,
		"slot_datetime":    slotDatetime,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnmarkQueueSlotFull removes the manually-full mark from a queue slot, making
// it available for new posts again. slotDatetime must be an ISO 8601 UTC string.
func (c *Client) UnmarkQueueSlotFull(ctx context.Context, profileUsername, slotDatetime string) (*SimpleResponse, error) {
	if profileUsername == "" {
		return nil, fmt.Errorf("uploadpost: profileUsername is required")
	}
	if slotDatetime == "" {
		return nil, fmt.Errorf("uploadpost: slotDatetime is required (ISO 8601 UTC)")
	}
	var resp SimpleResponse
	if err := c.deleteJSON(ctx, "/uploadposts/queue/slot-full", map[string]string{
		"profile_username": profileUsername,
		"slot_datetime":    slotDatetime,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNextQueueSlot returns the next available queue slot for a profile.
// When no slot is available within the next 30 days, NextSlot will be nil —
// this is not an error condition.
func (c *Client) GetNextQueueSlot(ctx context.Context, profileUsername string) (*NextQueueSlotResponse, error) {
	if profileUsername == "" {
		return nil, fmt.Errorf("uploadpost: profileUsername is required")
	}
	var resp NextQueueSlotResponse
	if err := c.getJSON(ctx, "/uploadposts/queue/next-slot", map[string]string{
		"profile_username": profileUsername,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// buildQueryURL appends query params to a path string.
func buildQueryURL(path string, params map[string]string) string {
	if len(params) == 0 {
		return path
	}
	sep := "?"
	for k, v := range params {
		if v != "" {
			path += sep + k + "=" + v
			sep = "&"
		}
	}
	return path
}
