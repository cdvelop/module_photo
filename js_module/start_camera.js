

 

let width = 0;
let height = 0; // Esto se calculará en función del flujo de entrada

let camera_is_closed = true;
// |transmisión| indica si estamos transmitiendo o no actualmente
// video de la cámara. Obviamente, empezamos en false.
// let streaming = false;


let video = module.querySelector('#video');
// let canvas = module.querySelector('#canvas');
let photo = module.querySelector('#photo');


function startupCapture() {
    const compute_style = window.getComputedStyle(form);
    width = parseInt(compute_style.getPropertyValue('width'));
    // if (showViewLiveResultButton()) {
    //     return;
    // }
    btn_add.disabled = true;
    camera_is_closed = false;

    btn_cancel.disabled = false;

    // console.log("BOTÓN INICIO: ", startbutton)

    navigator.mediaDevices
        .getUserMedia({ video: true, audio: false })
        .then((stream) => {
            video.srcObject = stream;
            video.play();
        })
        .catch((err) => {
            console.error(`An error occurred: ${err}`);
        });

    video.addEventListener(
        "canplay",
        (ev) => {
            // if (!streaming) {
                height = video.videoHeight / (video.videoWidth / width);

                // Firefox actualmente tiene un error en el que no se puede leer la altura
                // el video, por lo que haremos suposiciones si esto sucede.

                if (isNaN(height)) {
                    height = width / (4 / 3);
                }

                video.setAttribute("width", width);
                video.setAttribute("height", height);
                // canvas.setAttribute("width", width);
                // canvas.setAttribute("height", height);
                // streaming = true;
            // }
        },
        false
    );

    btn_save.addEventListener("click",takePicture,false);

    btn_save.disabled = false;
    btn_delete.disabled = true;

}



// Visit our article
//   href="https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Taking_still_photos">
