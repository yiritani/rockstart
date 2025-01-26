# go * gRPCの素振り
暇だからChatGPTに考えてもらったシステムを作る

# 使用技術羅列
## app

### frontend
- Waku
### backend
- go
- gRPC
- sqlx
### infra
- AWS(予定)

# お題
# シンプルな書籍管理システム (Book/Author) を gRPC で実装する

- **エンティティ1: Author**
  - 名前・メールアドレス・自己紹介などを保持
  - CRUD（Create / Read / Update / Delete）を用意

- **エンティティ2: Book**
  - タイトル・出版日・著者(Author)ID・在庫数などを保持
  - CRUD（Create / Read / Update / Delete）を用意

## 相互作用の例
- **Book** は必ず **存在する Author の ID** を持たないと作成できない  
- **Author** を削除すると、関連する **Book** にどう影響するか？

上記2つのエンティティを使えば、以下のように **2セット分の CRUD** が作れます。

- **AuthorService**
  - CreateAuthor / GetAuthor / UpdateAuthor / DeleteAuthor

- **BookService**
  - CreateBook / GetBook / UpdateBook / DeleteBook

さらに、**Author** と **Book** が相互に参照し合うようなロジック(**外部キー制約**など)があるため、DBアクセスやトランザクション管理も試せる題材になります。