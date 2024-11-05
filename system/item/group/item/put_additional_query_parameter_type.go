package item
type PutAdditionalQueryParameterType int

const (
    ACCESSPROFILEPROVISIONING_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    PROVGROUPS_PUTADDITIONALQUERYPARAMETERTYPE
    SERVICEACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"accessProfileProvisioning", "audit", "provgroups", "serviceAccounts"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACCESSPROFILEPROVISIONING_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accessProfileProvisioning":
            result = ACCESSPROFILEPROVISIONING_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "provgroups":
            result = PROVGROUPS_PUTADDITIONALQUERYPARAMETERTYPE
        case "serviceAccounts":
            result = SERVICEACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
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