(function () {

}())
var EditMode = false;

$(document).ready(function () {
    initViewEvent()
    $("#main_container").show();
    $('.form_datetime').datetimepicker({
        format: 'yyyy-mm-dd',
        autoclose: 1,
        todayHighlight: 1,
        startView: 2,
        minView: 2,
        forceParse: 0,
        showMeridian: 1,
    }).on('changeDate', function(e) {
        var date = $('#date').val();
        $("#main_container").empty();
        updatDate();
    });
    $('.form_datetime').datetimepicker("setDate", new Date());
    updatDate();
})

function initViewEvent() {
    $("#export").click(function () {
        $("#main_container").empty();
        if(EditMode) {
            EditMode = false;
            updatDate();
        }else {
            EditMode = true;
            updatDate();
        }
    });
}
function updatDate() {
    var date = $('#date').val();
    $.ajax({
        type: 'get',
        url: '/monitoring/ping',
        data:{
            "date": date,
        },
        success:function(result){
            buildDatas(result);
        }
    })
}

function buildDatas(result) {
    var Group = ""
    var html = ""
    for ( var prop in result.AllData) {
        var data = result.AllData[prop];
        var CategoryId = data.CategoryId; //分類 id int
        var ItemId = data.ItemId;  // IP id  int
        var Category = data.Category;  // 分類  string
        var Item = data.Item;   //IP Title  string
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
}
function EditStatus(parent_id, e) {

    var text = document.getElementById("tr"+parent_id);
    var tr = "#tr"+parent_id;
    var td = "#td"+e;
    var node = $(tr+" "+td).find("div");
    var status = node.attr("data-status");
    if(status == -1 || status == undefined) {
        node.find("button").attr('class', 'btn btn-outline-primary').text("正常");
        node.attr("data-status",1);
    } else if (status == 1){
        node.find("button").attr('class', 'btn btn-outline-danger').text("異常");
        node.attr("data-status",-1);
    }
    console.log(text);
}
function GenerateTestPlan(id , cases) {
    var node = $("#main_container").find("table#"+id);
    var caseTitle = '<thead><tr>' +
        '<th class="text-center">Id</th>' +
        '<th class="text-center">測試項目</th>';
    for ( var prop in cases) {
        caseTitle += '<th class="text-center">'+cases[prop]+'</th>'
    }
    caseTitle+= '</tr></thead>'
    node.append(caseTitle);
}
function GrenerateRow(parent_id,casetime) {
    var node = ""
    for ( var prop in casetime) {
        var timecase = casetime[prop];
        if(EditMode) {
            node += '<td id="td'+timecase+'">'+'<div class="text-center" id="d'+timecase+'"><button type="button" class="btn btn-outline-secondary" style="font-size: 5px" , onclick="EditStatus(\''+parent_id+'\',\''+timecase+'\')">尚無值</button>'+'</div></td>'
        }else {
            console.log("Editmode2:"+EditMode);
            node += '<td id="td'+timecase+'">'+"  "+'</td>'
        }
    }
    return node;
}
function GreneratePingStatus( resultData ) {

    for (var prop in resultData) {
        var statusItem = resultData[prop];
        var trID = "tr"+statusItem.ItemId;
        var tdID = "td"+statusItem.Time;
        var status = statusItem.Status;
        if(EditMode) {
            $("#"+trID+" #"+tdID+" div").attr("data-status",status);
            if(status == -1){
                $("#"+trID+" #"+tdID).find('button').text("異常").attr('class', 'btn btn-outline-danger');
            }
            else if(status == 1) {
                $("#"+trID+" #"+tdID).find('button').text("正常").attr('class', 'btn btn-outline-primary')
            }
        } else {
            if(status == -1)
                $("#"+trID+" #"+tdID).text("異常").css('color', 'red').attr('class','text-center');
            else if(status == 1) {
                $("#"+trID+" #"+tdID).text("正常").css('color', 'blue').attr('class','text-center');
            }
        }

    }
}