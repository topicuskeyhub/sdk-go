package item
import (
    "errors"
)
type PutAdditionalQueryParameterType int

const (
    ACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    ADMINISTEREDCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
    ADMINISTEREDSYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
    ADMINS_PUTADDITIONALQUERYPARAMETERTYPE
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    AUTHORIZEDGROUPS_PUTADDITIONALQUERYPARAMETERTYPE
    CLIENTPERMISSIONS_PUTADDITIONALQUERYPARAMETERTYPE
    CLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
    CONTENTADMINISTEREDSYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
    GROUPAUDITINGINFO_PUTADDITIONALQUERYPARAMETERTYPE
    GROUPINFO_PUTADDITIONALQUERYPARAMETERTYPE
    HELPDESK_PUTADDITIONALQUERYPARAMETERTYPE
    MARKERS_PUTADDITIONALQUERYPARAMETERTYPE
    MYACCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
    MYDELEGATEDACCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
    NESTEDGROUPS_PUTADDITIONALQUERYPARAMETERTYPE
    ONLYLINKEDGOS_PUTADDITIONALQUERYPARAMETERTYPE
    OWNEDCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
    OWNEDDIRECTORIES_PUTADDITIONALQUERYPARAMETERTYPE
    OWNEDGROUPSONSYSTEM_PUTADDITIONALQUERYPARAMETERTYPE
    OWNEDORGANIZATIONALUNITS_PUTADDITIONALQUERYPARAMETERTYPE
    OWNEDSYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
    RECENTAUDITS_PUTADDITIONALQUERYPARAMETERTYPE
    REQUESTSTATUS_PUTADDITIONALQUERYPARAMETERTYPE
    SERVICEACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
    SYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
    VAULT_PUTADDITIONALQUERYPARAMETERTYPE
    WEBHOOKS_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"accounts", "administeredClients", "administeredSystems", "admins", "audit", "authorizedGroups", "clientPermissions", "clients", "contentAdministeredSystems", "groupauditinginfo", "groupinfo", "helpdesk", "markers", "myaccount", "mydelegatedaccount", "nestedGroups", "onlyLinkedGOS", "ownedClients", "ownedDirectories", "ownedGroupsOnSystem", "ownedOrganizationalUnits", "ownedSystems", "recentAudits", "requeststatus", "serviceAccounts", "systems", "vault", "webhooks"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accounts":
            result = ACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "administeredClients":
            result = ADMINISTEREDCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "administeredSystems":
            result = ADMINISTEREDSYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
        case "admins":
            result = ADMINS_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "authorizedGroups":
            result = AUTHORIZEDGROUPS_PUTADDITIONALQUERYPARAMETERTYPE
        case "clientPermissions":
            result = CLIENTPERMISSIONS_PUTADDITIONALQUERYPARAMETERTYPE
        case "clients":
            result = CLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "contentAdministeredSystems":
            result = CONTENTADMINISTEREDSYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
        case "groupauditinginfo":
            result = GROUPAUDITINGINFO_PUTADDITIONALQUERYPARAMETERTYPE
        case "groupinfo":
            result = GROUPINFO_PUTADDITIONALQUERYPARAMETERTYPE
        case "helpdesk":
            result = HELPDESK_PUTADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_PUTADDITIONALQUERYPARAMETERTYPE
        case "myaccount":
            result = MYACCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
        case "mydelegatedaccount":
            result = MYDELEGATEDACCOUNT_PUTADDITIONALQUERYPARAMETERTYPE
        case "nestedGroups":
            result = NESTEDGROUPS_PUTADDITIONALQUERYPARAMETERTYPE
        case "onlyLinkedGOS":
            result = ONLYLINKEDGOS_PUTADDITIONALQUERYPARAMETERTYPE
        case "ownedClients":
            result = OWNEDCLIENTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "ownedDirectories":
            result = OWNEDDIRECTORIES_PUTADDITIONALQUERYPARAMETERTYPE
        case "ownedGroupsOnSystem":
            result = OWNEDGROUPSONSYSTEM_PUTADDITIONALQUERYPARAMETERTYPE
        case "ownedOrganizationalUnits":
            result = OWNEDORGANIZATIONALUNITS_PUTADDITIONALQUERYPARAMETERTYPE
        case "ownedSystems":
            result = OWNEDSYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
        case "recentAudits":
            result = RECENTAUDITS_PUTADDITIONALQUERYPARAMETERTYPE
        case "requeststatus":
            result = REQUESTSTATUS_PUTADDITIONALQUERYPARAMETERTYPE
        case "serviceAccounts":
            result = SERVICEACCOUNTS_PUTADDITIONALQUERYPARAMETERTYPE
        case "systems":
            result = SYSTEMS_PUTADDITIONALQUERYPARAMETERTYPE
        case "vault":
            result = VAULT_PUTADDITIONALQUERYPARAMETERTYPE
        case "webhooks":
            result = WEBHOOKS_PUTADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown PutAdditionalQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializePutAdditionalQueryParameterType(values []PutAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PutAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
