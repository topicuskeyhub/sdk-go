package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClientLdapClient struct {
    ClientClientApplication
    // The bindDn property
    bindDn *string
    // The clientCertificate property
    clientCertificate CertificateCertificatePrimerable
    // The sharedSecret property
    sharedSecret VaultVaultRecordPrimerable
    // The shareSecretInVault property
    shareSecretInVault *bool
    // The usedForProvisioning property
    usedForProvisioning *bool
}
// NewClientLdapClient instantiates a new ClientLdapClient and sets the default values.
func NewClientLdapClient()(*ClientLdapClient) {
    m := &ClientLdapClient{
        ClientClientApplication: *NewClientClientApplication(),
    }
    typeEscapedValue := "client.LdapClient"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientLdapClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClientLdapClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientLdapClient(), nil
}
// GetBindDn gets the bindDn property value. The bindDn property
// returns a *string when successful
func (m *ClientLdapClient) GetBindDn()(*string) {
    return m.bindDn
}
// GetClientCertificate gets the clientCertificate property value. The clientCertificate property
// returns a CertificateCertificatePrimerable when successful
func (m *ClientLdapClient) GetClientCertificate()(CertificateCertificatePrimerable) {
    return m.clientCertificate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClientLdapClient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClientClientApplication.GetFieldDeserializers()
    res["bindDn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindDn(val)
        }
        return nil
    }
    res["clientCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificatePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientCertificate(val.(CertificateCertificatePrimerable))
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
    res["usedForProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedForProvisioning(val)
        }
        return nil
    }
    return res
}
// GetSharedSecret gets the sharedSecret property value. The sharedSecret property
// returns a VaultVaultRecordPrimerable when successful
func (m *ClientLdapClient) GetSharedSecret()(VaultVaultRecordPrimerable) {
    return m.sharedSecret
}
// GetShareSecretInVault gets the shareSecretInVault property value. The shareSecretInVault property
// returns a *bool when successful
func (m *ClientLdapClient) GetShareSecretInVault()(*bool) {
    return m.shareSecretInVault
}
// GetUsedForProvisioning gets the usedForProvisioning property value. The usedForProvisioning property
// returns a *bool when successful
func (m *ClientLdapClient) GetUsedForProvisioning()(*bool) {
    return m.usedForProvisioning
}
// Serialize serializes information the current object
func (m *ClientLdapClient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClientClientApplication.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("clientCertificate", m.GetClientCertificate())
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
    return nil
}
// SetBindDn sets the bindDn property value. The bindDn property
func (m *ClientLdapClient) SetBindDn(value *string)() {
    m.bindDn = value
}
// SetClientCertificate sets the clientCertificate property value. The clientCertificate property
func (m *ClientLdapClient) SetClientCertificate(value CertificateCertificatePrimerable)() {
    m.clientCertificate = value
}
// SetSharedSecret sets the sharedSecret property value. The sharedSecret property
func (m *ClientLdapClient) SetSharedSecret(value VaultVaultRecordPrimerable)() {
    m.sharedSecret = value
}
// SetShareSecretInVault sets the shareSecretInVault property value. The shareSecretInVault property
func (m *ClientLdapClient) SetShareSecretInVault(value *bool)() {
    m.shareSecretInVault = value
}
// SetUsedForProvisioning sets the usedForProvisioning property value. The usedForProvisioning property
func (m *ClientLdapClient) SetUsedForProvisioning(value *bool)() {
    m.usedForProvisioning = value
}
type ClientLdapClientable interface {
    ClientClientApplicationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBindDn()(*string)
    GetClientCertificate()(CertificateCertificatePrimerable)
    GetSharedSecret()(VaultVaultRecordPrimerable)
    GetShareSecretInVault()(*bool)
    GetUsedForProvisioning()(*bool)
    SetBindDn(value *string)()
    SetClientCertificate(value CertificateCertificatePrimerable)()
    SetSharedSecret(value VaultVaultRecordPrimerable)()
    SetShareSecretInVault(value *bool)()
    SetUsedForProvisioning(value *bool)()
}
