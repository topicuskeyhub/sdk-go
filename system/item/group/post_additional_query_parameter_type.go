package group
type PostAdditionalQueryParameterType int

const (
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    PROVGROUPS_POSTADDITIONALQUERYPARAMETERTYPE
    SERVICEACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"audit", "provgroups", "serviceAccounts"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
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
