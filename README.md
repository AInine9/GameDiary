# golang-react-mysql-docker

## Stats

- Go
- React
- MySQL

## Tools

- SwaggerUI
- air (Hot reload for backend)

## Setup

Edit `variables.env`
`docker-compose up -d`

## Backend Architecture

- config  
  config management
- domain/model  
  ドメインモデル
- domain/repository  
  ドメインの振る舞い
- infrastructure/persistence  
  DB接続と処理
- interface/handler  
  HTTPリクエスト・レスポンスを扱う
- usecase  
  ユースケースに沿った処理
