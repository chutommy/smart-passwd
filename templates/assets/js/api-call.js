const Url='/api/passwd'

// post a request to a server and show the generated values
function generatePasswd() {
	
	// get values
	var	extra = parseInt($("#extra").text());
	var len = parseInt($("#len").text())-extra;
	var helper = document.getElementById("helper").value;

    // helper exists
    if (helper != "") {
        len = 5;
        console.log(len);
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
        }
    };
    var data = JSON.stringify({"len":len,"extra":extra,"helper":helper});
    xhr.send(data);
}
