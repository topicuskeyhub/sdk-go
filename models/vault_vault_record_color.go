package models
import (
    "errors"
)
type VaultVaultRecordColor int

const (
    NONE_VAULTVAULTRECORDCOLOR VaultVaultRecordColor = iota
    GREEN_VAULTVAULTRECORDCOLOR
    RED_VAULTVAULTRECORDCOLOR
    BLUE_VAULTVAULTRECORDCOLOR
    DARK_VAULTVAULTRECORDCOLOR
    PINK_LAVENDER_VAULTVAULTRECORDCOLOR
    CRIMSON_RED_VAULTVAULTRECORDCOLOR
    MIDDLE_YELLOW_VAULTVAULTRECORDCOLOR
    ANDROID_GREEN_VAULTVAULTRECORDCOLOR
    SAGE_VAULTVAULTRECORDCOLOR
    ARTICHOKE_VAULTVAULTRECORDCOLOR
)

func (i VaultVaultRecordColor) String() string {
    return []string{"NONE", "GREEN", "RED", "BLUE", "DARK", "PINK_LAVENDER", "CRIMSON_RED", "MIDDLE_YELLOW", "ANDROID_GREEN", "SAGE", "ARTICHOKE"}[i]
}
func ParseVaultVaultRecordColor(v string) (any, error) {
    result := NONE_VAULTVAULTRECORDCOLOR
    switch v {
        case "NONE":
            result = NONE_VAULTVAULTRECORDCOLOR
        case "GREEN":
            result = GREEN_VAULTVAULTRECORDCOLOR
        case "RED":
            result = RED_VAULTVAULTRECORDCOLOR
        case "BLUE":
            result = BLUE_VAULTVAULTRECORDCOLOR
        case "DARK":
            result = DARK_VAULTVAULTRECORDCOLOR
        case "PINK_LAVENDER":
            result = PINK_LAVENDER_VAULTVAULTRECORDCOLOR
        case "CRIMSON_RED":
            result = CRIMSON_RED_VAULTVAULTRECORDCOLOR
        case "MIDDLE_YELLOW":
            result = MIDDLE_YELLOW_VAULTVAULTRECORDCOLOR
        case "ANDROID_GREEN":
            result = ANDROID_GREEN_VAULTVAULTRECORDCOLOR
        case "SAGE":
            result = SAGE_VAULTVAULTRECORDCOLOR
        case "ARTICHOKE":
            result = ARTICHOKE_VAULTVAULTRECORDCOLOR
        default:
            return 0, errors.New("Unknown VaultVaultRecordColor value: " + v)
    }
    return &result, nil
}
func SerializeVaultVaultRecordColor(values []VaultVaultRecordColor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VaultVaultRecordColor) isMultiValue() bool {
    return false
}
