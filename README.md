# Displaying program versions with go modules

[go modules](https://github.com/golang/go/wiki/Modules) can be used to display
program versions instead of the usual linker trick that defines variable(s) at
build time:

```
go build -ldflags="-X main.version v0.1.1"
```

This is handy when you want `go get` to "just work" and derive the version from
a git tag. This is done with [ReadBuildInfo][rbi] from the `runtime/debug`
package.

```console
$ GO111MODULE=on go get github.com/dlespiau/go-main-module-version@v0.1.1
go: finding github.com/dlespiau/go-main-module-version v0.1.1
go: downloading github.com/dlespiau/go-main-module-version v0.1.1
go: extracting github.com/dlespiau/go-main-module-version v0.1.1

$ go-main-module-version
Version: v0.1.1
```

This works with semver tags starting with a `v`. `v0.1.1-rc4` will work, `foo`
will not.

[rbi]: https://tip.golang.org/pkg/runtime/debug/#BuildInfo
