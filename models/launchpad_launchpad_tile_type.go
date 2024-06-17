package models
type LaunchpadLaunchpadTileType int

const (
    MANUAL_LAUNCHPADLAUNCHPADTILETYPE LaunchpadLaunchpadTileType = iota
    SSO_APPLICATION_LAUNCHPADLAUNCHPADTILETYPE
    VAULT_RECORD_LAUNCHPADLAUNCHPADTILETYPE
)

func (i LaunchpadLaunchpadTileType) String() string {
    return []string{"MANUAL", "SSO_APPLICATION", "VAULT_RECORD"}[i]
}
func ParseLaunchpadLaunchpadTileType(v string) (any, error) {
    result := MANUAL_LAUNCHPADLAUNCHPADTILETYPE
    switch v {
        case "MANUAL":
            result = MANUAL_LAUNCHPADLAUNCHPADTILETYPE
        case "SSO_APPLICATION":
            result = SSO_APPLICATION_LAUNCHPADLAUNCHPADTILETYPE
        case "VAULT_RECORD":
            result = VAULT_RECORD_LAUNCHPADLAUNCHPADTILETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLaunchpadLaunchpadTileType(values []LaunchpadLaunchpadTileType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LaunchpadLaunchpadTileType) isMultiValue() bool {
    return false
}
