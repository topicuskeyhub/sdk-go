package models
type ProfileAccessProfileComputedAttributeStatus int

const (
    IN_SYNC_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS ProfileAccessProfileComputedAttributeStatus = iota
    OUT_OF_SYNC_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
    INVALID_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
    DUPLICATE_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
    ERROR_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
    UNKNOWN_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
)

func (i ProfileAccessProfileComputedAttributeStatus) String() string {
    return []string{"IN_SYNC", "OUT_OF_SYNC", "INVALID", "DUPLICATE", "ERROR", "UNKNOWN"}[i]
}
func ParseProfileAccessProfileComputedAttributeStatus(v string) (any, error) {
    result := IN_SYNC_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
    switch v {
        case "IN_SYNC":
            result = IN_SYNC_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
        case "OUT_OF_SYNC":
            result = OUT_OF_SYNC_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
        case "INVALID":
            result = INVALID_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
        case "DUPLICATE":
            result = DUPLICATE_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
        case "ERROR":
            result = ERROR_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
        case "UNKNOWN":
            result = UNKNOWN_PROFILEACCESSPROFILECOMPUTEDATTRIBUTESTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProfileAccessProfileComputedAttributeStatus(values []ProfileAccessProfileComputedAttributeStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProfileAccessProfileComputedAttributeStatus) isMultiValue() bool {
    return false
}
