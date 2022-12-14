let container = document.getElementById("container");
let amountInputField = document.getElementById("amount-field");
const urlSearchParams = new URLSearchParams(window.location.search);
const params = Object.fromEntries(urlSearchParams.entries());

let title = document.getElementById("title");
if (params.type == "invoice") {
    title.textContent = "Receive by Invoice"
} else {
    title.textContent = "Receive via BTC address" 
    receiveByBtcAddress()
}

function clearPage() {
    while (container.firstChild) {
        container.removeChild(container.lastChild);
    }
}

async function createInvoice() {
    if (!amountInputField.value) {
        return alert("Amount is mandatory");
    }
    clearPage()
    let response = await fetch("/receive/invoice", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            Amount: parseInt(amountInputField.value || "0")
        })
    });

    let jsonResponse = await response.json();
    let png = await b64toBlob(jsonResponse.qrcode, "image/png")
    let objectURL = URL.createObjectURL(png);
    let qrcode = document.createElement('img');
    qrcode.src = objectURL;
    let fee = document.createElement("div");
    fee.textContent = "Fee: " + jsonResponse.fee;
    let lnurl = document.createElement("div");
    lnurl.textContent = jsonResponse.invoice;
    container.appendChild(qrcode);
    container.appendChild(fee)
    container.appendChild(lnurl);
}

async function receiveByBtcAddress() {
    clearPage()
    let response = await fetch("/receive/btc", {
        method: 'GET'
    }); 
    let jsonResponse = await response.json();
    let png = await b64toBlob(jsonResponse.qrcode, "image/png")
    let objectURL = URL.createObjectURL(png);
    let qrcode = document.createElement('img');
    qrcode.src = objectURL;
    let address = document.createElement("div");
    address.textContent = jsonResponse.address
    container.appendChild(qrcode);
    container.appendChild(address)
}
