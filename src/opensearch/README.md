# 概要
OpenSearchを試すサンプル

# 操作

## 起動
docker compose up

## opensearchコンテナの起動確認
curl -XGET https://localhost:9200 -u 'admin:admin' --insecure

## ダッシュボード
http://localhost:5601/app/dashboards

## コンソール画面
http://localhost:5601/app/dev_tools#/console

## index一覧確認
`curl -XGET https://localhost:9200/_cat/indices -u 'admin:admin' --insecure`

# 参考
https://opensearch.org/docs/latest/
https://opensearch.org/docs/latest/clients/go/
https://sheltie-garage.xyz/tech/2023/04/go%E3%81%8B%E3%82%89opensearch%E3%82%92%E5%88%A9%E7%94%A8%E3%81%97%E3%81%A6%E3%81%BF%E3%82%8B/
https://dev.classmethod.jp/articles/how-to-build-opensearch-with-docker/
