function adder(url) {
    $.post(
        url = url,
        data = {
            a: Number($('#first-number').val()),
            b: Number($('#second-number').val())
        },
        type = 'json')
        .done(response => $('#result').html(response.code))
        .fail(response => showMessage(`${name} error: ${response.status}!`, response.statusText));
}