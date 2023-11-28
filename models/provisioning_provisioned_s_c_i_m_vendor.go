package models
import (
    "errors"
)
// 
type ProvisioningProvisionedSCIMVendor int

const (
    DEFAULTESCAPED_PROVISIONINGPROVISIONEDSCIMVENDOR ProvisioningProvisionedSCIMVendor = iota
    AWS_PROVISIONINGPROVISIONEDSCIMVENDOR
    KEYSTONE_PROVISIONINGPROVISIONEDSCIMVENDOR
)

func (i ProvisioningProvisionedSCIMVendor) String() string {
    return []string{"DEFAULT", "AWS", "KEYSTONE"}[i]
}
func ParseProvisioningProvisionedSCIMVendor(v string) (any, error) {
    result := DEFAULTESCAPED_PROVISIONINGPROVISIONEDSCIMVENDOR
    switch v {
        case "DEFAULT":
            result = DEFAULTESCAPED_PROVISIONINGPROVISIONEDSCIMVENDOR
        case "AWS":
            result = AWS_PROVISIONINGPROVISIONEDSCIMVENDOR
        case "KEYSTONE":
            result = KEYSTONE_PROVISIONINGPROVISIONEDSCIMVENDOR
        default:
            return 0, errors.New("Unknown ProvisioningProvisionedSCIMVendor value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningProvisionedSCIMVendor(values []ProvisioningProvisionedSCIMVendor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningProvisionedSCIMVendor) isMultiValue() bool {
    return false
}
