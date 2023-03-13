自作関数の説明

handler/ResFunc.go:
ResFlgCreate(ReSt int,ReRe string) (model.ResFlg){
    基本的にハンドラの処理の結果をレスポンスするのに使用する。
    使い方は成功か失敗が決定したところに
    ResFlg := ResFlgCreate(1,"succesful")//成功なら１，失敗なら２
    ResFlgをレスポンスにエンコードする
    handler/signup_hundler.goの42、44行目とかを参照
}

handler/mypage-edit_handler.go
Checkvariable(UsDt UpdataUser) (err string){
    限定的にマイページ編集画面での編集で中身が空かどうかを調べるのに使う
    なのでUpdataUserとかいうよくわからないのがある。
    UpdataUserは4つの引数の構造体
    それを設定して、変数で入れたら使える
    文字か数字かを判定できるが、限定的に使う際に文字しか使わないのでその機能は消した
}

