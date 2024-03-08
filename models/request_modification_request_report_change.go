package models
import (
    "errors"
)
type RequestModificationRequestReportChange int

const (
    ACCOUNT_REMOVED_FROM_GROUP_REQUESTMODIFICATIONREQUESTREPORTCHANGE RequestModificationRequestReportChange = iota
    MODIFICATION_REQUEST_CANCELLED_REQUESTMODIFICATIONREQUESTREPORTCHANGE
    PROVISIONING_GROUP_REMOVED_REQUESTMODIFICATIONREQUESTREPORTCHANGE
    SERVICE_ACCOUNT_GROUP_REMOVED_REQUESTMODIFICATIONREQUESTREPORTCHANGE
)

func (i RequestModificationRequestReportChange) String() string {
    return []string{"ACCOUNT_REMOVED_FROM_GROUP", "MODIFICATION_REQUEST_CANCELLED", "PROVISIONING_GROUP_REMOVED", "SERVICE_ACCOUNT_GROUP_REMOVED"}[i]
}
func ParseRequestModificationRequestReportChange(v string) (any, error) {
    result := ACCOUNT_REMOVED_FROM_GROUP_REQUESTMODIFICATIONREQUESTREPORTCHANGE
    switch v {
        case "ACCOUNT_REMOVED_FROM_GROUP":
            result = ACCOUNT_REMOVED_FROM_GROUP_REQUESTMODIFICATIONREQUESTREPORTCHANGE
        case "MODIFICATION_REQUEST_CANCELLED":
            result = MODIFICATION_REQUEST_CANCELLED_REQUESTMODIFICATIONREQUESTREPORTCHANGE
        case "PROVISIONING_GROUP_REMOVED":
            result = PROVISIONING_GROUP_REMOVED_REQUESTMODIFICATIONREQUESTREPORTCHANGE
        case "SERVICE_ACCOUNT_GROUP_REMOVED":
            result = SERVICE_ACCOUNT_GROUP_REMOVED_REQUESTMODIFICATIONREQUESTREPORTCHANGE
        default:
            return 0, errors.New("Unknown RequestModificationRequestReportChange value: " + v)
    }
    return &result, nil
}
func SerializeRequestModificationRequestReportChange(values []RequestModificationRequestReportChange) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequestModificationRequestReportChange) isMultiValue() bool {
    return false
}
