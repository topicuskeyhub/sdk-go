package group
type PostAdditionalQueryParameterType int

const (
    ACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE PostAdditionalQueryParameterType = iota
    ADMINISTEREDCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
    ADMINISTEREDSYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
    ADMINS_POSTADDITIONALQUERYPARAMETERTYPE
    AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
    AUTHORIZEDGROUPS_POSTADDITIONALQUERYPARAMETERTYPE
    CLIENTPERMISSIONS_POSTADDITIONALQUERYPARAMETERTYPE
    CLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
    CONTENTADMINISTEREDSYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
    GLOBALROLES_POSTADDITIONALQUERYPARAMETERTYPE
    GROUPACCESSINFO_POSTADDITIONALQUERYPARAMETERTYPE
    GROUPAUDITINGINFO_POSTADDITIONALQUERYPARAMETERTYPE
    GROUPINFO_POSTADDITIONALQUERYPARAMETERTYPE
    HELPDESK_POSTADDITIONALQUERYPARAMETERTYPE
    MARKERS_POSTADDITIONALQUERYPARAMETERTYPE
    MYACCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
    MYDELEGATEDACCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
    NESTEDGROUPS_POSTADDITIONALQUERYPARAMETERTYPE
    ONLYLINKEDGOS_POSTADDITIONALQUERYPARAMETERTYPE
    OWNEDCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
    OWNEDDIRECTORIES_POSTADDITIONALQUERYPARAMETERTYPE
    OWNEDGROUPSONSYSTEM_POSTADDITIONALQUERYPARAMETERTYPE
    OWNEDORGANIZATIONALUNITS_POSTADDITIONALQUERYPARAMETERTYPE
    OWNEDSYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
    RECENTAUDITS_POSTADDITIONALQUERYPARAMETERTYPE
    REQUESTSTATUS_POSTADDITIONALQUERYPARAMETERTYPE
    SERVICEACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
    SYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
    VAULT_POSTADDITIONALQUERYPARAMETERTYPE
    WEBHOOKS_POSTADDITIONALQUERYPARAMETERTYPE
)

func (i PostAdditionalQueryParameterType) String() string {
    return []string{"accounts", "administeredClients", "administeredSystems", "admins", "audit", "authorizedGroups", "clientPermissions", "clients", "contentAdministeredSystems", "globalRoles", "groupAccessInfo", "groupauditinginfo", "groupinfo", "helpdesk", "markers", "myaccount", "mydelegatedaccount", "nestedGroups", "onlyLinkedGOS", "ownedClients", "ownedDirectories", "ownedGroupsOnSystem", "ownedOrganizationalUnits", "ownedSystems", "recentAudits", "requeststatus", "serviceAccounts", "systems", "vault", "webhooks"}[i]
}
func ParsePostAdditionalQueryParameterType(v string) (any, error) {
    result := ACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "accounts":
            result = ACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "administeredClients":
            result = ADMINISTEREDCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "administeredSystems":
            result = ADMINISTEREDSYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
        case "admins":
            result = ADMINS_POSTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_POSTADDITIONALQUERYPARAMETERTYPE
        case "authorizedGroups":
            result = AUTHORIZEDGROUPS_POSTADDITIONALQUERYPARAMETERTYPE
        case "clientPermissions":
            result = CLIENTPERMISSIONS_POSTADDITIONALQUERYPARAMETERTYPE
        case "clients":
            result = CLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "contentAdministeredSystems":
            result = CONTENTADMINISTEREDSYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
        case "globalRoles":
            result = GLOBALROLES_POSTADDITIONALQUERYPARAMETERTYPE
        case "groupAccessInfo":
            result = GROUPACCESSINFO_POSTADDITIONALQUERYPARAMETERTYPE
        case "groupauditinginfo":
            result = GROUPAUDITINGINFO_POSTADDITIONALQUERYPARAMETERTYPE
        case "groupinfo":
            result = GROUPINFO_POSTADDITIONALQUERYPARAMETERTYPE
        case "helpdesk":
            result = HELPDESK_POSTADDITIONALQUERYPARAMETERTYPE
        case "markers":
            result = MARKERS_POSTADDITIONALQUERYPARAMETERTYPE
        case "myaccount":
            result = MYACCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
        case "mydelegatedaccount":
            result = MYDELEGATEDACCOUNT_POSTADDITIONALQUERYPARAMETERTYPE
        case "nestedGroups":
            result = NESTEDGROUPS_POSTADDITIONALQUERYPARAMETERTYPE
        case "onlyLinkedGOS":
            result = ONLYLINKEDGOS_POSTADDITIONALQUERYPARAMETERTYPE
        case "ownedClients":
            result = OWNEDCLIENTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "ownedDirectories":
            result = OWNEDDIRECTORIES_POSTADDITIONALQUERYPARAMETERTYPE
        case "ownedGroupsOnSystem":
            result = OWNEDGROUPSONSYSTEM_POSTADDITIONALQUERYPARAMETERTYPE
        case "ownedOrganizationalUnits":
            result = OWNEDORGANIZATIONALUNITS_POSTADDITIONALQUERYPARAMETERTYPE
        case "ownedSystems":
            result = OWNEDSYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
        case "recentAudits":
            result = RECENTAUDITS_POSTADDITIONALQUERYPARAMETERTYPE
        case "requeststatus":
            result = REQUESTSTATUS_POSTADDITIONALQUERYPARAMETERTYPE
        case "serviceAccounts":
            result = SERVICEACCOUNTS_POSTADDITIONALQUERYPARAMETERTYPE
        case "systems":
            result = SYSTEMS_POSTADDITIONALQUERYPARAMETERTYPE
        case "vault":
            result = VAULT_POSTADDITIONALQUERYPARAMETERTYPE
        case "webhooks":
            result = WEBHOOKS_POSTADDITIONALQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePostAdditionalQueryParameterType(values []PostAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PostAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
