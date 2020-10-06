function showMessage(title, message) {
    $('#response-title').html(title);
    $('#response-message').html(message);
    $('#response-modal').modal('show');
}

function request(name, url) {
    $.get(url)
        .done((response) => showMessage(`${name} response`, response))
        .fail((xhr, status) => showMessage(`${name} error: ${xhr.status}!`, status));
}

function requestToNode() {
    request("Node.js", "node");
}

function requestToGo() {
    request("GoLang", "go");
}
