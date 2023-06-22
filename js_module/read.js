
const photo_list = module.querySelector('.container-list-only')

// readPhotos()

function readPhotos(params) {
    // console.log("leyendo imágenes")

    NewRequest({
        method: "GET",
        contentType: "json",
        endpoint: "/file?folder_id=11111",
        response: Response
    });

}




function Response(data) {
    console.log(data);
    // console.log(`Total archivos encontrados: ${data.Data.length}`);
    const imagesHTML = createImagesHTML(data,true);
    photo_list.innerHTML += imagesHTML;

}


function createImagesHTML(data, append = true) {
    let html = "";
  
    data.Data.forEach((image) => {
      html += `<div class="col-img">
      <div class="checkbox-container">
        <label for="${image.id_file}">
          <img src="file?id_file=${image.id_file}" class="img-list">
          <input type="checkbox"  id="${image.id_file}">
        </label>
      </div>
    </div>`;
    });
  
    if (append) {
        return html;
      } else {
        return html.split("").reverse().join("");
      }

  }




