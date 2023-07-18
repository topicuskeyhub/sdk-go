package models
import (
    "errors"
)
// 
type GroupGroupRights int

const (
    MANAGER_GROUPGROUPRIGHTS GroupGroupRights = iota
    NORMAL_GROUPGROUPRIGHTS
)

func (i GroupGroupRights) String() string {
    return []string{"MANAGER", "NORMAL"}[i]
}
func ParseGroupGroupRights(v string) (any, error) {
    result := MANAGER_GROUPGROUPRIGHTS
    switch v {
        case "MANAGER":
            result = MANAGER_GROUPGROUPRIGHTS
        case "NORMAL":
            result = NORMAL_GROUPGROUPRIGHTS
        default:
            return 0, errors.New("Unknown GroupGroupRights value: " + v)
    }
    return &result, nil
}
func SerializeGroupGroupRights(values []GroupGroupRights) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
