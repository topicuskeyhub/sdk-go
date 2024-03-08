package models
import (
    "errors"
)
type ProvisioningCircuitBreakerState int

const (
    CLOSED_PROVISIONINGCIRCUITBREAKERSTATE ProvisioningCircuitBreakerState = iota
    OPEN_PROVISIONINGCIRCUITBREAKERSTATE
    HALF_OPEN_PROVISIONINGCIRCUITBREAKERSTATE
)

func (i ProvisioningCircuitBreakerState) String() string {
    return []string{"CLOSED", "OPEN", "HALF_OPEN"}[i]
}
func ParseProvisioningCircuitBreakerState(v string) (any, error) {
    result := CLOSED_PROVISIONINGCIRCUITBREAKERSTATE
    switch v {
        case "CLOSED":
            result = CLOSED_PROVISIONINGCIRCUITBREAKERSTATE
        case "OPEN":
            result = OPEN_PROVISIONINGCIRCUITBREAKERSTATE
        case "HALF_OPEN":
            result = HALF_OPEN_PROVISIONINGCIRCUITBREAKERSTATE
        default:
            return 0, errors.New("Unknown ProvisioningCircuitBreakerState value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningCircuitBreakerState(values []ProvisioningCircuitBreakerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningCircuitBreakerState) isMultiValue() bool {
    return false
}
