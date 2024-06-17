package profile
type PostAdditionalQueryParameterType int

const (
    ACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"accounts", "audit"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accounts":
            result = ACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        default:
            return nil, nil
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
