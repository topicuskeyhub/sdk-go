package models
type AuthPasswordMode int

const (
    UNSET_AUTHPASSWORDMODE AuthPasswordMode = iota
    ONE_PASSWORD_UNSYNCED_AUTHPASSWORDMODE
    ONE_PASSWORD_AUTHPASSWORDMODE
    TWO_PASSWORDS_AUTHPASSWORDMODE
)

func (i AuthPasswordMode) String() string {
    return []string{"UNSET", "ONE_PASSWORD_UNSYNCED", "ONE_PASSWORD", "TWO_PASSWORDS"}[i]
}
func ParseAuthPasswordMode(v string) (any, error) {
    result := UNSET_AUTHPASSWORDMODE
    switch v {
        case "UNSET":
            result = UNSET_AUTHPASSWORDMODE
        case "ONE_PASSWORD_UNSYNCED":
            result = ONE_PASSWORD_UNSYNCED_AUTHPASSWORDMODE
        case "ONE_PASSWORD":
            result = ONE_PASSWORD_AUTHPASSWORDMODE
        case "TWO_PASSWORDS":
            result = TWO_PASSWORDS_AUTHPASSWORDMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthPasswordMode(values []AuthPasswordMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthPasswordMode) isMultiValue() bool {
    return false
}
