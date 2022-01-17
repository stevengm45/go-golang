package CodigoFacilito

func (self *Curso) getTitulo() string { // el metodo es privado por que el nombre de la func comienza en minuscula
	return self.Titulo
}

func (self *Curso) ObtenerTitulo() string { // el emtodo es publico por que el nombre de la func comienza en mayuscula
	return self.getTitulo()
}
