package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestModificationRequestReportObjectChangeDetails struct {
    NonLinkable
    // The changeDetails property
    changeDetails []RequestModificationRequestReportChangeDetailsable
    // The name property
    name *string
}
// NewRequestModificationRequestReportObjectChangeDetails instantiates a new RequestModificationRequestReportObjectChangeDetails and sets the default values.
func NewRequestModificationRequestReportObjectChangeDetails()(*RequestModificationRequestReportObjectChangeDetails) {
    m := &RequestModificationRequestReportObjectChangeDetails{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.ModificationRequestReportObjectChangeDetails"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestModificationRequestReportObjectChangeDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestModificationRequestReportObjectChangeDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestModificationRequestReportObjectChangeDetails(), nil
}
// GetChangeDetails gets the changeDetails property value. The changeDetails property
// returns a []RequestModificationRequestReportChangeDetailsable when successful
func (m *RequestModificationRequestReportObjectChangeDetails) GetChangeDetails()([]RequestModificationRequestReportChangeDetailsable) {
    return m.changeDetails
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestModificationRequestReportObjectChangeDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["changeDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRequestModificationRequestReportChangeDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RequestModificationRequestReportChangeDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RequestModificationRequestReportChangeDetailsable)
                }
            }
            m.SetChangeDetails(res)
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
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *RequestModificationRequestReportObjectChangeDetails) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *RequestModificationRequestReportObjectChangeDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChangeDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChangeDetails()))
        for i, v := range m.GetChangeDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("changeDetails", cast)
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
// SetChangeDetails sets the changeDetails property value. The changeDetails property
func (m *RequestModificationRequestReportObjectChangeDetails) SetChangeDetails(value []RequestModificationRequestReportChangeDetailsable)() {
    m.changeDetails = value
}
// SetName sets the name property value. The name property
func (m *RequestModificationRequestReportObjectChangeDetails) SetName(value *string)() {
    m.name = value
}
type RequestModificationRequestReportObjectChangeDetailsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeDetails()([]RequestModificationRequestReportChangeDetailsable)
    GetName()(*string)
    SetChangeDetails(value []RequestModificationRequestReportChangeDetailsable)()
    SetName(value *string)()
}