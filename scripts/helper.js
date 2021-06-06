// Execute a function when the user releases a key.
document.getElementById("helper").addEventListener("keyup", function (event) {
  // submit on enter
  if (event.key === "Enter") {
    event.preventDefault(); // cancel default action
    document.getElementById("gen").click(); // trigger button element
  }
  // disable len slider
  if (this.value.length > 0) {
    document.getElementById('slider-len').disabled = true;
    document.getElementById("slider-len").value = this.value.length;
    document.getElementById("len").innerHTML = this.value.length + parseInt(document.getElementById("extra").innerHTML);
    document.getElementById("clear").disabled = false;
    $('[id="slider-len"], [id="len"]>span').css('filter', 'opacity(34%)');
    document.getElementById("slider-len").max = 60;
  } else {
    document.getElementById("clear").disabled = true;
    document.getElementById("len").innerHTML = (5 + parseInt(document.getElementById("extra").innerHTML)).toString();
    document.getElementById("slider-len").value = 0;
    document.getElementById('slider-len').disabled = false;
    $('[id="slider-len"], [id="len"]>span').css('filter', 'opacity(100%)');
    document.getElementById("slider-len").max = 27;
  }
});
