package models
type GroupGroupClassificationConformance int

const (
    MATCHES_ALL_CRITERIA_GROUPGROUPCLASSIFICATIONCONFORMANCE GroupGroupClassificationConformance = iota
    UNMATCHED_FIXABLE_CRITERIA_GROUPGROUPCLASSIFICATIONCONFORMANCE
    UNMATCHED_CRITERIA_GROUPGROUPCLASSIFICATIONCONFORMANCE
)

func (i GroupGroupClassificationConformance) String() string {
    return []string{"MATCHES_ALL_CRITERIA", "UNMATCHED_FIXABLE_CRITERIA", "UNMATCHED_CRITERIA"}[i]
}
func ParseGroupGroupClassificationConformance(v string) (any, error) {
    result := MATCHES_ALL_CRITERIA_GROUPGROUPCLASSIFICATIONCONFORMANCE
    switch v {
        case "MATCHES_ALL_CRITERIA":
            result = MATCHES_ALL_CRITERIA_GROUPGROUPCLASSIFICATIONCONFORMANCE
        case "UNMATCHED_FIXABLE_CRITERIA":
            result = UNMATCHED_FIXABLE_CRITERIA_GROUPGROUPCLASSIFICATIONCONFORMANCE
        case "UNMATCHED_CRITERIA":
            result = UNMATCHED_CRITERIA_GROUPGROUPCLASSIFICATIONCONFORMANCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGroupGroupClassificationConformance(values []GroupGroupClassificationConformance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupGroupClassificationConformance) isMultiValue() bool {
    return false
}
