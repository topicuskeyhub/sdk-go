package directory
type GetAdditionalQueryParameterType int

const (
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    MARKERS_GETADDITIONALQUERYPARAMETERTYPE
    STATUS_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"audit", "markers", "status"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_GETADDITIONALQUERYPARAMETERTYPE
        case "status":
            result = STATUS_GETADDITIONALQUERYPARAMETERTYPE
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
