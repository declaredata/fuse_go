package fuse

func Col(name string) Column {
	return &existingCol{name: name}
}
