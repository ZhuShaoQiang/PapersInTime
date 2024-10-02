// 查询论文（通过citename)

function getPaperDetailByCiteName(citename) {
    data = {
        "citeName": citename
    }
    $.ajax({
        url: '/papers/getCitesAndCited',
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function(res){
            // 返回一堆数据，有前有后
            console.log(res)
            if (res["status"] == true) {
                return res["data"];
            } else {
                return null;
            }

        },
        error: function () {
            return null;
        }
    });  // ajax
    
}
