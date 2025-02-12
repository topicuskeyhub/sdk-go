package models
type ProvisioningDeprovisionAction int

const (
    NONE_PROVISIONINGDEPROVISIONACTION ProvisioningDeprovisionAction = iota
    DEPROVISION_ACCOUNTS_PROVISIONINGDEPROVISIONACTION
    DESTROY_ACCOUNTS_PROVISIONINGDEPROVISIONACTION
    DESTROY_ALL_PROVISIONINGDEPROVISIONACTION
)

func (i ProvisioningDeprovisionAction) String() string {
    return []string{"NONE", "DEPROVISION_ACCOUNTS", "DESTROY_ACCOUNTS", "DESTROY_ALL"}[i]
}
func ParseProvisioningDeprovisionAction(v string) (any, error) {
    result := NONE_PROVISIONINGDEPROVISIONACTION
    switch v {
        case "NONE":
            result = NONE_PROVISIONINGDEPROVISIONACTION
        case "DEPROVISION_ACCOUNTS":
            result = DEPROVISION_ACCOUNTS_PROVISIONINGDEPROVISIONACTION
        case "DESTROY_ACCOUNTS":
            result = DESTROY_ACCOUNTS_PROVISIONINGDEPROVISIONACTION
        case "DESTROY_ALL":
            result = DESTROY_ALL_PROVISIONINGDEPROVISIONACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProvisioningDeprovisionAction(values []ProvisioningDeprovisionAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningDeprovisionAction) isMultiValue() bool {
    return false
}
