const cookies = document.cookie;
console.log("all-cookie", cookies);
const cookielist = cookies.split(';');

const getID = cookielist.filter((value)=>{
    const content = value.split('=');
    return content[0]==='id' && content[1]!=='';
})

const login_logout = document.getElementById('login_logout')
if (getID.length !== 0) {
    login_logout.innerText = 'ログアウト';
    login_logout.href = "index.html";
}
