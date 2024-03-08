package serviceaccount
import (
    "errors"
)
type PostAdditionalQueryParameterType int

const (
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    GROUPS_POSTADDITIONALQUERYPARAMETERTYPE
    SECRET_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"audit", "groups", "secret"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_POSTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_POSTADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown PostAdditionalQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializePostAdditionalQueryParameterType(values []PostAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PostAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
