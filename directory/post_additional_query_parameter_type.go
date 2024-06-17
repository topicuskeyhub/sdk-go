package directory
type PostAdditionalQueryParameterType int

const (
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    MARKERS_POSTADDITIONALQUERYPARAMETERTYPE
    STATUS_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"audit", "markers", "status"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_POSTADDITIONALQUERYPARAMETERTYPE
        case "status":
            result = STATUS_POSTADDITIONALQUERYPARAMETERTYPE
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
