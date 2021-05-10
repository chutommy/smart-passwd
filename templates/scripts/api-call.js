const URI = '/gen';

let statusCounter = 0;

// post a request to a server and show generated values
function generatePasswd() {
  // retrieve document values
  const extra = parseInt($("#extra").text());
  let len = parseInt($("#len").text()) - extra;
  const helper = document.getElementById("helper").value;

  const request = {
    len: len,
    extra: extra,
    helper: helper
  };

  const param = {
    headers: {
      "Content-Type": "application/json"
    },
    method: "GET",
    body: JSON.stringify(request)
  };

  fetch(URI, param)
    .then(data => data.json())
    .then(resp => {
      document.getElementById("passwd").value = resp.password;
      document.getElementById("helper-p").value = resp.helper;

      document.getElementById("hidBtn").disabled = false;
      document.getElementById("copyPasswd").disabled = false;
      document.getElementById("copyHelper").disabled = false;

      // status
      hidBtn.innerHTML = '<i class="fas fa-eye"></i>';
      document.getElementById("passwd").type = "text";

      statusCounter++;
      const temp = statusCounter;

      // set visibility timeout
      setTimeout(function () {
        if (temp === statusCounter) {
          if (hidBtn.innerHTML === '<i class="fas fa-eye"></i>') {
            togglePasswordVisibility();
          }
        }
      }, 3000);
    })
    .catch(error => console.log(error))
}
