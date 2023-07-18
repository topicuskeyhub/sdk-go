package models
import (
    "errors"
)
// 
type VaultAccountVaultStatus int

const (
    NOT_CREATED_VAULTACCOUNTVAULTSTATUS VaultAccountVaultStatus = iota
    INITIALIZING_VAULTACCOUNTVAULTSTATUS
    LOCKED_VAULTACCOUNTVAULTSTATUS
    UNLOCKED_VAULTACCOUNTVAULTSTATUS
)

func (i VaultAccountVaultStatus) String() string {
    return []string{"NOT_CREATED", "INITIALIZING", "LOCKED", "UNLOCKED"}[i]
}
func ParseVaultAccountVaultStatus(v string) (any, error) {
    result := NOT_CREATED_VAULTACCOUNTVAULTSTATUS
    switch v {
        case "NOT_CREATED":
            result = NOT_CREATED_VAULTACCOUNTVAULTSTATUS
        case "INITIALIZING":
            result = INITIALIZING_VAULTACCOUNTVAULTSTATUS
        case "LOCKED":
            result = LOCKED_VAULTACCOUNTVAULTSTATUS
        case "UNLOCKED":
            result = UNLOCKED_VAULTACCOUNTVAULTSTATUS
        default:
            return 0, errors.New("Unknown VaultAccountVaultStatus value: " + v)
    }
    return &result, nil
}
func SerializeVaultAccountVaultStatus(values []VaultAccountVaultStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
