package uploadpost

import (
	"encoding/json"
	"fmt"
)

// ParseWebhookEvent parses the raw JSON payload received on your webhook
// endpoint and returns a typed event struct. The underlying type can be
// asserted with a type switch:
//
//	switch evt := uploadpost.ParseWebhookEvent(body).(type) {
//	case *uploadpost.WebhookUploadCompleted:
//	    fmt.Println(evt.JobID, evt.Result.URL)
//	case *uploadpost.WebhookSocialAccountConnected:
//	    fmt.Println(evt.Platform, evt.AccountName)
//	case *uploadpost.WebhookSocialAccountDisconnected:
//	    fmt.Println(evt.Reason)
//	case *uploadpost.WebhookSocialAccountReauthRequired:
//	    fmt.Println(evt.ProfileUsername)
//	}
//
// Unknown event types are returned as a plain map[string]interface{} so that
// new events added by the API do not break existing code.
func ParseWebhookEvent(payload []byte) (interface{}, error) {
	// Extract the event type first.
	var base struct {
		Event WebhookEventType `json:"event"`
	}
	if err := json.Unmarshal(payload, &base); err != nil {
		return nil, fmt.Errorf("uploadpost: failed to parse webhook event: %w", err)
	}

	switch base.Event {
	case WebhookEventUploadCompleted:
		var evt WebhookUploadCompleted
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, fmt.Errorf("uploadpost: failed to parse upload_completed event: %w", err)
		}
		return &evt, nil

	case WebhookEventSocialAccountConnected:
		var evt WebhookSocialAccountConnected
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, fmt.Errorf("uploadpost: failed to parse social_account_connected event: %w", err)
		}
		return &evt, nil

	case WebhookEventSocialAccountDisconnected:
		var evt WebhookSocialAccountDisconnected
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, fmt.Errorf("uploadpost: failed to parse social_account_disconnected event: %w", err)
		}
		return &evt, nil

	case WebhookEventSocialAccountReauthRequired:
		var evt WebhookSocialAccountReauthRequired
		if err := json.Unmarshal(payload, &evt); err != nil {
			return nil, fmt.Errorf("uploadpost: failed to parse social_account_reauth_required event: %w", err)
		}
		return &evt, nil

	default:
		// Unknown event — parse generically so the caller is not broken.
		var m map[string]interface{}
		if err := json.Unmarshal(payload, &m); err != nil {
			return nil, fmt.Errorf("uploadpost: failed to parse unknown webhook event: %w", err)
		}
		return m, nil
	}
}
