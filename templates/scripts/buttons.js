// button to clear custom helper field
function resetHelper() {
  if (document.getElementById("helper").value !== "") {
    document.getElementById("helper").value = "";
    document.getElementById("len").innerHTML = (5 + parseInt(document.getElementById("extra").innerHTML)).toString();
    document.getElementById("slider-len").value = 0;
    document.getElementById("slider-len").disabled = false;
    $('[id="slider-len"], [id="len"]>span').css("filter", "hue-rotate(0deg)");
    document.getElementById("slider-len").max = 27;
  }
}

// button to copy a content of the element with the given id
function copyText(id) {
  const text = document.getElementById(id);
  navigator.clipboard.writeText(text.value)
    .then(() => $.notify("Successfully copied", "success"),
      () => $.notify("Unsuccessfully copied", "error"));
}

function copyDisabledElement(e) {
  e.disabled = false;
  e.select();
  e.setSelectionRange(0, 99999); /* For mobile devices */
  e.disabled = true;

  try {
    document.execCommand("copy");
    removeSelection();
    $.notify("Successfully copied", "success");
  } catch (err) {
    $.notify("Unsuccessfully copied", "error");
  }
}

// button to toggle the password's visibility
function togglePasswordVisibility() {
  const hideButton = document.getElementById("hideButton");
  const field = document.getElementById("passwd");

  if (hideButton.innerHTML === '<i class="fas fa-eye"></i>') {
    hideButton.innerHTML = '<i class="fas fa-eye-slash"></i>';
    field.type = "password";
  } else {
    hideButton.innerHTML = '<i class="fas fa-eye"></i>';
    field.type = "text";
  }
}

function removeSelection() {
  if (window.getSelection) {
    if (window.getSelection().empty) {  // Chrome
      window.getSelection().empty();
    } else if (window.getSelection().removeAllRanges) {  // Firefox
      window.getSelection().removeAllRanges();
    }
  } else if (document.selection) {  // IE?
    document.selection.empty();
  }
}