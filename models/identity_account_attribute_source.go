package models
type IdentityAccountAttributeSource int

const (
    SCIM_IDENTITYACCOUNTATTRIBUTESOURCE IdentityAccountAttributeSource = iota
    EXTERNAL_IDENTITYACCOUNTATTRIBUTESOURCE
    DIRECTORY_IDENTITYACCOUNTATTRIBUTESOURCE
    FORMULA_IDENTITYACCOUNTATTRIBUTESOURCE
    DEFAULT_IDENTITYACCOUNTATTRIBUTESOURCE
    SELF_SERVICE_IDENTITYACCOUNTATTRIBUTESOURCE
    OVERRIDE_IDENTITYACCOUNTATTRIBUTESOURCE
)

func (i IdentityAccountAttributeSource) String() string {
    return []string{"SCIM", "EXTERNAL", "DIRECTORY", "FORMULA", "DEFAULT", "SELF_SERVICE", "OVERRIDE"}[i]
}
func ParseIdentityAccountAttributeSource(v string) (any, error) {
    result := SCIM_IDENTITYACCOUNTATTRIBUTESOURCE
    switch v {
        case "SCIM":
            result = SCIM_IDENTITYACCOUNTATTRIBUTESOURCE
        case "EXTERNAL":
            result = EXTERNAL_IDENTITYACCOUNTATTRIBUTESOURCE
        case "DIRECTORY":
            result = DIRECTORY_IDENTITYACCOUNTATTRIBUTESOURCE
        case "FORMULA":
            result = FORMULA_IDENTITYACCOUNTATTRIBUTESOURCE
        case "DEFAULT":
            result = DEFAULT_IDENTITYACCOUNTATTRIBUTESOURCE
        case "SELF_SERVICE":
            result = SELF_SERVICE_IDENTITYACCOUNTATTRIBUTESOURCE
        case "OVERRIDE":
            result = OVERRIDE_IDENTITYACCOUNTATTRIBUTESOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentityAccountAttributeSource(values []IdentityAccountAttributeSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentityAccountAttributeSource) isMultiValue() bool {
    return false
}