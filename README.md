# prac-orm-transaction
## 動作方法(随時変更)
- Docker, Docker-compose, go の開発環境が整っていることを前提とする
1. DBを起動
```cassandraql
docker-compose up -d
```
2. アプリケーションを起動
```cassandraql
go run cmd/main.go
```

## 基本概要
#### リポジトリ概要
本リポジトリでは、情報記事管理サービスを想定し、普段あまり触れたことのないライブラリ（chi、orm）の使用と
DDD+CQRSを採用した開発になれることを目的としている．

#### 開発フロー
- git-flow

#### 使用AFW
- chi

#### ORM
- gorm

#### 使用liter(静的解析ツール)
- go vet  
go lint は非推奨のため未使用

#### ディレクトリ構成
```

```

#### 製作者
take-2405(Gitアカウント名)

### 参考情報
- chi 公式ドキュメント  
  https://github.com/go-chi/chi  
- go vet 公式ドキュメント  
  https://pkg.go.dev/cmd/vet