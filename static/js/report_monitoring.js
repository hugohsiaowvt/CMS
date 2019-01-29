(function () {

}())
var EditMode = false;

$(document).ready(function () {

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
        var date = $('#date').val();
        if (!CheckCurrentDate()) {
            alert("無法修改非當天報表");
            return
        }

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

function  CheckCurrentDate() {    //sDate1和sDate2是2002-12-18格式
    var date = $('#date').val();
    var dest_date = new Date(date);
    var now_date = new Date(convertDateToString (new Date()));
    // 如果跨日後還沒有超過早上六點就還能修改
    if (dest_date - now_date == -86400000) {
        var today=new Date();
        var h = (today.getHours()<10 ? '0' : '')+today.getHours();
        var m = (today.getMinutes()<10 ? '0' : '')+today.getMinutes();
        var currentDateTime = h + m;
        return currentDateTime <= "0600" ? true : false;
    } else {
        return dest_date - now_date < 0 ? false : true;
    }
}

function updatDate() {
    if(!CheckCurrentDate())
        EditMode = false;

    var date = $('#date').val();

    $.ajax({
        type: 'get',
        url: '/reportmonitoring/data',
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

        var resultData = result.Result[prop];
        var resultId = resultData.ResultId

        if (Category!=Group) {
            Group = Category;
            var table = '<table class="table table-striped table-sm" id="t'+Group+'"></table>';
            $("#main_container").append('<div style="padding:5px;background:#fafafa;width: 100%;"><h4>'+Group+'</h4></div>').append(table)
            GenerateSchema('t'+Group);
        }
        var node = $("#main_container").find("table#t"+Group);
        var nodeId = "tr"+ItemId;
        var extNode = GrenerateRow(ItemId, resultId);
        var itemnode =
            '<tr id="'+nodeId+'">' +
                '<td class="text-left">'+Item+'</td>' +
                extNode
            '</tr>';
        node.append(itemnode);
    }
    GreneratePingStatus(result.Result)
}
function EditStatus(parent_id) {
    var text = document.getElementById("tr" + parent_id);
    var tr = "#tr"+parent_id;
    var node = $(tr).find("div");
    var status = node.attr("data-status");
    var resultID = node.attr("data-resultid");
    var url = "/reportresult/status/edit";
    switch (status) {
        case "0":
            status = 1;
            node.find("button").attr('class', 'btn btn-outline-primary').text("正常");
            node.attr("data-status",status);
            break;
        case "-1":
            status = 1;
            node.find("button").attr('class', 'btn btn-outline-primary').text("正常");
            node.attr("data-status",status);
            break;
        case "1":
            status = -1;
            node.find("button").attr('class', 'btn btn-outline-danger').text("異常");
            node.attr("data-status",status);
            break
    }

    var data = {
        "result_id": resultID,
        "status": status,
    };

    ModifyStatus(url, data);
}

function EditNote(parent_id) {
    var tr = "#tr"+parent_id;
    note = $(tr+" #td-note").find('textarea').val().replace(/\n/g,'<br>');
    var node = $(tr).find("div");
    var resultID = node.attr("data-resultid");
    var data = {
        "result_id": resultID,
        "note": note,
    };

    $.ajax({
        type: 'post',
        url: '/reportresult/note/edit',
        data: data,
        success: function(result) {
            var response_status = result.Status;
            var msg = result.Msg;
            if (response_status < 0) {
                alert(msg);
            }
            console.log("response_status:"+response_status+" msg:"+msg)
        }
    })

}

function OnChange(parent_id) {
    // $("#tr"+parent_id+" #td-note").find('textarea');
}

function convertDateToString (date) {
    var MM = (date.getMonth()+1 < 10 ? "0" :"" )+ (date.getMonth()+1);
    var dd = (date.getDate() < 10 ? "0" :"") + date.getDate();
    return date.getFullYear()+"-"+MM+"-"+dd;
}

function ModifyStatus(url, data) {
    $.ajax({
        type: 'get',
        url: url,
        data: data,
        success: function(result) {
            var response_status = result.Status;
            var msg = result.Msg;
            if (response_status < 0) {
                alert(msg);
            }
            console.log("response_status:"+response_status+" msg:"+msg)
        }
    })
}
function GenerateSchema(id) {
    var node = $("#main_container").find("table#"+id);
    var caseTitle = '<thead><tr>' +
        '<th class="text-left" width="20%">監控項目</th>' +
        '<th class="text-left" width="10%">異常情況</th>' +
        '<th class="text-left" width="60%">備註</th>';
    caseTitle+= '</tr></thead>'
    node.append(caseTitle);
}
function GrenerateRow(parent_id) {
    var node = ""
    if(EditMode) {
        node += '<td id="td-status">'+
            '<div class="text-center">' +
                '<button type="button" class="btn btn-outline-secondary" style="font-size: 5px" , onclick="EditStatus(\''+parent_id+'\')">尚無值</button>'+
            '</div></td>'
        node += '<td id="td-note">'+
            '<div class="text-center">'+
                '<textarea onInput="OnChange(\''+parent_id+'\')" onChange="EditNote(\''+parent_id+'\')" style="width:100%;" rows="3">\n' +
                '</textarea>'+
            '</div></td>'
    } else {
        node += '<td class="text-center" id="td-status">'+""+'</td>'
        node += '<td class="text-center" id="td-note">'+""+'</td>'
    }
    return node;
}
function GreneratePingStatus( resultData ) {

    for (var prop in resultData) {
        var statusItem = resultData[prop];
        var trID = "tr"+statusItem.ItemId;
        var tdStatus = "td-status";
        var tdNote = "td-note";
        var resultID = statusItem.ResultId;
        var status = statusItem.Status;
        var note = statusItem.Note
        if(EditMode) {
            // 狀態
            $("#"+trID+" #"+tdStatus+" div").attr("data-status",status).attr("data-resultid",resultID);
            if(status == -1){
                $("#"+trID+" #"+tdStatus).find('button').text("異常").attr('class', 'text-left btn btn-outline-danger');
            }
            else if(status == 1) {
                $("#"+trID+" #"+tdStatus).find('button').text("正常").attr('class', 'text-left btn btn-outline-primary');
            }
            // 備註
            $("#"+trID+" #"+tdNote+" div").attr("data-resultid",resultID);

            note = note.replace(/\<br>/g,'\n');
            console.log(note)
            $("#"+trID+" #"+tdNote).find('textarea').text(note).attr('class', 'form-control').attr('rows', '6');
        } else {
            // 狀態
            if(status == -1)
                $("#"+trID+" #"+tdStatus).text("異常").css('color', 'red').attr('class','text-left');
            else if(status == 1) {
                $("#"+trID+" #"+tdStatus).text("正常").css('color', 'blue').attr('class','text-left');
            }
            // 備註
            $("#"+trID+" #"+tdNote).html(note).attr('class','text-left');
        }
    }
}