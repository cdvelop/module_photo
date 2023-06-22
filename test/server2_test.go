package test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func Servidor2(t *testing.T) {
	// Manejador para la ruta /enviar
	http.HandleFunc("/enviar", enviarFormulario)

	// Manejador para el resto de rutas (sirve el archivo contacto.html)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Iniciar el servidor en el puerto 1300
	go func() {
		if err := http.ListenAndServe(":1300", nil); err != nil {
			panic(err)
		}
	}()

	// Opciones de inicialización del contexto de chromedp
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Desactivar el modo headless
		chromedp.WindowSize(1920, 1080),
	)

	// Crear el contexto de chromedp con las opciones de inicialización
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Establecer las dimensiones de la ventana del navegador
	// ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf),
	// 	chromedp.EmulateViewport(1024, 768),
	// )

	// ctx, cancel = chromedp.NewContext(ctx,
	// 	chromedp.WithLogf(log.Printf),
	// 	chromedp.EmulateViewport(1024, 768),
	// )

	// defer cancel()

	// Abrir las herramientas de desarrollador
	devtools := chromedp.ActionFunc(func(ctx context.Context) error {
		wsURL := "http://localhost:1300"
		allocCtx, cancel := chromedp.NewRemoteAllocator(ctx, wsURL)
		defer cancel()

		// Abrir las herramientas de desarrollador
		_, err := chromedp.NewContext(allocCtx)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	// Navegar a la página del formulario
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("http://localhost:1300/contacto.html"),
		chromedp.Sleep(2 * time.Second),
		devtools, // Abrir las herramientas de desarrollador
	}); err != nil {
		panic(err)
	}

	// Completar los campos del formulario
	if err := chromedp.Run(ctx, chromedp.SetValue("#nombre", "Juan")); err != nil {
		panic(err)
	}
	if err := chromedp.Run(ctx, chromedp.SetValue("#email", "juan@example.com")); err != nil {
		panic(err)
	}
	if err := chromedp.Run(ctx, chromedp.SetValue("#mensaje", "Este es un mensaje de prueba")); err != nil {
		panic(err)
	}

	// Enviar el formulario
	if err := chromedp.Run(ctx, chromedp.Click("input[type=\"submit\"]")); err != nil {
		panic(err)
	}

	// Esperar a que se cargue la página de confirmación
	chromedp.Sleep(5 * time.Second)

}
