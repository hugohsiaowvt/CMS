(function () {

}())

$(document).ready(function () {

    $("#main_container").show();
    $.ajax({
        type:'get',
        url:'/monitoring/ips',
        success:function(result){
            buildDatas(result);
        }
    })
})

function addIP() {
    $("#dialog_div").dialog("open");
    $.ajax({
        type:'get',
        url:'/monitoring/add',
        data: {
            "group":"高防400",
            "title":"高防400_1",
            "ip":"192.168.1.1"
        }
        ,success:function(result){
            buildDatas(result);
        }
    })
}

function delIP(e) {
    alert("delIP"+e)
    console.log(e);
}

function SearchReport() {
    alert("SearchReport")
}

function buildDatas(result) {
    if(result==null)
        return;
    var html ="";
    for ( var prop in result) {
        var id = result[prop].id;
        var action = result[prop].action == 1 ? "Ping IP" : "查表1";
        html += " <tr id='tr_"+prop+"'>" +
            "<td class=\"text-center\"><div id='d_category_id' data-category=\""+result[prop].category_id+"\"></div>" +
            (parseInt(prop)+1)+"</td>" +
            "<td class=\"text-center\">"+result[prop].monitor_group+"</td>" +
            "<td class=\"text-center\"><input id=\"title\" type=\"text\" value=\""+result[prop].title+"\" style=\"display: none ; text-align: center\"><span>"+result[prop].title+"</span></td>" +
            "<td class=\"text-center\"><input id=\"ip\" type=\"text\" value=\""+result[prop].ip+"\" style=\"display: none ; text-align: center\"><span>"+result[prop].ip+"</span></td>" +
            "<td class=\"text-center\">" +
            "<button  id=\"btn_type\" type=\"button\" value=\""+ result[prop].action +"\" class=\"btn btn-outline-secondary\" style=\"font-size: 5px\" disabled>"+action+"</button>" +
            "<td><div class=\"text-center\"><button type=\"button\" class=\"btn btn-info\" style=\"font-size: 5px\"  onclick='edit("+prop+','+id+")'>編輯</button>&nbsp;" +
            "<button type=\"button\" class=\"btn btn-danger\" style=\"font-size: 5px\"  onclick='del("+id+")'>刪除</button>  </div></td>"+
            "</tr>"
    }
    $("#tbody_datas").html(html);
}

function edit(prop, id) {
    var root = $("#tr_"+prop);
    editmode = root.find("input[type=text]").css('display') == 'none' ? false : true;
    if(!editmode) {
        root.find("input[type=text]").css('display','');
        root.find("input[type=text]").change(function () {

        });
        root.find("button.btn.btn-outline-secondary").removeAttr('disabled').attr("class","btn btn-outline-success")
        root.find("span").css('display','none');
        root.find("button.btn.btn-info").text("儲存").attr("class","btn btn-warning");
        root.find("button.btn.btn-outline-success").on('click',function () {
            console.log("prop:"+$(this).val());
            if ($(this).val() == 1) {
                $(this).val(2);
                $(this).text("查表");
            } else {
                $(this).val(1);
                $(this).text("Ping IP");
            }

        });
    } else {

        root.find("input[type=text]").css('display','none  ');
        root.find("span").css('display','');
        root.find("button.btn-outline-success").attr('disabled', true).attr("class","btn btn-outline-secondary");
        root.find("button.btn.btn-outline-secondary").off('click')
        root.find("button.btn.btn-warning").text("編輯").attr("class","btn btn-info");

        var category_id = root.find("#d_category_id").attr("data-category");
        var title= root.find("#title").val();
        var ip=root.find("#ip").val();
        var type=root.find("#btn_type").val();

        console.log("category_id:"+category_id+" title:"+title+" ip:"+ip+" type:"+type +" id:"+id);

        $.ajax({
            type: 'get',
            url: '/monitoring/edit',
            data:{
                "id": id,
                "category_id":category_id,
                "title":title,
                "ip":ip,
                "type":type
            },
            success:function(result){
                var response_status = result.Status;
                var msg = result.Msg;
                if(response_status<0)
                    alert(msg)
                else {
                    root.find("input[type=text]").each(function (_index) {
                        var _val = this.value;
                        root.find("span").eq(_index).text(_val);
                    });
                }
                console.log("response_status:"+response_status+" msg:"+msg)

            }
        })
    }
}

function del(id) {

    var r = confirm("確定要刪除此筆資料？");
    if (r) {
        $.ajax({
            type: 'get',
            url: '/monitoring/del',
            data:{
                "id": id,
            },
            success:function(result){
                var response_status = result.Status;
                var msg = result.Msg;
                if(response_status<0)
                    alert(msg)
                else {
                    updatIps();
                }

            }
        })
    }
}

function updatIps() {
    $("#tbody_datas").empty();
    $.ajax({
        type:'get',
        url:'/monitoring/ips',
        success:function(result){
            buildDatas(result);
        }
    })
}