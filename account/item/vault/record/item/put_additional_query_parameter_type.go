package item
type PutAdditionalQueryParameterType int

const (
    ACTIVATIONSTATUS_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
    PARENT_PUTADDITIONALQUERYPARAMETERTYPE
    PASSWORDMETADATA_PUTADDITIONALQUERYPARAMETERTYPE
    SECRET_PUTADDITIONALQUERYPARAMETERTYPE
    SHARESUMMARY_PUTADDITIONALQUERYPARAMETERTYPE
    SHARES_PUTADDITIONALQUERYPARAMETERTYPE
    TILE_PUTADDITIONALQUERYPARAMETERTYPE
    VAULTHOLDER_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"activationStatus", "audit", "deleteTile", "parent", "passwordMetadata", "secret", "shareSummary", "shares", "tile", "vaultholder"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACTIVATIONSTATUS_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "activationStatus":
            result = ACTIVATIONSTATUS_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "deleteTile":
            result = DELETETILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "parent":
            result = PARENT_PUTADDITIONALQUERYPARAMETERTYPE
        case "passwordMetadata":
            result = PASSWORDMETADATA_PUTADDITIONALQUERYPARAMETERTYPE
        case "secret":
            result = SECRET_PUTADDITIONALQUERYPARAMETERTYPE
        case "shareSummary":
            result = SHARESUMMARY_PUTADDITIONALQUERYPARAMETERTYPE
        case "shares":
            result = SHARES_PUTADDITIONALQUERYPARAMETERTYPE
        case "tile":
            result = TILE_PUTADDITIONALQUERYPARAMETERTYPE
        case "vaultholder":
            result = VAULTHOLDER_PUTADDITIONALQUERYPARAMETERTYPE
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
