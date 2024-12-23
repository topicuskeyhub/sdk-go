package profile
type PostAdditionalQueryParameterType int

const (
    ACCOUNTSWITHATTRIBUTES_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    ATTRIBUTERULES_POSTADDITIONALQUERYPARAMETERTYPE
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    PROVISIONING_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"accountsWithAttributes", "attributeRules", "audit", "provisioning"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTSWITHATTRIBUTES_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accountsWithAttributes":
            result = ACCOUNTSWITHATTRIBUTES_POSTADDITIONALQUERYPARAMETERTYPE
        case "attributeRules":
            result = ATTRIBUTERULES_POSTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "provisioning":
            result = PROVISIONING_POSTADDITIONALQUERYPARAMETERTYPE
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
