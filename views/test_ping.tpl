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
<table class="tg">
    <colgroup>
        <col style="width: 10%">
        <col style="width: 18%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
        <col style="width: 6%">
    </colgroup>
    <tr>
        <th class="tg-p7ly" colspan="14">{{.Date}}夜間高防400G、800G各IP試ping結果</th>
    </tr>
    <tr>
        <td class="tg-s6z2" colspan="2">測試項目</td>
        <td class="tg-s6z2">PM6:30</td>
        <td class="tg-s6z2">PM7:45</td>
        <td class="tg-s6z2">PM8:15</td>
        <td class="tg-s6z2">PM9:30</td>
        <td class="tg-s6z2">PM10:45</td>
        <td class="tg-s6z2">PM11:15</td>
        <td class="tg-s6z2">AM00:30</td>
        <td class="tg-s6z2">AM01:45</td>
        <td class="tg-s6z2">AM02:15</td>
        <td class="tg-s6z2">AM03:30</td>
        <td class="tg-s6z2">AM04:45</td>
        <td class="tg-s6z2">AM05:55</td>
    </tr>
    {{$isFirst := true}}
    {{range $k, $v := .Count}}
        {{$isFirst = true}}
        {{range $k1, $v1 := $.Data}}
            {{if eq $k $v1.CategoryId}}
                {{if eq $isFirst true}}
                    <tr>
                        <td class="tg-s6z2" rowspan="{{$v}}">{{$v1.Category}}</td>
                        <td class="tg-s6z2">{{$v1.Item}}</td>

                        {{range $time := $.Times}}
                            {{$hasValue := false}}
                            {{range $k2, $v2 := $.Result}}
                                {{if eq $v1.ItemId $v2.ItemId}}
                                    {{if eq $time $v2.Time}}
                                        {{if eq $v2.Status 1}}
                                            <td class="tg-baqh-green"></td>
                                        {{else}}
                                            <td class="tg-baqh-red"></td>
                                        {{end}}
                                        {{$hasValue = true}}
                                    {{end}}
                                {{end}}
                            {{end}}
                            {{if eq $hasValue false}}
                                <td class="tg-baqh-gray"></td>
                            {{end}}
                        {{end}}
                    </tr>
                    {{$isFirst = false}}
                {{else}}
                    <tr>
                        <td class="tg-s6z2">{{$v1.Item}}</td>
                        {{range $time := $.Times}}
                            {{$hasValue := false}}
                            {{range $k2, $v2 := $.Result}}
                                {{if eq $v1.ItemId $v2.ItemId}}
                                    {{if eq $time $v2.Time}}
                                        {{if eq $v2.Status 1}}
                                            <td class="tg-baqh-green"></td>
                                        {{else}}
                                            <td class="tg-baqh-red"></td>
                                        {{end}}
                                        {{$hasValue = true}}
                                    {{end}}
                                {{end}}
                            {{end}}
                            {{if eq $hasValue false}}
                                <td class="tg-baqh-gray"></td>
                            {{end}}
                        {{end}}
                    </tr>
                {{end}}
            {{end}}
        {{end}}
    {{end}}
</table>