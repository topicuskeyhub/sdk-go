package profile
type GetAdditionalQueryParameterType int

const (
    ACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    PROVISIONING_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"accounts", "audit", "provisioning"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accounts":
            result = ACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "provisioning":
            result = PROVISIONING_GETADDITIONALQUERYPARAMETERTYPE
        default:
            return nil, nil
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
