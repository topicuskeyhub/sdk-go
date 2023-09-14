package models
import (
    "errors"
)
// 
type RequestGroupNestAccountAction int

const (
    REMOVE_REQUESTGROUPNESTACCOUNTACTION RequestGroupNestAccountAction = iota
    CONVERT_REQUESTGROUPNESTACCOUNTACTION
    RETAIN_REQUESTGROUPNESTACCOUNTACTION
)

func (i RequestGroupNestAccountAction) String() string {
    return []string{"REMOVE", "CONVERT", "RETAIN"}[i]
}
func ParseRequestGroupNestAccountAction(v string) (any, error) {
    result := REMOVE_REQUESTGROUPNESTACCOUNTACTION
    switch v {
        case "REMOVE":
            result = REMOVE_REQUESTGROUPNESTACCOUNTACTION
        case "CONVERT":
            result = CONVERT_REQUESTGROUPNESTACCOUNTACTION
        case "RETAIN":
            result = RETAIN_REQUESTGROUPNESTACCOUNTACTION
        default:
            return 0, errors.New("Unknown RequestGroupNestAccountAction value: " + v)
    }
    return &result, nil
}
func SerializeRequestGroupNestAccountAction(values []RequestGroupNestAccountAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequestGroupNestAccountAction) isMultiValue() bool {
    return false
}
