// 全局管理对象
gf = {
    // 刷新验证码
    reloadCaptcha: function() {
        $("img.captcha").attr("src","/captcha?v="+Math.random());
    },
    notice: function (msg, pageCls) {
        pageCls = pageCls || "";
        pageCls = (pageCls == "") ? ".notice" : (pageCls + " .notice");
        $(pageCls).html('<span class="iconfont">&#xe653;</span> ' + msg);
        $(pageCls).fadeIn();
    }
}

jQuery(function ($) {
    // 为必填字段添加提示
    $('.required').prepend('&nbsp;<span class="icon iconfont red">&#xe71b;</span>');

    // 回车搜索
    $("#search").keydown(function (e) {
        if (e.keyCode == 13) {
            gf.search();
            e.preventDefault();
        }
    });

    // 分页高亮
    $pageItem = $("ul.pagination li.page-item")
    if ($pageItem.length > 4) {
        $pageItem.each(function (index, element) {
            if (index < 2 || index > $pageItem.length - 3) {
                return
            }

            if ($(element).attr("class").indexOf("disabled") > -1) {
                $(element).removeClass("disabled").addClass("active");
                return
            }
        });
    }
})