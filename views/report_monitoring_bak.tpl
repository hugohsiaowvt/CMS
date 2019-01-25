<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Basic FileBox - jQuery EasyUI Demo</title>
    <link rel="stylesheet" type="text/css" href="/static/easyui/themes/default/easyui.css">
    <link rel="stylesheet" type="text/css" href="/static/easyui/themes/icon.css">
    <link rel="stylesheet" type="text/css" href="/static/css/datepicker.css">
    <style type="text/css">
        .tg  {border-collapse:collapse;border-spacing:0;border-color:#999;width:100%;}
        .tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#999;color:#444;background-color:#F7FDFA;}
        .tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#999;color:#fff;background-color:#26ADE4;}
        .tg .tg-s6z2{text-align:center}
        .tg .tg-baqh-gray{text-align:center;vertical-align:center;background-color:#DCDCDC;}
        .tg .tg-baqh-red{text-align:center;vertical-align:center;background-color:#FF3333;}
        .tg .tg-baqh-green{text-align:center;vertical-align:center;background-color:#32EF32;}
        .tg .tg-p7ly{font-weight:bold;font-size:20px;text-align:center}
    </style>
    <script type="text/javascript" src="/static/easyui/jquery.min.js"></script>
    <script type="text/javascript" src="/static/easyui/jquery.easyui.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-datepicker.js"></script>
    <script type="text/javascript" src="/static/js/report_monitoring.js"></script>
</head>
<body>
<div>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">夜間監控事項記錄</h1>

        <div class="btn-toolbar mb-2 mb-md-0" style="margin-right: 20px">
            <span style="align-self: center">選擇日期：</span> <div class="input-group date form_datetime" data-date="1979-09-16" data-date-format="yyyy-mm-dd">
                <input class="form-control" id="date" size="16" type="text" value="">
                <span class="input-group-addon"><span class="glyphicon glyphicon-remove"></span></span>
                <span class="input-group-addon"><span class="glyphicon glyphicon-th"></span></span>
            </div>
            <button type="button" class="btn btn-sm btn-outline-secondary" style="margin-left: 10px" id="export">修改狀態</button>
        </div>
    </div>
    <div id="main_container"></div>
</div>
</body>
<script></script>
</html>
