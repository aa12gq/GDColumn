package bootstrap

import (
	"GDColumn/pkg/snowflake"
	"fmt"
)

func SetupSnowflake() {

	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
}