package models
type IdentityAccountAttributeFormat int

const (
    INTERNAL_IDENTITYACCOUNTATTRIBUTEFORMAT IdentityAccountAttributeFormat = iota
    EMAIL_IDENTITYACCOUNTATTRIBUTEFORMAT
    TELEPHONE_IDENTITYACCOUNTATTRIBUTEFORMAT
    NUMBER_IDENTITYACCOUNTATTRIBUTEFORMAT
    DATE_IDENTITYACCOUNTATTRIBUTEFORMAT
    BOOLEAN_IDENTITYACCOUNTATTRIBUTEFORMAT
    TEXT_IDENTITYACCOUNTATTRIBUTEFORMAT
)

func (i IdentityAccountAttributeFormat) String() string {
    return []string{"INTERNAL", "EMAIL", "TELEPHONE", "NUMBER", "DATE", "BOOLEAN", "TEXT"}[i]
}
func ParseIdentityAccountAttributeFormat(v string) (any, error) {
    result := INTERNAL_IDENTITYACCOUNTATTRIBUTEFORMAT
    switch v {
        case "INTERNAL":
            result = INTERNAL_IDENTITYACCOUNTATTRIBUTEFORMAT
        case "EMAIL":
            result = EMAIL_IDENTITYACCOUNTATTRIBUTEFORMAT
        case "TELEPHONE":
            result = TELEPHONE_IDENTITYACCOUNTATTRIBUTEFORMAT
        case "NUMBER":
            result = NUMBER_IDENTITYACCOUNTATTRIBUTEFORMAT
        case "DATE":
            result = DATE_IDENTITYACCOUNTATTRIBUTEFORMAT
        case "BOOLEAN":
            result = BOOLEAN_IDENTITYACCOUNTATTRIBUTEFORMAT
        case "TEXT":
            result = TEXT_IDENTITYACCOUNTATTRIBUTEFORMAT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentityAccountAttributeFormat(values []IdentityAccountAttributeFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentityAccountAttributeFormat) isMultiValue() bool {
    return false
}
