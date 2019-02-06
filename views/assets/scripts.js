function togglePasswordView(e, eye) {
    $(this).toggleClass('opened');

    var input = $(this).siblings('input');

    if (input.attr("type") === "password") {
        input.attr("type", "text");
    } else {
        input.attr("type", "password");
    }
}

function resolveTimeZoneOptions() {
    var timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    var hrs = -(new Date().getTimezoneOffset() / 60);

    $('input[name=timezone]') = timezone;
    $('input[name=utc]') = hrs;
}

$(function () {
    $('.eye').click(togglePasswordView);

    resolveTimeZoneOptions();
});

function validatePassword() {
    let match = $("#password").val() === $("#confirm-password").val()
    if (!match) {
        alert("Passwords do not match!")
    }
    return match
}