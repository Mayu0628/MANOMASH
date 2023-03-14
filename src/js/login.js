//フォームに入力された値を取得する関数
const getLogInInfo = () => {
    const email = document.getElementById('email').value;
    const pass = document.getElementById('pass').value;

    console.log(email);
    console.log(pass);
    return {
        email: email,
        password: pass
    }
}

//Goのサーバーに取得した値を送って登録
const tryLogIn = async (obj) => {
    const body = JSON.stringify(obj)
    const url = "http://localhost:8080/login";
    const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json charset=utf-8'
    };
    const response = await fetch(url, {
      method: "POST",
      cache: "no-cache",
      headers: headers,
      body: body,
    });
    console.log(response)
    const logInInfo = await response.json();
    console.log(logInInfo);
    console.log("status_flg",logInInfo.status)
    return logInInfo
};

const displayLogInResult = async () => {
    obj = getLogInInfo()
    const alert = document.getElementById('alert')
    if (obj.email === '' | obj.password === '') {
        alert.innerHTML = `<p class ="alert">未入力の項目があります。</p>`
        return;
    } 

    logInInfo = await tryLogIn(obj)
    console.log(logInInfo.status)

    if (logInInfo.status === 0) {
       alert.innerHTML = `<p class ="mail-alert">メールアドレスかパスワードが間違っています。</p>`
        return ;
    } else if (logInInfo.status === 1) {
        document.cookie = "id=" + logInInfo.id
        window.location.href = '../src/mypage.html';
        return ;
    }
}