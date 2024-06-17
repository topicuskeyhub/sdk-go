package item
type GetAdditionalQueryParameterType int

const (
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    CREATEASPARENTOF_GETADDITIONALQUERYPARAMETERTYPE
    SETTINGS_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"audit", "createAsParentOf", "settings"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "createAsParentOf":
            result = CREATEASPARENTOF_GETADDITIONALQUERYPARAMETERTYPE
        case "settings":
            result = SETTINGS_GETADDITIONALQUERYPARAMETERTYPE
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
