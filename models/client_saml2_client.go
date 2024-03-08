package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClientSaml2Client struct {
    ClientClientApplication
    // The attributes property
    attributes ClientSaml2Client_attributesable
    // The metadata property
    metadata *string
    // The metadataUrl property
    metadataUrl *string
    // The subjectFormat property
    subjectFormat *ClientSubjectFormat
}
// NewClientSaml2Client instantiates a new ClientSaml2Client and sets the default values.
func NewClientSaml2Client()(*ClientSaml2Client) {
    m := &ClientSaml2Client{
        ClientClientApplication: *NewClientClientApplication(),
    }
    typeEscapedValue := "client.Saml2Client"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientSaml2ClientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClientSaml2ClientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientSaml2Client(), nil
}
// GetAttributes gets the attributes property value. The attributes property
// returns a ClientSaml2Client_attributesable when successful
func (m *ClientSaml2Client) GetAttributes()(ClientSaml2Client_attributesable) {
    return m.attributes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClientSaml2Client) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClientClientApplication.GetFieldDeserializers()
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientSaml2Client_attributesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributes(val.(ClientSaml2Client_attributesable))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
        }
        return nil
    }
    res["metadataUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataUrl(val)
        }
        return nil
    }
    res["subjectFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientSubjectFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectFormat(val.(*ClientSubjectFormat))
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. The metadata property
// returns a *string when successful
func (m *ClientSaml2Client) GetMetadata()(*string) {
    return m.metadata
}
// GetMetadataUrl gets the metadataUrl property value. The metadataUrl property
// returns a *string when successful
func (m *ClientSaml2Client) GetMetadataUrl()(*string) {
    return m.metadataUrl
}
// GetSubjectFormat gets the subjectFormat property value. The subjectFormat property
// returns a *ClientSubjectFormat when successful
func (m *ClientSaml2Client) GetSubjectFormat()(*ClientSubjectFormat) {
    return m.subjectFormat
}
// Serialize serializes information the current object
func (m *ClientSaml2Client) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadataUrl", m.GetMetadataUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectFormat() != nil {
        cast := (*m.GetSubjectFormat()).String()
        err = writer.WriteStringValue("subjectFormat", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttributes sets the attributes property value. The attributes property
func (m *ClientSaml2Client) SetAttributes(value ClientSaml2Client_attributesable)() {
    m.attributes = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *ClientSaml2Client) SetMetadata(value *string)() {
    m.metadata = value
}
// SetMetadataUrl sets the metadataUrl property value. The metadataUrl property
func (m *ClientSaml2Client) SetMetadataUrl(value *string)() {
    m.metadataUrl = value
}
// SetSubjectFormat sets the subjectFormat property value. The subjectFormat property
func (m *ClientSaml2Client) SetSubjectFormat(value *ClientSubjectFormat)() {
    m.subjectFormat = value
}
type ClientSaml2Clientable interface {
    ClientClientApplicationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributes()(ClientSaml2Client_attributesable)
    GetMetadata()(*string)
    GetMetadataUrl()(*string)
    GetSubjectFormat()(*ClientSubjectFormat)
    SetAttributes(value ClientSaml2Client_attributesable)()
    SetMetadata(value *string)()
    SetMetadataUrl(value *string)()
    SetSubjectFormat(value *ClientSubjectFormat)()
}
