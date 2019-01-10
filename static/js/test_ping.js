(function () {
    console.log("onload")
}())

$(document).ready(function () {
    console.log("ready")
    $("#main_container").show();
    $.ajax({
        type:'get',
        url:'/monitoring/ping',
        data:{
            "date":"2019-01-10",
        },
        success:function(result){
            buildDatas(result);
        }
    })
})

function buildDatas(result) {
    var Group = ""
    var html = ""
    for ( var prop in result.AllData) {
        var data = result.AllData[prop];
        var CategoryId = data.CategoryId;
        var ItemId = data.ItemId;
        var Category = data.Category;
        var Item = data.Item;
        if (Category!=Group) {
            Group = Category;

            var table = '<table class="table table-striped table-sm" id="t'+Group+'"></table>';
            $("#main_container").append('<div style="padding:5px;background:#fafafa;width: 100%;"><h4>'+Group+'</h4></div>').append(table)
            GenerateTestPlan('t'+Group , result.TestPlanCase);

        }
        var node = $("#main_container").find("table#t"+Group);
        var nodeId = "tr"+ItemId;
        var extNode = GrenerateRow(ItemId,result.Times);
        var itemnode = '<tr id="'+nodeId+'"><td>'+ItemId+'</td><td>'+Item+'</td>'+ extNode +'</tr>'
        node.append(itemnode);
        GreneratePingStatus(result.Result)
    }

    var Date = result.Date;
    console.log("Date:"+result.TestPlanCase.length)
}

function GenerateTestPlan(id , cases) {
    var node = $("#main_container").find("table#"+id);
    var caseTitle = '<thead><tr>' +
        '<th>Id</th>' +
        '<th>測試項目</th>';
    for ( var prop in cases) {
        caseTitle += '<th>'+cases[prop]+'</th>'
    }
    caseTitle+= '</tr></thead>'
    node.append(caseTitle);
}
function GrenerateRow(id,casetime) {
    var node = ""
    for ( var prop in casetime) {
        var timecase = casetime[prop];
        node += '<td id="td'+timecase+'">'+"  "+'</td>'
    }
    console.log("node:"+node);
    return node;
}
function GreneratePingStatus( resultData ) {

    for (var prop in resultData) {
        var statusItem = resultData[prop];
        var trID = "tr"+statusItem.ItemId;
        var tdID = "td"+statusItem.Time;
        var status = statusItem.Status;
        if(status == -1)
            $("#"+trID+" #"+tdID).text("false").css('color', 'red');
        else if(status == 1) {
            $("#"+trID+" #"+tdID).text("true").css('color', 'blue');
        }
    }
}