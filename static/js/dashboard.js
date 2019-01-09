(function () {
    $(document).ready(function () {
        $('.nav.flex-column li a').click(function () {
            var ctx = $(this).find("span").attr("data-feather");
            var old = $('.nav.flex-column li a.nav-link.active').find("span").attr("data-feather");
            if (old) {
                $('.nav.flex-column li a.nav-link.active').removeClass("active");
            }
            $(this).addClass('active');
            SelectPage(ctx)
        })
    });

    SelectPage("login.tpl");
}())

function SelectPage( sPage) {
    $.ajax({
        type:'get',
        url:'/dashboard/showPage',
        data:{
            "page":sPage
        },success:function(result){
            $(".table-responsive").html(result);
        }
    })
}
