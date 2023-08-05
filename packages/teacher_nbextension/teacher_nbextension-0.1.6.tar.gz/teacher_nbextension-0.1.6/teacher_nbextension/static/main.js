define([
  'jquery',
  'base/js/namespace',
  'base/js/events'
], function (
  $,
  Jupyter,
  events) {
    function getCookie(name) { var r = document.cookie.match("\\b" + name + "=([^;]*)\\b"); return r ? r[1] : undefined; }


    function execute_codecell_callback (evt, data) {
        var cell = data.cell;
        var _xsrf_cookie = getCookie('_xsrf');

        var url = location.href.split('/notebooks')[0] // TODO: not sure is it good idea
        $.ajax(
            {"url":  url + "/execute", 
            "type": "post",
            "headers": {"X-XSRFToken": _xsrf_cookie}, 
            "data": {'input': cell.toJSON(), 'output': cell.toJSON()['outputs'][0]}, 
            "success": function(d) {console.log(d)}, 
            "dataType":"json"})
    }

    var load_ipython_extension = function () {
        events.on('execute.CodeCell', execute_codecell_callback);
    };

    return {
        load_ipython_extension : load_ipython_extension,
    };
});
