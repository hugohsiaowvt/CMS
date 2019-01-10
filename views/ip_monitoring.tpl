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
            console.log("onload")
        }())

        $(document).ready(function () {
            console.log("ready")
            $("#main_container").show();
            $.ajax({
                type:'get',
                url:'/monitoring/add',
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
                html += " <tr><td>" +
                    result[prop].monitor_group+"</td>" +
                    "<td>"+result[prop].title+"</td>" +
                    "<td>"+result[prop].ip+"</td> "+
                    "</tr>"
            }
            console.log("html:"+html);
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
    <div class="table-responsive">
        <table class="table table-striped table-sm">
            <thead>
            <tr>
                <th>#</th>
                <th>測試項目</th>
                <th>IP</th>
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