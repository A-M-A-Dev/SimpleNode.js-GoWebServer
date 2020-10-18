function getUrlParameter(name) {
    name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
    var regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
    var results = regex.exec(location.search);
    return results === null ? `${location.protocol}//${location.host}${location.pathname}` : decodeURIComponent(results[1].replace(/\+/g, ' '));
};

$(document).ready(() => {
    $('#url').html(getUrlParameter('url'));
    console.log(getUrlParameter('url'));
});