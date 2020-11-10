gf = {
    login: function() {
        $('#login').show();
        $('#noLogin').hide();
        $('#loginModal').modal('toggle');
    },
    logout: function() {
        $('#login').hide();
        $('#noLogin').show();
    }
}

jQuery(function($) {
    // 为必填字段添加提示
    $('.required').append('&nbsp;<span class="icon iconfont red">&#xe71b;</span>');
})