package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zdz1715/ghttp"
	"github.com/zdz1715/go-gitee"
)

func main() {
	// OAuth密码模式
	// docs: https://gitee.com/api/v5/oauth_doc#/list-item-2
	credential := &gitee.PasswordCredential{
		Endpoint:     gitee.CloudEndpoint,
		ClientID:     "YourClientID",
		ClientSecret: "YourClientSecret",
		Username:     "YourUsername",
		Password:     "YourPassword",
		Scope:        "user_info projects emails pull_requests issues notes keys hook groups gists enterprises",
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
	// 无需手动获取token，执行下面方法会自动获取token，在有效期内不会重复请求获取token，当然你也可以手动获取token存起来
	// 获取邮箱
	emails, err := client.Email.List(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("emails: %+v\n", emails)

}
