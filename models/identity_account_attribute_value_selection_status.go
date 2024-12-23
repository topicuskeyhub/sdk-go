package models
type IdentityAccountAttributeValueSelectionStatus int

const (
    IN_SYNC_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS IdentityAccountAttributeValueSelectionStatus = iota
    OUT_OF_SYNC_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
    INVALID_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
    DUPLICATE_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
    ERROR_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
    MISSING_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
    UNKNOWN_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
)

func (i IdentityAccountAttributeValueSelectionStatus) String() string {
    return []string{"IN_SYNC", "OUT_OF_SYNC", "INVALID", "DUPLICATE", "ERROR", "MISSING", "UNKNOWN"}[i]
}
func ParseIdentityAccountAttributeValueSelectionStatus(v string) (any, error) {
    result := IN_SYNC_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
    switch v {
        case "IN_SYNC":
            result = IN_SYNC_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
        case "OUT_OF_SYNC":
            result = OUT_OF_SYNC_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
        case "INVALID":
            result = INVALID_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
        case "DUPLICATE":
            result = DUPLICATE_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
        case "ERROR":
            result = ERROR_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
        case "MISSING":
            result = MISSING_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
        case "UNKNOWN":
            result = UNKNOWN_IDENTITYACCOUNTATTRIBUTEVALUESELECTIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentityAccountAttributeValueSelectionStatus(values []IdentityAccountAttributeValueSelectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentityAccountAttributeValueSelectionStatus) isMultiValue() bool {
    return false
}