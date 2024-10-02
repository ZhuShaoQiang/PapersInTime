// 输入框的动作变化的js
function getStoredCiteName() {
    // 首先拿到输入框的输入内容
    var content = $("#citename").val();
    if(content.length <= 0){
        return;
    }
    var res = getPaperDetailByCiteName(content);
    if (res != null ){
        console.log(res);
    } else {
        console.log("get a null")
    }
}