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

        $(document).ready(function () {
            console.log("ready")
            $("#dialog_div").dialog({
                autoOpen: false,
            });
            $("#dialog_div").dialog("close");
        })

        function addIP() {
            $("#dialog_div").dialog("open");
            $.ajax({
                type:'get',
                url:'/monitoring/add'
                ,success:function(result){
                   console.log(result);
                }
            });
        }

        function delIP() {
            alert("delIP")
        }

        function SearchReport() {
            alert("SearchReport")
        }

        function onLoad() {

        }

    </script>
</head>
<body onload="onLoad()" style="display: none">
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
        <tbody>
        {{range $key, $val := .MonitoringItems}}
        <tr>
            <td>{{$val.Monitor_group}}</td>
            <td>{{$val.Title}}</td>
            <td>{{$val.Ip}}</td>
        </tr>
        {{end}}
        </tbody>
    </table>
</body>
</html>