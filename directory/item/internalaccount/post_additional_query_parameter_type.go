package internalaccount
type PostAdditionalQueryParameterType int

const (
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    REASON_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"audit", "reason"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "reason":
            result = REASON_POSTADDITIONALQUERYPARAMETERTYPE
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
