function copyFunc(id) {
    var text = document.getElementById(id).value;

    // copy
    navigator.clipboard.writeText(text).then(function() {
        $.notify("Successfully copied", "success");
    }, function(err) {
        $.notify("Unsuccessfully copied", "error");
    });
}
