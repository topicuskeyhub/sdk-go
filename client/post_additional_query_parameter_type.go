package client
type PostAdditionalQueryParameterType int

const (
    ACCESSPROFILECLIENTS_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    DELETETILE_POSTADDITIONALQUERYPARAMETERTYPE
    GROUPCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
    GROUPS_POSTADDITIONALQUERYPARAMETERTYPE
    ORGANIZATIONALUNITS_POSTADDITIONALQUERYPARAMETERTYPE
    SECRET_POSTADDITIONALQUERYPARAMETERTYPE
    TILE_POSTADDITIONALQUERYPARAMETERTYPE
    VAULTRECORDCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"accessprofileclients", "audit", "deleteTile", "groupclients", "groups", "organizationalUnits", "secret", "tile", "vaultRecordCount"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := ACCESSPROFILECLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accessprofileclients":
            result = ACCESSPROFILECLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_POSTADDITIONALQUERYPARAMETERTYPE
        case "groupclients":
            result = GROUPCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_POSTADDITIONALQUERYPARAMETERTYPE
        case "organizationalUnits":
            result = ORGANIZATIONALUNITS_POSTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_POSTADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_POSTADDITIONALQUERYPARAMETERTYPE
        case "vaultRecordCount":
            result = VAULTRECORDCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
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
