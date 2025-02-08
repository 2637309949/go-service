package registry

import (
	"fmt"
	"path"
	"strings"
)

// Translates /foo/bar/zool into api service micro.api.foo method Bar.Zool
// Translates /foo/bar into api service micro.api.foo method Foo.Bar
func apiRoute(base string, p string) (string, string) {
	p = path.Clean(p)
	p = strings.TrimPrefix(p, base)
	p = strings.TrimPrefix(p, "/")
	parts := strings.SplitN(p, "/", 2)
	if len(parts) == 1 {
		return parts[0], "/call"
	} else {
		return parts[0], fmt.Sprintf("/%+s", parts[1])
	}
}
