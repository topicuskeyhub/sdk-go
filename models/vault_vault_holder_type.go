package models
import (
    "errors"
)
type VaultVaultHolderType int

const (
    ACCOUNT_VAULTVAULTHOLDERTYPE VaultVaultHolderType = iota
    CLIENT_VAULTVAULTHOLDERTYPE
    GROUP_VAULTVAULTHOLDERTYPE
)

func (i VaultVaultHolderType) String() string {
    return []string{"ACCOUNT", "CLIENT", "GROUP"}[i]
}
func ParseVaultVaultHolderType(v string) (any, error) {
    result := ACCOUNT_VAULTVAULTHOLDERTYPE
    switch v {
        case "ACCOUNT":
            result = ACCOUNT_VAULTVAULTHOLDERTYPE
        case "CLIENT":
            result = CLIENT_VAULTVAULTHOLDERTYPE
        case "GROUP":
            result = GROUP_VAULTVAULTHOLDERTYPE
        default:
            return 0, errors.New("Unknown VaultVaultHolderType value: " + v)
    }
    return &result, nil
}
func SerializeVaultVaultHolderType(values []VaultVaultHolderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VaultVaultHolderType) isMultiValue() bool {
    return false
}
