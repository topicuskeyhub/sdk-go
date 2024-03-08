package models
import (
    "errors"
)
type ProvisioningAccountProvisioningResult int

const (
    OK_PROVISIONINGACCOUNTPROVISIONINGRESULT ProvisioningAccountProvisioningResult = iota
    CONFIGURATION_REQUIRED_PROVISIONINGACCOUNTPROVISIONINGRESULT
    ERROR_PROVISIONINGACCOUNTPROVISIONINGRESULT
)

func (i ProvisioningAccountProvisioningResult) String() string {
    return []string{"OK", "CONFIGURATION_REQUIRED", "ERROR"}[i]
}
func ParseProvisioningAccountProvisioningResult(v string) (any, error) {
    result := OK_PROVISIONINGACCOUNTPROVISIONINGRESULT
    switch v {
        case "OK":
            result = OK_PROVISIONINGACCOUNTPROVISIONINGRESULT
        case "CONFIGURATION_REQUIRED":
            result = CONFIGURATION_REQUIRED_PROVISIONINGACCOUNTPROVISIONINGRESULT
        case "ERROR":
            result = ERROR_PROVISIONINGACCOUNTPROVISIONINGRESULT
        default:
            return 0, errors.New("Unknown ProvisioningAccountProvisioningResult value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningAccountProvisioningResult(values []ProvisioningAccountProvisioningResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningAccountProvisioningResult) isMultiValue() bool {
    return false
}
