# revel_sample
吉田塾用レポジトリ

# 起動方法
docker-compose build
docker-compose up

# 起動確認
## revelフレームワーク
ブラウザから　http://localhost:9000/ に接続
→ウェブ画面が表示されること（初期は It Works!）

## elasticsearch
ブラウザから http://localhost:9200 に接続
→エラーとならないこと（json形式の文字が表示される）