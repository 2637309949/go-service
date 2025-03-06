package util

import (
	"context"
	"net/http"
	"net/textproto"
	"strings"

	"go-micro.dev/v5/metadata"
)

func FromRequest(r *http.Request) context.Context {
	ctx := r.Context()
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata)
	}
	for k, v := range r.Header {
		md[textproto.CanonicalMIMEHeaderKey(k)] = strings.Join(v, ",")
	}
	// pass http host
	md["Host"] = r.Host
	// pass http method
	md["Method"] = r.Method
	if r.URL != nil {
		md["URL"] = r.URL.String()
	}
	return metadata.NewContext(ctx, md)
}
