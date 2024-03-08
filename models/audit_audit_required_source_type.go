package models
import (
    "errors"
)
type AuditAuditRequiredSourceType int

const (
    CONFORM_AUDITAUDITREQUIREDSOURCETYPE AuditAuditRequiredSourceType = iota
    PERIODIC_AUDITAUDITREQUIREDSOURCETYPE
    REQUESTED_AUDITAUDITREQUIREDSOURCETYPE
)

func (i AuditAuditRequiredSourceType) String() string {
    return []string{"CONFORM", "PERIODIC", "REQUESTED"}[i]
}
func ParseAuditAuditRequiredSourceType(v string) (any, error) {
    result := CONFORM_AUDITAUDITREQUIREDSOURCETYPE
    switch v {
        case "CONFORM":
            result = CONFORM_AUDITAUDITREQUIREDSOURCETYPE
        case "PERIODIC":
            result = PERIODIC_AUDITAUDITREQUIREDSOURCETYPE
        case "REQUESTED":
            result = REQUESTED_AUDITAUDITREQUIREDSOURCETYPE
        default:
            return 0, errors.New("Unknown AuditAuditRequiredSourceType value: " + v)
    }
    return &result, nil
}
func SerializeAuditAuditRequiredSourceType(values []AuditAuditRequiredSourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuditAuditRequiredSourceType) isMultiValue() bool {
    return false
}
