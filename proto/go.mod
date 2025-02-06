module proto

go 1.22.0

replace go-micro.dev/v5 => ../comm/go-micro

require (
	go-micro.dev/v5 v5.0.0-00010101000000-000000000000
	google.golang.org/genproto/googleapis/api v0.0.0-20250204164813-702378808489
	google.golang.org/protobuf v1.36.4
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
