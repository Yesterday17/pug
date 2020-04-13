module github.com/Yesterday17/pug

go 1.14

replace github.com/Yesterday17/pug/api => ./api

require (
	github.com/Yesterday17/bili-archive v1.2.1
	github.com/Yesterday17/pug/api v1.0.6
	github.com/mitchellh/mapstructure v1.1.2
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/tidwall/gjson v1.3.5
	gopkg.in/yaml.v2 v2.2.8
)
