let commanInputField = document.getElementById("command-field");
commanInputField.addEventListener("keypress", function(event) {
    if (event.key === "Enter") {
      event.preventDefault();
      sendCommand();
    }
  });
let commandResponse = document.getElementById("json");


async function sendCommand() {
    if (!commanInputField.value) {
        return alert("Command is mandatory");
    }
    let response = await fetch("/command", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            Command: commanInputField.value
        })
    });

    let jsonResponse = await response.json();
    commandResponse.innerHTML = jsonResponse.command
}