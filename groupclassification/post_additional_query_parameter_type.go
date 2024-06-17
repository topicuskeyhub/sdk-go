package groupclassification
type PostAdditionalQueryParameterType int

const (
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    INFO_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"audit", "info"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "info":
            result = INFO_POSTADDITIONALQUERYPARAMETERTYPE
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
