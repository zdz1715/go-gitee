package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zdz1715/ghttp"
	"github.com/zdz1715/go-gitee"
)

func main() {
	// 直接设置token
	credential := &gitee.TokenCredential{
		// default endpoint: https://gitee.com
		//Endpoint:    gitee.CloudEndpoint,
		AccessToken: "token",
	}

	client, err := gitee.NewClient(credential, &gitee.Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 获取邮箱
	emails, err := client.Email.List(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("emails: %+v\n", emails)

}
