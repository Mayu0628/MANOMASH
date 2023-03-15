//フォームに入力された値を取得する関数
const user_name = document.getElementById('edit-user-name');
const introduce = document.getElementById('edit-user-intro');
const getLogInInfo = () => {
    return {
        user_name: user_name.value,
        introduce: introduce.value
    }
}

const getUserInfo = async () => {
    const key = 'id';
    const value = document.cookie.match(
        new RegExp(key+'\=([^\;]*)\;*')
    )[1];
    console.log(value);
    const url = "http://localhost:8080/mypage?id=" + value;
    const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json charset=utf-8'
    };
    const response = await fetch(url, {
      method: "GET",
      cache: "no-cache",
      headers: headers,
    });
    console.log(response)
    const userInfo = await response.json();
    console.log(userInfo);
    console.log("status_flg",userInfo.status)
    user_name.innerText = userInfo.UserName
    introduce.innerText = userInfo.Introduce
};

const updateUserInfo = async () => {
    const obj = getLogInInfo()
    const key = 'id';
    const value = document.cookie.match(
        new RegExp(key+'\=([^\;]*)\;*')
    )[1];
    console.log(value);
    const body = await JSON.stringify({
        user_id: Number(value),
        user_name: obj.user_name,
        introduce: obj.introduce
    })
    console.log(body)
    const url = "http://localhost:8080/mypage/edit"
    const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json charset=utf-8'
    };
    const response = await fetch(url, {
      method: "POST",
      cache: "no-cache",
      headers: headers,
      body: body
    });
    console.log(response)
    const userInfo = await response.json();
    console.log(userInfo);
};

onload.window = getUserInfo()
