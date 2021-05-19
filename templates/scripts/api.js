const URI = '/gen';

let statusCounter = 0;

// post a request to a server and show generated values
function generatePasswd() {
  // retrieve document values
  const extra = parseInt($("#extra").text());
  const helper = document.getElementById("helper").value;

  let len = parseInt($("#len").text()) - extra;
  if (helper !== "") {
    len = 0;
  }
  
  if (helper.trim() === "" && len < 5) {
    resetHelper();
    return
  }
  
  if (helper.trim() === "" && len < 5) {
    resetHelper();
    return
  }
  
  if (helper.trim() === "" && len < 5) {
    resetHelper();
    return
  }

  const request = {
    len: len,
    extra: extra,
    helper: helper.trim()
  };

  const param = {
    headers: {
      "Content-Type": "application/json"
    },
    method: "POST",
    body: JSON.stringify(request)
  };

  fetch(URI, param)
    .then(data => data.json())
    .then(resp => {
      document.getElementById("passwd").value = resp.password;
      document.getElementById("helper-p").value = resp.helper;

      document.getElementById("hideButton").disabled = false;
      document.getElementById("copyPasswd").disabled = false;
      document.getElementById("copyHelper").disabled = false;

      // status
      const hide = document.getElementById("hideButton");
      hide.innerHTML = '<i class="fas fa-eye"></i>';
      document.getElementById("passwd").type = "text";

      statusCounter++;
      const temp = statusCounter;

      // set visibility timeout
      setTimeout(function () {
        if (temp === statusCounter) {
          if (hide.innerHTML === '<i class="fas fa-eye"></i>') {
            togglePasswordVisibility();
          }
        }
      }, 3000);
    })
    .catch(error => console.log(error))
}
