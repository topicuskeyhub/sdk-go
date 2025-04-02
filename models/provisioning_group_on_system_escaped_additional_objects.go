package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningGroupOnSystem_additionalObjects struct {
    // The accessProfileProvisioning property
    accessProfileProvisioning ProfileAccessProfileProvisioningLinkableWrapperWithCountable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The audit property
    audit AuditInfoable
    // The provgroups property
    provgroups GroupProvisioningGroupLinkableWrapperWithCountable
    // The serviceAccounts property
    serviceAccounts ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable
}
// NewProvisioningGroupOnSystem_additionalObjects instantiates a new ProvisioningGroupOnSystem_additionalObjects and sets the default values.
func NewProvisioningGroupOnSystem_additionalObjects()(*ProvisioningGroupOnSystem_additionalObjects) {
    m := &ProvisioningGroupOnSystem_additionalObjects{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisioningGroupOnSystem_additionalObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningGroupOnSystem_additionalObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningGroupOnSystem_additionalObjects(), nil
}
// GetAccessProfileProvisioning gets the accessProfileProvisioning property value. The accessProfileProvisioning property
// returns a ProfileAccessProfileProvisioningLinkableWrapperWithCountable when successful
func (m *ProvisioningGroupOnSystem_additionalObjects) GetAccessProfileProvisioning()(ProfileAccessProfileProvisioningLinkableWrapperWithCountable) {
    return m.accessProfileProvisioning
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProvisioningGroupOnSystem_additionalObjects) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAudit gets the audit property value. The audit property
// returns a AuditInfoable when successful
func (m *ProvisioningGroupOnSystem_additionalObjects) GetAudit()(AuditInfoable) {
    return m.audit
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningGroupOnSystem_additionalObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessProfileProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfileProvisioningLinkableWrapperWithCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessProfileProvisioning(val.(ProfileAccessProfileProvisioningLinkableWrapperWithCountable))
        }
        return nil
    }
    res["audit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudit(val.(AuditInfoable))
        }
        return nil
    }
    res["provgroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupProvisioningGroupLinkableWrapperWithCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvgroups(val.(GroupProvisioningGroupLinkableWrapperWithCountable))
        }
        return nil
    }
    res["serviceAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccountPrimerLinkableWrapperWithCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccounts(val.(ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable))
        }
        return nil
    }
    return res
}
// GetProvgroups gets the provgroups property value. The provgroups property
// returns a GroupProvisioningGroupLinkableWrapperWithCountable when successful
func (m *ProvisioningGroupOnSystem_additionalObjects) GetProvgroups()(GroupProvisioningGroupLinkableWrapperWithCountable) {
    return m.provgroups
}
// GetServiceAccounts gets the serviceAccounts property value. The serviceAccounts property
// returns a ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable when successful
func (m *ProvisioningGroupOnSystem_additionalObjects) GetServiceAccounts()(ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable) {
    return m.serviceAccounts
}
// Serialize serializes information the current object
func (m *ProvisioningGroupOnSystem_additionalObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessProfileProvisioning", m.GetAccessProfileProvisioning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("audit", m.GetAudit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("provgroups", m.GetProvgroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("serviceAccounts", m.GetServiceAccounts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessProfileProvisioning sets the accessProfileProvisioning property value. The accessProfileProvisioning property
func (m *ProvisioningGroupOnSystem_additionalObjects) SetAccessProfileProvisioning(value ProfileAccessProfileProvisioningLinkableWrapperWithCountable)() {
    m.accessProfileProvisioning = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningGroupOnSystem_additionalObjects) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAudit sets the audit property value. The audit property
func (m *ProvisioningGroupOnSystem_additionalObjects) SetAudit(value AuditInfoable)() {
    m.audit = value
}
// SetProvgroups sets the provgroups property value. The provgroups property
func (m *ProvisioningGroupOnSystem_additionalObjects) SetProvgroups(value GroupProvisioningGroupLinkableWrapperWithCountable)() {
    m.provgroups = value
}
// SetServiceAccounts sets the serviceAccounts property value. The serviceAccounts property
func (m *ProvisioningGroupOnSystem_additionalObjects) SetServiceAccounts(value ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable)() {
    m.serviceAccounts = value
}
type ProvisioningGroupOnSystem_additionalObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessProfileProvisioning()(ProfileAccessProfileProvisioningLinkableWrapperWithCountable)
    GetAudit()(AuditInfoable)
    GetProvgroups()(GroupProvisioningGroupLinkableWrapperWithCountable)
    GetServiceAccounts()(ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable)
    SetAccessProfileProvisioning(value ProfileAccessProfileProvisioningLinkableWrapperWithCountable)()
    SetAudit(value AuditInfoable)()
    SetProvgroups(value GroupProvisioningGroupLinkableWrapperWithCountable)()
    SetServiceAccounts(value ServiceaccountServiceAccountPrimerLinkableWrapperWithCountable)()
}
