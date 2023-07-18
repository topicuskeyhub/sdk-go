package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupClassification 
type GroupGroupClassification struct {
    GroupGroupClassificationPrimer
    // The additionalObjects property
    additionalObjects GroupGroupClassification_additionalObjectsable
    // The authorizingGroupAuditingRequired property
    authorizingGroupAuditingRequired *bool
    // The authorizingGroupDelegationRequired property
    authorizingGroupDelegationRequired *bool
    // The authorizingGroupMembershipRequired property
    authorizingGroupMembershipRequired *bool
    // The authorizingGroupProvisioningRequired property
    authorizingGroupProvisioningRequired *bool
    // The defaultClassification property
    defaultClassification *bool
    // The description property
    description *string
    // The maximumAuditInterval property
    maximumAuditInterval *int32
    // The minimumNrManagers property
    minimumNrManagers *int32
    // The recordTrailRequired property
    recordTrailRequired *bool
    // The requiredMonths property
    requiredMonths []Month
    // The rotatingPasswordRequired property
    rotatingPasswordRequired *bool
    // The vaultRequiresActivation property
    vaultRequiresActivation *bool
}
// NewGroupGroupClassification instantiates a new groupGroupClassification and sets the default values.
func NewGroupGroupClassification()(*GroupGroupClassification) {
    m := &GroupGroupClassification{
        GroupGroupClassificationPrimer: *NewGroupGroupClassificationPrimer(),
    }
    typeEscapedValue := "group.GroupClassification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupClassificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupClassificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupClassification(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupClassification) GetAdditionalObjects()(GroupGroupClassification_additionalObjectsable) {
    return m.additionalObjects
}
// GetAuthorizingGroupAuditingRequired gets the authorizingGroupAuditingRequired property value. The authorizingGroupAuditingRequired property
func (m *GroupGroupClassification) GetAuthorizingGroupAuditingRequired()(*bool) {
    return m.authorizingGroupAuditingRequired
}
// GetAuthorizingGroupDelegationRequired gets the authorizingGroupDelegationRequired property value. The authorizingGroupDelegationRequired property
func (m *GroupGroupClassification) GetAuthorizingGroupDelegationRequired()(*bool) {
    return m.authorizingGroupDelegationRequired
}
// GetAuthorizingGroupMembershipRequired gets the authorizingGroupMembershipRequired property value. The authorizingGroupMembershipRequired property
func (m *GroupGroupClassification) GetAuthorizingGroupMembershipRequired()(*bool) {
    return m.authorizingGroupMembershipRequired
}
// GetAuthorizingGroupProvisioningRequired gets the authorizingGroupProvisioningRequired property value. The authorizingGroupProvisioningRequired property
func (m *GroupGroupClassification) GetAuthorizingGroupProvisioningRequired()(*bool) {
    return m.authorizingGroupProvisioningRequired
}
// GetDefaultClassification gets the defaultClassification property value. The defaultClassification property
func (m *GroupGroupClassification) GetDefaultClassification()(*bool) {
    return m.defaultClassification
}
// GetDescription gets the description property value. The description property
func (m *GroupGroupClassification) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupClassification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupGroupClassificationPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupClassification_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(GroupGroupClassification_additionalObjectsable))
        }
        return nil
    }
    res["authorizingGroupAuditingRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupAuditingRequired(val)
        }
        return nil
    }
    res["authorizingGroupDelegationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupDelegationRequired(val)
        }
        return nil
    }
    res["authorizingGroupMembershipRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupMembershipRequired(val)
        }
        return nil
    }
    res["authorizingGroupProvisioningRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizingGroupProvisioningRequired(val)
        }
        return nil
    }
    res["defaultClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultClassification(val)
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
    res["maximumAuditInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAuditInterval(val)
        }
        return nil
    }
    res["minimumNrManagers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumNrManagers(val)
        }
        return nil
    }
    res["recordTrailRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordTrailRequired(val)
        }
        return nil
    }
    res["requiredMonths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseMonth)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Month, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*Month))
                }
            }
            m.SetRequiredMonths(res)
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
// GetMaximumAuditInterval gets the maximumAuditInterval property value. The maximumAuditInterval property
func (m *GroupGroupClassification) GetMaximumAuditInterval()(*int32) {
    return m.maximumAuditInterval
}
// GetMinimumNrManagers gets the minimumNrManagers property value. The minimumNrManagers property
func (m *GroupGroupClassification) GetMinimumNrManagers()(*int32) {
    return m.minimumNrManagers
}
// GetRecordTrailRequired gets the recordTrailRequired property value. The recordTrailRequired property
func (m *GroupGroupClassification) GetRecordTrailRequired()(*bool) {
    return m.recordTrailRequired
}
// GetRequiredMonths gets the requiredMonths property value. The requiredMonths property
func (m *GroupGroupClassification) GetRequiredMonths()([]Month) {
    return m.requiredMonths
}
// GetRotatingPasswordRequired gets the rotatingPasswordRequired property value. The rotatingPasswordRequired property
func (m *GroupGroupClassification) GetRotatingPasswordRequired()(*bool) {
    return m.rotatingPasswordRequired
}
// GetVaultRequiresActivation gets the vaultRequiresActivation property value. The vaultRequiresActivation property
func (m *GroupGroupClassification) GetVaultRequiresActivation()(*bool) {
    return m.vaultRequiresActivation
}
// Serialize serializes information the current object
func (m *GroupGroupClassification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupGroupClassificationPrimer.Serialize(writer)
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
        err = writer.WriteBoolValue("authorizingGroupAuditingRequired", m.GetAuthorizingGroupAuditingRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizingGroupDelegationRequired", m.GetAuthorizingGroupDelegationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizingGroupMembershipRequired", m.GetAuthorizingGroupMembershipRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizingGroupProvisioningRequired", m.GetAuthorizingGroupProvisioningRequired())
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
    {
        err = writer.WriteInt32Value("maximumAuditInterval", m.GetMaximumAuditInterval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumNrManagers", m.GetMinimumNrManagers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("recordTrailRequired", m.GetRecordTrailRequired())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredMonths() != nil {
        err = writer.WriteCollectionOfStringValues("requiredMonths", SerializeMonth(m.GetRequiredMonths()))
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
        err = writer.WriteBoolValue("vaultRequiresActivation", m.GetVaultRequiresActivation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *GroupGroupClassification) SetAdditionalObjects(value GroupGroupClassification_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetAuthorizingGroupAuditingRequired sets the authorizingGroupAuditingRequired property value. The authorizingGroupAuditingRequired property
func (m *GroupGroupClassification) SetAuthorizingGroupAuditingRequired(value *bool)() {
    m.authorizingGroupAuditingRequired = value
}
// SetAuthorizingGroupDelegationRequired sets the authorizingGroupDelegationRequired property value. The authorizingGroupDelegationRequired property
func (m *GroupGroupClassification) SetAuthorizingGroupDelegationRequired(value *bool)() {
    m.authorizingGroupDelegationRequired = value
}
// SetAuthorizingGroupMembershipRequired sets the authorizingGroupMembershipRequired property value. The authorizingGroupMembershipRequired property
func (m *GroupGroupClassification) SetAuthorizingGroupMembershipRequired(value *bool)() {
    m.authorizingGroupMembershipRequired = value
}
// SetAuthorizingGroupProvisioningRequired sets the authorizingGroupProvisioningRequired property value. The authorizingGroupProvisioningRequired property
func (m *GroupGroupClassification) SetAuthorizingGroupProvisioningRequired(value *bool)() {
    m.authorizingGroupProvisioningRequired = value
}
// SetDefaultClassification sets the defaultClassification property value. The defaultClassification property
func (m *GroupGroupClassification) SetDefaultClassification(value *bool)() {
    m.defaultClassification = value
}
// SetDescription sets the description property value. The description property
func (m *GroupGroupClassification) SetDescription(value *string)() {
    m.description = value
}
// SetMaximumAuditInterval sets the maximumAuditInterval property value. The maximumAuditInterval property
func (m *GroupGroupClassification) SetMaximumAuditInterval(value *int32)() {
    m.maximumAuditInterval = value
}
// SetMinimumNrManagers sets the minimumNrManagers property value. The minimumNrManagers property
func (m *GroupGroupClassification) SetMinimumNrManagers(value *int32)() {
    m.minimumNrManagers = value
}
// SetRecordTrailRequired sets the recordTrailRequired property value. The recordTrailRequired property
func (m *GroupGroupClassification) SetRecordTrailRequired(value *bool)() {
    m.recordTrailRequired = value
}
// SetRequiredMonths sets the requiredMonths property value. The requiredMonths property
func (m *GroupGroupClassification) SetRequiredMonths(value []Month)() {
    m.requiredMonths = value
}
// SetRotatingPasswordRequired sets the rotatingPasswordRequired property value. The rotatingPasswordRequired property
func (m *GroupGroupClassification) SetRotatingPasswordRequired(value *bool)() {
    m.rotatingPasswordRequired = value
}
// SetVaultRequiresActivation sets the vaultRequiresActivation property value. The vaultRequiresActivation property
func (m *GroupGroupClassification) SetVaultRequiresActivation(value *bool)() {
    m.vaultRequiresActivation = value
}
// GroupGroupClassificationable 
type GroupGroupClassificationable interface {
    GroupGroupClassificationPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(GroupGroupClassification_additionalObjectsable)
    GetAuthorizingGroupAuditingRequired()(*bool)
    GetAuthorizingGroupDelegationRequired()(*bool)
    GetAuthorizingGroupMembershipRequired()(*bool)
    GetAuthorizingGroupProvisioningRequired()(*bool)
    GetDefaultClassification()(*bool)
    GetDescription()(*string)
    GetMaximumAuditInterval()(*int32)
    GetMinimumNrManagers()(*int32)
    GetRecordTrailRequired()(*bool)
    GetRequiredMonths()([]Month)
    GetRotatingPasswordRequired()(*bool)
    GetVaultRequiresActivation()(*bool)
    SetAdditionalObjects(value GroupGroupClassification_additionalObjectsable)()
    SetAuthorizingGroupAuditingRequired(value *bool)()
    SetAuthorizingGroupDelegationRequired(value *bool)()
    SetAuthorizingGroupMembershipRequired(value *bool)()
    SetAuthorizingGroupProvisioningRequired(value *bool)()
    SetDefaultClassification(value *bool)()
    SetDescription(value *string)()
    SetMaximumAuditInterval(value *int32)()
    SetMinimumNrManagers(value *int32)()
    SetRecordTrailRequired(value *bool)()
    SetRequiredMonths(value []Month)()
    SetRotatingPasswordRequired(value *bool)()
    SetVaultRequiresActivation(value *bool)()
}
