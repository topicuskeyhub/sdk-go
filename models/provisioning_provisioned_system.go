package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedSystem 
type ProvisioningProvisionedSystem struct {
    ProvisioningProvisionedSystemPrimer
    // The accountCount property
    accountCount *int32
    // The additionalObjects property
    additionalObjects ProvisioningProvisionedSystem_additionalObjectsable
    // The contentAdministrator property
    contentAdministrator GroupGroupPrimerable
    // The externalUuid property
    externalUuid *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The owner property
    owner GroupGroupPrimerable
    // The selfServiceExistingGroups property
    selfServiceExistingGroups *bool
    // The selfServiceNewGroups property
    selfServiceNewGroups *bool
    // The selfServiceNewNamespaces property
    selfServiceNewNamespaces *bool
    // The selfServiceServiceAccounts property
    selfServiceServiceAccounts *bool
    // The shouldDestroyUnknownAccounts property
    shouldDestroyUnknownAccounts *bool
    // The technicalAdministrator property
    technicalAdministrator GroupGroupPrimerable
    // The usernamePrefix property
    usernamePrefix *string
}
// NewProvisioningProvisionedSystem instantiates a new provisioningProvisionedSystem and sets the default values.
func NewProvisioningProvisionedSystem()(*ProvisioningProvisionedSystem) {
    m := &ProvisioningProvisionedSystem{
        ProvisioningProvisionedSystemPrimer: *NewProvisioningProvisionedSystemPrimer(),
    }
    typeEscapedValue := "provisioning.ProvisionedSystem"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningProvisionedSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "provisioning.AbstractProvisionedLDAP":
                        return NewProvisioningAbstractProvisionedLDAP(), nil
                    case "provisioning.ProvisionedAD":
                        return NewProvisioningProvisionedAD(), nil
                    case "provisioning.ProvisionedAzureOIDCDirectory":
                        return NewProvisioningProvisionedAzureOIDCDirectory(), nil
                    case "provisioning.ProvisionedAzureSyncLDAPDirectory":
                        return NewProvisioningProvisionedAzureSyncLDAPDirectory(), nil
                    case "provisioning.ProvisionedAzureTenant":
                        return NewProvisioningProvisionedAzureTenant(), nil
                    case "provisioning.ProvisionedInternalLDAP":
                        return NewProvisioningProvisionedInternalLDAP(), nil
                    case "provisioning.ProvisionedLDAP":
                        return NewProvisioningProvisionedLDAP(), nil
                    case "provisioning.ProvisionedLDAPDirectory":
                        return NewProvisioningProvisionedLDAPDirectory(), nil
                    case "provisioning.ProvisionedNamespace":
                        return NewProvisioningProvisionedNamespace(), nil
                }
            }
        }
    }
    return NewProvisioningProvisionedSystem(), nil
}
// GetAccountCount gets the accountCount property value. The accountCount property
func (m *ProvisioningProvisionedSystem) GetAccountCount()(*int32) {
    return m.accountCount
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *ProvisioningProvisionedSystem) GetAdditionalObjects()(ProvisioningProvisionedSystem_additionalObjectsable) {
    return m.additionalObjects
}
// GetContentAdministrator gets the contentAdministrator property value. The contentAdministrator property
func (m *ProvisioningProvisionedSystem) GetContentAdministrator()(GroupGroupPrimerable) {
    return m.contentAdministrator
}
// GetExternalUuid gets the externalUuid property value. The externalUuid property
func (m *ProvisioningProvisionedSystem) GetExternalUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.externalUuid
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningProvisionedSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystemPrimer.GetFieldDeserializers()
    res["accountCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountCount(val)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystem_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProvisioningProvisionedSystem_additionalObjectsable))
        }
        return nil
    }
    res["contentAdministrator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentAdministrator(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["externalUuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUuid(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["selfServiceExistingGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceExistingGroups(val)
        }
        return nil
    }
    res["selfServiceNewGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceNewGroups(val)
        }
        return nil
    }
    res["selfServiceNewNamespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceNewNamespaces(val)
        }
        return nil
    }
    res["selfServiceServiceAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceServiceAccounts(val)
        }
        return nil
    }
    res["shouldDestroyUnknownAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldDestroyUnknownAccounts(val)
        }
        return nil
    }
    res["technicalAdministrator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnicalAdministrator(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["usernamePrefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernamePrefix(val)
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The owner property
func (m *ProvisioningProvisionedSystem) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// GetSelfServiceExistingGroups gets the selfServiceExistingGroups property value. The selfServiceExistingGroups property
func (m *ProvisioningProvisionedSystem) GetSelfServiceExistingGroups()(*bool) {
    return m.selfServiceExistingGroups
}
// GetSelfServiceNewGroups gets the selfServiceNewGroups property value. The selfServiceNewGroups property
func (m *ProvisioningProvisionedSystem) GetSelfServiceNewGroups()(*bool) {
    return m.selfServiceNewGroups
}
// GetSelfServiceNewNamespaces gets the selfServiceNewNamespaces property value. The selfServiceNewNamespaces property
func (m *ProvisioningProvisionedSystem) GetSelfServiceNewNamespaces()(*bool) {
    return m.selfServiceNewNamespaces
}
// GetSelfServiceServiceAccounts gets the selfServiceServiceAccounts property value. The selfServiceServiceAccounts property
func (m *ProvisioningProvisionedSystem) GetSelfServiceServiceAccounts()(*bool) {
    return m.selfServiceServiceAccounts
}
// GetShouldDestroyUnknownAccounts gets the shouldDestroyUnknownAccounts property value. The shouldDestroyUnknownAccounts property
func (m *ProvisioningProvisionedSystem) GetShouldDestroyUnknownAccounts()(*bool) {
    return m.shouldDestroyUnknownAccounts
}
// GetTechnicalAdministrator gets the technicalAdministrator property value. The technicalAdministrator property
func (m *ProvisioningProvisionedSystem) GetTechnicalAdministrator()(GroupGroupPrimerable) {
    return m.technicalAdministrator
}
// GetUsernamePrefix gets the usernamePrefix property value. The usernamePrefix property
func (m *ProvisioningProvisionedSystem) GetUsernamePrefix()(*string) {
    return m.usernamePrefix
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystemPrimer.Serialize(writer)
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
        err = writer.WriteObjectValue("contentAdministrator", m.GetContentAdministrator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("selfServiceExistingGroups", m.GetSelfServiceExistingGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("selfServiceNewGroups", m.GetSelfServiceNewGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("selfServiceNewNamespaces", m.GetSelfServiceNewNamespaces())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("selfServiceServiceAccounts", m.GetSelfServiceServiceAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shouldDestroyUnknownAccounts", m.GetShouldDestroyUnknownAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("technicalAdministrator", m.GetTechnicalAdministrator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernamePrefix", m.GetUsernamePrefix())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountCount sets the accountCount property value. The accountCount property
func (m *ProvisioningProvisionedSystem) SetAccountCount(value *int32)() {
    m.accountCount = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProvisioningProvisionedSystem) SetAdditionalObjects(value ProvisioningProvisionedSystem_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetContentAdministrator sets the contentAdministrator property value. The contentAdministrator property
func (m *ProvisioningProvisionedSystem) SetContentAdministrator(value GroupGroupPrimerable)() {
    m.contentAdministrator = value
}
// SetExternalUuid sets the externalUuid property value. The externalUuid property
func (m *ProvisioningProvisionedSystem) SetExternalUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.externalUuid = value
}
// SetOwner sets the owner property value. The owner property
func (m *ProvisioningProvisionedSystem) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
// SetSelfServiceExistingGroups sets the selfServiceExistingGroups property value. The selfServiceExistingGroups property
func (m *ProvisioningProvisionedSystem) SetSelfServiceExistingGroups(value *bool)() {
    m.selfServiceExistingGroups = value
}
// SetSelfServiceNewGroups sets the selfServiceNewGroups property value. The selfServiceNewGroups property
func (m *ProvisioningProvisionedSystem) SetSelfServiceNewGroups(value *bool)() {
    m.selfServiceNewGroups = value
}
// SetSelfServiceNewNamespaces sets the selfServiceNewNamespaces property value. The selfServiceNewNamespaces property
func (m *ProvisioningProvisionedSystem) SetSelfServiceNewNamespaces(value *bool)() {
    m.selfServiceNewNamespaces = value
}
// SetSelfServiceServiceAccounts sets the selfServiceServiceAccounts property value. The selfServiceServiceAccounts property
func (m *ProvisioningProvisionedSystem) SetSelfServiceServiceAccounts(value *bool)() {
    m.selfServiceServiceAccounts = value
}
// SetShouldDestroyUnknownAccounts sets the shouldDestroyUnknownAccounts property value. The shouldDestroyUnknownAccounts property
func (m *ProvisioningProvisionedSystem) SetShouldDestroyUnknownAccounts(value *bool)() {
    m.shouldDestroyUnknownAccounts = value
}
// SetTechnicalAdministrator sets the technicalAdministrator property value. The technicalAdministrator property
func (m *ProvisioningProvisionedSystem) SetTechnicalAdministrator(value GroupGroupPrimerable)() {
    m.technicalAdministrator = value
}
// SetUsernamePrefix sets the usernamePrefix property value. The usernamePrefix property
func (m *ProvisioningProvisionedSystem) SetUsernamePrefix(value *string)() {
    m.usernamePrefix = value
}
// ProvisioningProvisionedSystemable 
type ProvisioningProvisionedSystemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemPrimerable
    GetAccountCount()(*int32)
    GetAdditionalObjects()(ProvisioningProvisionedSystem_additionalObjectsable)
    GetContentAdministrator()(GroupGroupPrimerable)
    GetExternalUuid()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetOwner()(GroupGroupPrimerable)
    GetSelfServiceExistingGroups()(*bool)
    GetSelfServiceNewGroups()(*bool)
    GetSelfServiceNewNamespaces()(*bool)
    GetSelfServiceServiceAccounts()(*bool)
    GetShouldDestroyUnknownAccounts()(*bool)
    GetTechnicalAdministrator()(GroupGroupPrimerable)
    GetUsernamePrefix()(*string)
    SetAccountCount(value *int32)()
    SetAdditionalObjects(value ProvisioningProvisionedSystem_additionalObjectsable)()
    SetContentAdministrator(value GroupGroupPrimerable)()
    SetExternalUuid(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetOwner(value GroupGroupPrimerable)()
    SetSelfServiceExistingGroups(value *bool)()
    SetSelfServiceNewGroups(value *bool)()
    SetSelfServiceNewNamespaces(value *bool)()
    SetSelfServiceServiceAccounts(value *bool)()
    SetShouldDestroyUnknownAccounts(value *bool)()
    SetTechnicalAdministrator(value GroupGroupPrimerable)()
    SetUsernamePrefix(value *string)()
}
