package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroup 
type GroupGroup struct {
    GroupGroupPrimer
    // The additionalObjects property
    additionalObjects GroupGroup_additionalObjectsable
    // The applicationAdministration property
    applicationAdministration *bool
    // The auditConfig property
    auditConfig GroupGroupAuditConfigable
    // The auditor property
    auditor *bool
    // The auditRequested property
    auditRequested *bool
    // The authorizingGroupAuditing property
    authorizingGroupAuditing GroupGroupPrimerable
    // The authorizingGroupDelegation property
    authorizingGroupDelegation GroupGroupPrimerable
    // The authorizingGroupMembership property
    authorizingGroupMembership GroupGroupPrimerable
    // The authorizingGroupProvisioning property
    authorizingGroupProvisioning GroupGroupPrimerable
    // The authorizingGroupTypes property
    authorizingGroupTypes []RequestAuthorizingGroupType
    // The classification property
    classification GroupGroupClassificationPrimerable
    // The description property
    description *string
    // The extendedAccess property
    extendedAccess *GroupGroupExtendedAccess
    // The hideAuditTrail property
    hideAuditTrail *bool
    // The nestedUnder property
    nestedUnder GroupGroupPrimerable
    // The organizationalUnit property
    organizationalUnit OrganizationOrganizationalUnitPrimerable
    // The privateGroup property
    privateGroup *bool
    // The recordTrail property
    recordTrail *bool
    // The rotatingPasswordRequired property
    rotatingPasswordRequired *bool
    // The singleManaged property
    singleManaged *bool
    // The vaultRecovery property
    vaultRecovery *GroupVaultRecoveryAvailability
    // The vaultRequiresActivation property
    vaultRequiresActivation *bool
}
// NewGroupGroup instantiates a new groupGroup and sets the default values.
func NewGroupGroup()(*GroupGroup) {
    m := &GroupGroup{
        GroupGroupPrimer: *NewGroupGroupPrimer(),
    }
    typeEscapedValue := "group.Group"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroup(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupGroup) GetAdditionalObjects()(GroupGroup_additionalObjectsable) {
    return m.additionalObjects
}
// GetApplicationAdministration gets the applicationAdministration property value. The applicationAdministration property
func (m *GroupGroup) GetApplicationAdministration()(*bool) {
    return m.applicationAdministration
}
// GetAuditConfig gets the auditConfig property value. The auditConfig property
func (m *GroupGroup) GetAuditConfig()(GroupGroupAuditConfigable) {
    return m.auditConfig
}
// GetAuditor gets the auditor property value. The auditor property
func (m *GroupGroup) GetAuditor()(*bool) {
    return m.auditor
}
// GetAuditRequested gets the auditRequested property value. The auditRequested property
func (m *GroupGroup) GetAuditRequested()(*bool) {
    return m.auditRequested
}
// GetAuthorizingGroupAuditing gets the authorizingGroupAuditing property value. The authorizingGroupAuditing property
func (m *GroupGroup) GetAuthorizingGroupAuditing()(GroupGroupPrimerable) {
    return m.authorizingGroupAuditing
}
// GetAuthorizingGroupDelegation gets the authorizingGroupDelegation property value. The authorizingGroupDelegation property
func (m *GroupGroup) GetAuthorizingGroupDelegation()(GroupGroupPrimerable) {
    return m.authorizingGroupDelegation
}
// GetAuthorizingGroupMembership gets the authorizingGroupMembership property value. The authorizingGroupMembership property
func (m *GroupGroup) GetAuthorizingGroupMembership()(GroupGroupPrimerable) {
    return m.authorizingGroupMembership
}
// GetAuthorizingGroupProvisioning gets the authorizingGroupProvisioning property value. The authorizingGroupProvisioning property
func (m *GroupGroup) GetAuthorizingGroupProvisioning()(GroupGroupPrimerable) {
    return m.authorizingGroupProvisioning
}
// GetAuthorizingGroupTypes gets the authorizingGroupTypes property value. The authorizingGroupTypes property
func (m *GroupGroup) GetAuthorizingGroupTypes()([]RequestAuthorizingGroupType) {
    return m.authorizingGroupTypes
}
// GetClassification gets the classification property value. The classification property
func (m *GroupGroup) GetClassification()(GroupGroupClassificationPrimerable) {
    return m.classification
}
// GetDescription gets the description property value. The description property
func (m *GroupGroup) GetDescription()(*string) {
    return m.description
}
// GetExtendedAccess gets the extendedAccess property value. The extendedAccess property
func (m *GroupGroup) GetExtendedAccess()(*GroupGroupExtendedAccess) {
    return m.extendedAccess
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupGroupPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroup_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupGroup_additionalObjectsable))
        }
        return nil
    }
    res["applicationAdministration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationAdministration(val)
        }
        return nil
    }
    res["auditConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupAuditConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditConfig(val.(GroupGroupAuditConfigable))
        }
        return nil
    }
    res["auditor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditor(val)
        }
        return nil
    }
    res["auditRequested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditRequested(val)
        }
        return nil
    }
    res["authorizingGroupAuditing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupAuditing(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["authorizingGroupDelegation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupDelegation(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["authorizingGroupMembership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupMembership(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["authorizingGroupProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupProvisioning(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["authorizingGroupTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRequestAuthorizingGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RequestAuthorizingGroupType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*RequestAuthorizingGroupType))
                }
            }
            m.SetAuthorizingGroupTypes(res)
        }
        return nil
    }
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupClassificationPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(GroupGroupClassificationPrimerable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["extendedAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupExtendedAccess)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtendedAccess(val.(*GroupGroupExtendedAccess))
        }
        return nil
    }
    res["hideAuditTrail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideAuditTrail(val)
        }
        return nil
    }
    res["nestedUnder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNestedUnder(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["organizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnit(val.(OrganizationOrganizationalUnitPrimerable))
        }
        return nil
    }
    res["privateGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateGroup(val)
        }
        return nil
    }
    res["recordTrail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordTrail(val)
        }
        return nil
    }
    res["rotatingPasswordRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRotatingPasswordRequired(val)
        }
        return nil
    }
    res["singleManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleManaged(val)
        }
        return nil
    }
    res["vaultRecovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupVaultRecoveryAvailability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultRecovery(val.(*GroupVaultRecoveryAvailability))
        }
        return nil
    }
    res["vaultRequiresActivation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultRequiresActivation(val)
        }
        return nil
    }
    return res
}
// GetHideAuditTrail gets the hideAuditTrail property value. The hideAuditTrail property
func (m *GroupGroup) GetHideAuditTrail()(*bool) {
    return m.hideAuditTrail
}
// GetNestedUnder gets the nestedUnder property value. The nestedUnder property
func (m *GroupGroup) GetNestedUnder()(GroupGroupPrimerable) {
    return m.nestedUnder
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizationalUnit property
func (m *GroupGroup) GetOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable) {
    return m.organizationalUnit
}
// GetPrivateGroup gets the privateGroup property value. The privateGroup property
func (m *GroupGroup) GetPrivateGroup()(*bool) {
    return m.privateGroup
}
// GetRecordTrail gets the recordTrail property value. The recordTrail property
func (m *GroupGroup) GetRecordTrail()(*bool) {
    return m.recordTrail
}
// GetRotatingPasswordRequired gets the rotatingPasswordRequired property value. The rotatingPasswordRequired property
func (m *GroupGroup) GetRotatingPasswordRequired()(*bool) {
    return m.rotatingPasswordRequired
}
// GetSingleManaged gets the singleManaged property value. The singleManaged property
func (m *GroupGroup) GetSingleManaged()(*bool) {
    return m.singleManaged
}
// GetVaultRecovery gets the vaultRecovery property value. The vaultRecovery property
func (m *GroupGroup) GetVaultRecovery()(*GroupVaultRecoveryAvailability) {
    return m.vaultRecovery
}
// GetVaultRequiresActivation gets the vaultRequiresActivation property value. The vaultRequiresActivation property
func (m *GroupGroup) GetVaultRequiresActivation()(*bool) {
    return m.vaultRequiresActivation
}
// Serialize serializes information the current object
func (m *GroupGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupGroupPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationAdministration", m.GetApplicationAdministration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("auditConfig", m.GetAuditConfig())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizingGroupAuditing", m.GetAuthorizingGroupAuditing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizingGroupDelegation", m.GetAuthorizingGroupDelegation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizingGroupMembership", m.GetAuthorizingGroupMembership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authorizingGroupProvisioning", m.GetAuthorizingGroupProvisioning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("classification", m.GetClassification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetExtendedAccess() != nil {
        cast := (*m.GetExtendedAccess()).String()
        err = writer.WriteStringValue("extendedAccess", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideAuditTrail", m.GetHideAuditTrail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("nestedUnder", m.GetNestedUnder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizationalUnit", m.GetOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("privateGroup", m.GetPrivateGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("recordTrail", m.GetRecordTrail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("rotatingPasswordRequired", m.GetRotatingPasswordRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("singleManaged", m.GetSingleManaged())
        if err != nil {
            return err
        }
    }
    if m.GetVaultRecovery() != nil {
        cast := (*m.GetVaultRecovery()).String()
        err = writer.WriteStringValue("vaultRecovery", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("vaultRequiresActivation", m.GetVaultRequiresActivation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupGroup) SetAdditionalObjects(value GroupGroup_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetApplicationAdministration sets the applicationAdministration property value. The applicationAdministration property
func (m *GroupGroup) SetApplicationAdministration(value *bool)() {
    m.applicationAdministration = value
}
// SetAuditConfig sets the auditConfig property value. The auditConfig property
func (m *GroupGroup) SetAuditConfig(value GroupGroupAuditConfigable)() {
    m.auditConfig = value
}
// SetAuditor sets the auditor property value. The auditor property
func (m *GroupGroup) SetAuditor(value *bool)() {
    m.auditor = value
}
// SetAuditRequested sets the auditRequested property value. The auditRequested property
func (m *GroupGroup) SetAuditRequested(value *bool)() {
    m.auditRequested = value
}
// SetAuthorizingGroupAuditing sets the authorizingGroupAuditing property value. The authorizingGroupAuditing property
func (m *GroupGroup) SetAuthorizingGroupAuditing(value GroupGroupPrimerable)() {
    m.authorizingGroupAuditing = value
}
// SetAuthorizingGroupDelegation sets the authorizingGroupDelegation property value. The authorizingGroupDelegation property
func (m *GroupGroup) SetAuthorizingGroupDelegation(value GroupGroupPrimerable)() {
    m.authorizingGroupDelegation = value
}
// SetAuthorizingGroupMembership sets the authorizingGroupMembership property value. The authorizingGroupMembership property
func (m *GroupGroup) SetAuthorizingGroupMembership(value GroupGroupPrimerable)() {
    m.authorizingGroupMembership = value
}
// SetAuthorizingGroupProvisioning sets the authorizingGroupProvisioning property value. The authorizingGroupProvisioning property
func (m *GroupGroup) SetAuthorizingGroupProvisioning(value GroupGroupPrimerable)() {
    m.authorizingGroupProvisioning = value
}
// SetAuthorizingGroupTypes sets the authorizingGroupTypes property value. The authorizingGroupTypes property
func (m *GroupGroup) SetAuthorizingGroupTypes(value []RequestAuthorizingGroupType)() {
    m.authorizingGroupTypes = value
}
// SetClassification sets the classification property value. The classification property
func (m *GroupGroup) SetClassification(value GroupGroupClassificationPrimerable)() {
    m.classification = value
}
// SetDescription sets the description property value. The description property
func (m *GroupGroup) SetDescription(value *string)() {
    m.description = value
}
// SetExtendedAccess sets the extendedAccess property value. The extendedAccess property
func (m *GroupGroup) SetExtendedAccess(value *GroupGroupExtendedAccess)() {
    m.extendedAccess = value
}
// SetHideAuditTrail sets the hideAuditTrail property value. The hideAuditTrail property
func (m *GroupGroup) SetHideAuditTrail(value *bool)() {
    m.hideAuditTrail = value
}
// SetNestedUnder sets the nestedUnder property value. The nestedUnder property
func (m *GroupGroup) SetNestedUnder(value GroupGroupPrimerable)() {
    m.nestedUnder = value
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizationalUnit property
func (m *GroupGroup) SetOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)() {
    m.organizationalUnit = value
}
// SetPrivateGroup sets the privateGroup property value. The privateGroup property
func (m *GroupGroup) SetPrivateGroup(value *bool)() {
    m.privateGroup = value
}
// SetRecordTrail sets the recordTrail property value. The recordTrail property
func (m *GroupGroup) SetRecordTrail(value *bool)() {
    m.recordTrail = value
}
// SetRotatingPasswordRequired sets the rotatingPasswordRequired property value. The rotatingPasswordRequired property
func (m *GroupGroup) SetRotatingPasswordRequired(value *bool)() {
    m.rotatingPasswordRequired = value
}
// SetSingleManaged sets the singleManaged property value. The singleManaged property
func (m *GroupGroup) SetSingleManaged(value *bool)() {
    m.singleManaged = value
}
// SetVaultRecovery sets the vaultRecovery property value. The vaultRecovery property
func (m *GroupGroup) SetVaultRecovery(value *GroupVaultRecoveryAvailability)() {
    m.vaultRecovery = value
}
// SetVaultRequiresActivation sets the vaultRequiresActivation property value. The vaultRequiresActivation property
func (m *GroupGroup) SetVaultRequiresActivation(value *bool)() {
    m.vaultRequiresActivation = value
}
// GroupGroupable 
type GroupGroupable interface {
    GroupGroupPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(GroupGroup_additionalObjectsable)
    GetApplicationAdministration()(*bool)
    GetAuditConfig()(GroupGroupAuditConfigable)
    GetAuditor()(*bool)
    GetAuditRequested()(*bool)
    GetAuthorizingGroupAuditing()(GroupGroupPrimerable)
    GetAuthorizingGroupDelegation()(GroupGroupPrimerable)
    GetAuthorizingGroupMembership()(GroupGroupPrimerable)
    GetAuthorizingGroupProvisioning()(GroupGroupPrimerable)
    GetAuthorizingGroupTypes()([]RequestAuthorizingGroupType)
    GetClassification()(GroupGroupClassificationPrimerable)
    GetDescription()(*string)
    GetExtendedAccess()(*GroupGroupExtendedAccess)
    GetHideAuditTrail()(*bool)
    GetNestedUnder()(GroupGroupPrimerable)
    GetOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable)
    GetPrivateGroup()(*bool)
    GetRecordTrail()(*bool)
    GetRotatingPasswordRequired()(*bool)
    GetSingleManaged()(*bool)
    GetVaultRecovery()(*GroupVaultRecoveryAvailability)
    GetVaultRequiresActivation()(*bool)
    SetAdditionalObjects(value GroupGroup_additionalObjectsable)()
    SetApplicationAdministration(value *bool)()
    SetAuditConfig(value GroupGroupAuditConfigable)()
    SetAuditor(value *bool)()
    SetAuditRequested(value *bool)()
    SetAuthorizingGroupAuditing(value GroupGroupPrimerable)()
    SetAuthorizingGroupDelegation(value GroupGroupPrimerable)()
    SetAuthorizingGroupMembership(value GroupGroupPrimerable)()
    SetAuthorizingGroupProvisioning(value GroupGroupPrimerable)()
    SetAuthorizingGroupTypes(value []RequestAuthorizingGroupType)()
    SetClassification(value GroupGroupClassificationPrimerable)()
    SetDescription(value *string)()
    SetExtendedAccess(value *GroupGroupExtendedAccess)()
    SetHideAuditTrail(value *bool)()
    SetNestedUnder(value GroupGroupPrimerable)()
    SetOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)()
    SetPrivateGroup(value *bool)()
    SetRecordTrail(value *bool)()
    SetRotatingPasswordRequired(value *bool)()
    SetSingleManaged(value *bool)()
    SetVaultRecovery(value *GroupVaultRecoveryAvailability)()
    SetVaultRequiresActivation(value *bool)()
}
