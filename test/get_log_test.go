package test

// "github.com/chromedp/cdproto/runtime"
// "github.com/chromedp/chromedp"

// func getJavascriptLogs(url string) ([]string, error) {
// 	// Crea una nueva instancia de Chrome DevTools
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	// Captura los logs de JavaScript
// 	var jsLogs []string
// 	chromedp.ListenTarget(ctx, func(event interface{}) {
// 		switch ev := event.(type) {
// 		case *runtime.EventConsoleAPICalled:
// 			// Convierte el valor de easyjson.RawMessage a un valor de tipo interface{}
// 			var arg interface{}
// 			if err := json.Unmarshal([]byte(ev.Args[0].Value), &arg); err != nil {
// 				fmt.Printf("Error al convertir el valor de easyjson.RawMessage a un valor de tipo interface{}: %v\n", err)
// 				return
// 			}

// 			// Obtiene el valor de tipo string
// 			s, ok := arg.(string)
// 			if !ok {
// 				fmt.Println("El valor no es de tipo string")
// 				return
// 			}

// 			jsLogs = append(jsLogs, s)
// 		}
// 	})

// 	// Accede a la URL y espera a que se cargue la página
// 	if err := chromedp.Run(ctx, chromedp.Navigate(url), chromedp.WaitReady("body")); err != nil {
// 		return nil, err
// 	}

// 	// Retorna los logs de JavaScript
// 	return jsLogs, nil
// }
