// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type ProfileAccessProfileAccountActivation int

const (
    INACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION ProfileAccessProfileAccountActivation = iota
    ACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
    FORCE_INACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
    FORCE_ACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
)

func (i ProfileAccessProfileAccountActivation) String() string {
    return []string{"INACTIVE", "ACTIVE", "FORCE_INACTIVE", "FORCE_ACTIVE"}[i]
}
func ParseProfileAccessProfileAccountActivation(v string) (any, error) {
    result := INACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
    switch v {
        case "INACTIVE":
            result = INACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
        case "ACTIVE":
            result = ACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
        case "FORCE_INACTIVE":
            result = FORCE_INACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
        case "FORCE_ACTIVE":
            result = FORCE_ACTIVE_PROFILEACCESSPROFILEACCOUNTACTIVATION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProfileAccessProfileAccountActivation(values []ProfileAccessProfileAccountActivation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProfileAccessProfileAccountActivation) isMultiValue() bool {
    return false
}
