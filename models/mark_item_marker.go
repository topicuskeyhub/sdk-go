package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MarkItemMarker struct {
    NonLinkable
    // The level property
    level *MarkItemMarkerLevel
    // The type property
    markItemMarkerType *MarkItemMarkerType
    // The parameters property
    parameters MarkItemMarker_parametersable
}
// NewMarkItemMarker instantiates a new MarkItemMarker and sets the default values.
func NewMarkItemMarker()(*MarkItemMarker) {
    m := &MarkItemMarker{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "mark.ItemMarker"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateMarkItemMarkerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarkItemMarkerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarkItemMarker(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarkItemMarker) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMarkItemMarkerLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val.(*MarkItemMarkerLevel))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMarkItemMarkerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarkItemMarkerType(val.(*MarkItemMarkerType))
        }
        return nil
    }
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMarkItemMarker_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(MarkItemMarker_parametersable))
        }
        return nil
    }
    return res
}
// GetLevel gets the level property value. The level property
// returns a *MarkItemMarkerLevel when successful
func (m *MarkItemMarker) GetLevel()(*MarkItemMarkerLevel) {
    return m.level
}
// GetMarkItemMarkerType gets the type property value. The type property
// returns a *MarkItemMarkerType when successful
func (m *MarkItemMarker) GetMarkItemMarkerType()(*MarkItemMarkerType) {
    return m.markItemMarkerType
}
// GetParameters gets the parameters property value. The parameters property
// returns a MarkItemMarker_parametersable when successful
func (m *MarkItemMarker) GetParameters()(MarkItemMarker_parametersable) {
    return m.parameters
}
// Serialize serializes information the current object
func (m *MarkItemMarker) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLevel() != nil {
        cast := (*m.GetLevel()).String()
        err = writer.WriteStringValue("level", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMarkItemMarkerType() != nil {
        cast := (*m.GetMarkItemMarkerType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLevel sets the level property value. The level property
func (m *MarkItemMarker) SetLevel(value *MarkItemMarkerLevel)() {
    m.level = value
}
// SetMarkItemMarkerType sets the type property value. The type property
func (m *MarkItemMarker) SetMarkItemMarkerType(value *MarkItemMarkerType)() {
    m.markItemMarkerType = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *MarkItemMarker) SetParameters(value MarkItemMarker_parametersable)() {
    m.parameters = value
}
type MarkItemMarkerable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLevel()(*MarkItemMarkerLevel)
    GetMarkItemMarkerType()(*MarkItemMarkerType)
    GetParameters()(MarkItemMarker_parametersable)
    SetLevel(value *MarkItemMarkerLevel)()
    SetMarkItemMarkerType(value *MarkItemMarkerType)()
    SetParameters(value MarkItemMarker_parametersable)()
}
