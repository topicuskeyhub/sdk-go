package models
type DirectoryOIDCVendor int

const (
    GENERIC_DIRECTORYOIDCVENDOR DirectoryOIDCVendor = iota
    GOOGLE_DIRECTORYOIDCVENDOR
    AZURE_AD_DIRECTORYOIDCVENDOR
)

func (i DirectoryOIDCVendor) String() string {
    return []string{"GENERIC", "GOOGLE", "AZURE_AD"}[i]
}
func ParseDirectoryOIDCVendor(v string) (any, error) {
    result := GENERIC_DIRECTORYOIDCVENDOR
    switch v {
        case "GENERIC":
            result = GENERIC_DIRECTORYOIDCVENDOR
        case "GOOGLE":
            result = GOOGLE_DIRECTORYOIDCVENDOR
        case "AZURE_AD":
            result = AZURE_AD_DIRECTORYOIDCVENDOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDirectoryOIDCVendor(values []DirectoryOIDCVendor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DirectoryOIDCVendor) isMultiValue() bool {
    return false
}
