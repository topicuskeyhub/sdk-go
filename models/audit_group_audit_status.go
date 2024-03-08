package models
import (
    "errors"
)
type AuditGroupAuditStatus int

const (
    NEW_AUDITGROUPAUDITSTATUS AuditGroupAuditStatus = iota
    DRAFT_AUDITGROUPAUDITSTATUS
    UNDER_REVIEW_AUDITGROUPAUDITSTATUS
    FINAL_AUDITGROUPAUDITSTATUS
)

func (i AuditGroupAuditStatus) String() string {
    return []string{"NEW", "DRAFT", "UNDER_REVIEW", "FINAL"}[i]
}
func ParseAuditGroupAuditStatus(v string) (any, error) {
    result := NEW_AUDITGROUPAUDITSTATUS
    switch v {
        case "NEW":
            result = NEW_AUDITGROUPAUDITSTATUS
        case "DRAFT":
            result = DRAFT_AUDITGROUPAUDITSTATUS
        case "UNDER_REVIEW":
            result = UNDER_REVIEW_AUDITGROUPAUDITSTATUS
        case "FINAL":
            result = FINAL_AUDITGROUPAUDITSTATUS
        default:
            return 0, errors.New("Unknown AuditGroupAuditStatus value: " + v)
    }
    return &result, nil
}
func SerializeAuditGroupAuditStatus(values []AuditGroupAuditStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuditGroupAuditStatus) isMultiValue() bool {
    return false
}
