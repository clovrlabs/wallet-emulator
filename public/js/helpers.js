function b64toBlob(base64, type = 'application/octet-stream') {
    return fetch(`data:${type};base64,${base64}`).then(res => res.blob())
}

window.b64toBlob = b64toBlob