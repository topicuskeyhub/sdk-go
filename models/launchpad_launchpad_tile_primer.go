package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LaunchpadLaunchpadTilePrimer struct {
    Linkable
}
// NewLaunchpadLaunchpadTilePrimer instantiates a new LaunchpadLaunchpadTilePrimer and sets the default values.
func NewLaunchpadLaunchpadTilePrimer()(*LaunchpadLaunchpadTilePrimer) {
    m := &LaunchpadLaunchpadTilePrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "launchpad.LaunchpadTilePrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLaunchpadLaunchpadTilePrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLaunchpadLaunchpadTilePrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "launchpad.LaunchpadTile":
                        return NewLaunchpadLaunchpadTile(), nil
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
    return NewLaunchpadLaunchpadTilePrimer(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LaunchpadLaunchpadTilePrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *LaunchpadLaunchpadTilePrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type LaunchpadLaunchpadTilePrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
