package item
type GetAdditionalQueryParameterType int

const (
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    GROUPS_GETADDITIONALQUERYPARAMETERTYPE
    SECRET_GETADDITIONALQUERYPARAMETERTYPE
    SUPPORTEDFEATURES_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"audit", "groups", "secret", "supportedFeatures"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_GETADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_GETADDITIONALQUERYPARAMETERTYPE
        case "supportedFeatures":
            result = SUPPORTEDFEATURES_GETADDITIONALQUERYPARAMETERTYPE
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
