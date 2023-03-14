const getMypageInfo = async () => {
    document.cookie = "Name=test@test.com";
    await console.log("cookie", document.cookie);
    const url = "http://localhost:8080/getcookie"
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
    
}

//window.onload = getMypageInfo()
const obj = {
    id: 1,
    username: 'momii',
    introduce: 'よろしくおねがいします！！'
}



const user_name = document.getElementById('user_name')
console.log(user_name)
// if not null
if (user_name) user_name.innerHTML = obj.username;

const user_id = document.getElementById('user_id')
console.log(user_id)
if (user_id) user_id.innerHTML = obj.id;

const user_comment = document.getElementById('user_comment')
console.log(user_comment)
if (user_comment) user_comment.innerHTML = obj.introduce;
