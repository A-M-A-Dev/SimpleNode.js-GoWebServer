function request(url) {
    $.get(url)
    	.done((response) => alert(`response: ${response}`))
        .fail((xhr, status) => alert(`error: ${xhr.status}: ${status}`));
}

function requestToNode() {
    request("node");
}

function requestToGo() {
    request("go");
}
