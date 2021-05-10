const Url='/api/passwd';
var counter = 0;

// post a request to a server and show the generated values
function generatePasswd() {

    // get values
    var	extra = parseInt($("#extra").text());
    var len = parseInt($("#len").text())-extra;
    var helper = document.getElementById("helper").value;

    // helper exists
    if (helper != "") {
        len = 5;
    }

    // set request
    var xhr = new XMLHttpRequest();
    xhr.open("POST", Url, true);
    xhr.setRequestHeader("Content-Type", "application/json");

    // handle response
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var resp = JSON.parse(xhr.responseText);

            // insert into vars
            document.getElementById("passwd").value = resp.password;
            document.getElementById("helper-p").value = resp.helper;

            // enable copy and hide buttons
            document.getElementById("hidBtn").disabled = false;
            document.getElementById("copyPasswd").disabled = false;
            document.getElementById("copyHelper").disabled = false;

            // set the status
            hidBtn.innerHTML = '<i class="fas fa-eye"></i>';
            document.getElementById("passwd").type = "text";

              counter++;
              var temp = counter;
              // hide after a while
              setTimeout(function() {
                  if (temp == counter) {
                      if (hidBtn.innerHTML == '<i class="fas fa-eye"></i>') {
                          togglePasswordVisibility();
                      }
                  }
              },3000);
    })
    .catch(error => {
      console.log(error)
    })
}
