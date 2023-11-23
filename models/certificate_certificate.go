package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateCertificate 
type CertificateCertificate struct {
    CertificateCertificatePrimer
    // The additionalObjects property
    additionalObjects CertificateCertificate_additionalObjectsable
    // The keyData property
    keyData *string
}
// NewCertificateCertificate instantiates a new certificateCertificate and sets the default values.
func NewCertificateCertificate()(*CertificateCertificate) {
    m := &CertificateCertificate{
        CertificateCertificatePrimer: *NewCertificateCertificatePrimer(),
    }
    typeEscapedValue := "certificate.Certificate"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateCertificateCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateCertificate(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
func (m *CertificateCertificate) GetAdditionalObjects()(CertificateCertificate_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CertificateCertificatePrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificate_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(CertificateCertificate_additionalObjectsable))
        }
        return nil
    }
    res["keyData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyData(val)
        }
        return nil
    }
    return res
}
// GetKeyData gets the keyData property value. The keyData property
func (m *CertificateCertificate) GetKeyData()(*string) {
    return m.keyData
}
// Serialize serializes information the current object
func (m *CertificateCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CertificateCertificatePrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyData", m.GetKeyData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *CertificateCertificate) SetAdditionalObjects(value CertificateCertificate_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetKeyData sets the keyData property value. The keyData property
func (m *CertificateCertificate) SetKeyData(value *string)() {
    m.keyData = value
}
// CertificateCertificateable 
type CertificateCertificateable interface {
    CertificateCertificatePrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(CertificateCertificate_additionalObjectsable)
    GetKeyData()(*string)
    SetAdditionalObjects(value CertificateCertificate_additionalObjectsable)()
    SetKeyData(value *string)()
}
