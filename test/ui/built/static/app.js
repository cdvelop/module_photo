'use strict';
const retryDelay = 1000; // tiempo de espera para volver a intentar la conexión
const maxRetries = 10; // número máximo de reintentos

let eventSource;
let retryCount = 0;

function connect() {
    eventSource = new EventSource('/reload');
    eventSource.onmessage = function(event) {
        // console.log(event.data);
        if (event.data === 'Reload') {
            // console.log("Reload message received. Reloading page...");
            location.reload();
        }
    };
    eventSource.onerror = function(error) {
        console.error(error);
        eventSource.close();
        reconnect();
    };
    eventSource.addEventListener('open', function(event) {
        console.log('Connected to server');
    });
}

function reconnect() {
    if (retryCount < maxRetries) {
        retryCount++;
        console.log(`Disconnected from server. Retrying in ${retryDelay / 1000} seconds...`);
        setTimeout(connect, retryDelay);
    } else {
        console.log('Max retries reached. Giving up...');
    }
}

connect();
