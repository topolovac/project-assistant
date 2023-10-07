package main

type Directory struct {
	Name        string
	Files       []File
	Directories []Directory
}

type File struct {
	Name    string
	Type    string
	Content string
}
