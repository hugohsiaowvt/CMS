<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Basic FileBox - jQuery EasyUI Demo</title>
    <link rel="stylesheet" type="text/css" href="/static/easyui/themes/default/easyui.css">
    <link rel="stylesheet" type="text/css" href="/static/easyui/themes/icon.css">

    <script type="text/javascript" src="/static/easyui/jquery.min.js"></script>
    <script type="text/javascript" src="/static/easyui/jquery.easyui.min.js"></script>
    <script type="text/javascript">
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

        function delIP() {
            alert("delIP")
        }

        function SearchReport() {
            alert("SearchReport")
        }

        function buildDatas(result) {
            if(result==null)
                return;
            var html ="";
            for ( var prop in result) {
                html += " <tr><td class=\"text-center\">" +
                    prop+"</td>" +
                    "<td class=\"text-center\">"+result[prop].monitor_group+"</td>" +
                    "<td class=\"text-center\">"+result[prop].title+"</td>" +
                    "<td class=\"text-center\">"+result[prop].ip+"</td> "+
                    "<td class=\"text-center\">"+result[prop].action+"</td> "+
                    "<td><div class=\"text-center\"><button type=\"button\" class=\"btn btn-info\" style=\"font-size: 5px\">編輯</button>&nbsp;<button type=\"button\" class=\"btn btn-danger\" style=\"font-size: 5px\">刪除</button>  </div></td>"+
                    "</tr>"
            }
            $("#tbody_datas").html(html);
        }

    </script>
</head>
<div id="main_container" style="display: none;" align="left">
    <h2>扛打排程</h2>
    <p>高防打監控排程.</p>
    <div style="padding:5px;background:#fafafa;width: 100%;border:1px solid #ccc">
        <a href="#" class="easyui-linkbutton" plain="true" iconCls="icon-add" onclick="addIP()">新增監控</a>
        <a href="#" class="easyui-linkbutton" plain="true" iconCls="icon-cancel" onclick="delIP()">刪除監控</a>
        <a href="#" class="easyui-linkbutton" plain="true" iconCls="icon-search" onclick="SearchReport()">查詢</a>
    </div>
    <div>
        <table class="table table-striped table-sm">
            <thead>
            <tr>
                <th class="text-center">#</th>
                <th class="text-center">群組</th>
                <th style="width:20%" class="text-center">測試項目</th>
                <th style="width:20%" class="text-center">IP</th>
                <th class="text-center">操作(Ping / 查表)</th>
                <th class="text-center">編輯</th>
            </tr>
            </thead>
            <tbody id="tbody_datas">
            </tbody>
        </table>
    </div>
</div>

</body>
<script></script>
</html>