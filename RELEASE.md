# Updating and releasing Topicus KeyHub SDK for Go

## 1. Updating

### 1.1 Dependencies

```Shell
go get -u
go mod tidy
```

### 1.2 KeyHub OpenAPI spec

Use the keyhub-openapi-transformer-cli to download and preprocess the openapi spec from a running KeyHub instance

```Shell
transform --in https://<KEYHUB_HOSTNAME>/keyhub/rest/v1/openapi.json --out /path/to/sdk-go --target go
```

Regenerate the SDK

```Shell
go generate .
```

### 1.3 Commit the results

```Shell
git add .
git commit -m "Upgrade to KeyHub 40"
git push
```

## 2. Tag and release

```Shell
git tag v0.40.0
git push origin v0.40.0
```
