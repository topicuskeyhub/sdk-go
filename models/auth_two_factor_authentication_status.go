package models
import (
    "errors"
)
type AuthTwoFactorAuthenticationStatus int

const (
    DISABLED_AUTHTWOFACTORAUTHENTICATIONSTATUS AuthTwoFactorAuthenticationStatus = iota
    EXTERNAL_AUTHTWOFACTORAUTHENTICATIONSTATUS
    TOTP_AUTHTWOFACTORAUTHENTICATIONSTATUS
    APP_AUTHTWOFACTORAUTHENTICATIONSTATUS
    WEBAUTHN_AUTHTWOFACTORAUTHENTICATIONSTATUS
    MIXED_AUTHTWOFACTORAUTHENTICATIONSTATUS
)

func (i AuthTwoFactorAuthenticationStatus) String() string {
    return []string{"DISABLED", "EXTERNAL", "TOTP", "APP", "WEBAUTHN", "MIXED"}[i]
}
func ParseAuthTwoFactorAuthenticationStatus(v string) (any, error) {
    result := DISABLED_AUTHTWOFACTORAUTHENTICATIONSTATUS
    switch v {
        case "DISABLED":
            result = DISABLED_AUTHTWOFACTORAUTHENTICATIONSTATUS
        case "EXTERNAL":
            result = EXTERNAL_AUTHTWOFACTORAUTHENTICATIONSTATUS
        case "TOTP":
            result = TOTP_AUTHTWOFACTORAUTHENTICATIONSTATUS
        case "APP":
            result = APP_AUTHTWOFACTORAUTHENTICATIONSTATUS
        case "WEBAUTHN":
            result = WEBAUTHN_AUTHTWOFACTORAUTHENTICATIONSTATUS
        case "MIXED":
            result = MIXED_AUTHTWOFACTORAUTHENTICATIONSTATUS
        default:
            return 0, errors.New("Unknown AuthTwoFactorAuthenticationStatus value: " + v)
    }
    return &result, nil
}
func SerializeAuthTwoFactorAuthenticationStatus(values []AuthTwoFactorAuthenticationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthTwoFactorAuthenticationStatus) isMultiValue() bool {
    return false
}
