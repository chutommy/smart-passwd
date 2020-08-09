function hidePasswd() {
    var hidBtn = document.getElementById("hidBtn");
    var field = document.getElementById("passwd");

    if (hidBtn.innerHTML == "Hide") {

        hidBtn.innerHTML = "Show";
        field.type = "password";

    } else {

        hidBtn.innerHTML = "Hide";
        field.type = "text";
    }

}
