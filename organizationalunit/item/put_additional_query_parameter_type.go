package item
type PutAdditionalQueryParameterType int

const (
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    CREATEASPARENTOF_PUTADDITIONALQUERYPARAMETERTYPE
    SETTINGS_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"audit", "createAsParentOf", "settings"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "createAsParentOf":
            result = CREATEASPARENTOF_PUTADDITIONALQUERYPARAMETERTYPE
        case "settings":
            result = SETTINGS_PUTADDITIONALQUERYPARAMETERTYPE
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
