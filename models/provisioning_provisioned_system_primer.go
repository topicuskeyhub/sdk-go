package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedSystemPrimer struct {
    Linkable
    // The active property
    active *bool
    // The adminPermissions property
    adminPermissions *bool
    // The contentAdminPermissions property
    contentAdminPermissions *bool
    // The name property
    name *string
    // The organizationalUnit property
    organizationalUnit OrganizationOrganizationalUnitPrimerable
    // The ownerPermissions property
    ownerPermissions *bool
    // The type property
    provisioningProvisionedSystemPrimerType *ProvisioningProvisionedSystemType
    // The uuid property
    uuid *string
}
// NewProvisioningProvisionedSystemPrimer instantiates a new ProvisioningProvisionedSystemPrimer and sets the default values.
func NewProvisioningProvisionedSystemPrimer()(*ProvisioningProvisionedSystemPrimer) {
    m := &ProvisioningProvisionedSystemPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "provisioning.ProvisionedSystemPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "provisioning.ProvisionedSCIM":
                        return NewProvisioningProvisionedSCIM(), nil
                    case "provisioning.ProvisionedSystem":
                        return NewProvisioningProvisionedSystem(), nil
                }
            }
        }
    }
    return NewProvisioningProvisionedSystemPrimer(), nil
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *ProvisioningProvisionedSystemPrimer) GetActive()(*bool) {
    return m.active
}
// GetAdminPermissions gets the adminPermissions property value. The adminPermissions property
// returns a *bool when successful
func (m *ProvisioningProvisionedSystemPrimer) GetAdminPermissions()(*bool) {
    return m.adminPermissions
}
// GetContentAdminPermissions gets the contentAdminPermissions property value. The contentAdminPermissions property
// returns a *bool when successful
func (m *ProvisioningProvisionedSystemPrimer) GetContentAdminPermissions()(*bool) {
    return m.contentAdminPermissions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedSystemPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["adminPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminPermissions(val)
        }
        return nil
    }
    res["contentAdminPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentAdminPermissions(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["ownerPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerPermissions(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningProvisionedSystemType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningProvisionedSystemPrimerType(val.(*ProvisioningProvisionedSystemType))
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ProvisioningProvisionedSystemPrimer) GetName()(*string) {
    return m.name
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizationalUnit property
// returns a OrganizationOrganizationalUnitPrimerable when successful
func (m *ProvisioningProvisionedSystemPrimer) GetOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable) {
    return m.organizationalUnit
}
// GetOwnerPermissions gets the ownerPermissions property value. The ownerPermissions property
// returns a *bool when successful
func (m *ProvisioningProvisionedSystemPrimer) GetOwnerPermissions()(*bool) {
    return m.ownerPermissions
}
// GetProvisioningProvisionedSystemPrimerType gets the type property value. The type property
// returns a *ProvisioningProvisionedSystemType when successful
func (m *ProvisioningProvisionedSystemPrimer) GetProvisioningProvisionedSystemPrimerType()(*ProvisioningProvisionedSystemType) {
    return m.provisioningProvisionedSystemPrimerType
}
// GetUuid gets the uuid property value. The uuid property
// returns a *string when successful
func (m *ProvisioningProvisionedSystemPrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedSystemPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
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
    if m.GetProvisioningProvisionedSystemPrimerType() != nil {
        cast := (*m.GetProvisioningProvisionedSystemPrimerType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActive sets the active property value. The active property
func (m *ProvisioningProvisionedSystemPrimer) SetActive(value *bool)() {
    m.active = value
}
// SetAdminPermissions sets the adminPermissions property value. The adminPermissions property
func (m *ProvisioningProvisionedSystemPrimer) SetAdminPermissions(value *bool)() {
    m.adminPermissions = value
}
// SetContentAdminPermissions sets the contentAdminPermissions property value. The contentAdminPermissions property
func (m *ProvisioningProvisionedSystemPrimer) SetContentAdminPermissions(value *bool)() {
    m.contentAdminPermissions = value
}
// SetName sets the name property value. The name property
func (m *ProvisioningProvisionedSystemPrimer) SetName(value *string)() {
    m.name = value
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizationalUnit property
func (m *ProvisioningProvisionedSystemPrimer) SetOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)() {
    m.organizationalUnit = value
}
// SetOwnerPermissions sets the ownerPermissions property value. The ownerPermissions property
func (m *ProvisioningProvisionedSystemPrimer) SetOwnerPermissions(value *bool)() {
    m.ownerPermissions = value
}
// SetProvisioningProvisionedSystemPrimerType sets the type property value. The type property
func (m *ProvisioningProvisionedSystemPrimer) SetProvisioningProvisionedSystemPrimerType(value *ProvisioningProvisionedSystemType)() {
    m.provisioningProvisionedSystemPrimerType = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *ProvisioningProvisionedSystemPrimer) SetUuid(value *string)() {
    m.uuid = value
}
type ProvisioningProvisionedSystemPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetAdminPermissions()(*bool)
    GetContentAdminPermissions()(*bool)
    GetName()(*string)
    GetOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable)
    GetOwnerPermissions()(*bool)
    GetProvisioningProvisionedSystemPrimerType()(*ProvisioningProvisionedSystemType)
    GetUuid()(*string)
    SetActive(value *bool)()
    SetAdminPermissions(value *bool)()
    SetContentAdminPermissions(value *bool)()
    SetName(value *string)()
    SetOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)()
    SetOwnerPermissions(value *bool)()
    SetProvisioningProvisionedSystemPrimerType(value *ProvisioningProvisionedSystemType)()
    SetUuid(value *string)()
}
