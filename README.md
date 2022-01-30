# prac-orm-transaction
## 動作方法(随時変更)
- Docker, Docker-compose, go の開発環境が整っていることを前提とする

**makeコマンドが使用できる場合**  
1.1. DBを起動
```cassandraql
make compose-up
```
2. アプリケーションを起動
```cassandraql
make run
```

**makeコマンドが使用できない場合**  
1DBを起動
```cassandraql
docker-compose up -d
```
2. アプリケーションを起動
```cassandraql
go run cmd/main.go
```

## 基本概要
#### リポジトリ概要
本リポジトリでは、Qiitaのような記事管理サービスを想定し、あまり触れたことのないライブラリ（chi、orm）の使用と
**DDD+CQRS**を採用した開発になれることを目的としている．

#### 基本仕様
- 各エンドポイントのリクエスト&レスポンスは「./docsのsystem-request-response.md」に記載
- 現時点では、「sign/up」「sign/in」「article/create」を実装ずみ

#### 開発フロー
- git-flow

#### 認証
- APIキー認証

#### 使用AFW
- chi

#### ORM
- gorm

#### 使用liter(静的解析ツール)
- go vet  
go lint は非推奨のため未使用

#### ディレクトリ構成
```
./
├── presentation
│   ├── controller
│   ├── middleware
│   ├── request
│   ├── response
│   └── router
├── usecase
├── cmd
├── config
├── di
├── domain
└── infrastructure
```

#### 製作者
take-2405(Gitアカウント名)

### 参考情報
- chi 公式ドキュメント  
  https://github.com/go-chi/chi  
- go vet 公式ドキュメント  
  https://pkg.go.dev/cmd/vet