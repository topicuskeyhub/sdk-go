package account
import (
    "errors"
)
type GetAdditionalQueryParameterType int

const (
    ACTIVELOGIN_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    GROUPS_GETADDITIONALQUERYPARAMETERTYPE
    PENDINGRECOVERYREQUESTS_GETADDITIONALQUERYPARAMETERTYPE
    SETTINGS_GETADDITIONALQUERYPARAMETERTYPE
    STOREDATTRIBUTES_GETADDITIONALQUERYPARAMETERTYPE
    VAULT_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"activeLogin", "audit", "groups", "pendingRecoveryRequests", "settings", "storedAttributes", "vault"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := ACTIVELOGIN_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "activeLogin":
            result = ACTIVELOGIN_GETADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_GETADDITIONALQUERYPARAMETERTYPE
        case "pendingRecoveryRequests":
            result = PENDINGRECOVERYREQUESTS_GETADDITIONALQUERYPARAMETERTYPE
        case "settings":
            result = SETTINGS_GETADDITIONALQUERYPARAMETERTYPE
        case "storedAttributes":
            result = STOREDATTRIBUTES_GETADDITIONALQUERYPARAMETERTYPE
        case "vault":
            result = VAULT_GETADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetAdditionalQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetAdditionalQueryParameterType(values []GetAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
