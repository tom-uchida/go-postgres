# go-postgres

## Tech stack

|Contents|Tech|
|:-:|:-:|
|Language|Go|
|DB|PostgreSQL|
|Migration|[Atlas](https://github.com/ariga/atlas)|
|ORM|[SQLBoiler](https://github.com/volatiletech/sqlboiler)|

## How to start

```bash
% source .env
% make up
% make atlas-hash
% make atlas-schema-init
% make sqlboiler
% make run
% make down
```

## Result

```
make run
go run ./
2024/03/10 14:20:59 Book: &{BookID:1 Title:Sample Book 1 AuthorID:1 PublisherID:1 Isbn:{String:1234567890 Valid:true} YearPublished:{Int:2021 Valid:true} R:<nil> L:{}}
2024/03/10 14:20:59 Book: &{BookID:2 Title:Sample Book 2 AuthorID:2 PublisherID:2 Isbn:{String:0987654321 Valid:true} YearPublished:{Int:2020 Valid:true} R:<nil> L:{}}
2024/03/10 14:20:59 Book: &{BookID:3 Title:Sample Book 3 AuthorID:3 PublisherID:3 Isbn:{String:1122334455 Valid:true} YearPublished:{Int:2022 Valid:true} R:<nil> L:{}}
Select: 高度なクエリでの書籍の取得
Count: 書籍の数を数える
Count: 0
Exists: 特定の条件に一致する書籍が存在するかを確認
Exists: false
Insert: 書籍の挿入
Update: 書籍の更新
Upsert: 書籍のアップサート
Delete: 書籍の削除
Reload: 書籍の再読み込み
Reload: 書籍が見つかりませんでした
```
