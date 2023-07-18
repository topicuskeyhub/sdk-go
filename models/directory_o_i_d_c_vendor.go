package models
import (
    "errors"
)
// 
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
            return 0, errors.New("Unknown DirectoryOIDCVendor value: " + v)
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
