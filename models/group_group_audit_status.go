package models
import (
    "errors"
)
// 
type GroupGroupAuditStatus int

const (
    NEW_GROUPGROUPAUDITSTATUS GroupGroupAuditStatus = iota
    DRAFT_GROUPGROUPAUDITSTATUS
    UNDER_REVIEW_GROUPGROUPAUDITSTATUS
    FINAL_GROUPGROUPAUDITSTATUS
)

func (i GroupGroupAuditStatus) String() string {
    return []string{"NEW", "DRAFT", "UNDER_REVIEW", "FINAL"}[i]
}
func ParseGroupGroupAuditStatus(v string) (any, error) {
    result := NEW_GROUPGROUPAUDITSTATUS
    switch v {
        case "NEW":
            result = NEW_GROUPGROUPAUDITSTATUS
        case "DRAFT":
            result = DRAFT_GROUPGROUPAUDITSTATUS
        case "UNDER_REVIEW":
            result = UNDER_REVIEW_GROUPGROUPAUDITSTATUS
        case "FINAL":
            result = FINAL_GROUPGROUPAUDITSTATUS
        default:
            return 0, errors.New("Unknown GroupGroupAuditStatus value: " + v)
    }
    return &result, nil
}
func SerializeGroupGroupAuditStatus(values []GroupGroupAuditStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupGroupAuditStatus) isMultiValue() bool {
    return false
}
