package models
import (
    "errors"
)
type GroupGroupAuditingStatus int

const (
    AUDITED_GROUPGROUPAUDITINGSTATUS GroupGroupAuditingStatus = iota
    UNAUDITED_GROUPGROUPAUDITINGSTATUS
    OVERDUE_GROUPGROUPAUDITINGSTATUS
)

func (i GroupGroupAuditingStatus) String() string {
    return []string{"AUDITED", "UNAUDITED", "OVERDUE"}[i]
}
func ParseGroupGroupAuditingStatus(v string) (any, error) {
    result := AUDITED_GROUPGROUPAUDITINGSTATUS
    switch v {
        case "AUDITED":
            result = AUDITED_GROUPGROUPAUDITINGSTATUS
        case "UNAUDITED":
            result = UNAUDITED_GROUPGROUPAUDITINGSTATUS
        case "OVERDUE":
            result = OVERDUE_GROUPGROUPAUDITINGSTATUS
        default:
            return 0, errors.New("Unknown GroupGroupAuditingStatus value: " + v)
    }
    return &result, nil
}
func SerializeGroupGroupAuditingStatus(values []GroupGroupAuditingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupGroupAuditingStatus) isMultiValue() bool {
    return false
}
