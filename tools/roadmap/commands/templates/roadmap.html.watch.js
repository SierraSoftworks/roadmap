const eventSource = new EventSource("/watch");
eventSource.onerror = function(event) {
    console.log("Got SSE error: ", event)
}
eventSource.onmessage = function(event) {
    console.log("Got SSE event: ", event)

}

eventSource.addEventListener("reload", _ => {
    console.log("Got reload event")
    window.location.reload()
})