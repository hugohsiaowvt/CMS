(function () {
    $(document).ready(function () {
        console.log("ready");
        $('.nav.flex-column li a').click(function () {
            var ctx = $(this).find("span").attr("data-feather");
            var old = $('.nav.flex-column li a.nav-link.active').find("span").attr("data-feather");
            if (old) {
                $('.nav.flex-column li a.nav-link.active').removeClass("active");
            }
            $(this).addClass('active');
            $.ajax({
                type:'get',
                url:'/showPage',
                data:{
                    "page":ctx
                },success:function(result){
                    $(".table-responsive").html(result);
                }
            })
        })
    });
}())
