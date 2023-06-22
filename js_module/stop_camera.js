
function stopCameraVideo() {

    DO_YOU_WANT_DELETE_THIS = false;
    btn_add.disabled = false;
    
    if (!camera_is_closed) {
        // console.log('cerrando cámara')
        const stream = video.srcObject;
        const tracks = stream.getTracks();

        tracks.forEach((track) => {
            track.stop();
        });

        video.srcObject = null;
        camera_is_closed = true

        btn_save.disabled = true;
        btn_cancel.disabled = true;
        btn_add.disabled = false



        btn_save.removeEventListener("click", takePicture, false);

    }
}