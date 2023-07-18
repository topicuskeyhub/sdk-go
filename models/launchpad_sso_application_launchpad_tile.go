package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LaunchpadSsoApplicationLaunchpadTile 
type LaunchpadSsoApplicationLaunchpadTile struct {
    LaunchpadLaunchpadTile
    // The uri property
    uri *string
}
// NewLaunchpadSsoApplicationLaunchpadTile instantiates a new launchpadSsoApplicationLaunchpadTile and sets the default values.
func NewLaunchpadSsoApplicationLaunchpadTile()(*LaunchpadSsoApplicationLaunchpadTile) {
    m := &LaunchpadSsoApplicationLaunchpadTile{
        LaunchpadLaunchpadTile: *NewLaunchpadLaunchpadTile(),
    }
    typeEscapedValue := "launchpad.SsoApplicationLaunchpadTile"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLaunchpadSsoApplicationLaunchpadTileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLaunchpadSsoApplicationLaunchpadTileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLaunchpadSsoApplicationLaunchpadTile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LaunchpadSsoApplicationLaunchpadTile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LaunchpadLaunchpadTile.GetFieldDeserializers()
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
// GetUri gets the uri property value. The uri property
func (m *LaunchpadSsoApplicationLaunchpadTile) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *LaunchpadSsoApplicationLaunchpadTile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LaunchpadLaunchpadTile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUri sets the uri property value. The uri property
func (m *LaunchpadSsoApplicationLaunchpadTile) SetUri(value *string)() {
    m.uri = value
}
// LaunchpadSsoApplicationLaunchpadTileable 
type LaunchpadSsoApplicationLaunchpadTileable interface {
    LaunchpadLaunchpadTileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUri()(*string)
    SetUri(value *string)()
}
