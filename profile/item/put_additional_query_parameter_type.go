// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package item
type PutAdditionalQueryParameterType int

const (
    ACCOUNTSWITHATTRIBUTES_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    ATTRIBUTERULES_PUTADDITIONALQUERYPARAMETERTYPE
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    CLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
    GROUPS_PUTADDITIONALQUERYPARAMETERTYPE
    PROVISIONING_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"accountsWithAttributes", "attributeRules", "audit", "clients", "groups", "provisioning"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTSWITHATTRIBUTES_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accountsWithAttributes":
            result = ACCOUNTSWITHATTRIBUTES_PUTADDITIONALQUERYPARAMETERTYPE
        case "attributeRules":
            result = ATTRIBUTERULES_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "clients":
            result = CLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_PUTADDITIONALQUERYPARAMETERTYPE
        case "provisioning":
            result = PROVISIONING_PUTADDITIONALQUERYPARAMETERTYPE
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
