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
		Endpoint:     gitee.CloudEndpoint,
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
