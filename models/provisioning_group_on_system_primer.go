package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningGroupOnSystemPrimer struct {
    Linkable
    // The displayName property
    displayName *string
    // The nameInSystem property
    nameInSystem *string
    // The type property
    provisioningGroupOnSystemPrimerType *ProvisioningGroupOnSystemType
    // The shortNameInSystem property
    shortNameInSystem *string
}
// NewProvisioningGroupOnSystemPrimer instantiates a new ProvisioningGroupOnSystemPrimer and sets the default values.
func NewProvisioningGroupOnSystemPrimer()(*ProvisioningGroupOnSystemPrimer) {
    m := &ProvisioningGroupOnSystemPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "provisioning.GroupOnSystemPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningGroupOnSystemPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningGroupOnSystemPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "provisioning.GroupOnSystem":
                        return NewProvisioningGroupOnSystem(), nil
                    case "serviceaccount.ServiceAccountGroup":
                        return NewServiceaccountServiceAccountGroup(), nil
                }
            }
        }
    }
    return NewProvisioningGroupOnSystemPrimer(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *ProvisioningGroupOnSystemPrimer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningGroupOnSystemPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["nameInSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameInSystem(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningGroupOnSystemType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningGroupOnSystemPrimerType(val.(*ProvisioningGroupOnSystemType))
        }
        return nil
    }
    res["shortNameInSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortNameInSystem(val)
        }
        return nil
    }
    return res
}
// GetNameInSystem gets the nameInSystem property value. The nameInSystem property
// returns a *string when successful
func (m *ProvisioningGroupOnSystemPrimer) GetNameInSystem()(*string) {
    return m.nameInSystem
}
// GetProvisioningGroupOnSystemPrimerType gets the type property value. The type property
// returns a *ProvisioningGroupOnSystemType when successful
func (m *ProvisioningGroupOnSystemPrimer) GetProvisioningGroupOnSystemPrimerType()(*ProvisioningGroupOnSystemType) {
    return m.provisioningGroupOnSystemPrimerType
}
// GetShortNameInSystem gets the shortNameInSystem property value. The shortNameInSystem property
// returns a *string when successful
func (m *ProvisioningGroupOnSystemPrimer) GetShortNameInSystem()(*string) {
    return m.shortNameInSystem
}
// Serialize serializes information the current object
func (m *ProvisioningGroupOnSystemPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("nameInSystem", m.GetNameInSystem())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningGroupOnSystemPrimerType() != nil {
        cast := (*m.GetProvisioningGroupOnSystemPrimerType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ProvisioningGroupOnSystemPrimer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetNameInSystem sets the nameInSystem property value. The nameInSystem property
func (m *ProvisioningGroupOnSystemPrimer) SetNameInSystem(value *string)() {
    m.nameInSystem = value
}
// SetProvisioningGroupOnSystemPrimerType sets the type property value. The type property
func (m *ProvisioningGroupOnSystemPrimer) SetProvisioningGroupOnSystemPrimerType(value *ProvisioningGroupOnSystemType)() {
    m.provisioningGroupOnSystemPrimerType = value
}
// SetShortNameInSystem sets the shortNameInSystem property value. The shortNameInSystem property
func (m *ProvisioningGroupOnSystemPrimer) SetShortNameInSystem(value *string)() {
    m.shortNameInSystem = value
}
type ProvisioningGroupOnSystemPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetNameInSystem()(*string)
    GetProvisioningGroupOnSystemPrimerType()(*ProvisioningGroupOnSystemType)
    GetShortNameInSystem()(*string)
    SetDisplayName(value *string)()
    SetNameInSystem(value *string)()
    SetProvisioningGroupOnSystemPrimerType(value *ProvisioningGroupOnSystemType)()
    SetShortNameInSystem(value *string)()
}
