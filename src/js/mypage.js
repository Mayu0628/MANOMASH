const viewProfile = async (value) => {
    document.cookie = await "oshi_id=" + value;
    console.log(document.cookie);
    window.location.href = '../src/profile.html';
}

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
    user_name.innerHTML = mypageInfo.UserName;

    const user_id = document.getElementById('user_id')
    console.log(user_id)
    user_id.innerHTML = mypageInfo.UserID;

    const user_comment = document.getElementById('user_comment')
    console.log(user_comment)
    user_comment.innerHTML = mypageInfo.Introduce;
    console.log("OshiID.length",mypageInfo.OshiID.length)
    for (let i = 0; i < mypageInfo.OshiID.length; i++) {
        console.log("OshiID",mypageInfo.OshiID[i])
        document.getElementById("oshis").insertAdjacentHTML(
            "afterend",
            `
                <button type="button" class="osiprofile hover-tab" id="oshi" onclick="viewProfile(${mypageInfo.OshiID[i]})">
                    <img src="./images/demo-oshi-icon.png"><br>
                    <div class="osiprofile-information">
                        <h2> ${mypageInfo.OshiName[i]} </h2>
                    </div>
                </button>
            `
        );
        
    } 
}

window.onload = getMypageInfo()
