package photo

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/platform"
)

type module struct {
}

var Add = model.Module{
	Theme: platform.Theme{},
	Name:  "photo",
	Title: "Captura de Imágenes",
	Icon: model.Icon{
		Id:      "icon-camera",
		ViewBox: "0 0 16 16",
		Paths:   []string{"M4.75 9.5c0 1.795 1.455 3.25 3.25 3.25s3.25-1.455 3.25-3.25-1.455-3.25-3.25-3.25-3.25 1.455-3.25 3.25zM15 4h-3.5c-0.25-1-0.5-2-1.5-2h-4c-1 0-1.25 1-1.5 2h-3.5c-0.55 0-1 0.45-1 1v9c0 0.55 0.45 1 1 1h14c0.55 0 1-0.45 1-1v-9c0-0.55-0.45-1-1-1zM8 13.938c-2.451 0-4.438-1.987-4.438-4.438s1.987-4.438 4.438-4.438c2.451 0 4.438 1.987 4.438 4.438s-1.987 4.438-4.438 4.438zM15 7h-2v-1h2v1z"},
	},
	UI:         module{},
	Areas:      []byte{},
	Components: []model.Component{},
	Objects:    []model.Object{},
}

func (m module) UserInterface(options ...string) string {

	// data, err := os.ReadFile("modules/" + m.Name + "/index.html")
	// if err != nil {
	// 	fmt.Println("Error al leer el archivo:", err)
	// }

	// return string(data)

	// return "<h1>Acceso " + store.NameArea(area_access) + " no Habilitado en modulo: " + m.MainName() + "</h1>"
	return "MODULO IMÁGENES"

}
