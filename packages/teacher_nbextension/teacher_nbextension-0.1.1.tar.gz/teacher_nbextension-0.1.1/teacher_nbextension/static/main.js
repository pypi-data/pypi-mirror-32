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

        $.ajax(
            {"url": "http://" + location.host + "/execute", 
            "type": "post",
            "headers": {"X-XSRFToken": _xsrf_cookie}, 
            "data": cell.toJSON(), 
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
