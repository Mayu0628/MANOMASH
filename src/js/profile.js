const key_id = 'id';
const id = document.cookie.match(
    new RegExp(key_id+'\=([^\;]*)\;*')
)[1];
console.log("id:",id);

const key_oshi = 'oshi_id';
const oshi_id = document.cookie.match(
    new RegExp(key_oshi+'\=([^\;]*)\;*')
)[1];
console.log("oshi_id:",oshi_id);

const getProfileInfo = async () => {
    const url = "http://localhost:8080/profile?oshi_id=" + oshi_id;
    const headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json charset=utf-8'
    }
    const response = await fetch(url, {
        method: "GET",
        cache: "no-cache",
        headers: headers,
    })
    const profileInfo = await response.json();
    console.log("profileInfo",profileInfo)
    document.getElementById('name').value = profileInfo.OshiName
    document.getElementById('first').value = profileInfo.OshiLike1
    document.getElementById('second').value = profileInfo.OshiLike2
    document.getElementById('tird').value = profileInfo.OshiLike3
    document.getElementById('free_space').value = profileInfo.Free_Space
    document.getElementById('interest').value = profileInfo.Interest
}

window.onload = getProfileInfo()

const updateProfile = () => {
    window.location.href = '../src/mypage.html';
}