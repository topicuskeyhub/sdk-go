package models
type RequestUpdateGroupMembershipType int

const (
    ADD_REQUESTUPDATEGROUPMEMBERSHIPTYPE RequestUpdateGroupMembershipType = iota
    MODIFY_REQUESTUPDATEGROUPMEMBERSHIPTYPE
    REMOVE_REQUESTUPDATEGROUPMEMBERSHIPTYPE
)

func (i RequestUpdateGroupMembershipType) String() string {
    return []string{"ADD", "MODIFY", "REMOVE"}[i]
}
func ParseRequestUpdateGroupMembershipType(v string) (any, error) {
    result := ADD_REQUESTUPDATEGROUPMEMBERSHIPTYPE
    switch v {
        case "ADD":
            result = ADD_REQUESTUPDATEGROUPMEMBERSHIPTYPE
        case "MODIFY":
            result = MODIFY_REQUESTUPDATEGROUPMEMBERSHIPTYPE
        case "REMOVE":
            result = REMOVE_REQUESTUPDATEGROUPMEMBERSHIPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRequestUpdateGroupMembershipType(values []RequestUpdateGroupMembershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RequestUpdateGroupMembershipType) isMultiValue() bool {
    return false
}
