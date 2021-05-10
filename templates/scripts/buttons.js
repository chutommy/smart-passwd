// button to clear custom helper field
function resetHelper() {
  if (document.getElementById("helper").value !== "") {
    document.getElementById("helper").value = "";
    document.getElementById("len").innerHTML = 5 + parseInt(document.getElementById("extra").innerHTML);
    document.getElementById("slider-len").value = 0;
    document.getElementById("slider-len").disabled = false;
    $('[id="slider-len"], [id="len"]>span').css("filter", "hue-rotate(0deg)");
    document.getElementById("slider-len").max = 27;
  }
}

// button to copy a content of the element with the given id
function copyText(id) {
  const text = document.getElementById(id).value;
  navigator.clipboard.writeText(text)
    .then(() => $.notify("Successfully copied", "success"),
      err => $.notify("Unsuccessfully copied", "error"));
}

// button to toggle the password's visibility
function togglePasswordVisibility() {
  const hidBtn = document.getElementById("hidBtn");
  const field = document.getElementById("passwd");

  if (hidBtn.innerHTML === '<i class="fas fa-eye"></i>') {
    hidBtn.innerHTML = '<i class="fas fa-eye-slash"></i>';
    field.type = "password";
  } else {
    hidBtn.innerHTML = '<i class="fas fa-eye"></i>';
    field.type = "text";
  }

}