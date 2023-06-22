
const canvas = module.querySelector('#canvas');
canvas.style.display = 'none';

function takePicture(e) {

  const context = canvas.getContext("2d");

  // Aumentar la resolución del canvas
  canvas.width = 800;
  canvas.height = 600;
  // canvas.width = video.videoWidth * 2;
  // canvas.height = video.videoHeight * 2;

  // Obtener la imagen del objeto de video
  // context.drawImage(video, 0, 0, width, height);
  context.drawImage(video, 0, 0, canvas.width, canvas.height)

  // Obtener la imagen reducida como objeto Blob
  canvas.toBlob((blob) => {
    // Obtener la extensión de la imagen
    const extension = blob.type.split("/")[1];

    // Headers de la petición
    const headers = new Headers();
    headers.append("Action-Type", "create");
    headers.append("module_name", "medicalhistory");
    headers.append("field_name", "endoscopia");
    headers.append("folder_id", "11111");
    headers.append("description", "none");

    // Objeto con los datos de la petición
    const formData = new FormData();
    formData.append("endoscopia", blob, `endoscopia.${extension}`);

    // Opciones de la petición
    const options = {
      method: "POST",
      headers: headers,
      body: formData,
    };

    // Enviar la petición
    fetch("/file", options)
      .then((response) => response.json())
      .then((data) => {
        data.Message = "Nueva Imagen Capturada";
        ShowMessageType(data);
        Response(data);
        
      })
      .catch((err) => {
        ShowMessageType({"Type":"error","Message":"No hay conexión con el servidor ("+err.toString()+")"});
        stopCameraVideo()
      });
  }, "image/jpeg", .9);


  context.clearRect(0, 0, canvas.width, canvas.height);

}


