//フォームに入力された値を取得する関数
const getSignUpInfo = () => {
    const user_name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const pass = document.getElementById('pass').value;
    const confirmation = document.getElementById('confirmation').value;

    console.log(user_name);
    console.log(email);
    console.log(pass);
    console.log(confirmation);

    return {
        name: user_name,
        email: email,
        password: pass,
        confirmation: confirmation
    }
}

//Goのサーバーに取得した値を送って登録
const PostSignUp = async (obj) => {
    const body = JSON.stringify(obj)
    const url = "http://localhost:8080/sign-up";
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
    const postFlg = await response.json();
    console.log(postFlg);
    console.log("status_flg",postFlg.status_flg)
    return postFlg
};

const displaySignUpStatus = async () => {
    obj = getSignUpInfo()
    const alert = document.getElementById('alert')
    if (obj.name === '' | obj.email === '' | obj.pass === '') {
        alert.innerHTML = `<p class="alert">未入力の項目があります。</p>`
        return false;
    } else if(obj.pass != obj.confirmation){
        alert.innerHTML = `<p class="confirmation">パスワードが一致しません</p>`
        return false;
    }

    postFlg = await PostSignUp(obj)
    console.log(postFlg.status_flg)

    if (postFlg.status_flg === 0) {
       alert.innerHTML = `<p class ="mail-alert">すでに使われているメールアドレスです。</p>`
        return false;
    } else if (postFlg.status_flg === 1) {
        window.location.href = '../src/login.html';
        return true;
    }
}



