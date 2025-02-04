class Event {
  // Each Event needs a Type
  // The payload is not required
  constructor(type, payload) {
    this.type = type;
    this.payload = payload;
  }
}

function routeEvent(event) {
  if (event.type === undefined) {
    alert("no 'type' field in event");
  }
  switch (event.type) {
    case "new_message":
      console.log("new message");
      break;
    default:
      alert("unsupported message type");
      break;
  }
}

function hostGame(event) {
  //Request room from web server that another player can join
}

window.onload = function () {
  // Check if the browser supports WebSocket
  if (window["WebSocket"]) {
    console.log("supports websockets");
    // Connect to websocket
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    //read anything sent to us from the web socket server
    conn.onmessage = function (evt) {
      console.log(evt);
      const eventData = JSON.parse(evt.data);
      //console.log(eventData)
      const event = Object.assign(new Event(), eventData);
      routeEvent(event);
    };
  } else {
    alert("Not supporting websockets");
  }
};

//event listener that will send connection to socket
document.addEventListener(
  "keydown",
  function (event) {
    switch (event.key) {
      case "ArrowLeft":
        event.preventDefault(); // Stops page scrolling
        sendEvent("move", "ArrowLeft");
        break;
      case "ArrowRight":
        event.preventDefault();
        sendEvent("move", "ArrowRight");
        break;
      case "ArrowUp":
        event.preventDefault();
        sendEvent("move", "ArrowUp");
        break;
      case "ArrowDown":
        event.preventDefault();
        sendEvent("move", "ArrowDown");
        break;
    }
  },
  true
);

const divElement = document.querySelector("");

function createRoom() {
  console.log("hello world");
}

function sendEvent(eventName, payload) {
  // Create a event Object with a event named send_message
  const event = new Event(eventName, payload);
  // Format as JSON and send
  conn.send(JSON.stringify(event));
}
