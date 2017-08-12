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

# データ登録手順
* indexの作成 
curl -XPUT http://localhost:9200/employee_info
* データ定義（types）の登録 
cd import_data 
curl -XPUT http://localhost:9200/employee_info/info/_mapping -d @init_info.json 
* データの登録
for f in `find . -type f`
do
  curl -XPOST "http://localhost:9200/employee_info/info/_bulk" --data-binary @$f
done 
*確認方法
** index確認
curl -XGET localhost:9200/_aliases?pretty 
※"employee_info"が存在すること 
** types確認
curl -XGET http://localhost:9200/employee_info/_mapping/info?pretty=true
※"init_info.json"の内容が登録されていること
** 登録データ確認（10件） 
curl -XGET http://localhost:9200/employee_info/info/_search 
※データが登録されていること（全件検索はウェブ参照）
