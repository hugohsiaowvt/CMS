<head>
    <meta charset="UTF-8">
    <title></title>
    <link rel="stylesheet" type="text/css" href="/static/css/report_monitoring.css">
    <script>
        function addLog(index) {
            var root = $("#tlog_" + index);
            var intput_root = $("#logcase"+index);
            var log = intput_root.find("input[data-type='log']").val();
            var note = intput_root.find("input[data-type='note']").val();
            console.log(""+intput_root);
            var html = "<tr>\n" +
                "<td></td>\n" +
                "<td>"+log+"</td>\n" +
                "<td colspan=\"2\">"+note+"</td>\n" +
                "</tr>";
            root.append(html);
            intput_root.find("input[data-type='log']").val("");
            intput_root.find("input[data-type='note']").val("");
        }

        function switchbtn(e) {
            if($(e).text() == "新增紀錄..")
                $(e).text("關閉");
            else
                $(e).text("新增紀錄..");
        }
    </script>
</head>
<div class="container">
    <div class="row">
        <div class="col">
            <div class="card card-table" style="padding: 0px 0px 10px 0px;">
                <div class="card-header">
                    <div>遊戲測試</div>
                </div>
                <div class="card-body table-responsive" style="padding: 0px;">
                    <ul class="nav nav-tabs nav-justified">
                        <li class="nav-item">
                            <a class="nav-link active" data-toggle="tab" href="#android">Android</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" data-toggle="tab" href="#ios">IOS</a>
                        </li>
                    </ul>
                    <div class="tab-content">
                        <div class="tab-pane active container" id="android">
                            <table class="table">
                                <thead>
                                <tr>
                                    <th style="width: 1%"></th>
                                    <th style="width: 65%;">紀錄說明</th>
                                    <th style="width: 15%;">備註</th>
                                    <th style="width: 15%;"></th>
                                </tr>
                                </thead>
                                <thead class="thead-light">
                                <tr>
                                    <th colspan="4">下載安裝測試情況 (Android系統) - www.gotend.cn測試</th>
                                </tr>
                                </thead>

                                <tbody>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.gotend.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.gotend.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.gotend.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.gotend.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                </tbody>
                                <thead class="thead-light">
                                <tr>
                                    <th colspan="4">下載安裝測試情況 (Android系統) - www.chinawaygo.com測試</th>
                                </tr>
                                </thead>

                                <tbody>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.chinawaygo.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.chinawaygo.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.chinawaygo.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                <tr>
                                    <td></td>
                                    <td>PM 06:32 www.chinawaygo.cn網頁點選android系統測試下載無異常。</td>
                                    <td colspan="2">AM 06:00 通報Shark 及Hsiang今日測試情況。</td>
                                </tr>
                                </tbody>

                            </table>
                        </div>
                        <div class="tab-pane container" id="ios">
                            <table class="table">
                                <thead>
                                <tr>
                                    <th style="width: 1%"></th>
                                    <th style="width: 65%;">紀錄說明</th>
                                    <th style="width: 15%;">備註</th>
                                    <th style="width: 15%;"></th>
                                </tr>
                                </thead>
                                <thead class="thead-light">
                                <tr>
                                    <th colspan="3">下載安裝測試情況 (IOS系統) -  www.gotend.cn測試</th>
                                    <th style="padding: 4px ; vertical-align: center">
                                        <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#logcase0" onclick="switchbtn(this)">新增紀錄..</button></th>
                                </tr>
                                <tr id="logcase0" class="collapse">
                                    <td colspan="4">
                                        <div class="input-group mb-3">
                                            <input type="text" class="form-control" style="width: 60%" data-type="log" placeholder="請寫下紀錄..">
                                            <input type="text" class="form-control" style="width: 30%" data-type="note" placeholder="備註..">
                                            <div class="input-group-append" style="width: 10%">
                                                <button class="btn btn-primary small-btn" type="button" onclick="addLog(0)">提交</button>
                                            </div>
                                        </div>
                                    </td>
                                </tr>
                                </thead>
                                <tbody  id="tlog_0"></tbody>
                                <thead class="thead-light">
                                <tr>
                                    <th colspan="3">下載安裝測試情況 (IOS系統) -  www.chinawaygo.com測試</th>
                                    <th style="padding: 4px ; vertical-align: center">
                                        <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#logcase1" onclick="switchbtn(this)">新增紀錄..</button></th>
                                </tr>
                                <tr id="logcase1" class="collapse">
                                    <td colspan="4">
                                        <div class="input-group mb-3">
                                            <input type="text" class="form-control" style="width: 60%" data-type="log" placeholder="請寫下紀錄..">
                                            <input type="text" class="form-control" style="width: 30%" data-type="note" placeholder="備註..">
                                            <div class="input-group-append" style="width: 10%">
                                                <button class="btn btn-primary small-btn" type="button" onclick="addLog(1)">提交</button>
                                            </div>
                                        </div>
                                    </td>
                                </tr>
                                </thead>
                                <tbody  id="tlog_1"></tbody>
                                <thead class="thead-light">
                                <tr>
                                    <th colspan="3">下載安裝測試情況 (IOS系統) -  測IOS系統 超飞娱乐</th>
                                    <th style="padding: 4px ; vertical-align: center">
                                        <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#logcase2" onclick="switchbtn(this)">新增紀錄..</button></th>
                                </tr>
                                <tr id="logcase2" class="collapse">
                                    <td colspan="4">
                                        <div class="input-group mb-3">
                                            <input type="text" class="form-control" style="width: 60%" data-type="log" placeholder="請寫下紀錄..">
                                            <input type="text" class="form-control" style="width: 30%" data-type="note" placeholder="備註..">
                                            <div class="input-group-append" style="width: 10%">
                                                <button class="btn btn-primary small-btn" type="button" onclick="addLog(2)">提交</button>
                                            </div>
                                        </div>
                                    </td>
                                </tr>
                                </thead>
                                <tbody  id="tlog_2"></tbody>
                                <thead class="thead-light">
                                <tr>
                                    <th colspan="3">下載安裝測試情況 (IOS系統) -  大成棋牌</th>
                                    <th style="padding: 4px ; vertical-align: center">
                                        <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#logcase3" onclick="switchbtn(this)">新增紀錄..</button></th>
                                </tr>
                                <tr id="logcase3" class="collapse">
                                    <td colspan="4">
                                        <div class="input-group mb-3">
                                            <input type="text" class="form-control" style="width: 60%" data-type="log" placeholder="請寫下紀錄..">
                                            <input type="text" class="form-control" style="width: 30%" data-type="note" placeholder="備註..">
                                            <div class="input-group-append" style="width: 10%">
                                                <button class="btn btn-primary small-btn" type="button" onclick="addLog(3)">提交</button>
                                            </div>
                                        </div>
                                    </td>
                                </tr>
                                </thead>
                                <tbody  id="tlog_3"></tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <br><br>
    <div class="row">
        <div class="col">
            <div class="card card-table" style="padding: 0px">
                <div class="card-header">
                    <div>扛打監控</div>
                </div>
                <div class="card-body table-responsive" style="padding: 0px;">
                    <table class="table table-striped table-borderless">
                        <thead>
                        <tr>
                            <th style="width: 60%;">監控項目</th>
                            <th style="width: 25%;">紀錄說明</th>
                            <th style="width: 15%;">備註</th>
                        </tr>
                        </thead>
                        <tbody class="tbody no-border-x ">
                        <tr>
                            <td>攻擊狀況</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#case1">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="case1">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        <tr>
                            <td>扛打機器調度情況</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#case2">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="case2">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        <tr>
                            <td>機器存活及死亡情況</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#case3">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="case3">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
        <div class="col">
            <div class="card card-table" style="padding: 0px">
                <div class="card-header">
                    <div>IM系統測試</div>
                </div>
                <div class="card-body table-responsive" style="padding: 0px;">
                    <table class="table table-striped table-borderless">
                        <thead>
                        <tr>
                            <th style="width: 60%;">監控項目</th>
                            <th style="width: 25%;">紀錄說明</th>
                            <th style="width: 15%;">備註</th>
                        </tr>
                        </thead>
                        <tbody class="tbody no-border-x ">
                        <tr>
                            <td>HugoGram和VCtalk&旺聊 訊息發送及接收</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#g2_case1">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="g2_case1">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        <tr>
                            <td>VCtalk網頁情況</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#g2_case2">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="g2_case2">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        <tr>
                            <td>HugoGram和VCtalk&旺聊 支付及轉帳功能情況</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#g2_case3">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="g2_case3">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        <tr>
                            <td>語音功能情況</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#g2_case4">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="g2_case4">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        <tr>
                            <td>版本升級情況</td>
                            <td><span class="text-success">無異常</span></td>
                            <td style="padding: 4px ; vertical-align: center">
                                <button class="btn btn-secondary small-btn" data-toggle="collapse" data-target="#g2_case5">詳情..
                                </button>
                            </td>
                        </tr>
                        <tr class="collapse" id="g2_case5">
                            <td colspan="3">Lorem ipsum dolor sit amet, consectetur adipisicing elit,
                                sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
                                veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                            </td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</div>