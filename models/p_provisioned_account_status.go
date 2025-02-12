package models
type PProvisionedAccountStatus int

const (
    NON_EXISTENT_PPROVISIONEDACCOUNTSTATUS PProvisionedAccountStatus = iota
    EXISTS_PASSWORD_SET_PPROVISIONEDACCOUNTSTATUS
    EXISTS_PASSWORD_UNSET_PPROVISIONEDACCOUNTSTATUS
)

func (i PProvisionedAccountStatus) String() string {
    return []string{"NON_EXISTENT", "EXISTS_PASSWORD_SET", "EXISTS_PASSWORD_UNSET"}[i]
}
func ParsePProvisionedAccountStatus(v string) (any, error) {
    result := NON_EXISTENT_PPROVISIONEDACCOUNTSTATUS
    switch v {
        case "NON_EXISTENT":
            result = NON_EXISTENT_PPROVISIONEDACCOUNTSTATUS
        case "EXISTS_PASSWORD_SET":
            result = EXISTS_PASSWORD_SET_PPROVISIONEDACCOUNTSTATUS
        case "EXISTS_PASSWORD_UNSET":
            result = EXISTS_PASSWORD_UNSET_PPROVISIONEDACCOUNTSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePProvisionedAccountStatus(values []PProvisionedAccountStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PProvisionedAccountStatus) isMultiValue() bool {
    return false
}
