# システムでの想定リクエストとレスポンス
本ドキュメントはSwaggerの代替として使用している。

以下に想定される機能を記載する(面倒なので、部分的に記載する)
また、あくまでも目的はライブラリと設計に慣れることであるため、不足する要素などがあることは前提としている。

### アカウント作成(ユーザー登録)
##### リクエスト
method:POST  

**ボディ**
```cassandraql
{
    "name": "string" ,
    "password": "string",
    "age":18,
    "email":"string"
}
```

##### レスポンス
```cassandraql
{
"token": "string" 
}
```

### ログイン
##### リクエスト
method:POST  

**ボディ**
```cassandraql
{
    "name": "string" ,
    "password": "string" 
}
```

##### レスポンス
```cassandraql
{
"token": "string" 
}
```

### ユーザー情報変更
##### リクエスト
method:PUT  
**ヘッダー**  
token:string  

**ボディ**
```cassandraql
{
    "name": "string" ,
    "password": "string",
    "age":18,
    "email":"string"
}
```
変更したいものだけをリクエストに組み込む

##### レスポンス
```cassandraql
{
"status":200 
}
```

### 記事の投稿機能
##### リクエスト
**ヘッダー**  
token:string  
method:POST  

**ボディ**
```cassandraql
{
  "title": "string",
  "description": "string",
  "images": [
    {
      "image": "string"
    }
  ],
  "tags": [
    {
      "tag": "string"
    }
  ]
}
```
##### レスポンス
```cassandraql
{
"articleID":"string" 
}
```

### 記事の修正機能
##### リクエスト
**ヘッダー**  
token:string  
method:PUT  
**ボディ**
```cassandraql
{
  "title": "string",
  "description": "string",
  "images": [
    {
      "image": "string"
    }
  ],
  "tags": [
    {
      "tag": "string"
    }
  ]
}
```
##### レスポンス
```cassandraql
{
"articleID":"string" 
}
```
### 記事の検索機能(複数の記事を提供)

##### リクエスト
**ヘッダー**  
token:string  
method:GET  
PathParam:articlesNumber

##### レスポンス
```cassandraql
{
  "articles": [
    {
      "articleID": "string",
      "title": "string",
      "image": "string",
      "description": "string"
    }
  ]
}
```
### 記事閲覧機能(１つの記事の詳細な内容)
##### リクエスト
**ヘッダー**  
token:string  
method:GET  
PathParam:articleID

##### レスポンス
```cassandraql
{
  "title": "string",
  "description": "string",
  "images": [
    {
      "image": "string"
    }
  ],
  "tags": [
    {
      "tag": "string"
    }
  ],
  "nice":0,
  "comments": [
    {
      "comment": "string",
      "name": "string"
    }
  ]
}
```

### ユーザーのプロフィール表示機能
### 記事に対する「高評価」機能
##### リクエスト
**ヘッダー**  
token:string  
method:POST  
PathParam:articleID
##### レスポンス
```cassandraql
{
 "nice"0
}
```
### 記事に対するコメント機能
##### リクエスト
**ヘッダー**  
token:string  
method:POST  
PathParam:articleID  
**ボディ**
```cassandraql
{
  "content": "string"
  }
```
##### レスポンス
```cassandraql
{
 "success":200
}
```