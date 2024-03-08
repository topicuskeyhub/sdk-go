package item
import (
    "errors"
)
type GetAdditionalQueryParameterType int

const (
    ACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE GetAdditionalQueryParameterType = iota
    ADMINISTEREDCLIENTS_GETADDITIONALQUERYPARAMETERTYPE
    ADMINISTEREDSYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
    ADMINS_GETADDITIONALQUERYPARAMETERTYPE
    AUDIT_GETADDITIONALQUERYPARAMETERTYPE
    AUTHORIZEDGROUPS_GETADDITIONALQUERYPARAMETERTYPE
    CLIENTPERMISSIONS_GETADDITIONALQUERYPARAMETERTYPE
    CLIENTS_GETADDITIONALQUERYPARAMETERTYPE
    CONTENTADMINISTEREDSYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
    GROUPAUDITINGINFO_GETADDITIONALQUERYPARAMETERTYPE
    GROUPINFO_GETADDITIONALQUERYPARAMETERTYPE
    HELPDESK_GETADDITIONALQUERYPARAMETERTYPE
    MARKERS_GETADDITIONALQUERYPARAMETERTYPE
    MYACCOUNT_GETADDITIONALQUERYPARAMETERTYPE
    MYDELEGATEDACCOUNT_GETADDITIONALQUERYPARAMETERTYPE
    NESTEDGROUPS_GETADDITIONALQUERYPARAMETERTYPE
    ONLYLINKEDGOS_GETADDITIONALQUERYPARAMETERTYPE
    OWNEDCLIENTS_GETADDITIONALQUERYPARAMETERTYPE
    OWNEDDIRECTORIES_GETADDITIONALQUERYPARAMETERTYPE
    OWNEDGROUPSONSYSTEM_GETADDITIONALQUERYPARAMETERTYPE
    OWNEDORGANIZATIONALUNITS_GETADDITIONALQUERYPARAMETERTYPE
    OWNEDSYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
    RECENTAUDITS_GETADDITIONALQUERYPARAMETERTYPE
    REQUESTSTATUS_GETADDITIONALQUERYPARAMETERTYPE
    SERVICEACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
    SYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
    VAULT_GETADDITIONALQUERYPARAMETERTYPE
    WEBHOOKS_GETADDITIONALQUERYPARAMETERTYPE
)

func (i GetAdditionalQueryParameterType) String() string {
    return []string{"accounts", "administeredClients", "administeredSystems", "admins", "audit", "authorizedGroups", "clientPermissions", "clients", "contentAdministeredSystems", "groupauditinginfo", "groupinfo", "helpdesk", "markers", "myaccount", "mydelegatedaccount", "nestedGroups", "onlyLinkedGOS", "ownedClients", "ownedDirectories", "ownedGroupsOnSystem", "ownedOrganizationalUnits", "ownedSystems", "recentAudits", "requeststatus", "serviceAccounts", "systems", "vault", "webhooks"}[i]
}
func ParseGetAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accounts":
            result = ACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
        case "administeredClients":
            result = ADMINISTEREDCLIENTS_GETADDITIONALQUERYPARAMETERTYPE
        case "administeredSystems":
            result = ADMINISTEREDSYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
        case "admins":
            result = ADMINS_GETADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_GETADDITIONALQUERYPARAMETERTYPE
        case "authorizedGroups":
            result = AUTHORIZEDGROUPS_GETADDITIONALQUERYPARAMETERTYPE
        case "clientPermissions":
            result = CLIENTPERMISSIONS_GETADDITIONALQUERYPARAMETERTYPE
        case "clients":
            result = CLIENTS_GETADDITIONALQUERYPARAMETERTYPE
        case "contentAdministeredSystems":
            result = CONTENTADMINISTEREDSYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
        case "groupauditinginfo":
            result = GROUPAUDITINGINFO_GETADDITIONALQUERYPARAMETERTYPE
        case "groupinfo":
            result = GROUPINFO_GETADDITIONALQUERYPARAMETERTYPE
        case "helpdesk":
            result = HELPDESK_GETADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_GETADDITIONALQUERYPARAMETERTYPE
        case "myaccount":
            result = MYACCOUNT_GETADDITIONALQUERYPARAMETERTYPE
        case "mydelegatedaccount":
            result = MYDELEGATEDACCOUNT_GETADDITIONALQUERYPARAMETERTYPE
        case "nestedGroups":
            result = NESTEDGROUPS_GETADDITIONALQUERYPARAMETERTYPE
        case "onlyLinkedGOS":
            result = ONLYLINKEDGOS_GETADDITIONALQUERYPARAMETERTYPE
        case "ownedClients":
            result = OWNEDCLIENTS_GETADDITIONALQUERYPARAMETERTYPE
        case "ownedDirectories":
            result = OWNEDDIRECTORIES_GETADDITIONALQUERYPARAMETERTYPE
        case "ownedGroupsOnSystem":
            result = OWNEDGROUPSONSYSTEM_GETADDITIONALQUERYPARAMETERTYPE
        case "ownedOrganizationalUnits":
            result = OWNEDORGANIZATIONALUNITS_GETADDITIONALQUERYPARAMETERTYPE
        case "ownedSystems":
            result = OWNEDSYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
        case "recentAudits":
            result = RECENTAUDITS_GETADDITIONALQUERYPARAMETERTYPE
        case "requeststatus":
            result = REQUESTSTATUS_GETADDITIONALQUERYPARAMETERTYPE
        case "serviceAccounts":
            result = SERVICEACCOUNTS_GETADDITIONALQUERYPARAMETERTYPE
        case "systems":
            result = SYSTEMS_GETADDITIONALQUERYPARAMETERTYPE
        case "vault":
            result = VAULT_GETADDITIONALQUERYPARAMETERTYPE
        case "webhooks":
            result = WEBHOOKS_GETADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetAdditionalQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetAdditionalQueryParameterType(values []GetAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
