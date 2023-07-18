package models
import (
    "errors"
)
// 
type GroupGroupRequestStatus int

const (
    MEMBER_GROUPGROUPREQUESTSTATUS GroupGroupRequestStatus = iota
    REQUEST_OPEN_GROUPGROUPREQUESTSTATUS
    AVAILABLE_GROUPGROUPREQUESTSTATUS
)

func (i GroupGroupRequestStatus) String() string {
    return []string{"MEMBER", "REQUEST_OPEN", "AVAILABLE"}[i]
}
func ParseGroupGroupRequestStatus(v string) (any, error) {
    result := MEMBER_GROUPGROUPREQUESTSTATUS
    switch v {
        case "MEMBER":
            result = MEMBER_GROUPGROUPREQUESTSTATUS
        case "REQUEST_OPEN":
            result = REQUEST_OPEN_GROUPGROUPREQUESTSTATUS
        case "AVAILABLE":
            result = AVAILABLE_GROUPGROUPREQUESTSTATUS
        default:
            return 0, errors.New("Unknown GroupGroupRequestStatus value: " + v)
    }
    return &result, nil
}
func SerializeGroupGroupRequestStatus(values []GroupGroupRequestStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
