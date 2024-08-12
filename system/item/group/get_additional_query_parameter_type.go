package group
type GetAdditionalQueryParameterType int

const (
    ACCESSPROFILEPROVISIONING_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    PROVGROUPS_GETADDITIONALQUERYPARAMETERTYPE
    SERVICEACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"accessProfileProvisioning", "audit", "provgroups", "serviceAccounts"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := ACCESSPROFILEPROVISIONING_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accessProfileProvisioning":
            result = ACCESSPROFILEPROVISIONING_GETADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "provgroups":
            result = PROVGROUPS_GETADDITIONALQUERYPARAMETERTYPE
        case "serviceAccounts":
            result = SERVICEACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
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
