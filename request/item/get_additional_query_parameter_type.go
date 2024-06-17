package item
type GetAdditionalQueryParameterType int

const (
    ACCEPTPARAMS_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    RESETSTATUS_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"acceptParams", "audit", "resetStatus"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := ACCEPTPARAMS_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "acceptParams":
            result = ACCEPTPARAMS_GETADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "resetStatus":
            result = RESETSTATUS_GETADDITIONALQUERYPARAMETERTYPE
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
