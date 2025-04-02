package models
type VaultMoveVaultRecordAction int

const (
    SHARE_VAULTMOVEVAULTRECORDACTION VaultMoveVaultRecordAction = iota
    COPY_VAULTMOVEVAULTRECORDACTION
    MOVE_VAULTMOVEVAULTRECORDACTION
)

func (i VaultMoveVaultRecordAction) String() string {
    return []string{"SHARE", "COPY", "MOVE"}[i]
}
func ParseVaultMoveVaultRecordAction(v string) (any, error) {
    result := SHARE_VAULTMOVEVAULTRECORDACTION
    switch v {
        case "SHARE":
            result = SHARE_VAULTMOVEVAULTRECORDACTION
        case "COPY":
            result = COPY_VAULTMOVEVAULTRECORDACTION
        case "MOVE":
            result = MOVE_VAULTMOVEVAULTRECORDACTION
        default:
            return nil, nil
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
func (i VaultMoveVaultRecordAction) isMultiValue() bool {
    return false
}
