package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthInternalAccount 
type AuthInternalAccount struct {
    AuthAccount
    // The activationCode property
    activationCode *string
    // The activationDeadline property
    activationDeadline *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The company property
    company *string
    // The firstName property
    firstName *string
    // The lastName property
    lastName *string
    // The status property
    status *AuthInternalAccountStatus
    // The telephone property
    telephone *string
}
// NewAuthInternalAccount instantiates a new authInternalAccount and sets the default values.
func NewAuthInternalAccount()(*AuthInternalAccount) {
    m := &AuthInternalAccount{
        AuthAccount: *NewAuthAccount(),
    }
    typeEscapedValue := "auth.InternalAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthInternalAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthInternalAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthInternalAccount(), nil
}
// GetActivationCode gets the activationCode property value. The activationCode property
func (m *AuthInternalAccount) GetActivationCode()(*string) {
    return m.activationCode
}
// GetActivationDeadline gets the activationDeadline property value. The activationDeadline property
func (m *AuthInternalAccount) GetActivationDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activationDeadline
}
// GetCompany gets the company property value. The company property
func (m *AuthInternalAccount) GetCompany()(*string) {
    return m.company
}
// GetFieldDeserializers the deserialization information for the current model
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
    res["firstName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["lastName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
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
    res["telephone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTelephone(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the firstName property value. The firstName property
func (m *AuthInternalAccount) GetFirstName()(*string) {
    return m.firstName
}
// GetLastName gets the lastName property value. The lastName property
func (m *AuthInternalAccount) GetLastName()(*string) {
    return m.lastName
}
// GetStatus gets the status property value. The status property
func (m *AuthInternalAccount) GetStatus()(*AuthInternalAccountStatus) {
    return m.status
}
// GetTelephone gets the telephone property value. The telephone property
func (m *AuthInternalAccount) GetTelephone()(*string) {
    return m.telephone
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
    {
        err = writer.WriteStringValue("firstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastName", m.GetLastName())
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
    {
        err = writer.WriteStringValue("telephone", m.GetTelephone())
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
// SetFirstName sets the firstName property value. The firstName property
func (m *AuthInternalAccount) SetFirstName(value *string)() {
    m.firstName = value
}
// SetLastName sets the lastName property value. The lastName property
func (m *AuthInternalAccount) SetLastName(value *string)() {
    m.lastName = value
}
// SetStatus sets the status property value. The status property
func (m *AuthInternalAccount) SetStatus(value *AuthInternalAccountStatus)() {
    m.status = value
}
// SetTelephone sets the telephone property value. The telephone property
func (m *AuthInternalAccount) SetTelephone(value *string)() {
    m.telephone = value
}
// AuthInternalAccountable 
type AuthInternalAccountable interface {
    AuthAccountable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationCode()(*string)
    GetActivationDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCompany()(*string)
    GetFirstName()(*string)
    GetLastName()(*string)
    GetStatus()(*AuthInternalAccountStatus)
    GetTelephone()(*string)
    SetActivationCode(value *string)()
    SetActivationDeadline(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCompany(value *string)()
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetStatus(value *AuthInternalAccountStatus)()
    SetTelephone(value *string)()
}
