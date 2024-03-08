package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestPasswordResetRequestStatus struct {
    NonLinkable
    // The cooldownEnd property
    cooldownEnd *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managers property
    managers []string
    // The nrAccepted property
    nrAccepted *int64
    // The resetViaMail property
    resetViaMail *RequestPasswordResetEmailType
}
// NewRequestPasswordResetRequestStatus instantiates a new RequestPasswordResetRequestStatus and sets the default values.
func NewRequestPasswordResetRequestStatus()(*RequestPasswordResetRequestStatus) {
    m := &RequestPasswordResetRequestStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "request.PasswordResetRequestStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestPasswordResetRequestStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestPasswordResetRequestStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestPasswordResetRequestStatus(), nil
}
// GetCooldownEnd gets the cooldownEnd property value. The cooldownEnd property
// returns a *Time when successful
func (m *RequestPasswordResetRequestStatus) GetCooldownEnd()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.cooldownEnd
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestPasswordResetRequestStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["cooldownEnd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCooldownEnd(val)
        }
        return nil
    }
    res["managers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetManagers(res)
        }
        return nil
    }
    res["nrAccepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrAccepted(val)
        }
        return nil
    }
    res["resetViaMail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestPasswordResetEmailType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetViaMail(val.(*RequestPasswordResetEmailType))
        }
        return nil
    }
    return res
}
// GetManagers gets the managers property value. The managers property
// returns a []string when successful
func (m *RequestPasswordResetRequestStatus) GetManagers()([]string) {
    return m.managers
}
// GetNrAccepted gets the nrAccepted property value. The nrAccepted property
// returns a *int64 when successful
func (m *RequestPasswordResetRequestStatus) GetNrAccepted()(*int64) {
    return m.nrAccepted
}
// GetResetViaMail gets the resetViaMail property value. The resetViaMail property
// returns a *RequestPasswordResetEmailType when successful
func (m *RequestPasswordResetRequestStatus) GetResetViaMail()(*RequestPasswordResetEmailType) {
    return m.resetViaMail
}
// Serialize serializes information the current object
func (m *RequestPasswordResetRequestStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("cooldownEnd", m.GetCooldownEnd())
        if err != nil {
            return err
        }
    }
    if m.GetManagers() != nil {
        err = writer.WriteCollectionOfStringValues("managers", m.GetManagers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("nrAccepted", m.GetNrAccepted())
        if err != nil {
            return err
        }
    }
    if m.GetResetViaMail() != nil {
        cast := (*m.GetResetViaMail()).String()
        err = writer.WriteStringValue("resetViaMail", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCooldownEnd sets the cooldownEnd property value. The cooldownEnd property
func (m *RequestPasswordResetRequestStatus) SetCooldownEnd(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.cooldownEnd = value
}
// SetManagers sets the managers property value. The managers property
func (m *RequestPasswordResetRequestStatus) SetManagers(value []string)() {
    m.managers = value
}
// SetNrAccepted sets the nrAccepted property value. The nrAccepted property
func (m *RequestPasswordResetRequestStatus) SetNrAccepted(value *int64)() {
    m.nrAccepted = value
}
// SetResetViaMail sets the resetViaMail property value. The resetViaMail property
func (m *RequestPasswordResetRequestStatus) SetResetViaMail(value *RequestPasswordResetEmailType)() {
    m.resetViaMail = value
}
type RequestPasswordResetRequestStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCooldownEnd()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagers()([]string)
    GetNrAccepted()(*int64)
    GetResetViaMail()(*RequestPasswordResetEmailType)
    SetCooldownEnd(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagers(value []string)()
    SetNrAccepted(value *int64)()
    SetResetViaMail(value *RequestPasswordResetEmailType)()
}
