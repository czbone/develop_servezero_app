■ビルド準備
・jsディレクトリ直下にアプリケーションで使用するJavascriptファイルを置く。(app.js)
・scssディレクトリ直下にBootstrapデフォルトから更新するscssファイルを置く。

■ビルド処理
・scssファイルからassets/css/app.cssとbuild/assets/css/app.min.cssが生成される。
・Bootstrapで使用するJavascriptはすべてjs/vendorディレクトリにコピーされる。

