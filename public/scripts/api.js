// const URI = '/gen';

let statusCounter = 0;

// post a request to a server and show generated values
async function generatePasswd() {
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

  const request = {
    len: len,
    extra: extra,
    helper: helper.trim()
  };

  let result = await password(request);
  updatePasswdFields(result[0], result[1]);

  // const param = {
  //   headers: {
  //     "Content-Type": "application/json"
  //   },
  //   method: "POST",
  //   body: JSON.stringify(request)
  // };
  //
  // fetch(URI, param)
  //   .then(data => data.json())
  //   .then(resp => {
  //     updatePasswdFields(resp.password, resp.helper)
  //   })
  //   .catch(error => console.log(error))
}

async function password(request) {
  const helper = await parseHelper(request);
  return generate(request.extra, helper);
}

async function parseHelper(request) {
  if (request.helper === "") {
    let newHelper = [];

    let dist = distribute(request.len);
    for (let i = 0; i < dist.length; i++) {
      newHelper.push(await randomWord(dist[i]));
    }

    request.helper = newHelper.join(" ")
  }

  return request.helper
}

function updatePasswdFields(password, helper) {
  document.getElementById("passwd").value = password;
  document.getElementById("helper-p").value = helper;

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
}