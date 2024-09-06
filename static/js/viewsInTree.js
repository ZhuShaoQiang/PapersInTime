function GetRawDataToDrawByData (data) {  
    // 先拿到现有的画图的数据
    var drawflowdiv = $("#forDrawUseOnly");  // 存储画图用的内容
    var record = $("#forRecordUseOnly")
    var curRawData = record.val();  // 获取当前的pair
    var curDrawData = drawflowdiv.val();  // 获取当前画图的url
    console.log("cur: ",curRawData)
    if (curDrawData == ""){
        // 第一次执行
        // 引用了谁被谁指，被谁引用指向谁
        curDrawData += "graph TD\n";
    }
    // 后面添加内容
    var splitedByLine = curRawData.split("\n");  // 把原始数据切开，切成列表
    if (!data["cited"].length == 0){  // 必须===判空
        for (var ele of data["cited"]){  // 拿到引用自己的数据，这个数据是一个完整的paper数据
            tmp = data["citeName"] + " --> " + ele.citeName;  // 组合这一个引用对
            if (!splitedByLine.includes(tmp)){  // 如果这个引用对不存在
                // 下面添加一个画图用的块
                tmpDrawData = data["citeName"]+`[<a href='#' id='${data["citeName"]}'>${data["title"]}</a>]`
                tmpDrawData += " --> "
                // 这里有个bug，必须单引号
                tmpDrawData += ele.citeName + `[<a href='#' id='${ele.citeName}'>${ele.title}</a>]`
                tmpDrawData += "\n"

                // 添加数据
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
                tmpDrawData = ele.citeName+`[<a href='#' id='${ele.citeName}'>${ele.title}</a>]`
                tmpDrawData += " --> "
                tmpDrawData += data["citeName"]+`[<a href='#' id='${data["citeName"]}'>${data["title"]}</a>]`
                tmpDrawData += "\n"

                curRawData += tmp + "\n";
                curDrawData += tmpDrawData;
            }
        }
    }

    record.val(curRawData)
    drawflowdiv.val(curDrawData)
    return curDrawData
}