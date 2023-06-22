

let DO_YOU_WANT_DELETE_THIS = false;
let MESSAGE_TO_DELETE = ''

console.log("COLOR SELECCIONADO: ",__col_select); // imprime "red"

function ActionButtonDELETE() {
  // console.log("TARGET A ELIMINAR: ",MAIN_TARGET_SELECTED.OBJECT," LEFT DATA: ",left_data," TEXTO: ",left_text)

  
  if (DO_YOU_WANT_DELETE_THIS) {
   
    deleteImages()
    DO_YOU_WANT_DELETE_THIS = false;

    btn_delete.setAttribute('disabled', 'true');
		btn_delete.style.removeProperty('background-color');
    btn_cancel.disabled = true;
    btn_add.disabled = false;
  } else {

    ShowMessageType({"Type":"del","Message":"Deseas eliminar de la base de datos la(s) imagen(es) Seleccionada(s)?"})

    // MSJ.innerHTML = '<H4 class="del">¿' + MESSAGE_TO_DELETE + '?</H4>';
    DO_YOU_WANT_DELETE_THIS = true;
    btn_delete.style.backgroundColor = __col_select;
    btn_add.disabled = true;
    btn_cancel.disabled = false;

  }
};



function deleteImages() {

  // realizar la solicitud de eliminación a l servidor


    for (let i = 0; i < images_selected.length; i++) {
        const id = images_selected[i];
        const img = module.querySelector(`img[src="file?id_file=${id}"]`);
        if (img) {
          const div = img.parentNode.parentNode.parentNode;

          console.log("DIV A ELIMINAR: ",div)

          div.parentNode.removeChild(div);
        }
      }

    //   vaciar el array de ids
      images_selected.splice(0, images_selected.length);
}
