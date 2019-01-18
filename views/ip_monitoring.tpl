<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Basic FileBox - jQuery EasyUI Demo</title>
    <script src="/static/js/ip_monitoring.js"></script>



</head>
<div id="main_container" style="display: none;" align="left">
    <h2>扛打排程</h2>
    <div id="accordion">

        <div class="card">
            <div class="card-header">
                <a class="collapsed card-link" data-toggle="collapse" href="#collapseTwo">
                   新增排程
                </a>
            </div>
            <div id="collapseTwo" class="collapse" data-parent="#accordion">
                <div class="card-body">
                    <form style="max-width: 500px ;">
                        <div class="form-group" id="group_input_mode" style="display:none ">
                            <label for="pwd">群組:</label>
                            <input type="text" class="form-control" id="group_add" placeholder="新增群組">
                            <div style="margin-top: 10px">
                                <button type="button" class="btn btn-primary " onclick="onAddGroup()">新增</button>
                                <button type="button" class="btn btn-danger " onclick="switchMode()">取消</button>
                            </div>

                        </div>
                        <div class="form-group" id="group_select_mode" style="display: ">
                            <label for="group">群組:</label>
                            <div class="dropdown">
                                <button type="button" class="btn btn-light dropdown-toggle" data-toggle="dropdown">
                                    <span id="onSelectGroup">請選擇或新增群組</span>
                                </button>
                                <div class="dropdown-menu">
                                    <div id="select_group">
                                    </div>

                                    <div class="dropdown-divider"></div>
                                    <a class="dropdown-item" href="#" onclick="switchMode()">新增群組</a>
                                </div>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="pwd">測試項目:</label>
                            <input type="text" class="form-control" id="title" placeholder="輸入測試項目標題">
                        </div>
                        <div class="form-group">
                            <label for="pwd">IP:</label>
                            <input type="text" class="form-control" id="ip" placeholder="輸入排程IP">
                        </div>
                        <div class="form-group">
                            <label for="pwd">操作:</label>
                            <div class="dropdown">
                                <button type="button" class="btn btn-light dropdown-toggle" data-toggle="dropdown">
                                    <span id="onSelectAction">請選擇操作</span>
                                </button>
                                <div class="dropdown-menu">
                                    <a class="dropdown-item" href="#" onclick="actionSelect(this)" value="1">Ping</a>
                                    <a class="dropdown-item" href="#" onclick="actionSelect(this)" value="2">查表</a>
                                </div>
                            </div>
                        </div>

                        <button type="submit" class="btn btn-primary" onclick="CreateSchedule()">建立排程</button>
                        <button type="submit" class="btn btn-secondary" onclick="CloseSchedule()">取消</button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>
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


</body>
<script></script>
</html>