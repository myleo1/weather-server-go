//go:generate pkger -include /swagger-ui
package main

import (
	_ "github.com/lib/pq"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"weather-server-go/cmd"
)

func main() {
	initkit.LoadConfig()
	cmd.Execute()
}
