package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestModificationRequestReportChangeDetails struct {
    NonLinkable
    // The change property
    change *RequestModificationRequestReportChange
    // The objectNames property
    objectNames []string
}
// NewRequestModificationRequestReportChangeDetails instantiates a new RequestModificationRequestReportChangeDetails and sets the default values.
func NewRequestModificationRequestReportChangeDetails()(*RequestModificationRequestReportChangeDetails) {
    m := &RequestModificationRequestReportChangeDetails{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.ModificationRequestReportChangeDetails"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestModificationRequestReportChangeDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestModificationRequestReportChangeDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestModificationRequestReportChangeDetails(), nil
}
// GetChange gets the change property value. The change property
// returns a *RequestModificationRequestReportChange when successful
func (m *RequestModificationRequestReportChangeDetails) GetChange()(*RequestModificationRequestReportChange) {
    return m.change
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestModificationRequestReportChangeDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["change"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestModificationRequestReportChange)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChange(val.(*RequestModificationRequestReportChange))
        }
        return nil
    }
    res["objectNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetObjectNames(res)
        }
        return nil
    }
    return res
}
// GetObjectNames gets the objectNames property value. The objectNames property
// returns a []string when successful
func (m *RequestModificationRequestReportChangeDetails) GetObjectNames()([]string) {
    return m.objectNames
}
// Serialize serializes information the current object
func (m *RequestModificationRequestReportChangeDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChange() != nil {
        cast := (*m.GetChange()).String()
        err = writer.WriteStringValue("change", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetObjectNames() != nil {
        err = writer.WriteCollectionOfStringValues("objectNames", m.GetObjectNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChange sets the change property value. The change property
func (m *RequestModificationRequestReportChangeDetails) SetChange(value *RequestModificationRequestReportChange)() {
    m.change = value
}
// SetObjectNames sets the objectNames property value. The objectNames property
func (m *RequestModificationRequestReportChangeDetails) SetObjectNames(value []string)() {
    m.objectNames = value
}
type RequestModificationRequestReportChangeDetailsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChange()(*RequestModificationRequestReportChange)
    GetObjectNames()([]string)
    SetChange(value *RequestModificationRequestReportChange)()
    SetObjectNames(value []string)()
}
