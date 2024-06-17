package models
type ProvisioningADSamAccountNameScheme int

const (
    OMIT_PROVISIONINGADSAMACCOUNTNAMESCHEME ProvisioningADSamAccountNameScheme = iota
    TRUNCATE_PROVISIONINGADSAMACCOUNTNAMESCHEME
    TRANSFER_PROVISIONINGADSAMACCOUNTNAMESCHEME
    TRANSFER_TRUNCATE_PROVISIONINGADSAMACCOUNTNAMESCHEME
    USERNAME_PROVISIONINGADSAMACCOUNTNAMESCHEME
)

func (i ProvisioningADSamAccountNameScheme) String() string {
    return []string{"OMIT", "TRUNCATE", "TRANSFER", "TRANSFER_TRUNCATE", "USERNAME"}[i]
}
func ParseProvisioningADSamAccountNameScheme(v string) (any, error) {
    result := OMIT_PROVISIONINGADSAMACCOUNTNAMESCHEME
    switch v {
        case "OMIT":
            result = OMIT_PROVISIONINGADSAMACCOUNTNAMESCHEME
        case "TRUNCATE":
            result = TRUNCATE_PROVISIONINGADSAMACCOUNTNAMESCHEME
        case "TRANSFER":
            result = TRANSFER_PROVISIONINGADSAMACCOUNTNAMESCHEME
        case "TRANSFER_TRUNCATE":
            result = TRANSFER_TRUNCATE_PROVISIONINGADSAMACCOUNTNAMESCHEME
        case "USERNAME":
            result = USERNAME_PROVISIONINGADSAMACCOUNTNAMESCHEME
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProvisioningADSamAccountNameScheme(values []ProvisioningADSamAccountNameScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningADSamAccountNameScheme) isMultiValue() bool {
    return false
}
