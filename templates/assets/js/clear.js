function resetHelper() {
    document.getElementById('helper').value = ''
    document.getElementById("len").innerHTML = 5 + parseInt(document.getElementById("extra").innerHTML);
    document.getElementById("slider-len").value = 0;
    document.getElementById('slider-len').disabled = false;
    $('[id="slider-len"], [id="len"]>span').css('filter', 'hue-rotate(0deg)');
}
