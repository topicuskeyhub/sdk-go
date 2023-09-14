package models
import (
    "errors"
)
// 
type VaultVaultSecretType int

const (
    PASSWORD_VAULTVAULTSECRETTYPE VaultVaultSecretType = iota
    FILE_VAULTVAULTSECRETTYPE
    TOTP_VAULTVAULTSECRETTYPE
    COMMENT_VAULTVAULTSECRETTYPE
)

func (i VaultVaultSecretType) String() string {
    return []string{"PASSWORD", "FILE", "TOTP", "COMMENT"}[i]
}
func ParseVaultVaultSecretType(v string) (any, error) {
    result := PASSWORD_VAULTVAULTSECRETTYPE
    switch v {
        case "PASSWORD":
            result = PASSWORD_VAULTVAULTSECRETTYPE
        case "FILE":
            result = FILE_VAULTVAULTSECRETTYPE
        case "TOTP":
            result = TOTP_VAULTVAULTSECRETTYPE
        case "COMMENT":
            result = COMMENT_VAULTVAULTSECRETTYPE
        default:
            return 0, errors.New("Unknown VaultVaultSecretType value: " + v)
    }
    return &result, nil
}
func SerializeVaultVaultSecretType(values []VaultVaultSecretType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VaultVaultSecretType) isMultiValue() bool {
    return false
}
