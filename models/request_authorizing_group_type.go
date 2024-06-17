package models
type RequestAuthorizingGroupType int

const (
    PROVISIONING_REQUESTAUTHORIZINGGROUPTYPE RequestAuthorizingGroupType = iota
    MEMBERSHIP_REQUESTAUTHORIZINGGROUPTYPE
    DELEGATION_REQUESTAUTHORIZINGGROUPTYPE
    AUDITING_REQUESTAUTHORIZINGGROUPTYPE
)

func (i RequestAuthorizingGroupType) String() string {
    return []string{"PROVISIONING", "MEMBERSHIP", "DELEGATION", "AUDITING"}[i]
}
func ParseRequestAuthorizingGroupType(v string) (any, error) {
    result := PROVISIONING_REQUESTAUTHORIZINGGROUPTYPE
    switch v {
        case "PROVISIONING":
            result = PROVISIONING_REQUESTAUTHORIZINGGROUPTYPE
        case "MEMBERSHIP":
            result = MEMBERSHIP_REQUESTAUTHORIZINGGROUPTYPE
        case "DELEGATION":
            result = DELEGATION_REQUESTAUTHORIZINGGROUPTYPE
        case "AUDITING":
            result = AUDITING_REQUESTAUTHORIZINGGROUPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRequestAuthorizingGroupType(values []RequestAuthorizingGroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequestAuthorizingGroupType) isMultiValue() bool {
    return false
}
