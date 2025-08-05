# UnIaC Mappings Examples

This directory contains sample mappings to guide contributors in adding support for new cloud providers.

## AWS Sample
- File: `example-mapper.go`
- Description: A sample mapper that generates Terraform HCL based on the `World` model.
- Usage: Copy this file, modify it for your provider, and build as a plugin.

## Contributing
1. Fork the repo.
2. Create a `*.so` plugin under `~/.uniac/plugins/`.
3. Implement the `mappings.Mapper` interface.
4. Compile with  `go build -buildmode=plugin -o output.so plugin-file.go`
5. Share your plugin.