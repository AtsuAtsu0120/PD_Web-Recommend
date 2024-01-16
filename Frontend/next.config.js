// next.config.js

module.exports = {
  // ...

  // リバースプロキシの設定
  rewrites: [
    {
      // 元のパス
      source: "/api/:path",
      // 転送先のパス
      destination: "http://localhost:1323/:path",
    },
  ],
};