package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MarkItemMarkers struct {
    NonLinkable
    // The markers property
    markers []MarkItemMarkerable
}
// NewMarkItemMarkers instantiates a new MarkItemMarkers and sets the default values.
func NewMarkItemMarkers()(*MarkItemMarkers) {
    m := &MarkItemMarkers{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "mark.ItemMarkers"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateMarkItemMarkersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarkItemMarkersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarkItemMarkers(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarkItemMarkers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["markers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMarkItemMarkerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MarkItemMarkerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MarkItemMarkerable)
                }
            }
            m.SetMarkers(res)
        }
        return nil
    }
    return res
}
// GetMarkers gets the markers property value. The markers property
// returns a []MarkItemMarkerable when successful
func (m *MarkItemMarkers) GetMarkers()([]MarkItemMarkerable) {
    return m.markers
}
// Serialize serializes information the current object
func (m *MarkItemMarkers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMarkers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMarkers()))
        for i, v := range m.GetMarkers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("markers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMarkers sets the markers property value. The markers property
func (m *MarkItemMarkers) SetMarkers(value []MarkItemMarkerable)() {
    m.markers = value
}
type MarkItemMarkersable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMarkers()([]MarkItemMarkerable)
    SetMarkers(value []MarkItemMarkerable)()
}
