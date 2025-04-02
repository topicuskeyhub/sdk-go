package item
type PutAdditionalQueryParameterType int

const (
    ACCESSPROFILECLIENTS_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
    GROUPCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
    GROUPS_PUTADDITIONALQUERYPARAMETERTYPE
    ORGANIZATIONALUNITS_PUTADDITIONALQUERYPARAMETERTYPE
    SECRET_PUTADDITIONALQUERYPARAMETERTYPE
    TILE_PUTADDITIONALQUERYPARAMETERTYPE
    VAULTRECORDCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"accessprofileclients", "audit", "deleteTile", "groupclients", "groups", "organizationalUnits", "secret", "tile", "vaultRecordCount"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACCESSPROFILECLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accessprofileclients":
            result = ACCESSPROFILECLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "groupclients":
            result = GROUPCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "groups":
            result = GROUPS_PUTADDITIONALQUERYPARAMETERTYPE
        case "organizationalUnits":
            result = ORGANIZATIONALUNITS_PUTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_PUTADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "vaultRecordCount":
            result = VAULTRECORDCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
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
