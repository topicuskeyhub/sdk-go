package models
type AuthAccountValidity int

const (
    VALID_AUTHACCOUNTVALIDITY AuthAccountValidity = iota
    NOT_APPLICABLE_AUTHACCOUNTVALIDITY
    REREGISTRATION_REQUIRED_AUTHACCOUNTVALIDITY
    DISABLED_AUTHACCOUNTVALIDITY
    DIRECTORY_DISABLED_AUTHACCOUNTVALIDITY
    INVALID_IN_DIRECTORY_AUTHACCOUNTVALIDITY
)

func (i AuthAccountValidity) String() string {
    return []string{"VALID", "NOT_APPLICABLE", "REREGISTRATION_REQUIRED", "DISABLED", "DIRECTORY_DISABLED", "INVALID_IN_DIRECTORY"}[i]
}
func ParseAuthAccountValidity(v string) (any, error) {
    result := VALID_AUTHACCOUNTVALIDITY
    switch v {
        case "VALID":
            result = VALID_AUTHACCOUNTVALIDITY
        case "NOT_APPLICABLE":
            result = NOT_APPLICABLE_AUTHACCOUNTVALIDITY
        case "REREGISTRATION_REQUIRED":
            result = REREGISTRATION_REQUIRED_AUTHACCOUNTVALIDITY
        case "DISABLED":
            result = DISABLED_AUTHACCOUNTVALIDITY
        case "DIRECTORY_DISABLED":
            result = DIRECTORY_DISABLED_AUTHACCOUNTVALIDITY
        case "INVALID_IN_DIRECTORY":
            result = INVALID_IN_DIRECTORY_AUTHACCOUNTVALIDITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAuthAccountValidity(values []AuthAccountValidity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthAccountValidity) isMultiValue() bool {
    return false
}
