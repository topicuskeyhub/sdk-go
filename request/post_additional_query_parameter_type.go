package request
type PostAdditionalQueryParameterType int

const (
    ACCEPTPARAMS_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    RESETSTATUS_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"acceptParams", "audit", "resetStatus"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := ACCEPTPARAMS_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "acceptParams":
            result = ACCEPTPARAMS_POSTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "resetStatus":
            result = RESETSTATUS_POSTADDITIONALQUERYPARAMETERTYPE
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
