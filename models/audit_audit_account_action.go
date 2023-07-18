package models
import (
    "errors"
)
// 
type AuditAuditAccountAction int

const (
    APPROVE_AUDITAUDITACCOUNTACTION AuditAuditAccountAction = iota
    CHANGE_TO_MANAGER_AUDITAUDITACCOUNTACTION
    CHANGE_TO_NORMAL_AUDITAUDITACCOUNTACTION
    REMOVE_AUDITAUDITACCOUNTACTION
    CONNECT_NESTED_AUDITAUDITACCOUNTACTION
)

func (i AuditAuditAccountAction) String() string {
    return []string{"APPROVE", "CHANGE_TO_MANAGER", "CHANGE_TO_NORMAL", "REMOVE", "CONNECT_NESTED"}[i]
}
func ParseAuditAuditAccountAction(v string) (any, error) {
    result := APPROVE_AUDITAUDITACCOUNTACTION
    switch v {
        case "APPROVE":
            result = APPROVE_AUDITAUDITACCOUNTACTION
        case "CHANGE_TO_MANAGER":
            result = CHANGE_TO_MANAGER_AUDITAUDITACCOUNTACTION
        case "CHANGE_TO_NORMAL":
            result = CHANGE_TO_NORMAL_AUDITAUDITACCOUNTACTION
        case "REMOVE":
            result = REMOVE_AUDITAUDITACCOUNTACTION
        case "CONNECT_NESTED":
            result = CONNECT_NESTED_AUDITAUDITACCOUNTACTION
        default:
            return 0, errors.New("Unknown AuditAuditAccountAction value: " + v)
    }
    return &result, nil
}
func SerializeAuditAuditAccountAction(values []AuditAuditAccountAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
