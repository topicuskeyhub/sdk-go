package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IdentitysourceIdentitySourcePrimer struct {
    Linkable
    // The active property
    active *bool
    // The type property
    identitysourceIdentitySourcePrimerType *IdentitysourceIdentitySourceType
    // The name property
    name *string
    // The uuid property
    uuid *string
}
// NewIdentitysourceIdentitySourcePrimer instantiates a new IdentitysourceIdentitySourcePrimer and sets the default values.
func NewIdentitysourceIdentitySourcePrimer()(*IdentitysourceIdentitySourcePrimer) {
    m := &IdentitysourceIdentitySourcePrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "identitysource.IdentitySourcePrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateIdentitysourceIdentitySourcePrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIdentitysourceIdentitySourcePrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "identitysource.AFASIdentitySource":
                        return NewIdentitysourceAFASIdentitySource(), nil
                    case "identitysource.IdentitySource":
                        return NewIdentitysourceIdentitySource(), nil
                }
            }
        }
    }
    return NewIdentitysourceIdentitySourcePrimer(), nil
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *IdentitysourceIdentitySourcePrimer) GetActive()(*bool) {
    return m.active
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IdentitysourceIdentitySourcePrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentitysourceIdentitySourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitysourceIdentitySourcePrimerType(val.(*IdentitysourceIdentitySourceType))
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
// GetIdentitysourceIdentitySourcePrimerType gets the type property value. The type property
// returns a *IdentitysourceIdentitySourceType when successful
func (m *IdentitysourceIdentitySourcePrimer) GetIdentitysourceIdentitySourcePrimerType()(*IdentitysourceIdentitySourceType) {
    return m.identitysourceIdentitySourcePrimerType
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *IdentitysourceIdentitySourcePrimer) GetName()(*string) {
    return m.name
}
// GetUuid gets the uuid property value. The uuid property
// returns a *string when successful
func (m *IdentitysourceIdentitySourcePrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *IdentitysourceIdentitySourcePrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetIdentitysourceIdentitySourcePrimerType() != nil {
        cast := (*m.GetIdentitysourceIdentitySourcePrimerType()).String()
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
// SetActive sets the active property value. The active property
func (m *IdentitysourceIdentitySourcePrimer) SetActive(value *bool)() {
    m.active = value
}
// SetIdentitysourceIdentitySourcePrimerType sets the type property value. The type property
func (m *IdentitysourceIdentitySourcePrimer) SetIdentitysourceIdentitySourcePrimerType(value *IdentitysourceIdentitySourceType)() {
    m.identitysourceIdentitySourcePrimerType = value
}
// SetName sets the name property value. The name property
func (m *IdentitysourceIdentitySourcePrimer) SetName(value *string)() {
    m.name = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *IdentitysourceIdentitySourcePrimer) SetUuid(value *string)() {
    m.uuid = value
}
type IdentitysourceIdentitySourcePrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActive()(*bool)
    GetIdentitysourceIdentitySourcePrimerType()(*IdentitysourceIdentitySourceType)
    GetName()(*string)
    GetUuid()(*string)
    SetActive(value *bool)()
    SetIdentitysourceIdentitySourcePrimerType(value *IdentitysourceIdentitySourceType)()
    SetName(value *string)()
    SetUuid(value *string)()
}
