package item
type GetAdditionalQueryParameterType int

const (
    ACCESSPROFILECLIENTS_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    DELETETILE_GETADDITIONALQUERYPARAMETERTYPE
    GROUPCLIENTS_GETADDITIONALQUERYPARAMETERTYPE
    GROUPS_GETADDITIONALQUERYPARAMETERTYPE
    ORGANIZATIONALUNITS_GETADDITIONALQUERYPARAMETERTYPE
    SECRET_GETADDITIONALQUERYPARAMETERTYPE
    TILE_GETADDITIONALQUERYPARAMETERTYPE
    VAULTRECORDCOUNT_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"accessprofileclients", "audit", "deleteTile", "groupclients", "groups", "organizationalUnits", "secret", "tile", "vaultRecordCount"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := ACCESSPROFILECLIENTS_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accessprofileclients":
            result = ACCESSPROFILECLIENTS_GETADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_GETADDITIONALQUERYPARAMETERTYPE
        case "groupclients":
            result = GROUPCLIENTS_GETADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_GETADDITIONALQUERYPARAMETERTYPE
        case "organizationalUnits":
            result = ORGANIZATIONALUNITS_GETADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_GETADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_GETADDITIONALQUERYPARAMETERTYPE
        case "vaultRecordCount":
            result = VAULTRECORDCOUNT_GETADDITIONALQUERYPARAMETERTYPE
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
