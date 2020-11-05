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