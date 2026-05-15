<img src="banner.png" alt="Nomad Language Server">

A simple language server for [nomad](https://developer.hashicorp.com/nomad).

### Features

- Autocomplete
- Diagnostics
- Formatting
- Hover information
- Driver support (docker, exec, raw_exec, qemu, java)

### Building

```shell
$ make
```

### Editor Extensions

- [Zed](https://github.com/loczek/zed-nomad-extension)
- [VSCode](https://github.com/loczek/vscode-nomad-extension)

### Development setup with Zed

If you are already using the [Zed Nomad Extension](https://github.com/loczek/zed-nomad-extension) and want your changes in this repo to be reflected, please

1. Run `go install` in this repo to install the updated binary to your `$GOPATH`.
1. Open Zed's command palette (`Cmd+Shift+P`) and run `editor: restart language server`
