function createSingleLinkInSVG(citeName, title){
    return citeName+`[<a href='#' id='${citeName}'>${title}</a>]\n`
}
function createALinkInSVG(parentCiteName, parentTitle, childCiteName, childTitle){
    // 下面添加一个画图用的块
    tmpDrawData = parentCiteName+`[<a href='#' id='${parentCiteName}'>${parentTitle}</a>]`
    tmpDrawData += " --> "
    // 这里有个bug，必须单引号
    tmpDrawData += childCiteName + `[<a href='#' id='${childCiteName}'>${childTitle}</a>]`
    tmpDrawData += "\n"
    
    return tmpDrawData
}
function GetRawDataToDrawByData (data) {  
    // 先拿到现有的画图的数据
    var drawflowdiv = $("#forDrawUseOnly");  // 存储画图用的内容
    var record = $("#forRecordUseOnly")
    var curRawData = record.val();  // 获取当前的pair
    var curDrawData = drawflowdiv.val();  // 获取当前画图的url
    if (curDrawData == ""){
        // 第一次执行
        // 引用了谁被谁指，被谁引用指向谁
        curDrawData += "graph TD\n";
    }
    // 如果这个没有前置引用和后置引用，好像不能画单个的
    if (data["cited"].length == 0 && data["cite"].length == 0){
        return // 提前返回
    }
    // 后面添加内容
    var splitedByLine = curRawData.split("\n");  // 把原始数据切开，切成列表
    if (!data["cited"].length == 0){  // 必须===判空
        for (var ele of data["cited"]){  // 拿到引用自己的数据，这个数据是一个完整的paper数据

            tmp = data["citeName"] + " --> " + ele.citeName;  // 组合这一个引用对
            if (!splitedByLine.includes(tmp)){  // 如果这个引用对不存在
                tmpDrawData = createALinkInSVG(data["citeName"], data["title"], ele.citeName, ele.title);
                // 添加数据到里面
                curRawData += tmp + "\n";
                curDrawData += tmpDrawData;
            }
        }
    }
    if (!data["cite"].length == 0 ){
        for (var ele of data["cite"]){
            tmp = ele.citeName + " --> " + data["citeName"];
            if (!splitedByLine.includes(tmp)){  // 如果不存在才添加
                // 下面添加一个画图用的块
                tmpDrawData = createALinkInSVG(ele.citeName, ele.title, data["citeName"], data["title"]);

                curRawData += tmp + "\n";
                curDrawData += tmpDrawData;
            }
        }
    }

    record.val(curRawData)
    drawflowdiv.val(curDrawData)
    return curDrawData
}

// 更新引用树的js
function GetCitesAndCitedAjax(basepaper){
    console.log("base: ", basepaper)
    data = {
        "citeName": basepaper
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
                const output = document.querySelector(".flowchart");
                const code = GetRawDataToDrawByData(res["data"]);
                let insert = function (code) {
                    output.innerHTML = code;
                };
                console.log("output")
                console.log(output)
                // 每次需要不同的id，要不然不行，并且有最短长度，不能低于3个
                mermaid.render("preparedScheme"+Date.now(), code, insert);
            }
        },
        error: function (res) {
            alert("发生错误：", res["msg"]);
        }
    });  // ajax
}
