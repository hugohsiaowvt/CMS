(function () {

}())
var EditMode = false;
var TodayDate ;

$(document).ready(function () {
    TodayDate  = new Date().getMilliseconds();
    console.log("TodayDate:"+TodayDate)

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
        $("#main_container").empty();
        updatDate();
    });
    initViewEvent()
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
/*    var tmp = new Date();
    var _time = (new Date(date)) - tmp ;*/
    //console.log("tmp:"+tmp+" date:"+date+ " data:"+_time );
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
        var itemnode = '<tr id="'+nodeId+'"><td class="text-center">'+ItemId+'</td><td class="text-center">'+Item+'</td>'+ extNode +'</tr>'
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
    var resultID = node.attr("data-resultid");
    var url = "/pingresult/edit";
    switch (status) {
        case undefined:
            var url = "/pingresult/add";
        case "-1":
            status = 1
            node.find("button").attr('class', 'btn btn-outline-primary').text("正常");
            node.attr("data-status",status);
            break;
        case "1":
            status = -1
            node.find("button").attr('class', 'btn btn-outline-danger').text("異常");
            node.attr("data-status",status);
            break
    }
    console.log("time:"+e);
    var _date = $("#date").val();
    if (e <= 1200) {
        //am 所以加一天
        var date_format = new Date($("#date").val())
        date_format.setDate(date_format.getDate()+1)

        var MM = (date_format.getMonth()+1 < 10 ? "0" :"" )+ (date_format.getMonth()+1);
        var dd = (date_format.getDate() < 10 ? "0" :"") + date_format.getDate();
        _date = date_format.getFullYear()+"-"+MM+"-"+dd;

    }
    console.log("_date:"+_date);
    var data = {
        "result_id":resultID,
        "item_id":parent_id,
        "date":_date,
        "time":e,
        "status" : status,
    }
    ModifyStatus(url,data);
}

function ModifyStatus(url,data) {

    $.ajax({
        type: 'get',
        url: url,
        data:data,
        success:function(result){
            var response_status = result.Status;
            var msg = result.Msg;
            console.log("response_status:"+response_status+" msg:"+msg)
        }
    })
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
        var resultID = statusItem.ResultId;
        var status = statusItem.Status;
        if(EditMode) {
            $("#"+trID+" #"+tdID+" div").attr("data-status",status).attr("data-resultid",resultID);
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