



async function putTeam() {

    let va = document.getElementById('putTeam').value

    let woo = document.getElementById('putPart').value


    alert("http://127.0.0.1:8080/Allentown?Part=" + woo + "&TeamNum=" + va)



    const res = await fetch("http://127.0.0.1:8080/Allentown?Part=" +woo+ "&TeamNum=" +va, {
        method: 'POST',
        mode: "cors",

    })
    const text = await res.text()

    console.log(text)

  }

  async function delTeam() {

    let woo = document.getElementById('delPart').value

    const res = await fetch("http://127.0.0.1:8080/AllentownRemove?Part=" +woo, {
        method: 'POST',
        mode: "cors",

    })
    const text = await res.text()

    console.log(text)

  }