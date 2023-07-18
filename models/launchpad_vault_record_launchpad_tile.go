package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LaunchpadVaultRecordLaunchpadTile 
type LaunchpadVaultRecordLaunchpadTile struct {
    LaunchpadLaunchpadTile
}
// NewLaunchpadVaultRecordLaunchpadTile instantiates a new launchpadVaultRecordLaunchpadTile and sets the default values.
func NewLaunchpadVaultRecordLaunchpadTile()(*LaunchpadVaultRecordLaunchpadTile) {
    m := &LaunchpadVaultRecordLaunchpadTile{
        LaunchpadLaunchpadTile: *NewLaunchpadLaunchpadTile(),
    }
    typeEscapedValue := "launchpad.VaultRecordLaunchpadTile"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLaunchpadVaultRecordLaunchpadTileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLaunchpadVaultRecordLaunchpadTileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLaunchpadVaultRecordLaunchpadTile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LaunchpadVaultRecordLaunchpadTile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LaunchpadLaunchpadTile.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *LaunchpadVaultRecordLaunchpadTile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LaunchpadLaunchpadTile.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// LaunchpadVaultRecordLaunchpadTileable 
type LaunchpadVaultRecordLaunchpadTileable interface {
    LaunchpadLaunchpadTileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
