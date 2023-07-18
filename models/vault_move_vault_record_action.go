package models
import (
    "errors"
)
// 
type VaultMoveVaultRecordAction int

const (
    MOVE_VAULTMOVEVAULTRECORDACTION VaultMoveVaultRecordAction = iota
    COPY_VAULTMOVEVAULTRECORDACTION
    SHARE_VAULTMOVEVAULTRECORDACTION
)

func (i VaultMoveVaultRecordAction) String() string {
    return []string{"MOVE", "COPY", "SHARE"}[i]
}
func ParseVaultMoveVaultRecordAction(v string) (any, error) {
    result := MOVE_VAULTMOVEVAULTRECORDACTION
    switch v {
        case "MOVE":
            result = MOVE_VAULTMOVEVAULTRECORDACTION
        case "COPY":
            result = COPY_VAULTMOVEVAULTRECORDACTION
        case "SHARE":
            result = SHARE_VAULTMOVEVAULTRECORDACTION
        default:
            return 0, errors.New("Unknown VaultMoveVaultRecordAction value: " + v)
    }
    return &result, nil
}
func SerializeVaultMoveVaultRecordAction(values []VaultMoveVaultRecordAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
