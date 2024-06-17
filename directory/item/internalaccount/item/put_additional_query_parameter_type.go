package item
type PutAdditionalQueryParameterType int

const (
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    REASON_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"audit", "reason"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "reason":
            result = REASON_PUTADDITIONALQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePutAdditionalQueryParameterType(values []PutAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PutAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
