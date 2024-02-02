# go-dingtalk

本项目为钉钉开放平台服务端接口的golang版本sdk

- 对接了部分接口

### Useage

```
package main

import (
	"context"
	"fmt"
	"github.com/yrzs/go-dingtalk/dingtalk"
	"log"
)

func main() {
	var config = &dingtalk.DingConfig{
		ClientId:     "xxx",
		ClientSecret: "xxx",
	}
	ding, err := dingtalk.NewClient(config)
	if err != nil {
		log.Fatalf("new dingtalk client err:%v", err)
	}
	ctx := context.Background()
	res, err := ding.ListDepartment(ctx, "667109893")
	fmt.Println(res, err)
}

```
