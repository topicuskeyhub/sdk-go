package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestModificationRequestReportGroupChangeDetails struct {
    NonLinkable
    // The changeDetails property
    changeDetails []RequestModificationRequestReportChangeDetailsable
    // The groupName property
    groupName *string
}
// NewRequestModificationRequestReportGroupChangeDetails instantiates a new RequestModificationRequestReportGroupChangeDetails and sets the default values.
func NewRequestModificationRequestReportGroupChangeDetails()(*RequestModificationRequestReportGroupChangeDetails) {
    m := &RequestModificationRequestReportGroupChangeDetails{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.ModificationRequestReportGroupChangeDetails"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestModificationRequestReportGroupChangeDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestModificationRequestReportGroupChangeDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestModificationRequestReportGroupChangeDetails(), nil
}
// GetChangeDetails gets the changeDetails property value. The changeDetails property
// returns a []RequestModificationRequestReportChangeDetailsable when successful
func (m *RequestModificationRequestReportGroupChangeDetails) GetChangeDetails()([]RequestModificationRequestReportChangeDetailsable) {
    return m.changeDetails
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestModificationRequestReportGroupChangeDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["groupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    return res
}
// GetGroupName gets the groupName property value. The groupName property
// returns a *string when successful
func (m *RequestModificationRequestReportGroupChangeDetails) GetGroupName()(*string) {
    return m.groupName
}
// Serialize serializes information the current object
func (m *RequestModificationRequestReportGroupChangeDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChangeDetails sets the changeDetails property value. The changeDetails property
func (m *RequestModificationRequestReportGroupChangeDetails) SetChangeDetails(value []RequestModificationRequestReportChangeDetailsable)() {
    m.changeDetails = value
}
// SetGroupName sets the groupName property value. The groupName property
func (m *RequestModificationRequestReportGroupChangeDetails) SetGroupName(value *string)() {
    m.groupName = value
}
type RequestModificationRequestReportGroupChangeDetailsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeDetails()([]RequestModificationRequestReportChangeDetailsable)
    GetGroupName()(*string)
    SetChangeDetails(value []RequestModificationRequestReportChangeDetailsable)()
    SetGroupName(value *string)()
}
