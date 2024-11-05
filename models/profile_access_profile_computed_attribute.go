package models
type ProfileAccessProfileComputedAttribute int

const (
    USERNAME_PROFILEACCESSPROFILECOMPUTEDATTRIBUTE ProfileAccessProfileComputedAttribute = iota
    EMAIL_PROFILEACCESSPROFILECOMPUTEDATTRIBUTE
    LICENSE_PROFILEACCESSPROFILECOMPUTEDATTRIBUTE
)

func (i ProfileAccessProfileComputedAttribute) String() string {
    return []string{"USERNAME", "EMAIL", "LICENSE"}[i]
}
func ParseProfileAccessProfileComputedAttribute(v string) (any, error) {
    result := USERNAME_PROFILEACCESSPROFILECOMPUTEDATTRIBUTE
    switch v {
        case "USERNAME":
            result = USERNAME_PROFILEACCESSPROFILECOMPUTEDATTRIBUTE
        case "EMAIL":
            result = EMAIL_PROFILEACCESSPROFILECOMPUTEDATTRIBUTE
        case "LICENSE":
            result = LICENSE_PROFILEACCESSPROFILECOMPUTEDATTRIBUTE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProfileAccessProfileComputedAttribute(values []ProfileAccessProfileComputedAttribute) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProfileAccessProfileComputedAttribute) isMultiValue() bool {
    return false
}
