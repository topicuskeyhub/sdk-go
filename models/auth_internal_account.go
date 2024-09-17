package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthInternalAccount struct {
    AuthAccount
    // The activationCode property
    activationCode *string
    // The activationDeadline property
    activationDeadline *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The company property
    company *string
    // The status property
    status *AuthInternalAccountStatus
}
// NewAuthInternalAccount instantiates a new AuthInternalAccount and sets the default values.
func NewAuthInternalAccount()(*AuthInternalAccount) {
    m := &AuthInternalAccount{
        AuthAccount: *NewAuthAccount(),
    }
    typeEscapedValue := "auth.InternalAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthInternalAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthInternalAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthInternalAccount(), nil
}
// GetActivationCode gets the activationCode property value. The activationCode property
// returns a *string when successful
func (m *AuthInternalAccount) GetActivationCode()(*string) {
    return m.activationCode
}
// GetActivationDeadline gets the activationDeadline property value. The activationDeadline property
// returns a *Time when successful
func (m *AuthInternalAccount) GetActivationDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activationDeadline
}
// GetCompany gets the company property value. The company property
// returns a *string when successful
func (m *AuthInternalAccount) GetCompany()(*string) {
    return m.company
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthInternalAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthAccount.GetFieldDeserializers()
    res["activationCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationCode(val)
        }
        return nil
    }
    res["activationDeadline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationDeadline(val)
        }
        return nil
    }
    res["company"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthInternalAccountStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AuthInternalAccountStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
// returns a *AuthInternalAccountStatus when successful
func (m *AuthInternalAccount) GetStatus()(*AuthInternalAccountStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *AuthInternalAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthAccount.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationCode sets the activationCode property value. The activationCode property
func (m *AuthInternalAccount) SetActivationCode(value *string)() {
    m.activationCode = value
}
// SetActivationDeadline sets the activationDeadline property value. The activationDeadline property
func (m *AuthInternalAccount) SetActivationDeadline(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activationDeadline = value
}
// SetCompany sets the company property value. The company property
func (m *AuthInternalAccount) SetCompany(value *string)() {
    m.company = value
}
// SetStatus sets the status property value. The status property
func (m *AuthInternalAccount) SetStatus(value *AuthInternalAccountStatus)() {
    m.status = value
}
type AuthInternalAccountable interface {
    AuthAccountable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationCode()(*string)
    GetActivationDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCompany()(*string)
    GetStatus()(*AuthInternalAccountStatus)
    SetActivationCode(value *string)()
    SetActivationDeadline(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCompany(value *string)()
    SetStatus(value *AuthInternalAccountStatus)()
}
