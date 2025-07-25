// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package item
type PutAdditionalQueryParameterType int

const (
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"audit"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
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
