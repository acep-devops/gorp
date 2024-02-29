# Go Reverse Proxy (gorp)

Dead simple http reverse proxy

## Build using

### Building for default OS

```bash
mkdir build
go build -o build/gorp
```

### Building for Windows
```bash
mkdir build
env GOOS=windows GOARCH=amd64 go build -o build/gorp.exe
```