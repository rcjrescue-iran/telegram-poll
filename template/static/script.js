function sendToServer() {
    var result = {};
    result.level = getRadioItems(1);
    result.refree = getRadioItems(2);
    result.proportionality = getRadioItems(3);
    result.timing = getRadioItems(4);
    result.morality = getRadioItems(5);
    result.idea = getRadioItems(6);
    result.quality = getRadioItems(7);
    result.partition = getRadioItems(8);
    result.broadcast = getRadioItems(9);    
    result.points = getRadioItems(10);
    result.text = $("#text").val();
    $.ajax({
        url: "/api/submit",
        type: "POST",
        data: JSON.stringify(result),
        beforeSend: function (request) {
            toast("لطفا منتظر بمانید")
            request.setRequestHeader("X-CSRF-Token", $("input[name='gorilla.csrf.Token']").val());
        },
        success: function () {
            toast("نظر شما برای ما ثبت شد ، با تشکر از همکاری شما")
        },
        error: function () {
            toast("در ثبت نظر شما خطایی رخ داده است")
        }
    })
}

function toast(data) {
    $('.toast').remove()
    Materialize.toast(data, 3000)
}

function getRadioItems(id) {
    for (var i = 1; i <= 3; i++) {
        if ($("#item" + id + "-" + i).is(':checked')) {
            return i;
        }
    }
    return 0;
}