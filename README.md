## アカウント管理アプリ
一連のユーザー管理機能を提供するアプリケーションです。

詳しくはドキュメントを参照。

### 大まかな機能一覧
- ユーザー登録機能
- パスワード変更機能
- トークン発行機能
- トークン検証機能

### 環境
- docker
#### 構築手順
- 環境変数ファイルの作成
- ```account-management/root/config```に```development.yml```を作成する

ファイル内容
```
server:
    port: ":3000" // 利用するdockerコンテナ内のポート番号
database:
    dns: postgres://postgres:postgres@account-management-postgres:5432/account_management?sslmode=disable
jwt:
  secret: "H&6^fMgXmQqYb4StLLHY3iCY%dBei^FTD&BJTkBF&2YDZ#@GvVuT%5RbtGanWcoqEAFyrAHapjJkW&Tn&fj2iuxRBtzL5fYL$Kn5gVcPJLsfM6XdjhQX@UMB$qqq8YRH"
  expire_second: 2592000
```

構築コマンド
```
docker-compose up -d --build
```

### テスト実行
```
go test -v test/...
```
