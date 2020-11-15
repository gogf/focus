// 全局管理对象
gf = {
    // 刷新验证码
    reloadCaptcha: function() {
        $("img.captcha").attr("src","/captcha?v="+Math.random());
    },
    // 用户模块
    user: {
        // 退出登录
        logout: function () {
            swal({
                title:   "注销登录",
                text:    "您确定需要注销当前登录状态吗？",
                icon:    "warning",
                buttons: ["取消", "确定"]
            }).then((value) => {
                if (value) {
                    window.location.href = "/user/logout";
                }
            });
        },
    },
    // 内容模块
    content: {
        delete: function (id) {
            swal({
                title:   "删除内容",
                text:    "您确定要删除该内容吗？",
                icon:    "warning",
                buttons: ["取消", "确定"]
            }).then((value) => {
                if (value) {
                    jQuery.ajax({
                        type: 'POST',
                        url : '/content/do-delete',
                        data: {
                            id: id
                        },
                        sync: true,
                        success: function (data) {
                            swal({
                                title:   "删除完成",
                                text:    "3秒后自动跳转到",
                                icon:    "success",
                                timer:   2000,
                                buttons: false
                            }).then((value) => {
                                window.location.href = "/";
                            })
                        }
                    });
                }
            });
        }
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
    let pageItem = $("ul.pagination li.page-item")
    if (pageItem.length > 4) {
        pageItem.each(function (index, element) {
            if (index < 2 || index > pageItem.length - 3) {
                return
            }
            if ($(element).attr("class").indexOf("disabled") > -1) {
                $(element).removeClass("disabled").addClass("active");
                return
            }
        });
    }
})