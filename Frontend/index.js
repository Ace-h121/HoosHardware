
getParts()


async function putTeam() {

    let va = document.getElementById('putTeam').value

    let woo = document.getElementById('putPart').value

    let event = await getEvent()
    alert("http://127.0.0.1:8080/"+ "Bensalem" + "?Part=" +woo+ "&TeamNum=" +va)
    if (!(woo == "")){
        const res = await fetch("http://127.0.0.1:8080/"+ event +"?Part=" +woo+ "&TeamNum=" +va, {
            method: 'POST',
            mode: "cors",

        })
        const text = await res.text()

    console.log(text)
    }

}

async function delTeam() {

    let woo = document.getElementById('delPart').value


    let event = await getEvent()
    if (!(woo == "")){
        const res = await fetch("http://127.0.0.1:8080/"+ event +"Remove?Part=" +woo, {
            method: 'POST',
            mode: "cors",

        })
        const text = await res.text()

    console.log(text)
    }
}
async function getEvent(){
    
    let team = document.getElementById('Events').value;
    alert(team)
    return team
}

async function getParts(){
    let event = await getEvent()

    const res = await fetch("http://127.0.0.1:8080/" + event, {
            method: 'GET',
            mode: "cors",

        })
        const text = await res.text()
        console.log(text)
        document.getElementById("list").value = text
}