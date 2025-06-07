# FUTO マーチングダッシュボード

FUTO マーチングバンドのスケジュールとタスクを管理するためのダッシュボードアプリケーションです。

## 機能

- ユーザー認証（ログイン/ログアウト）
- ユーザー管理（作成、編集、削除）
- 役割管理（管理者と一般ユーザー）
- カレンダー機能
- タスク管理
- 練習メニュー管理
- 時間追跡

## 技術スタック

### バックエンド
- Go 1.24.2
- Echo Web Framework
- MongoDB

### フロントエンド
- Next.js
- Tailwind CSS
- TypeScript

## はじめに

### 前提条件
- Docker and Docker Compose
- Node.js 18+
- Go 1.24+
- MongoDB

### Docker Composeで実行

```bash
# リポジトリをクローン
git clone https://github.com/kynmh69/futo-marching-dashboad.git
cd futo-marching-dashboad

# アプリケーションを開始
docker-compose up -d
```

アプリケーションは以下のURLでアクセスできます：
- フロントエンド: http://localhost:3000
- バックエンドAPI: http://localhost:8080

### 開発環境セットアップ

#### バックエンド
```bash
cd backend
cp .env.example .env  # 環境変数を設定
go mod download
go run cmd/server/main.go
```

#### フロントエンド
```bash
cd frontend
npm install
npm run dev
```

## テスト

このリポジトリには、バックエンドとフロントエンドの両方のコンポーネントのユニットテストが含まれています。

### バックエンドテスト
```bash
cd backend
go test -v ./...
```

### フロントエンドテスト
```bash
cd frontend
npm test
```

## GitHub Actions

このプロジェクトはCI/CDにGitHub Actionsを使用しています。プルリクエストとメインブランチへのプッシュで自動的にテストが実行されます。