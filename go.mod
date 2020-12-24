module github.com/ffddorf/unms-exporter

go 1.15

replace github.com/go-swagger/go-swagger => github.com/mraerino/go-swagger v0.25.1-0.20201224025314-541cd837ee0e

require (
	github.com/go-openapi/errors v0.19.9
	github.com/go-openapi/runtime v0.19.24
	github.com/go-openapi/strfmt v0.19.11
	github.com/go-openapi/swag v0.19.12
	github.com/go-openapi/validate v0.20.0
	github.com/go-swagger/go-swagger v0.25.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/prometheus/client_golang v1.9.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	golang.org/x/sys v0.0.0-20201223074533-0d417f636930 // indirect
)
