package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryAccountDirectoryPrimer 
type DirectoryAccountDirectoryPrimer struct {
    Linkable
    // The accountValiditySupported property
    accountValiditySupported *bool
    // The active property
    active *bool
    // The type property
    directoryAccountDirectoryPrimerType *DirectoryAccountDirectoryType
    // The name property
    name *string
    // The uuid property
    uuid *string
}
// NewDirectoryAccountDirectoryPrimer instantiates a new directoryAccountDirectoryPrimer and sets the default values.
func NewDirectoryAccountDirectoryPrimer()(*DirectoryAccountDirectoryPrimer) {
    m := &DirectoryAccountDirectoryPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "directory.AccountDirectoryPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryAccountDirectoryPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryAccountDirectoryPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "directory.AccountDirectory":
                        return NewDirectoryAccountDirectory(), nil
                    case "directory.InternalDirectory":
                        return NewDirectoryInternalDirectory(), nil
                    case "directory.LDAPDirectory":
                        return NewDirectoryLDAPDirectory(), nil
                    case "directory.MaintenanceDirectory":
                        return NewDirectoryMaintenanceDirectory(), nil
                    case "directory.OIDCDirectory":
                        return NewDirectoryOIDCDirectory(), nil
                }
            }
        }
    }
    return NewDirectoryAccountDirectoryPrimer(), nil
}
// GetAccountValiditySupported gets the accountValiditySupported property value. The accountValiditySupported property
func (m *DirectoryAccountDirectoryPrimer) GetAccountValiditySupported()(*bool) {
    return m.accountValiditySupported
}
// GetActive gets the active property value. The active property
func (m *DirectoryAccountDirectoryPrimer) GetActive()(*bool) {
    return m.active
}
// GetDirectoryAccountDirectoryPrimerType gets the type property value. The type property
func (m *DirectoryAccountDirectoryPrimer) GetDirectoryAccountDirectoryPrimerType()(*DirectoryAccountDirectoryType) {
    return m.directoryAccountDirectoryPrimerType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryAccountDirectoryPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accountValiditySupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountValiditySupported(val)
        }
        return nil
    }
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryAccountDirectoryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryAccountDirectoryPrimerType(val.(*DirectoryAccountDirectoryType))
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
func (m *DirectoryAccountDirectoryPrimer) GetName()(*string) {
    return m.name
}
// GetUuid gets the uuid property value. The uuid property
func (m *DirectoryAccountDirectoryPrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *DirectoryAccountDirectoryPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDirectoryAccountDirectoryPrimerType() != nil {
        cast := (*m.GetDirectoryAccountDirectoryPrimerType()).String()
        err = writer.WriteStringValue("type", &cast)
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
    return nil
}
// SetAccountValiditySupported sets the accountValiditySupported property value. The accountValiditySupported property
func (m *DirectoryAccountDirectoryPrimer) SetAccountValiditySupported(value *bool)() {
    m.accountValiditySupported = value
}
// SetActive sets the active property value. The active property
func (m *DirectoryAccountDirectoryPrimer) SetActive(value *bool)() {
    m.active = value
}
// SetDirectoryAccountDirectoryPrimerType sets the type property value. The type property
func (m *DirectoryAccountDirectoryPrimer) SetDirectoryAccountDirectoryPrimerType(value *DirectoryAccountDirectoryType)() {
    m.directoryAccountDirectoryPrimerType = value
}
// SetName sets the name property value. The name property
func (m *DirectoryAccountDirectoryPrimer) SetName(value *string)() {
    m.name = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *DirectoryAccountDirectoryPrimer) SetUuid(value *string)() {
    m.uuid = value
}
// DirectoryAccountDirectoryPrimerable 
type DirectoryAccountDirectoryPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountValiditySupported()(*bool)
    GetActive()(*bool)
    GetDirectoryAccountDirectoryPrimerType()(*DirectoryAccountDirectoryType)
    GetName()(*string)
    GetUuid()(*string)
    SetAccountValiditySupported(value *bool)()
    SetActive(value *bool)()
    SetDirectoryAccountDirectoryPrimerType(value *DirectoryAccountDirectoryType)()
    SetName(value *string)()
    SetUuid(value *string)()
}
