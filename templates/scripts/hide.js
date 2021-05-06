function hidePasswd() {
    var hidBtn = document.getElementById("hidBtn");
    var field = document.getElementById("passwd");

    if (hidBtn.innerHTML == '<i class="fas fa-eye"></i>') {

        hidBtn.innerHTML = '<i class="fas fa-eye-slash"></i>';
        field.type = "password";

    } else {

        hidBtn.innerHTML = '<i class="fas fa-eye"></i>';
        field.type = "text";
    }

}
