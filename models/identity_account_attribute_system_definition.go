package models
type IdentityAccountAttributeSystemDefinition int

const (
    USERNAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION IdentityAccountAttributeSystemDefinition = iota
    DISPLAY_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    GIVEN_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    MIDDLE_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    FAMILY_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    EMAIL_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    LICENSE_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    TELEPHONE_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    CUSTOM_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
)

func (i IdentityAccountAttributeSystemDefinition) String() string {
    return []string{"USERNAME", "DISPLAY_NAME", "GIVEN_NAME", "MIDDLE_NAME", "FAMILY_NAME", "EMAIL", "LICENSE", "TELEPHONE", "CUSTOM"}[i]
}
func ParseIdentityAccountAttributeSystemDefinition(v string) (any, error) {
    result := USERNAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
    switch v {
        case "USERNAME":
            result = USERNAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "DISPLAY_NAME":
            result = DISPLAY_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "GIVEN_NAME":
            result = GIVEN_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "MIDDLE_NAME":
            result = MIDDLE_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "FAMILY_NAME":
            result = FAMILY_NAME_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "EMAIL":
            result = EMAIL_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "LICENSE":
            result = LICENSE_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "TELEPHONE":
            result = TELEPHONE_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        case "CUSTOM":
            result = CUSTOM_IDENTITYACCOUNTATTRIBUTESYSTEMDEFINITION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentityAccountAttributeSystemDefinition(values []IdentityAccountAttributeSystemDefinition) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentityAccountAttributeSystemDefinition) isMultiValue() bool {
    return false
}
