// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type ClientOAuth2ClientPermissionType int

const (
    ACCOUNTS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE ClientOAuth2ClientPermissionType = iota
    ACCOUNTS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    ACCOUNTS_REMOVE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    CLIENTS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    CLIENTS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUPONSYSTEM_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUPS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUPS_GRANT_PERMISSIONS_AFTER_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUPS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUPS_SET_CLASSIFICATION_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUP_FULL_VAULT_ACCESS_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUP_LAUNCHPADTILES_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUP_READ_CONTENTS_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    GROUP_SET_AUTHORIZATION_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    PROVISIONEDSYSTEMS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    SERVICE_ACCOUNTS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    SERVICE_ACCOUNTS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    SERVICE_ACCOUNTS_UPDATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
)

func (i ClientOAuth2ClientPermissionType) String() string {
    return []string{"ACCOUNTS_CREATE", "ACCOUNTS_QUERY", "ACCOUNTS_REMOVE", "CLIENTS_CREATE", "CLIENTS_QUERY", "GROUPONSYSTEM_CREATE", "GROUPS_CREATE", "GROUPS_GRANT_PERMISSIONS_AFTER_CREATE", "GROUPS_QUERY", "GROUPS_SET_CLASSIFICATION", "GROUP_FULL_VAULT_ACCESS", "GROUP_LAUNCHPADTILES", "GROUP_READ_CONTENTS", "GROUP_SET_AUTHORIZATION", "PROVISIONEDSYSTEMS_QUERY", "SERVICE_ACCOUNTS_CREATE", "SERVICE_ACCOUNTS_QUERY", "SERVICE_ACCOUNTS_UPDATE"}[i]
}
func ParseClientOAuth2ClientPermissionType(v string) (any, error) {
    result := ACCOUNTS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
    switch v {
        case "ACCOUNTS_CREATE":
            result = ACCOUNTS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "ACCOUNTS_QUERY":
            result = ACCOUNTS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "ACCOUNTS_REMOVE":
            result = ACCOUNTS_REMOVE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "CLIENTS_CREATE":
            result = CLIENTS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "CLIENTS_QUERY":
            result = CLIENTS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUPONSYSTEM_CREATE":
            result = GROUPONSYSTEM_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUPS_CREATE":
            result = GROUPS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUPS_GRANT_PERMISSIONS_AFTER_CREATE":
            result = GROUPS_GRANT_PERMISSIONS_AFTER_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUPS_QUERY":
            result = GROUPS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUPS_SET_CLASSIFICATION":
            result = GROUPS_SET_CLASSIFICATION_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUP_FULL_VAULT_ACCESS":
            result = GROUP_FULL_VAULT_ACCESS_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUP_LAUNCHPADTILES":
            result = GROUP_LAUNCHPADTILES_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUP_READ_CONTENTS":
            result = GROUP_READ_CONTENTS_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "GROUP_SET_AUTHORIZATION":
            result = GROUP_SET_AUTHORIZATION_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "PROVISIONEDSYSTEMS_QUERY":
            result = PROVISIONEDSYSTEMS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "SERVICE_ACCOUNTS_CREATE":
            result = SERVICE_ACCOUNTS_CREATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "SERVICE_ACCOUNTS_QUERY":
            result = SERVICE_ACCOUNTS_QUERY_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        case "SERVICE_ACCOUNTS_UPDATE":
            result = SERVICE_ACCOUNTS_UPDATE_CLIENTOAUTH2CLIENTPERMISSIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeClientOAuth2ClientPermissionType(values []ClientOAuth2ClientPermissionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClientOAuth2ClientPermissionType) isMultiValue() bool {
    return false
}
