module weather-server-go

go 1.15

require (
	github.com/lib/pq v1.8.0
	github.com/mizuki1412/go-core-kit v0.0.0-20200929142440-99ec274d4764
	github.com/spf13/cobra v1.1.1
	github.com/tidwall/gjson v1.6.1
	linkortech/framework/lkt-service-core v0.0.1
)

replace github.com/mizuki1412/go-core-kit v0.0.0-20200929142440-99ec274d4764 => ../mizuki/framework/core-kit

replace linkortech/framework/lkt-service-core v0.0.1 => ../linkortech/framework/lkt-service-core
