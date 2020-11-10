gf = {
    login: function () {
        this.noticeClear("#loginModal");

        var loginName = $('#login-username').val();
        var loginPasswd = $('#login-passwd').val();
        if (loginName == '') {
            this.notice('账号不能为空！');
            return;
        }
        var regexEmail = /^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$/;
        if (!regexEmail.test(loginName)) {
            gf.notice('账号格式不正确！', "#loginModal");
            return;
        }
        if (loginPasswd == '') {
            gf.notice('密码不能为空！', "#loginModal");
            return;
        }
        if (loginPasswd.length < 6) {
            gf.notice('账号或密码错误', "#loginModal");
            return;
        }
        if (loginPasswd.length > 20) {
            gf.notice('账号或密码错误', "#loginModal");
            return;
        }

        jQuery.ajax({
            type: 'POST',
            url: '/user/do-login',
            data: {
                "passport": loginName,
                "password": hex_md5(loginPasswd + loginPasswd)
            },
            success: function (data) {
                if (data.code == 0) {
                    gf.notice('登录成功', "#loginModal");
                    setTimeout(function () {
                        window.top.location.href = "/";
                    }, 500);

                } else {
                    // loadPicimageCode();
                    gf.notice('登录失败：' + data.message, "#loginModal");
                }
            },
            error: function (html) {
                var flag = (typeof console != 'undefined');
                if (flag) console.log("服务器忙，提交数据失败，代码:" + html.status + "，请联系管理员！");
                alert("服务器忙，提交数据失败，请联系管理员！");
            }
        });
    },
    loginGithub: function () {
        this.noticeClear("#loginModal");
        this.notice("暂未开通Github登录，尽情期待", "#loginModal")
    },
    logout: function () {
        jQuery.ajax({
            type: 'POST',
            url: '/user/logout',
            data: {},
            success: function (data) {
                if (data.code == 0) {
                    window.top.location.href = "/";
                } else {
                    // loadPicimageCode();
                   console.log('登录失败：' + data.message);
                }
            },
            error: function (html) {
                var flag = (typeof console != 'undefined');
                if (flag) console.log("服务器忙，提交数据失败，代码:" + html.status + "，请联系管理员！");
                alert("服务器忙，提交数据失败，请联系管理员！");
            }
        });
    },
    notice: function (msg, pageCls) {
        pageCls = pageCls || "";
        pageCls = (pageCls == "") ? ".notice" : (pageCls + " .notice");
        $(pageCls).html('<span class="iconfont">&#xe653;</span> ' + msg);
        $(pageCls).fadeIn();
    },
    noticeClear: function (pageCls) {
        pageCls = pageCls || "";
        pageCls = (pageCls == "") ? ".notice" : (pageCls + " .notice");
        $(pageCls).hide();
    }
}

jQuery(function ($) {
    // 为必填字段添加提示
    $('.required').prepend('&nbsp;<span class="icon iconfont red">&#xe71b;</span>');

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