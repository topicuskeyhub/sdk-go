package models
import (
    "errors"
)
// 
type WebhookWebhookAuthenticationScheme int

const (
    NONE_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME WebhookWebhookAuthenticationScheme = iota
    BASIC_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
    BEARER_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
    CUSTOM_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
)

func (i WebhookWebhookAuthenticationScheme) String() string {
    return []string{"NONE", "BASIC", "BEARER", "CUSTOM"}[i]
}
func ParseWebhookWebhookAuthenticationScheme(v string) (any, error) {
    result := NONE_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
    switch v {
        case "NONE":
            result = NONE_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
        case "BASIC":
            result = BASIC_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
        case "BEARER":
            result = BEARER_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
        case "CUSTOM":
            result = CUSTOM_WEBHOOKWEBHOOKAUTHENTICATIONSCHEME
        default:
            return 0, errors.New("Unknown WebhookWebhookAuthenticationScheme value: " + v)
    }
    return &result, nil
}
func SerializeWebhookWebhookAuthenticationScheme(values []WebhookWebhookAuthenticationScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WebhookWebhookAuthenticationScheme) isMultiValue() bool {
    return false
}
