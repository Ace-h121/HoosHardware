logMovies()

async function logMovies() {


    const res = await fetch('http://127.0.0.1:8080/AllentownRemove?Part=neo550', {
        method: 'POST',
        mode: "cors",

    })
    const text = await res.text()

    console.log(text)

  }