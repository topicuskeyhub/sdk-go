package models
import (
    "errors"
)
// 
type AuthInternalAccountStatus int

const (
    UNVERIFIED_AUTHINTERNALACCOUNTSTATUS AuthInternalAccountStatus = iota
    NOT_ACTIVATED_AUTHINTERNALACCOUNTSTATUS
    ACTIVATION_EXPIRED_AUTHINTERNALACCOUNTSTATUS
    ACTIVATED_AUTHINTERNALACCOUNTSTATUS
)

func (i AuthInternalAccountStatus) String() string {
    return []string{"UNVERIFIED", "NOT_ACTIVATED", "ACTIVATION_EXPIRED", "ACTIVATED"}[i]
}
func ParseAuthInternalAccountStatus(v string) (any, error) {
    result := UNVERIFIED_AUTHINTERNALACCOUNTSTATUS
    switch v {
        case "UNVERIFIED":
            result = UNVERIFIED_AUTHINTERNALACCOUNTSTATUS
        case "NOT_ACTIVATED":
            result = NOT_ACTIVATED_AUTHINTERNALACCOUNTSTATUS
        case "ACTIVATION_EXPIRED":
            result = ACTIVATION_EXPIRED_AUTHINTERNALACCOUNTSTATUS
        case "ACTIVATED":
            result = ACTIVATED_AUTHINTERNALACCOUNTSTATUS
        default:
            return 0, errors.New("Unknown AuthInternalAccountStatus value: " + v)
    }
    return &result, nil
}
func SerializeAuthInternalAccountStatus(values []AuthInternalAccountStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
