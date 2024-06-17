package models
type GroupGroupVisibility int

const (
    ALL_GROUPGROUPVISIBILITY GroupGroupVisibility = iota
    PUBLIC_GROUPGROUPVISIBILITY
    PUBLIC_OR_MEMBER_GROUPGROUPVISIBILITY
)

func (i GroupGroupVisibility) String() string {
    return []string{"ALL", "PUBLIC", "PUBLIC_OR_MEMBER"}[i]
}
func ParseGroupGroupVisibility(v string) (any, error) {
    result := ALL_GROUPGROUPVISIBILITY
    switch v {
        case "ALL":
            result = ALL_GROUPGROUPVISIBILITY
        case "PUBLIC":
            result = PUBLIC_GROUPGROUPVISIBILITY
        case "PUBLIC_OR_MEMBER":
            result = PUBLIC_OR_MEMBER_GROUPGROUPVISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGroupGroupVisibility(values []GroupGroupVisibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupGroupVisibility) isMultiValue() bool {
    return false
}
