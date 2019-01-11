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
}
function EditStatus(e) {

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
        if(EditMode) {
            console.log("Editmode1:"+EditMode);
           /* node += '<td id="td'+timecase+'">'+'<div id="d'+timecase+'"><button type="button" class="btn btn-outline-secondary" style="font-size: 5px" , onclick="EditStatus(\''+timecase+'\')">尚無值</button>'+'</div></td>'*/
            node += '<td id="td'+timecase+'">'+'<div id="d'+timecase+'"><button type="button" class="btn btn-outline-secondary" style="font-size: 5px" , onclick="EditStatus(this)">尚無值</button>'+'</div></td>'
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
            if(status == -1){
                $("#"+trID+" #"+tdID+" div").attr("data-id",statusItem.ItemId);
                $("#"+trID+" #"+tdID).find('button').text("異常").attr('class', 'btn btn-outline-danger').attr('data-id',statusItem.ItemId);}
            else if(status == 1) {
                $("#"+trID+" #"+tdID).find('button').text("正常").attr('class', 'btn btn-outline-primary').attr('data-id',statusItem.ItemId);
            }
        } else {
            if(status == -1)
                $("#"+trID+" #"+tdID).text("異常").css('color', 'red');
            else if(status == 1) {
                $("#"+trID+" #"+tdID).text("正常").css('color', 'blue');
            }
        }

    }
}