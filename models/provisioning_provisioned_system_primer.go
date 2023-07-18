package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningProvisionedSystemPrimer 
type ProvisioningProvisionedSystemPrimer struct {
    Linkable
    // The active property
    active *bool
    // The name property
    name *string
    // The type property
    provisioningProvisionedSystemPrimerType *ProvisioningProvisionedSystemType
    // The uuid property
    uuid *string
}
// NewProvisioningProvisionedSystemPrimer instantiates a new provisioningProvisionedSystemPrimer and sets the default values.
func NewProvisioningProvisionedSystemPrimer()(*ProvisioningProvisionedSystemPrimer) {
    m := &ProvisioningProvisionedSystemPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "provisioning.ProvisionedSystemPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
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
                    case "provisioning.ProvisionedSystem":
                        return NewProvisioningProvisionedSystem(), nil
                }
            }
        }
    }
    return NewProvisioningProvisionedSystemPrimer(), nil
}
// GetActive gets the active property value. The active property
func (m *ProvisioningProvisionedSystemPrimer) GetActive()(*bool) {
    return m.active
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *ProvisioningProvisionedSystemPrimer) GetName()(*string) {
    return m.name
}
// GetProvisioningProvisionedSystemPrimerType gets the type property value. The type property
func (m *ProvisioningProvisionedSystemPrimer) GetProvisioningProvisionedSystemPrimerType()(*ProvisioningProvisionedSystemType) {
    return m.provisioningProvisionedSystemPrimerType
}
// GetUuid gets the uuid property value. The uuid property
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
// SetName sets the name property value. The name property
func (m *ProvisioningProvisionedSystemPrimer) SetName(value *string)() {
    m.name = value
}
// SetProvisioningProvisionedSystemPrimerType sets the type property value. The type property
func (m *ProvisioningProvisionedSystemPrimer) SetProvisioningProvisionedSystemPrimerType(value *ProvisioningProvisionedSystemType)() {
    m.provisioningProvisionedSystemPrimerType = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *ProvisioningProvisionedSystemPrimer) SetUuid(value *string)() {
    m.uuid = value
}
// ProvisioningProvisionedSystemPrimerable 
type ProvisioningProvisionedSystemPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetName()(*string)
    GetProvisioningProvisionedSystemPrimerType()(*ProvisioningProvisionedSystemType)
    GetUuid()(*string)
    SetActive(value *bool)()
    SetName(value *string)()
    SetProvisioningProvisionedSystemPrimerType(value *ProvisioningProvisionedSystemType)()
    SetUuid(value *string)()
}
