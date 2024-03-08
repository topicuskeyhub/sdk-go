package models
import (
    "errors"
)
type AuditAuditNestedGroupAction int

const (
    INCLUDE_AUDITAUDITNESTEDGROUPACTION AuditAuditNestedGroupAction = iota
    EXCLUDE_AUDITAUDITNESTEDGROUPACTION
    MISMATCH_AUDITAUDITNESTEDGROUPACTION
)

func (i AuditAuditNestedGroupAction) String() string {
    return []string{"INCLUDE", "EXCLUDE", "MISMATCH"}[i]
}
func ParseAuditAuditNestedGroupAction(v string) (any, error) {
    result := INCLUDE_AUDITAUDITNESTEDGROUPACTION
    switch v {
        case "INCLUDE":
            result = INCLUDE_AUDITAUDITNESTEDGROUPACTION
        case "EXCLUDE":
            result = EXCLUDE_AUDITAUDITNESTEDGROUPACTION
        case "MISMATCH":
            result = MISMATCH_AUDITAUDITNESTEDGROUPACTION
        default:
            return 0, errors.New("Unknown AuditAuditNestedGroupAction value: " + v)
    }
    return &result, nil
}
func SerializeAuditAuditNestedGroupAction(values []AuditAuditNestedGroupAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuditAuditNestedGroupAction) isMultiValue() bool {
    return false
}
