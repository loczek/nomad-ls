## Overview

A simple [nomad](https://developer.hashicorp.com/nomad) language server

### Features

- Autocomplete
- Diagnostics
- Hover information
- Driver support (docker, exec, raw_exec, qemu, java)

### Building

```shell
$ make
```

### Editor Extensions

- [Zed](https://github.com/loczek/zed-nomad-extension) 
- VSCode (todo)

### Updating the language server while using it for the Zed Nomad Extension
If you are already using the [Zed Nomad Extension](https://github.com/loczek/zed-nomad-extension) and want your changes in this repo to be reflected, please
1. Run `go install` in this repo to install the updated binary to your `$GOPATH`.
2. Open Zed's command palette (`Cmd+Shift+P`) and run `workspace: reload`
