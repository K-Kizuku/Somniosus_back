# オンデンカップ バックエンド

## getting started

### ツールのインストール

`make setup-tools`

使用ツール一覧

- git-cz
- scaffdog
- wire
- gomock

### 各種コマンドの説明

```sh
make c-service DIR=<サービス名>
```

マイクロサービスの雛形の作成  
`apps/<サービス名>`に作成される  

```sh
make commit
```

gitのコミットを行う(git-czを使用)
`{type}{scope}: {emoji}{subject}`のフォーマットでコミットメッセージが作成される

```sh
make adr
```

ADRのテンプレートファイルを`docs/adr`に作成する  
<details>
<summary>ADRとは</summary>

Architectural Decision Recordの略  
なぜその技術選定を行なったのかの**根拠**を残すためのもの  
プロジェクトの意思決定を記録として残せるだけでなく，開発者間のナレッジの共有にも繋がる
</details>

## ディレクトリ構成

### 全体像

```sh
.
├── .github
│   ├── labeler.yml # PRにタグを自動で付与する設定
│   └── workflows
│       └── label.yml # PRにタグを自動で付与するworkflows
├── .gitignore
├── .scaffdog
│   ├── ADR.md # ADRのテンプレート
│   └── config.js # scaffdogの設定
├── Makefile
├── README.md
├── apps # マイクロサービス群
│   └── temp
│       ├── app
│       ├── cmd
│       ├── config
│       └── test
├── changelog.config.js # git-czの設定
├── docs # 各種ドキュメント
│   └── adr # ADRの格納
├── go.mod
├── lib # 汎用関数群
├── protobuf # protocol buffers
└── scripts # 各種スクリプト
    └── create_service.sh # マイクロサービスのテンプレート作成のスクリプト
```

### マイクロサービス内部

マイクロサービス内部はクリーンアーキテクチャを意識しています

```sh
.
├── app
│   ├── domain # ドメインロジック
│   │   ├── model # ドメインモデル
│   │   ├── repository # repositoryのインターフェイス
│   │   └── service # ドメインサービス
│   ├── infra
│   │   ├── dao # repositoryの実装
│   │   │   └── query # readの処理
│   │   ├── db
│   │   │   ├── ddl # テーブル定義
│   │   │   ├── migrations # マイグレーション
│   │   │   └── seed # シード
│   │   └── dto # structの変換
│   ├── ui
│   │   └── grpc # grpcのアダプター
│   └── usecase
│       └── query # getの処理
├── cmd
├── config # 各種設定
└── test
    ├── mock_domain # ドメインのモック
    │   └── mockgen.go
    ├── mock_infra # インフラのモック
    │   └── mockgen.go
    └── mock_usecase # ユースケースのモック
        └── mockgen.go
```
