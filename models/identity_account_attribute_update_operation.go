package models
type IdentityAccountAttributeUpdateOperation int

const (
    SYNC_ONLY_IDENTITYACCOUNTATTRIBUTEUPDATEOPERATION IdentityAccountAttributeUpdateOperation = iota
    UPDATE_AUTOMATICALLY_IDENTITYACCOUNTATTRIBUTEUPDATEOPERATION
    UPDATE_FULLY_IDENTITYACCOUNTATTRIBUTEUPDATEOPERATION
)

func (i IdentityAccountAttributeUpdateOperation) String() string {
    return []string{"SYNC_ONLY", "UPDATE_AUTOMATICALLY", "UPDATE_FULLY"}[i]
}
func ParseIdentityAccountAttributeUpdateOperation(v string) (any, error) {
    result := SYNC_ONLY_IDENTITYACCOUNTATTRIBUTEUPDATEOPERATION
    switch v {
        case "SYNC_ONLY":
            result = SYNC_ONLY_IDENTITYACCOUNTATTRIBUTEUPDATEOPERATION
        case "UPDATE_AUTOMATICALLY":
            result = UPDATE_AUTOMATICALLY_IDENTITYACCOUNTATTRIBUTEUPDATEOPERATION
        case "UPDATE_FULLY":
            result = UPDATE_FULLY_IDENTITYACCOUNTATTRIBUTEUPDATEOPERATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentityAccountAttributeUpdateOperation(values []IdentityAccountAttributeUpdateOperation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentityAccountAttributeUpdateOperation) isMultiValue() bool {
    return false
}
