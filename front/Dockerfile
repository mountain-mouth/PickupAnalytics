# ベースイメージを指定します。Node.js LTS (最新の安定版) を使用します。
FROM node:20-alpine

# コンテナ内での作業ディレクトリを指定します。
WORKDIR /front

# ホストの package.json と package-lock.json をコンテナの作業ディレクトリにコピーします。
COPY package.json .
COPY package-lock.json .

# npm パッケージをインストールします。
RUN npm install

# ホストのソースコードをコンテナの作業ディレクトリにコピーします。
COPY . .

# ポート 3000 を外部に公開します。
EXPOSE 3000

# アプリケーションを実行します。
CMD ["npm", "run", "dev"]
