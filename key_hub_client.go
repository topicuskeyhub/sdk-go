package sdkgo

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i25381d874a7e3097702763676eb8a598478a75d07191e3563cf4ef48e5ac1fc6 "github.com/topicuskeyhub/sdk-go/serviceaccount"
    i29b864faf50a3a1afdf661366bb4234f149662638d0e3841e35acccd9352cf20 "github.com/topicuskeyhub/sdk-go/provisioninggroup"
    i3786de14b5891e0b6237ed5a04286048bb7ccab8227911414f7bb8f1e165a7c0 "github.com/topicuskeyhub/sdk-go/groupclassification"
    i385775f507ab02e04cf3e40e8696ee406fb1371e877539b6dfbde0c8e4b123b2 "github.com/topicuskeyhub/sdk-go/launchpadtile"
    i43af1840a4616958e60c4cbee17a55c9c7345f03d798c4b93f71ba6ea8a8fed7 "github.com/topicuskeyhub/sdk-go/groupfolder"
    i4e9df38273eb83f44b000e660fb1ceeb2fdcb7939eb1e594bf0e1c6dd130c8e2 "github.com/topicuskeyhub/sdk-go/request"
    i4f6f8a97ade4bcd9f00e48fb22660ba2cb3b474e860f677cfc37ab0f9b6d8370 "github.com/topicuskeyhub/sdk-go/groupclient"
    i6f1edc8cb6d249770cb1634d5c07b8945f2349df7d873c8e22898a5bfd8123aa "github.com/topicuskeyhub/sdk-go/group"
    i7f1810e47506a1087d42711bae10d2eacb3eb432b8e9e68ef97ff738e3bb6945 "github.com/topicuskeyhub/sdk-go/webhook"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i98ee0f07cd6158c2879ae08ef6a731b5a05b3610b774bf729ac7c7bf7d1efa2c "github.com/topicuskeyhub/sdk-go/numberseq"
    i9a416ff0973cadc86a9386e3f28ad7c475076105c80d51f978001646d6a93eb8 "github.com/topicuskeyhub/sdk-go/audit"
    i9acdc2253f5a58228075fc1e46bdd8aca4f69569179560e0bf2e87d0811bfc18 "github.com/topicuskeyhub/sdk-go/deletedvaultholder"
    i9c60bc0bc864eda4e8fcd5447d4e572d9b6ba65365f60dea8e1097309b0531da "github.com/topicuskeyhub/sdk-go/vaultrecord"
    ia4ccd640f04375f8459d703d9bd3e36f5d6a57b48d83c18fcc686ac92a251806 "github.com/topicuskeyhub/sdk-go/certificate"
    ia9c84b329cdbc7e9961b248161459b4dfb0d53b98e1b471cd7c10b6d027a90d3 "github.com/topicuskeyhub/sdk-go/info"
    ib269aad83a59eeaa00948e8e0cf61b0f8955fcb99e4910b745aa70768d6d63c4 "github.com/topicuskeyhub/sdk-go/account"
    ib7bf253b601139c5fca70bba4f4843e6731e3522bee75715df4b77d2107eff28 "github.com/topicuskeyhub/sdk-go/directory"
    ibdde25a8196c6c3406e5b0b02de36d833849b28ee37dca06694dc2ec8a725bdc "github.com/topicuskeyhub/sdk-go/system"
    ifb310ad74101a00ed1acdc17b384ce2e8964f5a38f9d9302b22aaaa28ff4d3e4 "github.com/topicuskeyhub/sdk-go/organizationalunit"
    ifdfb91cbbf9b2d09fc5fb9b2405c3a72348d071234970eca4305212177f72cf2 "github.com/topicuskeyhub/sdk-go/client"
)

// KeyHubClient the main entry point of the SDK, exposes the configuration and the fluent API.
type KeyHubClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Account the account property
func (m *KeyHubClient) Account()(*ib269aad83a59eeaa00948e8e0cf61b0f8955fcb99e4910b745aa70768d6d63c4.AccountRequestBuilder) {
    return ib269aad83a59eeaa00948e8e0cf61b0f8955fcb99e4910b745aa70768d6d63c4.NewAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Audit the audit property
func (m *KeyHubClient) Audit()(*i9a416ff0973cadc86a9386e3f28ad7c475076105c80d51f978001646d6a93eb8.AuditRequestBuilder) {
    return i9a416ff0973cadc86a9386e3f28ad7c475076105c80d51f978001646d6a93eb8.NewAuditRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Certificate the certificate property
func (m *KeyHubClient) Certificate()(*ia4ccd640f04375f8459d703d9bd3e36f5d6a57b48d83c18fcc686ac92a251806.CertificateRequestBuilder) {
    return ia4ccd640f04375f8459d703d9bd3e36f5d6a57b48d83c18fcc686ac92a251806.NewCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Client the client property
func (m *KeyHubClient) Client()(*ifdfb91cbbf9b2d09fc5fb9b2405c3a72348d071234970eca4305212177f72cf2.ClientRequestBuilder) {
    return ifdfb91cbbf9b2d09fc5fb9b2405c3a72348d071234970eca4305212177f72cf2.NewClientRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewKeyHubClient instantiates a new KeyHubClient and sets the default values.
func NewKeyHubClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeyHubClient) {
    m := &KeyHubClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://keyhub.example.com/keyhub/rest/v1")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Deletedvaultholder the deletedvaultholder property
func (m *KeyHubClient) Deletedvaultholder()(*i9acdc2253f5a58228075fc1e46bdd8aca4f69569179560e0bf2e87d0811bfc18.DeletedvaultholderRequestBuilder) {
    return i9acdc2253f5a58228075fc1e46bdd8aca4f69569179560e0bf2e87d0811bfc18.NewDeletedvaultholderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Directory the directory property
func (m *KeyHubClient) Directory()(*ib7bf253b601139c5fca70bba4f4843e6731e3522bee75715df4b77d2107eff28.DirectoryRequestBuilder) {
    return ib7bf253b601139c5fca70bba4f4843e6731e3522bee75715df4b77d2107eff28.NewDirectoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Group the group property
func (m *KeyHubClient) Group()(*i6f1edc8cb6d249770cb1634d5c07b8945f2349df7d873c8e22898a5bfd8123aa.GroupRequestBuilder) {
    return i6f1edc8cb6d249770cb1634d5c07b8945f2349df7d873c8e22898a5bfd8123aa.NewGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Groupclassification the groupclassification property
func (m *KeyHubClient) Groupclassification()(*i3786de14b5891e0b6237ed5a04286048bb7ccab8227911414f7bb8f1e165a7c0.GroupclassificationRequestBuilder) {
    return i3786de14b5891e0b6237ed5a04286048bb7ccab8227911414f7bb8f1e165a7c0.NewGroupclassificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Groupclient the groupclient property
func (m *KeyHubClient) Groupclient()(*i4f6f8a97ade4bcd9f00e48fb22660ba2cb3b474e860f677cfc37ab0f9b6d8370.GroupclientRequestBuilder) {
    return i4f6f8a97ade4bcd9f00e48fb22660ba2cb3b474e860f677cfc37ab0f9b6d8370.NewGroupclientRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Groupfolder the groupfolder property
func (m *KeyHubClient) Groupfolder()(*i43af1840a4616958e60c4cbee17a55c9c7345f03d798c4b93f71ba6ea8a8fed7.GroupfolderRequestBuilder) {
    return i43af1840a4616958e60c4cbee17a55c9c7345f03d798c4b93f71ba6ea8a8fed7.NewGroupfolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Info the info property
func (m *KeyHubClient) Info()(*ia9c84b329cdbc7e9961b248161459b4dfb0d53b98e1b471cd7c10b6d027a90d3.InfoRequestBuilder) {
    return ia9c84b329cdbc7e9961b248161459b4dfb0d53b98e1b471cd7c10b6d027a90d3.NewInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Launchpadtile the launchpadtile property
func (m *KeyHubClient) Launchpadtile()(*i385775f507ab02e04cf3e40e8696ee406fb1371e877539b6dfbde0c8e4b123b2.LaunchpadtileRequestBuilder) {
    return i385775f507ab02e04cf3e40e8696ee406fb1371e877539b6dfbde0c8e4b123b2.NewLaunchpadtileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Numberseq the numberseq property
func (m *KeyHubClient) Numberseq()(*i98ee0f07cd6158c2879ae08ef6a731b5a05b3610b774bf729ac7c7bf7d1efa2c.NumberseqRequestBuilder) {
    return i98ee0f07cd6158c2879ae08ef6a731b5a05b3610b774bf729ac7c7bf7d1efa2c.NewNumberseqRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizationalunit the organizationalunit property
func (m *KeyHubClient) Organizationalunit()(*ifb310ad74101a00ed1acdc17b384ce2e8964f5a38f9d9302b22aaaa28ff4d3e4.OrganizationalunitRequestBuilder) {
    return ifb310ad74101a00ed1acdc17b384ce2e8964f5a38f9d9302b22aaaa28ff4d3e4.NewOrganizationalunitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Provisioninggroup the provisioninggroup property
func (m *KeyHubClient) Provisioninggroup()(*i29b864faf50a3a1afdf661366bb4234f149662638d0e3841e35acccd9352cf20.ProvisioninggroupRequestBuilder) {
    return i29b864faf50a3a1afdf661366bb4234f149662638d0e3841e35acccd9352cf20.NewProvisioninggroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Request the request property
func (m *KeyHubClient) Request()(*i4e9df38273eb83f44b000e660fb1ceeb2fdcb7939eb1e594bf0e1c6dd130c8e2.RequestRequestBuilder) {
    return i4e9df38273eb83f44b000e660fb1ceeb2fdcb7939eb1e594bf0e1c6dd130c8e2.NewRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Serviceaccount the serviceaccount property
func (m *KeyHubClient) Serviceaccount()(*i25381d874a7e3097702763676eb8a598478a75d07191e3563cf4ef48e5ac1fc6.ServiceaccountRequestBuilder) {
    return i25381d874a7e3097702763676eb8a598478a75d07191e3563cf4ef48e5ac1fc6.NewServiceaccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// System the system property
func (m *KeyHubClient) System()(*ibdde25a8196c6c3406e5b0b02de36d833849b28ee37dca06694dc2ec8a725bdc.SystemRequestBuilder) {
    return ibdde25a8196c6c3406e5b0b02de36d833849b28ee37dca06694dc2ec8a725bdc.NewSystemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Vaultrecord the vaultrecord property
func (m *KeyHubClient) Vaultrecord()(*i9c60bc0bc864eda4e8fcd5447d4e572d9b6ba65365f60dea8e1097309b0531da.VaultrecordRequestBuilder) {
    return i9c60bc0bc864eda4e8fcd5447d4e572d9b6ba65365f60dea8e1097309b0531da.NewVaultrecordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Webhook the webhook property
func (m *KeyHubClient) Webhook()(*i7f1810e47506a1087d42711bae10d2eacb3eb432b8e9e68ef97ff738e3bb6945.WebhookRequestBuilder) {
    return i7f1810e47506a1087d42711bae10d2eacb3eb432b8e9e68ef97ff738e3bb6945.NewWebhookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
