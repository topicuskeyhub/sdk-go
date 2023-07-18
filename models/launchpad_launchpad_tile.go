package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LaunchpadLaunchpadTile 
type LaunchpadLaunchpadTile struct {
    LaunchpadLaunchpadTilePrimer
    // The additionalObjects property
    additionalObjects LaunchpadLaunchpadTile_additionalObjectsable
    // The application property
    application ClientClientApplicationPrimerable
    // The group property
    group GroupGroupPrimerable
    // The identiconCode property
    identiconCode *int32
    // The type property
    launchpadLaunchpadTileType *LaunchpadLaunchpadTileType
    // The logo property
    logo []string
    // The vaultRecord property
    vaultRecord VaultVaultRecordPrimerable
}
// NewLaunchpadLaunchpadTile instantiates a new launchpadLaunchpadTile and sets the default values.
func NewLaunchpadLaunchpadTile()(*LaunchpadLaunchpadTile) {
    m := &LaunchpadLaunchpadTile{
        LaunchpadLaunchpadTilePrimer: *NewLaunchpadLaunchpadTilePrimer(),
    }
    typeEscapedValue := "launchpad.LaunchpadTile"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLaunchpadLaunchpadTileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLaunchpadLaunchpadTileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "launchpad.ManualLaunchpadTile":
                        return NewLaunchpadManualLaunchpadTile(), nil
                    case "launchpad.SsoApplicationLaunchpadTile":
                        return NewLaunchpadSsoApplicationLaunchpadTile(), nil
                    case "launchpad.VaultRecordLaunchpadTile":
                        return NewLaunchpadVaultRecordLaunchpadTile(), nil
                }
            }
        }
    }
    return NewLaunchpadLaunchpadTile(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *LaunchpadLaunchpadTile) GetAdditionalObjects()(LaunchpadLaunchpadTile_additionalObjectsable) {
    return m.additionalObjects
}
// GetApplication gets the application property value. The application property
func (m *LaunchpadLaunchpadTile) GetApplication()(ClientClientApplicationPrimerable) {
    return m.application
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LaunchpadLaunchpadTile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LaunchpadLaunchpadTilePrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLaunchpadLaunchpadTile_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(LaunchpadLaunchpadTile_additionalObjectsable))
        }
        return nil
    }
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplicationPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(ClientClientApplicationPrimerable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["identiconCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdenticonCode(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLaunchpadLaunchpadTileType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLaunchpadLaunchpadTileType(val.(*LaunchpadLaunchpadTileType))
        }
        return nil
    }
    res["logo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetLogo(res)
        }
        return nil
    }
    res["vaultRecord"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultRecord(val.(VaultVaultRecordPrimerable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *LaunchpadLaunchpadTile) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// GetIdenticonCode gets the identiconCode property value. The identiconCode property
func (m *LaunchpadLaunchpadTile) GetIdenticonCode()(*int32) {
    return m.identiconCode
}
// GetLaunchpadLaunchpadTileType gets the type property value. The type property
func (m *LaunchpadLaunchpadTile) GetLaunchpadLaunchpadTileType()(*LaunchpadLaunchpadTileType) {
    return m.launchpadLaunchpadTileType
}
// GetLogo gets the logo property value. The logo property
func (m *LaunchpadLaunchpadTile) GetLogo()([]string) {
    return m.logo
}
// GetVaultRecord gets the vaultRecord property value. The vaultRecord property
func (m *LaunchpadLaunchpadTile) GetVaultRecord()(VaultVaultRecordPrimerable) {
    return m.vaultRecord
}
// Serialize serializes information the current object
func (m *LaunchpadLaunchpadTile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LaunchpadLaunchpadTilePrimer.Serialize(writer)
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
        err = writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("identiconCode", m.GetIdenticonCode())
        if err != nil {
            return err
        }
    }
    if m.GetLaunchpadLaunchpadTileType() != nil {
        cast := (*m.GetLaunchpadLaunchpadTileType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLogo() != nil {
        err = writer.WriteCollectionOfStringValues("logo", m.GetLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vaultRecord", m.GetVaultRecord())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *LaunchpadLaunchpadTile) SetAdditionalObjects(value LaunchpadLaunchpadTile_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetApplication sets the application property value. The application property
func (m *LaunchpadLaunchpadTile) SetApplication(value ClientClientApplicationPrimerable)() {
    m.application = value
}
// SetGroup sets the group property value. The group property
func (m *LaunchpadLaunchpadTile) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
// SetIdenticonCode sets the identiconCode property value. The identiconCode property
func (m *LaunchpadLaunchpadTile) SetIdenticonCode(value *int32)() {
    m.identiconCode = value
}
// SetLaunchpadLaunchpadTileType sets the type property value. The type property
func (m *LaunchpadLaunchpadTile) SetLaunchpadLaunchpadTileType(value *LaunchpadLaunchpadTileType)() {
    m.launchpadLaunchpadTileType = value
}
// SetLogo sets the logo property value. The logo property
func (m *LaunchpadLaunchpadTile) SetLogo(value []string)() {
    m.logo = value
}
// SetVaultRecord sets the vaultRecord property value. The vaultRecord property
func (m *LaunchpadLaunchpadTile) SetVaultRecord(value VaultVaultRecordPrimerable)() {
    m.vaultRecord = value
}
// LaunchpadLaunchpadTileable 
type LaunchpadLaunchpadTileable interface {
    LaunchpadLaunchpadTilePrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(LaunchpadLaunchpadTile_additionalObjectsable)
    GetApplication()(ClientClientApplicationPrimerable)
    GetGroup()(GroupGroupPrimerable)
    GetIdenticonCode()(*int32)
    GetLaunchpadLaunchpadTileType()(*LaunchpadLaunchpadTileType)
    GetLogo()([]string)
    GetVaultRecord()(VaultVaultRecordPrimerable)
    SetAdditionalObjects(value LaunchpadLaunchpadTile_additionalObjectsable)()
    SetApplication(value ClientClientApplicationPrimerable)()
    SetGroup(value GroupGroupPrimerable)()
    SetIdenticonCode(value *int32)()
    SetLaunchpadLaunchpadTileType(value *LaunchpadLaunchpadTileType)()
    SetLogo(value []string)()
    SetVaultRecord(value VaultVaultRecordPrimerable)()
}
