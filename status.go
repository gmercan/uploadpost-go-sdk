package uploadpost

import "context"

// GetStatus returns the current status of an asynchronous upload by request ID.
func (c *Client) GetStatus(ctx context.Context, requestID string) (*StatusResponse, error) {
	var resp StatusResponse
	err := c.getJSON(ctx, "/uploadposts/status", map[string]string{
		"request_id": requestID,
	}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJobStatus returns the status of a scheduled or queued post by job ID.
func (c *Client) GetJobStatus(ctx context.Context, jobID string) (*StatusResponse, error) {
	var resp StatusResponse
	err := c.getJSON(ctx, "/uploadposts/status", map[string]string{
		"job_id": jobID,
	}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHistory returns paginated upload history.
// Pass a zero-value HistoryOptions{} for defaults (page 1, limit 20).
func (c *Client) GetHistory(ctx context.Context, opts HistoryOptions) (*HistoryResponse, error) {
	page := opts.Page
	if page < 1 {
		page = 1
	}
	limit := opts.Limit
	if limit == 0 {
		limit = 20
	}

	var resp HistoryResponse
	err := c.getJSON(ctx, "/uploadposts/history", map[string]string{
		"page":  itoa(page),
		"limit": itoa(limit),
	}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func itoa(n int) string {
	if n == 0 {
		return ""
	}
	s := ""
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	if neg {
		s = "-" + s
	}
	return s
}