## アカウント管理アプリ

### 環境
- docker
#### 構築手順
- 環境変数ファイルの作成
- ```account-management/root/config```に```development.yml```を作成する

ファイル内容
```
server:
    port: ":3000" // 利用するdockerコンテナ内のポート番号
```

構築コマンド
```
docker-compose up -d --build
```
