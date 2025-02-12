package main

import (
	"context"
	"reflect"
	"strings"

	"go-micro.dev/v5/client"
	"go-micro.dev/v5/metadata"
	regi "go-micro.dev/v5/registry"
)

func modifyRsp(cf client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *regi.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
		err := cf(ctx, node, req, rsp, opts)
		if err != nil {
			return err
		}
		traceID := ""
		if md, ok := metadata.FromContext(ctx); ok {
			traceID = md["Uber-Trace-Id"]
			traceID = strings.Split(traceID, ":")[0]
		}
		val := reflect.ValueOf(rsp)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		if !val.CanInterface() {
			return err
		}

		v, ok := val.Interface().(map[string]interface{})
		if !ok {
			return err
		}

		if dataField, ok := v["data"]; ok {
			if _, ok := dataField.([]interface{}); ok {
				v["code"] = 200
				if len(traceID) > 0 {
					v["request_id"] = traceID
				}
			}
		} else {
			newRsp := make(map[string]interface{})
			newRsp["data"] = v
			newRsp["code"] = 200
			if len(traceID) > 0 {
				newRsp["request_id"] = traceID
			}
			if val.CanSet() {
				val.Set(reflect.ValueOf(newRsp))
			}
		}

		return err
	}
}
