package group
type PostAdditionalQueryParameterType int

const (
    ACCESSPROFILEPROVISIONING_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    PROVGROUPS_POSTADDITIONALQUERYPARAMETERTYPE
    SERVICEACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"accessProfileProvisioning", "audit", "provgroups", "serviceAccounts"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := ACCESSPROFILEPROVISIONING_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accessProfileProvisioning":
            result = ACCESSPROFILEPROVISIONING_POSTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "provgroups":
            result = PROVGROUPS_POSTADDITIONALQUERYPARAMETERTYPE
        case "serviceAccounts":
            result = SERVICEACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
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
