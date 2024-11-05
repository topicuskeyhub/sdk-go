package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthStoredUserSession struct {
    Linkable
    // The ipAddress property
    ipAddress *string
    // The lastUsed property
    lastUsed *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The location property
    location *string
    // The loginDate property
    loginDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The sessionId property
    sessionId *string
    // The userAgent property
    userAgent *string
    // The username property
    username *string
}
// NewAuthStoredUserSession instantiates a new AuthStoredUserSession and sets the default values.
func NewAuthStoredUserSession()(*AuthStoredUserSession) {
    m := &AuthStoredUserSession{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "auth.StoredUserSession"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthStoredUserSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthStoredUserSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthStoredUserSession(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthStoredUserSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["lastUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["loginDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginDate(val)
        }
        return nil
    }
    res["sessionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
        }
        return nil
    }
    res["userAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAgent(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. The ipAddress property
// returns a *string when successful
func (m *AuthStoredUserSession) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetLastUsed gets the lastUsed property value. The lastUsed property
// returns a *Time when successful
func (m *AuthStoredUserSession) GetLastUsed()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUsed
}
// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *AuthStoredUserSession) GetLocation()(*string) {
    return m.location
}
// GetLoginDate gets the loginDate property value. The loginDate property
// returns a *Time when successful
func (m *AuthStoredUserSession) GetLoginDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.loginDate
}
// GetSessionId gets the sessionId property value. The sessionId property
// returns a *string when successful
func (m *AuthStoredUserSession) GetSessionId()(*string) {
    return m.sessionId
}
// GetUserAgent gets the userAgent property value. The userAgent property
// returns a *string when successful
func (m *AuthStoredUserSession) GetUserAgent()(*string) {
    return m.userAgent
}
// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *AuthStoredUserSession) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *AuthStoredUserSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUsed", m.GetLastUsed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("loginDate", m.GetLoginDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sessionId", m.GetSessionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userAgent", m.GetUserAgent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIpAddress sets the ipAddress property value. The ipAddress property
func (m *AuthStoredUserSession) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetLastUsed sets the lastUsed property value. The lastUsed property
func (m *AuthStoredUserSession) SetLastUsed(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUsed = value
}
// SetLocation sets the location property value. The location property
func (m *AuthStoredUserSession) SetLocation(value *string)() {
    m.location = value
}
// SetLoginDate sets the loginDate property value. The loginDate property
func (m *AuthStoredUserSession) SetLoginDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.loginDate = value
}
// SetSessionId sets the sessionId property value. The sessionId property
func (m *AuthStoredUserSession) SetSessionId(value *string)() {
    m.sessionId = value
}
// SetUserAgent sets the userAgent property value. The userAgent property
func (m *AuthStoredUserSession) SetUserAgent(value *string)() {
    m.userAgent = value
}
// SetUsername sets the username property value. The username property
func (m *AuthStoredUserSession) SetUsername(value *string)() {
    m.username = value
}
type AuthStoredUserSessionable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIpAddress()(*string)
    GetLastUsed()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocation()(*string)
    GetLoginDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSessionId()(*string)
    GetUserAgent()(*string)
    GetUsername()(*string)
    SetIpAddress(value *string)()
    SetLastUsed(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocation(value *string)()
    SetLoginDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSessionId(value *string)()
    SetUserAgent(value *string)()
    SetUsername(value *string)()
}
