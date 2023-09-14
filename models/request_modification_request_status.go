package models
import (
    "errors"
)
// 
type RequestModificationRequestStatus int

const (
    REQUESTED_REQUESTMODIFICATIONREQUESTSTATUS RequestModificationRequestStatus = iota
    ALLOWED_REQUESTMODIFICATIONREQUESTSTATUS
    DISALLOWED_REQUESTMODIFICATIONREQUESTSTATUS
    CANCELLED_REQUESTMODIFICATIONREQUESTSTATUS
)

func (i RequestModificationRequestStatus) String() string {
    return []string{"REQUESTED", "ALLOWED", "DISALLOWED", "CANCELLED"}[i]
}
func ParseRequestModificationRequestStatus(v string) (any, error) {
    result := REQUESTED_REQUESTMODIFICATIONREQUESTSTATUS
    switch v {
        case "REQUESTED":
            result = REQUESTED_REQUESTMODIFICATIONREQUESTSTATUS
        case "ALLOWED":
            result = ALLOWED_REQUESTMODIFICATIONREQUESTSTATUS
        case "DISALLOWED":
            result = DISALLOWED_REQUESTMODIFICATIONREQUESTSTATUS
        case "CANCELLED":
            result = CANCELLED_REQUESTMODIFICATIONREQUESTSTATUS
        default:
            return 0, errors.New("Unknown RequestModificationRequestStatus value: " + v)
    }
    return &result, nil
}
func SerializeRequestModificationRequestStatus(values []RequestModificationRequestStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequestModificationRequestStatus) isMultiValue() bool {
    return false
}
