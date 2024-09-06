function paperModalByPaperIDOnly(paperID){
    console.log(paperID)
    //先从后台读取一个论文
    dataToBack = {
        "citeName": paperID
    }
    $.ajax({
        url: "/papers/getPaperByID",
        method: "post",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify(dataToBack),
        success: function(res){
            // 这个地方应该不会有成功，因为成功就跳转了
            if(res["status"] == true) {
                paperModal(res["data"])
            } else {
                alert("发生错误：", res["msg"])
            }
        },
        error: function (res) {
            alert("发生错误：", res["msg"])
        }
    });

    // 修改标题
    $("#modal_title").text(paper.title)
    // 拿到当前我们访问的url
    var curURL = window.location.origin;
    var pdfPath = paper.path;
    console.log(curURL);
    // 展示论文
    $("#iframeInPaperDetailModal").attr("src", curURL+"/"+pdfPath)
    // 展示modal
    $("#paperDetailModal").show()
}
function paperModal(paper){
    console.log(paper)
    // 修改标题
    $("#modal_title").text(paper.title)
    // 拿到当前我们访问的url
    var curURL = window.location.origin;
    var pdfPath = paper.path;
    console.log(curURL);
    // 展示论文
    $("#iframeInPaperDetailModal").attr("src", curURL+"/"+pdfPath)
    // 防止一个标记，记录当前论文
    $("#paperCiteNameCur").val(paper.citeName);
    // 展示modal
    $("#paperDetailModal").show()
}
function DetailCallback(t) {
    // 跳转查看详情
    var citeName = t.attr("id").replace("detail_", "")
    dataToBack = {
        "citeName": citeName
    }
    $.ajax({
        url: "/papers/detail",
        method: "post",
        dataType: "json",
        contentType: "application/json",
        data: JSON.stringify(dataToBack),
        success: function(res){
            // 这个地方应该不会有成功，因为成功就跳转了
            if(res["status"] == true) {
                paperModal(res["data"])
            } else {
                alert("发生错误：", res["msg"])
            }
        },
        error: function (res) {
            alert("发生错误：", res["msg"])
        }
    });
    // 这个地方不跳转详情，拉起一个大幕让人看，然后还能随手缩下去。
}
function DownloadCallback(t) {
    // 
    alert(t.attr("id"));
}
function DeleteCallback(t) {
    // 
    alert(t.attr("id"));
}
function CreateButtonAppendToDivWithCallback(btnId, btnClass, btnText, targetDIV, callback) {
    $(`
    <button id="${btnId}" type="button" class="${btnClass}">${btnText}</button>
    `).appendTo(targetDIV).bind("click", function(){
        callback($(this))
    })
}
