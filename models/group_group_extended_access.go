package models
import (
    "errors"
)
type GroupGroupExtendedAccess int

const (
    NOT_ALLOWED_GROUPGROUPEXTENDEDACCESS GroupGroupExtendedAccess = iota
    ONE_WEEK_GROUPGROUPEXTENDEDACCESS
    TWO_WEEKS_GROUPGROUPEXTENDEDACCESS
    TWO_WEEKS_NO_CONFIRM_GROUPGROUPEXTENDEDACCESS
)

func (i GroupGroupExtendedAccess) String() string {
    return []string{"NOT_ALLOWED", "ONE_WEEK", "TWO_WEEKS", "TWO_WEEKS_NO_CONFIRM"}[i]
}
func ParseGroupGroupExtendedAccess(v string) (any, error) {
    result := NOT_ALLOWED_GROUPGROUPEXTENDEDACCESS
    switch v {
        case "NOT_ALLOWED":
            result = NOT_ALLOWED_GROUPGROUPEXTENDEDACCESS
        case "ONE_WEEK":
            result = ONE_WEEK_GROUPGROUPEXTENDEDACCESS
        case "TWO_WEEKS":
            result = TWO_WEEKS_GROUPGROUPEXTENDEDACCESS
        case "TWO_WEEKS_NO_CONFIRM":
            result = TWO_WEEKS_NO_CONFIRM_GROUPGROUPEXTENDEDACCESS
        default:
            return 0, errors.New("Unknown GroupGroupExtendedAccess value: " + v)
    }
    return &result, nil
}
func SerializeGroupGroupExtendedAccess(values []GroupGroupExtendedAccess) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupGroupExtendedAccess) isMultiValue() bool {
    return false
}
