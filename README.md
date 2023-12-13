# go-gitee
Gitee Go SDK

## Contents
- [Installation](#Installation)
- [Quick start](#quick-start)
- [ToDo](#todo)


## Installation
```shell
go get -u github.com/zdz1715/go-gitee@latest
```

## Quick start
### OAuth授权码模式
```go
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/zdz1715/ghttp"
	"github.com/zdz1715/go-gitee"
)

func main() {
	// OAuth授权码模式
	// docs: https://gitee.com/api/v5/oauth_doc#/list-item-2
	clientID := "YourClientID"
	clientSecret := "YourClientSecret"
	redirectURI := "http://127.0.0.1"
	credential := &gitee.OAuthCredential{
		// default endpoint: https://gitee.com
		//Endpoint:    gitee.CloudEndpoint,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
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
	// 先生成页面,获取code
	authURL := client.OAuth.GenerateAuthorizeURL(clientID, redirectURI, "user_info projects pull_requests issues notes keys hook groups gists enterprises")
	fmt.Printf("click url: %s", authURL)

	// 监听输出，手动输入获取的code
	codeChan := make(chan string, 1)
	go func() {
		buf := bufio.NewScanner(os.Stdin)
		fmt.Print("\ninput code: ")
		for buf.Scan() {
			codeChan <- buf.Text()
		}
	}()

	select {
	case code := <-codeChan:
		_ = os.Stdin.Close()
		// 通过code先手动获取一次token，获取之后在token有效期内，请求别的接口会自动带上token
		fmt.Printf("auth by code: %+v\n", code)
		tk, err := client.OAuth.GetAccessToken(context.Background(), &gitee.GetAccessTokenOptions{
			Code: code,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("auth token: %+v\n", tk)
		// 获取邮箱
		emails, err := client.Email.List(context.Background())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("emails: %+v\n", emails)
		// 若是想刷新token
		tk, err = client.OAuth.GetAccessToken(context.Background(), &gitee.GetAccessTokenOptions{
			RefreshToken: tk.RefreshToken,
		})
		fmt.Printf("RefreshToken: %+v\n", tk)
	}

}

```
### OAuth密码模式
```go
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
		// default endpoint: https://gitee.com
		//Endpoint:    gitee.CloudEndpoint,
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

```
### 直接设置token
```go
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


```

## ToDo
> [!NOTE]
> 现在提供的方法不多，会逐渐完善，也欢迎来贡献代码，只需要编写参数结构体、响应结构体就可以很快添加一个方法，参考下方代码。
```go
type Email struct {
	Email string   `json:"email"`
	State string   `json:"state"`
	Scope []string `json:"scope"`
}

func (s *EmailsService) List(ctx context.Context) ([]*Email, error) {
	const apiEndpoint = "/api/v5/emails"
	var respBody []*Email
	if err := s.client.InvokeByCredential(ctx, http.MethodGet, apiEndpoint, nil, &respBody); err != nil {
		return nil, err
	}
	return respBody, nil
}
```