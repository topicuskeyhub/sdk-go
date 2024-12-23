package models
type IdentityAccountAttributeValueStatus int

const (
    VALID_IDENTITYACCOUNTATTRIBUTEVALUESTATUS IdentityAccountAttributeValueStatus = iota
    INVALID_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
    DUPLICATE_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
    ERROR_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
)

func (i IdentityAccountAttributeValueStatus) String() string {
    return []string{"VALID", "INVALID", "DUPLICATE", "ERROR"}[i]
}
func ParseIdentityAccountAttributeValueStatus(v string) (any, error) {
    result := VALID_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
    switch v {
        case "VALID":
            result = VALID_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
        case "INVALID":
            result = INVALID_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
        case "DUPLICATE":
            result = DUPLICATE_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
        case "ERROR":
            result = ERROR_IDENTITYACCOUNTATTRIBUTEVALUESTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentityAccountAttributeValueStatus(values []IdentityAccountAttributeValueStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentityAccountAttributeValueStatus) isMultiValue() bool {
    return false
}
