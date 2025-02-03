const canvas = document.getElementById("gameCanvas");
const ctx = canvas.getContext("2d");

const ws = new WebSocket("ws://localhost:8080/ws");
let gameObjects = [];
let playerID = null;
let gameStarted = false;

// WebSocket connection handlers
ws.onopen = () => {
    console.log("Connected to game server");
};

ws.onmessage = (event) => {
    const data = JSON.parse(event.data);

    // Check if this is a player assignment message
    if (data.player_id !== undefined) {
        playerID = data.player_id;
        console.log(`Assigned Player ID: ${playerID}`);
        
        if (playerID !== null) {
            updateConnectionStatus();
        }
        return;
    }

    // Regular game state update
    if (data.objects) {
        gameObjects = data.objects;
        draw();
    }
};

ws.onerror = (error) => {
    console.error("WebSocket Error:", error);
};

ws.onclose = () => {
    console.log("Disconnected from game server");
};

// Input handling
function sendInput(action) {
    if (playerID !== null && gameStarted) {
        ws.send(JSON.stringify({ 
            player_id: playerID, 
            action: action 
        }));
    }
}

document.addEventListener("keydown", (event) => {
    switch (event.key) {
        case "ArrowUp": 
        case "w":
            sendInput("up"); 
            break;
        case "ArrowDown": 
        case "s":
            sendInput("down"); 
            break;
        case "ArrowLeft": 
        case "a":
            sendInput("left"); 
            break;
        case "ArrowRight": 
        case "d":
            sendInput("right"); 
            break;
    }
});

// Rendering function
function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    
    // Draw background elements
    ctx.fillStyle = "lightgreen";
    ctx.fillRect(0, canvas.height - 20, canvas.width, 20); // Ground
    
    // Draw game objects
    gameObjects.forEach((obj, index) => {
        ctx.fillStyle = index === 0 ? "blue" : "red";
        
        // Assuming Shape has MinX, MinY, and potentially width/height
        const shape = obj.Shape;
        ctx.fillRect(
            shape.MinX, 
            shape.MinY, 
            shape.Width || 15, 
            shape.Height || 15
        );
    });

    // Display player status
    ctx.fillStyle = "black";
    ctx.font = "12px Arial";
    ctx.fillText(`Player: ${playerID !== null ? playerID : "Not Connected"}`, 10, 20);
    ctx.fillText(`Game Status: ${gameStarted ? "Active" : "Waiting"}`, 10, 40);
}

// Connection status update
function updateConnectionStatus() {
    if (playerID !== null) {
        draw(); // Initial draw to show player info
    }
}

// Initial draw to set up canvas
draw();