// Get the input field
var input = document.getElementById("helper");

// Execute a function when the user releases a key
input.addEventListener("keyup", function(event) {

    // submit on enter
    if (event.keyCode === 13) {
        // Cancel the default action, if needed
        event.preventDefault();
        // Trigger the button element with a click
        document.getElementById("gen").click();
    }

    // if custom helper, disable len slider
    if(this.value.length > 0) {

        document.getElementById('slider-len').disabled = true;
        document.getElementById("slider-len").value = this.value.length;
        document.getElementById("len").innerHTML = this.value.length + parseInt(document.getElementById("extra").innerHTML);

    } else {

        document.getElementById("len").innerHTML = 5 + parseInt(document.getElementById("extra").innerHTML);
        document.getElementById("slider-len").value = 0;
        document.getElementById('slider-len').disabled = false;
    }
});
