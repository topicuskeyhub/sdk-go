package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateCertificatePrimer 
type CertificateCertificatePrimer struct {
    Linkable
    // The alias property
    alias *string
    // The type property
    certificateCertificatePrimerType *CertificateCertificateType
    // The certificateData property
    certificateData []string
    // The expiration property
    expiration *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The fingerprintSha1 property
    fingerprintSha1 *string
    // The fingerprintSha256 property
    fingerprintSha256 *string
    // The global property
    global *bool
    // The subjectDN property
    subjectDN *string
    // The uuid property
    uuid *string
}
// NewCertificateCertificatePrimer instantiates a new certificateCertificatePrimer and sets the default values.
func NewCertificateCertificatePrimer()(*CertificateCertificatePrimer) {
    m := &CertificateCertificatePrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "certificate.CertificatePrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateCertificateCertificatePrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateCertificatePrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "certificate.Certificate":
                        return NewCertificateCertificate(), nil
                }
            }
        }
    }
    return NewCertificateCertificatePrimer(), nil
}
// GetAlias gets the alias property value. The alias property
func (m *CertificateCertificatePrimer) GetAlias()(*string) {
    return m.alias
}
// GetCertificateCertificatePrimerType gets the type property value. The type property
func (m *CertificateCertificatePrimer) GetCertificateCertificatePrimerType()(*CertificateCertificateType) {
    return m.certificateCertificatePrimerType
}
// GetCertificateData gets the certificateData property value. The certificateData property
func (m *CertificateCertificatePrimer) GetCertificateData()([]string) {
    return m.certificateData
}
// GetExpiration gets the expiration property value. The expiration property
func (m *CertificateCertificatePrimer) GetExpiration()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateCertificatePrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["alias"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlias(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateCertificateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateCertificatePrimerType(val.(*CertificateCertificateType))
        }
        return nil
    }
    res["certificateData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCertificateData(res)
        }
        return nil
    }
    res["expiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiration(val)
        }
        return nil
    }
    res["fingerprintSha1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintSha1(val)
        }
        return nil
    }
    res["fingerprintSha256"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintSha256(val)
        }
        return nil
    }
    res["global"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobal(val)
        }
        return nil
    }
    res["subjectDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectDN(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetFingerprintSha1 gets the fingerprintSha1 property value. The fingerprintSha1 property
func (m *CertificateCertificatePrimer) GetFingerprintSha1()(*string) {
    return m.fingerprintSha1
}
// GetFingerprintSha256 gets the fingerprintSha256 property value. The fingerprintSha256 property
func (m *CertificateCertificatePrimer) GetFingerprintSha256()(*string) {
    return m.fingerprintSha256
}
// GetGlobal gets the global property value. The global property
func (m *CertificateCertificatePrimer) GetGlobal()(*bool) {
    return m.global
}
// GetSubjectDN gets the subjectDN property value. The subjectDN property
func (m *CertificateCertificatePrimer) GetSubjectDN()(*string) {
    return m.subjectDN
}
// GetUuid gets the uuid property value. The uuid property
func (m *CertificateCertificatePrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *CertificateCertificatePrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alias", m.GetAlias())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateCertificatePrimerType() != nil {
        cast := (*m.GetCertificateCertificatePrimerType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertificateData() != nil {
        err = writer.WriteCollectionOfStringValues("certificateData", m.GetCertificateData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("global", m.GetGlobal())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlias sets the alias property value. The alias property
func (m *CertificateCertificatePrimer) SetAlias(value *string)() {
    m.alias = value
}
// SetCertificateCertificatePrimerType sets the type property value. The type property
func (m *CertificateCertificatePrimer) SetCertificateCertificatePrimerType(value *CertificateCertificateType)() {
    m.certificateCertificatePrimerType = value
}
// SetCertificateData sets the certificateData property value. The certificateData property
func (m *CertificateCertificatePrimer) SetCertificateData(value []string)() {
    m.certificateData = value
}
// SetExpiration sets the expiration property value. The expiration property
func (m *CertificateCertificatePrimer) SetExpiration(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiration = value
}
// SetFingerprintSha1 sets the fingerprintSha1 property value. The fingerprintSha1 property
func (m *CertificateCertificatePrimer) SetFingerprintSha1(value *string)() {
    m.fingerprintSha1 = value
}
// SetFingerprintSha256 sets the fingerprintSha256 property value. The fingerprintSha256 property
func (m *CertificateCertificatePrimer) SetFingerprintSha256(value *string)() {
    m.fingerprintSha256 = value
}
// SetGlobal sets the global property value. The global property
func (m *CertificateCertificatePrimer) SetGlobal(value *bool)() {
    m.global = value
}
// SetSubjectDN sets the subjectDN property value. The subjectDN property
func (m *CertificateCertificatePrimer) SetSubjectDN(value *string)() {
    m.subjectDN = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *CertificateCertificatePrimer) SetUuid(value *string)() {
    m.uuid = value
}
// CertificateCertificatePrimerable 
type CertificateCertificatePrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlias()(*string)
    GetCertificateCertificatePrimerType()(*CertificateCertificateType)
    GetCertificateData()([]string)
    GetExpiration()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFingerprintSha1()(*string)
    GetFingerprintSha256()(*string)
    GetGlobal()(*bool)
    GetSubjectDN()(*string)
    GetUuid()(*string)
    SetAlias(value *string)()
    SetCertificateCertificatePrimerType(value *CertificateCertificateType)()
    SetCertificateData(value []string)()
    SetExpiration(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFingerprintSha1(value *string)()
    SetFingerprintSha256(value *string)()
    SetGlobal(value *bool)()
    SetSubjectDN(value *string)()
    SetUuid(value *string)()
}
