function loadPapers() {
    $.ajax({
        url: '/getpapers',
        method: 'GET',
        success: function(data) {
            if (data.length == 0) {
                console.log("没有论文");
                return;
            }
            const cardsObject = $("#containerForCards");  // 拿到需要放置到的位置
            cardsObject.html("");  // 清空现有表格

            data.forEach(paper => {
                // 创建内部container，专门放置button
                var divForButton = $(`<div>`);
                divForButton.addClass(".container");
                divForButton.addClass("row");
                CreateButtonAppendToDivWithCallback("detail_"+paper.citeName, "btn btn-primary", "查看详情", divForButton, DetailCallback);
                CreateButtonAppendToDivWithCallback("download_"+paper.citeName, "btn btn-info", "下载论文", divForButton, DownloadCallback);
                CreateButtonAppendToDivWithCallback("delete_"+paper.citeName, "btn btn-danger", "删除论文", divForButton, DeleteCallback);

                // 创建card-body并添加按钮
                var cardBody = $(`
                <div class="card-body">
                    <h5 class="card-title">${paper.author}</h5>
                    <h6 class="card-subtitle mb-2 text-muted">${paper.year}</h6>
                    <p class="card-text">${paper.title}</p>
                </div>
                `)
                divForButton.appendTo(cardBody)

                // 创建一个card，添加body并且添加到页面
                $(`
                    <div class="card" style="width: 18rem;" id="${paper.citeName}">
                    </div>
                `).append(cardBody).appendTo(cardsObject)
                
            });
        },
        error: function() {
            alert('加载论文列表失败');
        }
    });
}