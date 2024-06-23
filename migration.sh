#!/bin/bash

# リトライ回数の設定
MAX_RETRIES=3
CURRENT_RETRIES=0

while [ ${CURRENT_RETRIES} -lt ${MAX_RETRIES} ]; do
    # Dockerコンテナをバックグラウンドで起動する
    docker-compose run --rm migrate_app -path=/api/migrations/ -database "postgres://postgres:postgres@db:5432/pickup_analytics?sslmode=disable" up
    TABLE_EXISTS=$(docker-compose exec -T db psql -U postgres -d pickup_analytics -c "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'reports' AND table_name = 'reports');")

    # 結果を表示
    if echo "${TABLE_EXISTS}" | grep -q "t"; then
        echo "テーブルが存在します。"
        break
    else
        echo "テーブルが存在しません。"
         ((CURRENT_RETRIES++))
        sleep 5  # エラーが続いている場合、適宜調整する
    fi
done

# リトライ回数を超えた場合の処理
if [ ${CURRENT_RETRIES} -eq ${MAX_RETRIES} ]; then
    echo "リトライ回数を超えました。手動で対処してください。"
fi
