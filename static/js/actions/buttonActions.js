
function CloseModalWithID(ID) {
    var modal = $("#"+ID)
    modal.hide()  // 这个是关闭，hide也行
    return false;  // 阻止页面跳转
}

// 在大幕中点击跳转到引用树
function JumpToCiteTreeInModal(){
    baseURL = window.location.origin;
    targetURL = baseURL + "/view/intree";

    // 把这个内容放入localstorage
    var citeName = $("#paperCiteNameCur").val();
    window.localStorage.setItem("basepaper", citeName);

    // 拼接一个基础跳转
    window.location.href = targetURL
    return false;  // 阻止页面跳转
}