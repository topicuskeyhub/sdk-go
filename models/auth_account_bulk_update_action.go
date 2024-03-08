package models
import (
    "errors"
)
type AuthAccountBulkUpdateAction int

const (
    CHANGE_LICENSE_ROLE_TO_PRO_AUTHACCOUNTBULKUPDATEACTION AuthAccountBulkUpdateAction = iota
    CHANGE_LICENSE_ROLE_TO_BUSINESS_AUTHACCOUNTBULKUPDATEACTION
    MARK_FOR_RENAME_AUTHACCOUNTBULKUPDATEACTION
)

func (i AuthAccountBulkUpdateAction) String() string {
    return []string{"CHANGE_LICENSE_ROLE_TO_PRO", "CHANGE_LICENSE_ROLE_TO_BUSINESS", "MARK_FOR_RENAME"}[i]
}
func ParseAuthAccountBulkUpdateAction(v string) (any, error) {
    result := CHANGE_LICENSE_ROLE_TO_PRO_AUTHACCOUNTBULKUPDATEACTION
    switch v {
        case "CHANGE_LICENSE_ROLE_TO_PRO":
            result = CHANGE_LICENSE_ROLE_TO_PRO_AUTHACCOUNTBULKUPDATEACTION
        case "CHANGE_LICENSE_ROLE_TO_BUSINESS":
            result = CHANGE_LICENSE_ROLE_TO_BUSINESS_AUTHACCOUNTBULKUPDATEACTION
        case "MARK_FOR_RENAME":
            result = MARK_FOR_RENAME_AUTHACCOUNTBULKUPDATEACTION
        default:
            return 0, errors.New("Unknown AuthAccountBulkUpdateAction value: " + v)
    }
    return &result, nil
}
func SerializeAuthAccountBulkUpdateAction(values []AuthAccountBulkUpdateAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthAccountBulkUpdateAction) isMultiValue() bool {
    return false
}
