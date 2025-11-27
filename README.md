# GO Builder(GOBD)

- A Golang build tool that helps you easily build files for multiple operating systems and architectures

## Usage

```
# build local runtime pairs without debug and cgo
gobd -no_debug -no_cgo

# build linux-amd64 pair
gobd -os linux -arch amd64 -no_debug -no_cgo

# build linux-mipsle pair with env 'GOMIPS=softfloat'
gobd -os linux -arch mipsle -no_debug -no_cgo -envs 'GOMIPS=softfloat'

# build all os-arch pairs and show the build process
gobd -os linux -arch amd64 -d bin -no_debug -no_cgo -opts '-v'

# build all supported architectures for windows, macos, linux and freebsd
gobd -main -d bin -no_debug -no_cgo
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/gobd https://github.com/unix755/gobd/releases/latest/download/gobd-linux-amd64
chmod +x /usr/local/bin/gobd

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/gobd https://github.com/unix755/gobd/releases/latest/download/gobd-freebsd-amd64
chmod +x /usr/local/bin/gobd
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/unix755/gobd.git
cd gobd
export CGO_ENABLED=0 
go build -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
