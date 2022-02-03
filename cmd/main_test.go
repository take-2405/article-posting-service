package main

import (
	"fmt"
	"github.com/pkg/profile"
	"io/ioutil"
	"net/http"
	"os"
	"prac-orm-transaction/config"
	"prac-orm-transaction/di"
	router2 "prac-orm-transaction/presentation/router"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	defer profile.Start(profile.CPUProfile).Stop()

	setUp()
	status := m.Run()
	os.Exit(status)
}

// 前処理用の関数です。関数名は他の名前でも問題ありません。
func setUp() {
	//ルーターを初期化
	router := router2.NewServer()
	//ルーティングとDI
	di.InsertUserDI(router)
	port := config.GetServerPort()
	//ルーター起動
	go http.ListenAndServe(port, router.Router)
}

func TestCreateUser(t *testing.T) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// bodyにJSONを設定したリクエストの生成
	json := strings.NewReader(`{"id":"user","password":"pass"}`)
	req, _ := http.NewRequest("POST", "http://localhost:8080/sign/up", json)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
