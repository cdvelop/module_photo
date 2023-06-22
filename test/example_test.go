package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"

	cdpruntime "github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

const look_browser = true

func TestServidor(t *testing.T) {
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

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	if look_browser {
		opts := append(
			// select all the elements after the third element
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false), // Desactivar el modo headless

			// chromedp.DefaultExecAllocatorOptions[3:],
			// chromedp.NoFirstRun,
			// chromedp.NoDefaultBrowserCheck,
			// chromedp.Flag("auto-open-devtools-for-tabs", true),
			// chromedp.Flag("--webview-log-js-console-messages", true),
			chromedp.WindowSize(1080, 720),
		)

		// create chromedp's context
		parentCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()

		ctx, cancel = chromedp.NewContext(parentCtx)
		defer cancel()
	} else {

		// Abrir el navegador y navegar a la página del formulario
		ctx, cancel = chromedp.NewContext(context.Background())
		defer cancel()
	}

	// capturar los logs de JavaScript
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *cdpruntime.EventConsoleAPICalled:
			fmt.Println(">>>> frontend logs")
			for _, arg := range ev.Args {
				s, err := strconv.Unquote(string(arg.Value))
				if err != nil {
					log.Println(err)
					continue
				}
				fmt.Printf("console:%s\n", s)
			}
			fmt.Println("frontend logs <<<<")
		}
	})

	if err := chromedp.Run(ctx, chromedp.Navigate("http://localhost:1300/contacto.html")); err != nil {
		panic(err)
	}

	// Esperar a que se cargue el formulario
	if err := chromedp.Run(ctx, chromedp.Sleep(2*time.Second)); err != nil {
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

	// Esperar a que se cargue la página de confirmación
	if err := chromedp.Run(ctx, chromedp.Sleep(2*time.Second)); err != nil {
		panic(err)
	}

	// Enviar el formulario
	if err := chromedp.Run(ctx, chromedp.Click("input[type=\"submit\"]")); err != nil {
		panic(err)
	}

	fmt.Println("Formulario enviado correctamente!")

	// jsLogs, err := getJavascriptLogs("http://localhost:1300")
	// if err != nil {
	// 	fmt.Printf("Error al obtener los logs de JavaScript: %v", err)
	// } else {
	// 	fmt.Println("Logs de JavaScript:", jsLogs)
	// }

	// Esperar x segundos antes de finalizar la prueba
	time.Sleep(3 * time.Second)
}
