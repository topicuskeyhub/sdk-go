package item
type PutAdditionalQueryParameterType int

const (
    ACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    PROVISIONING_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"accounts", "audit", "provisioning"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accounts":
            result = ACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
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
