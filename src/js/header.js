const cookies = document.cookie;
console.log("all-cookie", cookies);
const cookielist = cookies.split(';');

const getID = cookielist.filter((value)=>{
    const content = value.split('=');
    return content[0]==='id' && content[1]!=='';
})

const login_logout = document.getElementById('login_logout')
if (getID.length !== 0) {
    login_logout.innerHTML = `
        <a id="login_logout" class=login-link href="login.html" onclick="logOut()">ログアウト</a>
    `
} else {
    login_logout.innerHTML = `
        <a id="login_logout" class=login-link href="login.html">ログイン</a>
    `
}

const logOut = () => {
    document.cookie = "id="
    login_logout.innerHTML = `
        <a id="login_logout" class=login-link href="login.html">ログイン</a>
    `
}
