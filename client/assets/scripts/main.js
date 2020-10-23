import "./bootstrap-input-spinner.js";

$(document).ready(() => {
    $('#nav-brand').hover(() => {
        $('#dev')
            .removeClass('bg-warning')
            .addClass('bg-transparent')
            .removeClass('text-dark')
            .addClass('text-white');
        $('#ama')
            .addClass('bg-warning')
            .removeClass('bg-transparent')
            .addClass('text-dark')
            .removeClass('text-white');

        $('#nav-icon').addClass('text-warning');
    }, () => {
        $('#ama')
            .removeClass('bg-warning')
            .addClass('bg-transparent')
            .removeClass('text-dark')
            .addClass('text-white');
        $('#dev')
            .addClass('bg-warning')
            .removeClass('bg-transparent')
            .addClass('text-dark')
            .removeClass('text-white');

        $('#nav-icon').removeClass('text-warning')
    });

    $("#line-number").inputSpinner({
        template: `
        <div class="input-group \${groupClass}">
            <div class="input-group-prepend"></div>
            <input type="text" inputmode="decimal" style="text-align: \${textAlign}" class="form-control"/>
            <button style="min-width: \${buttonsWidth}" class="btn btn-decrement rounded-0 btn-outline-danger" type="button">\${decrementButton}</button>
            <button style="min-width: \${buttonsWidth}; margin-right: 2.2px;" class="btn btn-increment rounded-0 btn-outline-success" type="button">\${incrementButton}</button>
        </div>`
    })
});

function showMessage(title, message) {
    $('#response-title').html(title);
    $('#response-message').html(message);
    $('#response-modal').modal('show');
}

window.request = (name, url) => {
    $.get(url)
        .done(response => showMessage(`${name} response`, response))
        .fail(response => showMessage(`${name} error: ${response.status}!`, response.statusText));
};

window.writeFile = url => {
    $.ajax({
        url: url,
        data: {
            line: $('#line-number').val()
        },
        headers: {
            Accept: "text/plain"
        }
    })
        .done(response => $('#write-result').html(response))
        .fail(response => showMessage(`${name} error: ${response.status}!`, response.statusText));
};

window.adder = url => {
    $.ajax({
        method: 'POST',
        url: url,
        contentType: 'application/json',
        data: JSON.stringify({
            a: Number($('#first-number').val()),
            b: Number($('#second-number').val())
        }),
        headers: {
            Accept: "application/json"
        }
    })
        .done(response => $('#result').html(response.sha256))
        .fail(jqXHR => showMessage(`${name} error: ${jqXHR.status}!`, `${jqXHR.statusText}: ${JSON.parse(jqXHR.responseText).message}`));
}