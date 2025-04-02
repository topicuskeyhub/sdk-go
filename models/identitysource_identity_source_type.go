package models
type IdentitysourceIdentitySourceType int

const (
    AFAS_IDENTITYSOURCEIDENTITYSOURCETYPE IdentitysourceIdentitySourceType = iota
)

func (i IdentitysourceIdentitySourceType) String() string {
    return []string{"AFAS"}[i]
}
func ParseIdentitysourceIdentitySourceType(v string) (any, error) {
    result := AFAS_IDENTITYSOURCEIDENTITYSOURCETYPE
    switch v {
        case "AFAS":
            result = AFAS_IDENTITYSOURCEIDENTITYSOURCETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIdentitysourceIdentitySourceType(values []IdentitysourceIdentitySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IdentitysourceIdentitySourceType) isMultiValue() bool {
    return false
}
