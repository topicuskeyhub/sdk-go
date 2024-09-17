package models
type ProvisioningGroupOnSystemProvisioning int

const (
    ALWAYS_PROVISIONED_PROVISIONINGGROUPONSYSTEMPROVISIONING ProvisioningGroupOnSystemProvisioning = iota
    PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED_PROVISIONINGGROUPONSYSTEMPROVISIONING
    NOT_PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED_PROVISIONINGGROUPONSYSTEMPROVISIONING
    PROVISIONED_BY_DEFAULT_PROVISIONINGGROUPONSYSTEMPROVISIONING
    NOT_PROVISIONED_BY_DEFAULT_PROVISIONINGGROUPONSYSTEMPROVISIONING
)

func (i ProvisioningGroupOnSystemProvisioning) String() string {
    return []string{"ALWAYS_PROVISIONED", "PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED", "NOT_PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED", "PROVISIONED_BY_DEFAULT", "NOT_PROVISIONED_BY_DEFAULT"}[i]
}
func ParseProvisioningGroupOnSystemProvisioning(v string) (any, error) {
    result := ALWAYS_PROVISIONED_PROVISIONINGGROUPONSYSTEMPROVISIONING
    switch v {
        case "ALWAYS_PROVISIONED":
            result = ALWAYS_PROVISIONED_PROVISIONINGGROUPONSYSTEMPROVISIONING
        case "PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED":
            result = PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED_PROVISIONINGGROUPONSYSTEMPROVISIONING
        case "NOT_PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED":
            result = NOT_PROVISIONED_BY_DEFAULT_CAN_ONLY_BE_ENABLED_PROVISIONINGGROUPONSYSTEMPROVISIONING
        case "PROVISIONED_BY_DEFAULT":
            result = PROVISIONED_BY_DEFAULT_PROVISIONINGGROUPONSYSTEMPROVISIONING
        case "NOT_PROVISIONED_BY_DEFAULT":
            result = NOT_PROVISIONED_BY_DEFAULT_PROVISIONINGGROUPONSYSTEMPROVISIONING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProvisioningGroupOnSystemProvisioning(values []ProvisioningGroupOnSystemProvisioning) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningGroupOnSystemProvisioning) isMultiValue() bool {
    return false
}
