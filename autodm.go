package uploadpost

import (
	"context"
	"fmt"
)

// StartAutoDM creates an AutoDM monitor that automatically sends private DMs
// to users who comment on the given Instagram post.
func (c *Client) StartAutoDM(ctx context.Context, opts AutoDMOptions) (*SimpleResponse, error) {
	if opts.PostURL == "" {
		return nil, fmt.Errorf("uploadpost: AutoDMOptions.PostURL is required")
	}
	if opts.ReplyMessage == "" {
		return nil, fmt.Errorf("uploadpost: AutoDMOptions.ReplyMessage is required")
	}
	if opts.ProfileUsername == "" {
		return nil, fmt.Errorf("uploadpost: AutoDMOptions.ProfileUsername is required")
	}

	payload := map[string]interface{}{
		"post_url":         opts.PostURL,
		"reply_message":    opts.ReplyMessage,
		"profile_username": opts.ProfileUsername,
	}
	if opts.MonitoringInterval != nil {
		payload["monitoring_interval"] = *opts.MonitoringInterval
	}
	if len(opts.TriggerKeywords) > 0 {
		payload["trigger_keywords"] = opts.TriggerKeywords
	}

	var resp SimpleResponse
	if err := c.postJSON(ctx, "/autodms", payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAutoDMStatus returns the status of AutoDM monitors.
// Set includeInactive to true to also include stopped and expired monitors.
func (c *Client) GetAutoDMStatus(ctx context.Context, includeInactive bool) (*AutoDMStatusResponse, error) {
	params := map[string]string{}
	if includeInactive {
		params["include_inactive"] = "true"
	}

	var resp AutoDMStatusResponse
	if err := c.getJSON(ctx, "/autodms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAutoDMLogs returns the activity log for a specific AutoDM monitor.
func (c *Client) GetAutoDMLogs(ctx context.Context, monitorID string) (*AutoDMLogsResponse, error) {
	var resp AutoDMLogsResponse
	if err := c.getJSON(ctx, "/autodms/logs", map[string]string{
		"monitor_id": monitorID,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PauseAutoDM temporarily pauses a running monitor.
func (c *Client) PauseAutoDM(ctx context.Context, monitorID string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/autodms/pause", map[string]string{"monitor_id": monitorID}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResumeAutoDM resumes a paused monitor.
func (c *Client) ResumeAutoDM(ctx context.Context, monitorID string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/autodms/resume", map[string]string{"monitor_id": monitorID}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopAutoDM deactivates a monitor (data is preserved, can be re-listed with
// includeInactive=true).
func (c *Client) StopAutoDM(ctx context.Context, monitorID string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/autodms/stop", map[string]string{"monitor_id": monitorID}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAutoDM permanently deletes a monitor and all its data.
func (c *Client) DeleteAutoDM(ctx context.Context, monitorID string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.deleteJSON(ctx, "/autodms", map[string]string{"monitor_id": monitorID}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}