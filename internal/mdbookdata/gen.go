//go:build ignore

package main

import (
	"fmt"

	"github.com/jcbhmr/go-mdbook/internal/version"
)

// 1. Download all the release archives (.tar.gz & .zip).
// 2. Extract the SEA binaries from the archives.
// 3. Save the SEA binaries to "$bin.$GOOS-$GOARCH(.exe)".
func main() {
	artifacts := map[string]string{
		"linux/arm64": fmt.Sprintf("https://github.com/rust-lang/mdBook/releases/download/v%[1]s/mdbook-v%[1]s-aarch64-unknown-linux-musl.tar.gz", version.Version),
		"darwin/amd64": fmt.Sprintf("https://github.com/rust-lang/mdBook/releases/download/v%[1]s/mdbook-v%[1]s-x86_64-apple-darwin.tar.gz", version.Version),
		"windows/amd64": fmt.Sprintf("https://github.com/rust-lang/mdBook/releases/download/v%[1]s/mdbook-v%[1]s-x86_64-pc-windows-msvc.zip", version.Version),
		"linux/amd64": fmt.Sprintf("https://github.com/rust-lang/mdBook/releases/download/v%[1]s/mdbook-v%[1]s-x86_64-unknown-linux-musl.tar.gz", version.Version),
	}
}
