
let controllerIndex = null;
let previousButtonState = [];

window.addEventListener("gamepadconnected", (event) => {
    const gamepad = event.gamepad;
    controllerIndex = gamepad.index;
    console.log("gamepad connected");
});

window.addEventListener("gamepaddisconnected", (event) => {
    controllerIndex = null;
    console.log("gamepad disconnected");
});

function gameLoop() {
    if (controllerIndex !== null) {
        const gamepad = navigator.getGamepads()[controllerIndex];
        handleButtons(gamepad.buttons);
    }
    requestAnimationFrame(gameLoop);
}


function handleButtons(buttons) {
    buttons.forEach((button, index) => {
        if (button.pressed && !previousButtonState[index]) {
            // console.log(`TOMANDO FOTO CON MANDO ${index}`);
            // console.log("ESTADO CÁMARA ESTA CERRADA?:",camera_is_closed)
            turnOnCamera()
        }
        previousButtonState[index] = button.pressed;
    });

}
    gameLoop();


    function turnOnCamera() {
        if (camera_is_closed) {
            startupCapture()
        }else{
            takePicture();
        }

    }