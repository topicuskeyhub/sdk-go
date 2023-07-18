package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VaultVaultRecordPrimer 
type VaultVaultRecordPrimer struct {
    Linkable
    // The color property
    color *VaultVaultRecordColor
    // The name property
    name *string
    // The shareEndTime property
    shareEndTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The uuid property
    uuid *string
}
// NewVaultVaultRecordPrimer instantiates a new vaultVaultRecordPrimer and sets the default values.
func NewVaultVaultRecordPrimer()(*VaultVaultRecordPrimer) {
    m := &VaultVaultRecordPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "vault.VaultRecordPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultRecordPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVaultVaultRecordPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "vault.VaultRecord":
                        return NewVaultVaultRecord(), nil
                }
            }
        }
    }
    return NewVaultVaultRecordPrimer(), nil
}
// GetColor gets the color property value. The color property
func (m *VaultVaultRecordPrimer) GetColor()(*VaultVaultRecordColor) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VaultVaultRecordPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVaultVaultRecordColor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*VaultVaultRecordColor))
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
    res["shareEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareEndTime(val)
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
func (m *VaultVaultRecordPrimer) GetName()(*string) {
    return m.name
}
// GetShareEndTime gets the shareEndTime property value. The shareEndTime property
func (m *VaultVaultRecordPrimer) GetShareEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.shareEndTime
}
// GetUuid gets the uuid property value. The uuid property
func (m *VaultVaultRecordPrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *VaultVaultRecordPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
        err = writer.WriteStringValue("color", &cast)
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
// SetColor sets the color property value. The color property
func (m *VaultVaultRecordPrimer) SetColor(value *VaultVaultRecordColor)() {
    m.color = value
}
// SetName sets the name property value. The name property
func (m *VaultVaultRecordPrimer) SetName(value *string)() {
    m.name = value
}
// SetShareEndTime sets the shareEndTime property value. The shareEndTime property
func (m *VaultVaultRecordPrimer) SetShareEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.shareEndTime = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *VaultVaultRecordPrimer) SetUuid(value *string)() {
    m.uuid = value
}
// VaultVaultRecordPrimerable 
type VaultVaultRecordPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*VaultVaultRecordColor)
    GetName()(*string)
    GetShareEndTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUuid()(*string)
    SetColor(value *VaultVaultRecordColor)()
    SetName(value *string)()
    SetShareEndTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUuid(value *string)()
}
