package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LaunchpadDisplayedLaunchpadTile 
type LaunchpadDisplayedLaunchpadTile struct {
    Linkable
    // The group property
    group GroupGroupPrimerable
    // The identiconCode property
    identiconCode *int32
    // The logo property
    logo *string
    // The tile property
    tile LaunchpadLaunchpadTilePrimerable
    // The title property
    title *string
    // The uri property
    uri *string
}
// NewLaunchpadDisplayedLaunchpadTile instantiates a new launchpadDisplayedLaunchpadTile and sets the default values.
func NewLaunchpadDisplayedLaunchpadTile()(*LaunchpadDisplayedLaunchpadTile) {
    m := &LaunchpadDisplayedLaunchpadTile{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "launchpad.DisplayedLaunchpadTile"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLaunchpadDisplayedLaunchpadTileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLaunchpadDisplayedLaunchpadTileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLaunchpadDisplayedLaunchpadTile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LaunchpadDisplayedLaunchpadTile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
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
    res["logo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogo(val)
        }
        return nil
    }
    res["tile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLaunchpadLaunchpadTilePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTile(val.(LaunchpadLaunchpadTilePrimerable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *LaunchpadDisplayedLaunchpadTile) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// GetIdenticonCode gets the identiconCode property value. The identiconCode property
func (m *LaunchpadDisplayedLaunchpadTile) GetIdenticonCode()(*int32) {
    return m.identiconCode
}
// GetLogo gets the logo property value. The logo property
func (m *LaunchpadDisplayedLaunchpadTile) GetLogo()(*string) {
    return m.logo
}
// GetTile gets the tile property value. The tile property
func (m *LaunchpadDisplayedLaunchpadTile) GetTile()(LaunchpadLaunchpadTilePrimerable) {
    return m.tile
}
// GetTitle gets the title property value. The title property
func (m *LaunchpadDisplayedLaunchpadTile) GetTitle()(*string) {
    return m.title
}
// GetUri gets the uri property value. The uri property
func (m *LaunchpadDisplayedLaunchpadTile) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *LaunchpadDisplayedLaunchpadTile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
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
    {
        err = writer.WriteStringValue("logo", m.GetLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tile", m.GetTile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The group property
func (m *LaunchpadDisplayedLaunchpadTile) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
// SetIdenticonCode sets the identiconCode property value. The identiconCode property
func (m *LaunchpadDisplayedLaunchpadTile) SetIdenticonCode(value *int32)() {
    m.identiconCode = value
}
// SetLogo sets the logo property value. The logo property
func (m *LaunchpadDisplayedLaunchpadTile) SetLogo(value *string)() {
    m.logo = value
}
// SetTile sets the tile property value. The tile property
func (m *LaunchpadDisplayedLaunchpadTile) SetTile(value LaunchpadLaunchpadTilePrimerable)() {
    m.tile = value
}
// SetTitle sets the title property value. The title property
func (m *LaunchpadDisplayedLaunchpadTile) SetTitle(value *string)() {
    m.title = value
}
// SetUri sets the uri property value. The uri property
func (m *LaunchpadDisplayedLaunchpadTile) SetUri(value *string)() {
    m.uri = value
}
// LaunchpadDisplayedLaunchpadTileable 
type LaunchpadDisplayedLaunchpadTileable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroup()(GroupGroupPrimerable)
    GetIdenticonCode()(*int32)
    GetLogo()(*string)
    GetTile()(LaunchpadLaunchpadTilePrimerable)
    GetTitle()(*string)
    GetUri()(*string)
    SetGroup(value GroupGroupPrimerable)()
    SetIdenticonCode(value *int32)()
    SetLogo(value *string)()
    SetTile(value LaunchpadLaunchpadTilePrimerable)()
    SetTitle(value *string)()
    SetUri(value *string)()
}
