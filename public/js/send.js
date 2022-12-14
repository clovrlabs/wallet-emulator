let container = document.getElementById("container");
let addressInputField = document.getElementById("address-field");
let amountField = document.getElementById("amount-field");
let sendButton = document.getElementById("send-button");

const urlSearchParams = new URLSearchParams(window.location.search);
const params = Object.fromEntries(urlSearchParams.entries());

let invoice = true;
function addAmountField() {
    if (params.type == "btc") {
        amountField.style.display = "inline";
        invoice = false
        return;
    }
    if (addressInputField?.value?.toLowerCase().startsWith("lnbcrt")) {
        amountField.style.display = "none";
        invoice = true;
    } else {
        amountField.style.display = "inline";
        invoice = false;
    }
}

let title = document.getElementById("title");
if (params.type == "invoice") {
    title.textContent = "Paste Invoice or ID";
    addressInputField.placeholder = "Input the address or id to send";
} else {
    title.textContent = "Send to BTC address";
    addressInputField.placeholder = "Input the btc address";
}

function clearPage() {
    while (container.firstChild) {
        container.removeChild(container.lastChild);
    }
}

async function sendPayment() {
    if (!addressInputField.value) {
        return alert("Address is mandatory");
    }
    if (!amountField.value && !invoice) {
        return alert("Amount is mandatory");
    }
    
    clearPage();
    let response = await fetch(`/send/${params.type}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            PaymentRequest: addressInputField.value,
            AmountSatoshi: invoice ? 0 : parseInt(amountField.value || "0"),
            Fee: -1
        })
    });

    let jsonResponse = await response.json();

    if (jsonResponse.error) {
        let error = document.createElement("div");
        error.textContent = jsonResponse.error;
        container.appendChild(error);
    } else {
        history.back();
    }
}

window.onpageshow = function () {
    if (params.type == "btc") {
        amountField.style.display = "inline";
        return;
    }
    amountField.style.display = "none";
};
