package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientOAuth2Client 
type ClientOAuth2Client struct {
    ClientClientApplication
    // The accountPermissions property
    accountPermissions []AuthPermissionable
    // The attributes property
    attributes ClientOAuth2Client_attributesable
    // The callbackURI property
    callbackURI *string
    // The debugMode property
    debugMode *bool
    // The idTokenClaims property
    idTokenClaims *string
    // The initiateLoginURI property
    initiateLoginURI *string
    // The profile property
    profile *ClientOAuth2ClientProfile
    // The resourceURIs property
    resourceURIs *string
    // The sharedSecret property
    sharedSecret VaultVaultRecordPrimerable
    // The shareSecretInVault property
    shareSecretInVault *bool
    // The showLandingPage property
    showLandingPage *bool
    // The useClientCredentials property
    useClientCredentials *bool
}
// NewClientOAuth2Client instantiates a new clientOAuth2Client and sets the default values.
func NewClientOAuth2Client()(*ClientOAuth2Client) {
    m := &ClientOAuth2Client{
        ClientClientApplication: *NewClientClientApplication(),
    }
    typeEscapedValue := "client.OAuth2Client"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientOAuth2ClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientOAuth2ClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientOAuth2Client(), nil
}
// GetAccountPermissions gets the accountPermissions property value. The accountPermissions property
func (m *ClientOAuth2Client) GetAccountPermissions()([]AuthPermissionable) {
    return m.accountPermissions
}
// GetAttributes gets the attributes property value. The attributes property
func (m *ClientOAuth2Client) GetAttributes()(ClientOAuth2Client_attributesable) {
    return m.attributes
}
// GetCallbackURI gets the callbackURI property value. The callbackURI property
func (m *ClientOAuth2Client) GetCallbackURI()(*string) {
    return m.callbackURI
}
// GetDebugMode gets the debugMode property value. The debugMode property
func (m *ClientOAuth2Client) GetDebugMode()(*bool) {
    return m.debugMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientOAuth2Client) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClientClientApplication.GetFieldDeserializers()
    res["accountPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthPermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthPermissionable)
                }
            }
            m.SetAccountPermissions(res)
        }
        return nil
    }
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientOAuth2Client_attributesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributes(val.(ClientOAuth2Client_attributesable))
        }
        return nil
    }
    res["callbackURI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackURI(val)
        }
        return nil
    }
    res["debugMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebugMode(val)
        }
        return nil
    }
    res["idTokenClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdTokenClaims(val)
        }
        return nil
    }
    res["initiateLoginURI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiateLoginURI(val)
        }
        return nil
    }
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientOAuth2ClientProfile)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val.(*ClientOAuth2ClientProfile))
        }
        return nil
    }
    res["resourceURIs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceURIs(val)
        }
        return nil
    }
    res["sharedSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedSecret(val.(VaultVaultRecordPrimerable))
        }
        return nil
    }
    res["shareSecretInVault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareSecretInVault(val)
        }
        return nil
    }
    res["showLandingPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLandingPage(val)
        }
        return nil
    }
    res["useClientCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseClientCredentials(val)
        }
        return nil
    }
    return res
}
// GetIdTokenClaims gets the idTokenClaims property value. The idTokenClaims property
func (m *ClientOAuth2Client) GetIdTokenClaims()(*string) {
    return m.idTokenClaims
}
// GetInitiateLoginURI gets the initiateLoginURI property value. The initiateLoginURI property
func (m *ClientOAuth2Client) GetInitiateLoginURI()(*string) {
    return m.initiateLoginURI
}
// GetProfile gets the profile property value. The profile property
func (m *ClientOAuth2Client) GetProfile()(*ClientOAuth2ClientProfile) {
    return m.profile
}
// GetResourceURIs gets the resourceURIs property value. The resourceURIs property
func (m *ClientOAuth2Client) GetResourceURIs()(*string) {
    return m.resourceURIs
}
// GetSharedSecret gets the sharedSecret property value. The sharedSecret property
func (m *ClientOAuth2Client) GetSharedSecret()(VaultVaultRecordPrimerable) {
    return m.sharedSecret
}
// GetShareSecretInVault gets the shareSecretInVault property value. The shareSecretInVault property
func (m *ClientOAuth2Client) GetShareSecretInVault()(*bool) {
    return m.shareSecretInVault
}
// GetShowLandingPage gets the showLandingPage property value. The showLandingPage property
func (m *ClientOAuth2Client) GetShowLandingPage()(*bool) {
    return m.showLandingPage
}
// GetUseClientCredentials gets the useClientCredentials property value. The useClientCredentials property
func (m *ClientOAuth2Client) GetUseClientCredentials()(*bool) {
    return m.useClientCredentials
}
// Serialize serializes information the current object
func (m *ClientOAuth2Client) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClientClientApplication.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("attributes", m.GetAttributes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callbackURI", m.GetCallbackURI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("debugMode", m.GetDebugMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("idTokenClaims", m.GetIdTokenClaims())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiateLoginURI", m.GetInitiateLoginURI())
        if err != nil {
            return err
        }
    }
    if m.GetProfile() != nil {
        cast := (*m.GetProfile()).String()
        err = writer.WriteStringValue("profile", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceURIs", m.GetResourceURIs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedSecret", m.GetSharedSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shareSecretInVault", m.GetShareSecretInVault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showLandingPage", m.GetShowLandingPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useClientCredentials", m.GetUseClientCredentials())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountPermissions sets the accountPermissions property value. The accountPermissions property
func (m *ClientOAuth2Client) SetAccountPermissions(value []AuthPermissionable)() {
    m.accountPermissions = value
}
// SetAttributes sets the attributes property value. The attributes property
func (m *ClientOAuth2Client) SetAttributes(value ClientOAuth2Client_attributesable)() {
    m.attributes = value
}
// SetCallbackURI sets the callbackURI property value. The callbackURI property
func (m *ClientOAuth2Client) SetCallbackURI(value *string)() {
    m.callbackURI = value
}
// SetDebugMode sets the debugMode property value. The debugMode property
func (m *ClientOAuth2Client) SetDebugMode(value *bool)() {
    m.debugMode = value
}
// SetIdTokenClaims sets the idTokenClaims property value. The idTokenClaims property
func (m *ClientOAuth2Client) SetIdTokenClaims(value *string)() {
    m.idTokenClaims = value
}
// SetInitiateLoginURI sets the initiateLoginURI property value. The initiateLoginURI property
func (m *ClientOAuth2Client) SetInitiateLoginURI(value *string)() {
    m.initiateLoginURI = value
}
// SetProfile sets the profile property value. The profile property
func (m *ClientOAuth2Client) SetProfile(value *ClientOAuth2ClientProfile)() {
    m.profile = value
}
// SetResourceURIs sets the resourceURIs property value. The resourceURIs property
func (m *ClientOAuth2Client) SetResourceURIs(value *string)() {
    m.resourceURIs = value
}
// SetSharedSecret sets the sharedSecret property value. The sharedSecret property
func (m *ClientOAuth2Client) SetSharedSecret(value VaultVaultRecordPrimerable)() {
    m.sharedSecret = value
}
// SetShareSecretInVault sets the shareSecretInVault property value. The shareSecretInVault property
func (m *ClientOAuth2Client) SetShareSecretInVault(value *bool)() {
    m.shareSecretInVault = value
}
// SetShowLandingPage sets the showLandingPage property value. The showLandingPage property
func (m *ClientOAuth2Client) SetShowLandingPage(value *bool)() {
    m.showLandingPage = value
}
// SetUseClientCredentials sets the useClientCredentials property value. The useClientCredentials property
func (m *ClientOAuth2Client) SetUseClientCredentials(value *bool)() {
    m.useClientCredentials = value
}
// ClientOAuth2Clientable 
type ClientOAuth2Clientable interface {
    ClientClientApplicationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountPermissions()([]AuthPermissionable)
    GetAttributes()(ClientOAuth2Client_attributesable)
    GetCallbackURI()(*string)
    GetDebugMode()(*bool)
    GetIdTokenClaims()(*string)
    GetInitiateLoginURI()(*string)
    GetProfile()(*ClientOAuth2ClientProfile)
    GetResourceURIs()(*string)
    GetSharedSecret()(VaultVaultRecordPrimerable)
    GetShareSecretInVault()(*bool)
    GetShowLandingPage()(*bool)
    GetUseClientCredentials()(*bool)
    SetAccountPermissions(value []AuthPermissionable)()
    SetAttributes(value ClientOAuth2Client_attributesable)()
    SetCallbackURI(value *string)()
    SetDebugMode(value *bool)()
    SetIdTokenClaims(value *string)()
    SetInitiateLoginURI(value *string)()
    SetProfile(value *ClientOAuth2ClientProfile)()
    SetResourceURIs(value *string)()
    SetSharedSecret(value VaultVaultRecordPrimerable)()
    SetShareSecretInVault(value *bool)()
    SetShowLandingPage(value *bool)()
    SetUseClientCredentials(value *bool)()
}
