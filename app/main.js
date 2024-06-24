const form = document.getElementById("form").addEventListener("submit", (e)=> {
  e.preventDefault();
  const input = document.getElementById("input")

  let form_data = new FormData(e.target);
  const barcode = form_data.get("barcode");
  console.log(barcode);


  const jsonn = {
    result: barcode
  }

  fetch("http://10.93.87.180:9090/api", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(jsonn)
  }).then((res)=> {
    if (res.ok) {
      console.log(res.statusText)
      input.value = ""
    }
  })
})