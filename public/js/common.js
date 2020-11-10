gf = {
    login: function () {
        $('#login').show();
        $('#noLogin').hide();
        $('#loginModal').modal('toggle');
    },
    logout: function () {
        $('#login').hide();
        $('#noLogin').show();
    }
}

jQuery(function ($) {
    // 为必填字段添加提示
    $('.required').append('&nbsp;<span class="icon iconfont red">&#xe71b;</span>');

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