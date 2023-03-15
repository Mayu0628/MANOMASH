const getMypageInfo = async () => {
    const key = 'id';
    const value = document.cookie.match(
        new RegExp(key+'\=([^\;]*)\;*')
    )[1];
    console.log(value);
    const url = "http://localhost:8080/mypage?id=" + value;
    const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json charset=utf-8'
    }
    const response = await fetch(url, {
        method: "GET",
        cache: "no-cache",
        headers: headers,
    })
    console.log(response)
    const mypageInfo = await response.json();
    console.log("mypageInfo",mypageInfo)
    const user_name = document.getElementById('user_name')
    console.log(user_name)
    // if not null
    user_name.innerHTML = mypageInfo.UserName;

    const user_id = document.getElementById('user_id')
    console.log(user_id)
    user_id.innerHTML = mypageInfo.UserID;

    const user_comment = document.getElementById('user_comment')
    console.log(user_comment)
    user_comment.innerHTML = mypageInfo.Introduce;
}
window.onload = getMypageInfo()
