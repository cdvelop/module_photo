
const btn_add = module.querySelector('button[name="btn_add"]');
const btn_save = module.querySelector('button[name="btn_save"]');
const btn_cancel = module.querySelector('button[name="btn_cancel"]');
const btn_delete = module.querySelector('button[name="btn_delete"]');


const form = module.querySelector('.form-distributed-fields');

if (!navigator.mediaDevices?.enumerateDevices) {
    console.log("enumerateDevices() not supported.");
}

