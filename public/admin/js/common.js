
jQuery(function ($) {
    // 为必填字段添加提示
    $('.required').prepend('<span class="required-mark">*</span>');

    // 初始化select2选择框
    $('.select2').select2()
})