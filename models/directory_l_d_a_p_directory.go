package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryLDAPDirectory 
type DirectoryLDAPDirectory struct {
    DirectoryAccountDirectory
    // The attributesToStore property
    attributesToStore *string
    // The baseDN property
    baseDN *string
    // The clientCertificate property
    clientCertificate CertificateCertificatePrimerable
    // The dialect property
    dialect *DirectoryLDAPDialect
    // The failoverHost property
    failoverHost *string
    // The failoverTrustedCertificate property
    failoverTrustedCertificate CertificateCertificatePrimerable
    // The host property
    host *string
    // The passwordRecovery property
    passwordRecovery *DirectoryLDAPDirectoryPasswordRecovery
    // The port property
    port *int32
    // The searchBindDN property
    searchBindDN *string
    // The searchBindPassword property
    searchBindPassword *string
    // The searchFilter property
    searchFilter *string
    // The tls property
    tls *TLSLevel
    // The trustedCertificate property
    trustedCertificate CertificateCertificatePrimerable
}
// NewDirectoryLDAPDirectory instantiates a new directoryLDAPDirectory and sets the default values.
func NewDirectoryLDAPDirectory()(*DirectoryLDAPDirectory) {
    m := &DirectoryLDAPDirectory{
        DirectoryAccountDirectory: *NewDirectoryAccountDirectory(),
    }
    typeEscapedValue := "directory.LDAPDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryLDAPDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryLDAPDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryLDAPDirectory(), nil
}
// GetAttributesToStore gets the attributesToStore property value. The attributesToStore property
func (m *DirectoryLDAPDirectory) GetAttributesToStore()(*string) {
    return m.attributesToStore
}
// GetBaseDN gets the baseDN property value. The baseDN property
func (m *DirectoryLDAPDirectory) GetBaseDN()(*string) {
    return m.baseDN
}
// GetClientCertificate gets the clientCertificate property value. The clientCertificate property
func (m *DirectoryLDAPDirectory) GetClientCertificate()(CertificateCertificatePrimerable) {
    return m.clientCertificate
}
// GetDialect gets the dialect property value. The dialect property
func (m *DirectoryLDAPDirectory) GetDialect()(*DirectoryLDAPDialect) {
    return m.dialect
}
// GetFailoverHost gets the failoverHost property value. The failoverHost property
func (m *DirectoryLDAPDirectory) GetFailoverHost()(*string) {
    return m.failoverHost
}
// GetFailoverTrustedCertificate gets the failoverTrustedCertificate property value. The failoverTrustedCertificate property
func (m *DirectoryLDAPDirectory) GetFailoverTrustedCertificate()(CertificateCertificatePrimerable) {
    return m.failoverTrustedCertificate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryLDAPDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryAccountDirectory.GetFieldDeserializers()
    res["attributesToStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributesToStore(val)
        }
        return nil
    }
    res["baseDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseDN(val)
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
    res["dialect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryLDAPDialect)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialect(val.(*DirectoryLDAPDialect))
        }
        return nil
    }
    res["failoverHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailoverHost(val)
        }
        return nil
    }
    res["failoverTrustedCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificatePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailoverTrustedCertificate(val.(CertificateCertificatePrimerable))
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val)
        }
        return nil
    }
    res["passwordRecovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryLDAPDirectoryPasswordRecovery)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRecovery(val.(*DirectoryLDAPDirectoryPasswordRecovery))
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["searchBindDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchBindDN(val)
        }
        return nil
    }
    res["searchBindPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchBindPassword(val)
        }
        return nil
    }
    res["searchFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchFilter(val)
        }
        return nil
    }
    res["tls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTLSLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTls(val.(*TLSLevel))
        }
        return nil
    }
    res["trustedCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificatePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustedCertificate(val.(CertificateCertificatePrimerable))
        }
        return nil
    }
    return res
}
// GetHost gets the host property value. The host property
func (m *DirectoryLDAPDirectory) GetHost()(*string) {
    return m.host
}
// GetPasswordRecovery gets the passwordRecovery property value. The passwordRecovery property
func (m *DirectoryLDAPDirectory) GetPasswordRecovery()(*DirectoryLDAPDirectoryPasswordRecovery) {
    return m.passwordRecovery
}
// GetPort gets the port property value. The port property
func (m *DirectoryLDAPDirectory) GetPort()(*int32) {
    return m.port
}
// GetSearchBindDN gets the searchBindDN property value. The searchBindDN property
func (m *DirectoryLDAPDirectory) GetSearchBindDN()(*string) {
    return m.searchBindDN
}
// GetSearchBindPassword gets the searchBindPassword property value. The searchBindPassword property
func (m *DirectoryLDAPDirectory) GetSearchBindPassword()(*string) {
    return m.searchBindPassword
}
// GetSearchFilter gets the searchFilter property value. The searchFilter property
func (m *DirectoryLDAPDirectory) GetSearchFilter()(*string) {
    return m.searchFilter
}
// GetTls gets the tls property value. The tls property
func (m *DirectoryLDAPDirectory) GetTls()(*TLSLevel) {
    return m.tls
}
// GetTrustedCertificate gets the trustedCertificate property value. The trustedCertificate property
func (m *DirectoryLDAPDirectory) GetTrustedCertificate()(CertificateCertificatePrimerable) {
    return m.trustedCertificate
}
// Serialize serializes information the current object
func (m *DirectoryLDAPDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryAccountDirectory.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("attributesToStore", m.GetAttributesToStore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("baseDN", m.GetBaseDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clientCertificate", m.GetClientCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetDialect() != nil {
        cast := (*m.GetDialect()).String()
        err = writer.WriteStringValue("dialect", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failoverHost", m.GetFailoverHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("failoverTrustedCertificate", m.GetFailoverTrustedCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRecovery() != nil {
        cast := (*m.GetPasswordRecovery()).String()
        err = writer.WriteStringValue("passwordRecovery", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("searchBindDN", m.GetSearchBindDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("searchBindPassword", m.GetSearchBindPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("searchFilter", m.GetSearchFilter())
        if err != nil {
            return err
        }
    }
    if m.GetTls() != nil {
        cast := (*m.GetTls()).String()
        err = writer.WriteStringValue("tls", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trustedCertificate", m.GetTrustedCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttributesToStore sets the attributesToStore property value. The attributesToStore property
func (m *DirectoryLDAPDirectory) SetAttributesToStore(value *string)() {
    m.attributesToStore = value
}
// SetBaseDN sets the baseDN property value. The baseDN property
func (m *DirectoryLDAPDirectory) SetBaseDN(value *string)() {
    m.baseDN = value
}
// SetClientCertificate sets the clientCertificate property value. The clientCertificate property
func (m *DirectoryLDAPDirectory) SetClientCertificate(value CertificateCertificatePrimerable)() {
    m.clientCertificate = value
}
// SetDialect sets the dialect property value. The dialect property
func (m *DirectoryLDAPDirectory) SetDialect(value *DirectoryLDAPDialect)() {
    m.dialect = value
}
// SetFailoverHost sets the failoverHost property value. The failoverHost property
func (m *DirectoryLDAPDirectory) SetFailoverHost(value *string)() {
    m.failoverHost = value
}
// SetFailoverTrustedCertificate sets the failoverTrustedCertificate property value. The failoverTrustedCertificate property
func (m *DirectoryLDAPDirectory) SetFailoverTrustedCertificate(value CertificateCertificatePrimerable)() {
    m.failoverTrustedCertificate = value
}
// SetHost sets the host property value. The host property
func (m *DirectoryLDAPDirectory) SetHost(value *string)() {
    m.host = value
}
// SetPasswordRecovery sets the passwordRecovery property value. The passwordRecovery property
func (m *DirectoryLDAPDirectory) SetPasswordRecovery(value *DirectoryLDAPDirectoryPasswordRecovery)() {
    m.passwordRecovery = value
}
// SetPort sets the port property value. The port property
func (m *DirectoryLDAPDirectory) SetPort(value *int32)() {
    m.port = value
}
// SetSearchBindDN sets the searchBindDN property value. The searchBindDN property
func (m *DirectoryLDAPDirectory) SetSearchBindDN(value *string)() {
    m.searchBindDN = value
}
// SetSearchBindPassword sets the searchBindPassword property value. The searchBindPassword property
func (m *DirectoryLDAPDirectory) SetSearchBindPassword(value *string)() {
    m.searchBindPassword = value
}
// SetSearchFilter sets the searchFilter property value. The searchFilter property
func (m *DirectoryLDAPDirectory) SetSearchFilter(value *string)() {
    m.searchFilter = value
}
// SetTls sets the tls property value. The tls property
func (m *DirectoryLDAPDirectory) SetTls(value *TLSLevel)() {
    m.tls = value
}
// SetTrustedCertificate sets the trustedCertificate property value. The trustedCertificate property
func (m *DirectoryLDAPDirectory) SetTrustedCertificate(value CertificateCertificatePrimerable)() {
    m.trustedCertificate = value
}
// DirectoryLDAPDirectoryable 
type DirectoryLDAPDirectoryable interface {
    DirectoryAccountDirectoryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributesToStore()(*string)
    GetBaseDN()(*string)
    GetClientCertificate()(CertificateCertificatePrimerable)
    GetDialect()(*DirectoryLDAPDialect)
    GetFailoverHost()(*string)
    GetFailoverTrustedCertificate()(CertificateCertificatePrimerable)
    GetHost()(*string)
    GetPasswordRecovery()(*DirectoryLDAPDirectoryPasswordRecovery)
    GetPort()(*int32)
    GetSearchBindDN()(*string)
    GetSearchBindPassword()(*string)
    GetSearchFilter()(*string)
    GetTls()(*TLSLevel)
    GetTrustedCertificate()(CertificateCertificatePrimerable)
    SetAttributesToStore(value *string)()
    SetBaseDN(value *string)()
    SetClientCertificate(value CertificateCertificatePrimerable)()
    SetDialect(value *DirectoryLDAPDialect)()
    SetFailoverHost(value *string)()
    SetFailoverTrustedCertificate(value CertificateCertificatePrimerable)()
    SetHost(value *string)()
    SetPasswordRecovery(value *DirectoryLDAPDirectoryPasswordRecovery)()
    SetPort(value *int32)()
    SetSearchBindDN(value *string)()
    SetSearchBindPassword(value *string)()
    SetSearchFilter(value *string)()
    SetTls(value *TLSLevel)()
    SetTrustedCertificate(value CertificateCertificatePrimerable)()
}
