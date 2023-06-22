
const images_selected = [];


function targetClickPhotoList(e) {

    const clickedElement = e.target;
    if (clickedElement.tagName === 'IMG') { // si el click fue en una imagen
        const checkbox = clickedElement.nextElementSibling;
        const imageId = clickedElement.src.split('=')[1];
        if (checkbox.checked) { // si el checkbox estaba marcado
            const index = images_selected.indexOf(imageId);
            if (index > -1) {
                images_selected.splice(index, 1);
            }
        } else { // si el checkbox no estaba marcado
            if (!images_selected.includes(imageId)) {
                images_selected.push(imageId);
            }
        }
        console.log(images_selected);

        // activar botón de eliminar si 
        if (images_selected.length != 0) {
            btn_delete.disabled = false;
        }else{
            btn_delete.disabled = true;
        }
    }

    console.log(e.target)
}