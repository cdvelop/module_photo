package module_photo

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/platform"
)

type module struct {
}

func Add() *model.Module {

	in := module{}

	m := model.Module{
		Theme:      platform.Theme{},
		ModuleName: "photo",
		Title:      "Captura de Imágenes",
		IconID:     "icon-camera",
		UI:         in,
		Areas:      []byte{},
		Objects:    []*model.Object{},
	}

	return &m
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
