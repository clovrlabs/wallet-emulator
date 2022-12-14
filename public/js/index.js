let bottomPanel = document.getElementById("bottom-panel");
let popup = document.getElementById("popup");
let drawer = document.getElementById("drawer")
let balance = document.getElementById("balance")

/* BOTTOM PANEL */
let closeBottomPanel = document.createElement("button");
closeBottomPanel.textContent = "Cancel";
closeBottomPanel.onclick = function clearBottomPanel() {
    bottomPanel.style.visibility = "hidden"
    while (bottomPanel.firstChild) {
        bottomPanel.removeChild(bottomPanel.lastChild);
    }
};

// Receive bottom panel
let receiveViaInvoiceButton = document.createElement("button");
receiveViaInvoiceButton.textContent = "Receive via invoice";
receiveViaInvoiceButton.onclick = function() {
    let titleParams = new URLSearchParams({ type: "invoice" })
    window.location.href = `/receive?${titleParams.toString()}`
}
let receiveViaBtcAddressButton = document.createElement("button");
receiveViaBtcAddressButton.textContent = "Receive via BTC address";
receiveViaBtcAddressButton.onclick = function() {
    let titleParams = new URLSearchParams({ type: "btc" })
    window.location.href = `/receive?${titleParams.toString()}`
}

function receive() {
    bottomPanel.appendChild(receiveViaInvoiceButton);
    bottomPanel.appendChild(receiveViaBtcAddressButton);
    bottomPanel.appendChild(closeBottomPanel);
    bottomPanel.style.visibility = "visible"
}

// Send bottom panel
let sendViaInvoiceOrIdButton = document.createElement("button");
sendViaInvoiceOrIdButton.textContent = "Paste Invoice or ID";
sendViaInvoiceOrIdButton.onclick = function() {
    let titleParams = new URLSearchParams({ type: "invoice" })
    window.location.href = `/send?${titleParams.toString()}`
}
let sendToBtcAddressButton = document.createElement("button");
sendToBtcAddressButton.textContent = "Send to BTC address";
sendToBtcAddressButton.onclick = function() {
    let titleParams = new URLSearchParams({ type: "btc" })
    window.location.href = `/send?${titleParams.toString()}`
}

function send() {
    bottomPanel.appendChild(sendViaInvoiceOrIdButton)
    bottomPanel.appendChild(sendToBtcAddressButton)
    bottomPanel.appendChild(closeBottomPanel)
    bottomPanel.style.visibility = "visible"
}
/* END BOTTOM PANEL */


/* DRAWER */
let openDrawerButton= document.getElementById("drawer-button");
openDrawerButton.onclick = function openDrawer() {
    drawer.appendChild(openDevelopersButton)
    drawer.appendChild(closeDrawerButton)
    drawer.style.visibility = "visible"
}

let closeDrawerButton = document.createElement("button");
closeDrawerButton.textContent = "Cancel";
closeDrawerButton.onclick = function closeDrawer() {
    drawer.style.visibility = "hidden"
    while (drawer.firstChild) {
        drawer.removeChild(drawer.lastChild);
    }
};

let openDevelopersButton = document.createElement("button");
openDevelopersButton.textContent = "Developers";
openDevelopersButton.onclick = function() {
    window.location.href = `/developers`
}
/* END DRAWER */

async function getAccountInfo() {
    let response = await fetch("/account/info", {
        method: 'GET',
    }); 
    let account = await response.json()
    balance.textContent = account.balance + " sats"
}

window.onpageshow = async function() {
    console.log("updating balance");
    await getAccountInfo()
}
