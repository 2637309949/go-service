package redis

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"go-micro.dev/v5/cache"
)

func Test_newUniversalClient(t *testing.T) {
	type fields struct {
		options cache.Options
	}
	type wantValues struct {
		username string
		password string
		address  string
	}

	tests := []struct {
		name   string
		fields fields
		want   wantValues
	}{
		{name: "No Url", fields: fields{options: cache.Options{}},
			want: wantValues{
				username: "",
				password: "",
				address:  "127.0.0.1:6379",
			}},
		{name: "legacy Url", fields: fields{options: cache.Options{Address: "127.0.0.1:6379"}},
			want: wantValues{
				username: "",
				password: "",
				address:  "127.0.0.1:6379",
			}},
		{name: "New Url", fields: fields{options: cache.Options{Address: "redis://127.0.0.1:6379"}},
			want: wantValues{
				username: "",
				password: "",
				address:  "127.0.0.1:6379",
			}},
		{name: "Url with Pwd", fields: fields{options: cache.Options{Address: "redis://:password@redis:6379"}},
			want: wantValues{
				username: "",
				password: "password",
				address:  "redis:6379",
			}},
		{name: "Url with username and Pwd", fields: fields{
			options: cache.Options{Address: "redis://username:password@redis:6379"}},
			want: wantValues{
				username: "username",
				password: "password",
				address:  "redis:6379",
			}},

		{name: "Sentinel Failover client", fields: fields{
			options: cache.Options{
				Context: context.WithValue(
					context.TODO(), redisOptionsContextKey{},
					redis.UniversalOptions{MasterName: "master-name"}),
			}},
			want: wantValues{
				username: "",
				password: "",
				address:  "FailoverClient", // <- Placeholder set by NewFailoverClient
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			univClient := newUniversalClient(tt.fields.options)
			client, ok := univClient.(*redis.Client)
			if !ok {
				t.Errorf("newUniversalClient() expect a *redis.Client")
				return
			}
			if client.Options().Addr != tt.want.address {
				t.Errorf("newUniversalClient() Address = %v, want address %v", client.Options().Addr, tt.want.address)
			}
			if client.Options().Password != tt.want.password {
				t.Errorf("newUniversalClient() password = %v, want password %v", client.Options().Password, tt.want.password)
			}
			if client.Options().Username != tt.want.username {
				t.Errorf("newUniversalClient() username = %v, want username %v", client.Options().Username, tt.want.username)
			}
		})
	}
}

func Test_newUniversalClientCluster(t *testing.T) {
	type fields struct {
		options cache.Options
	}
	type wantValues struct {
		username string
		password string
		addrs    []string
	}

	tests := []struct {
		name   string
		fields fields
		want   wantValues
	}{
		{name: "Addrs in redis options", fields: fields{
			options: cache.Options{
				Address: "127.0.0.1:6379", // <- ignored
				Context: context.WithValue(
					context.TODO(), redisOptionsContextKey{},
					redis.UniversalOptions{Addrs: []string{"127.0.0.1:6381", "127.0.0.1:6382"}}),
			}},
			want: wantValues{
				username: "",
				password: "",
				addrs:    []string{"127.0.0.1:6381", "127.0.0.1:6382"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			univClient := newUniversalClient(tt.fields.options)
			client, ok := univClient.(*redis.ClusterClient)
			if !ok {
				t.Errorf("newUniversalClient() expect a *redis.ClusterClient")
				return
			}
			if !reflect.DeepEqual(client.Options().Addrs, tt.want.addrs) {
				t.Errorf("newUniversalClient() Addrs = %v, want addrs %v", client.Options().Addrs, tt.want.addrs)
			}
			if client.Options().Password != tt.want.password {
				t.Errorf("newUniversalClient() password = %v, want password %v", client.Options().Password, tt.want.password)
			}
			if client.Options().Username != tt.want.username {
				t.Errorf("newUniversalClient() username = %v, want username %v", client.Options().Username, tt.want.username)
			}
		})
	}
}
