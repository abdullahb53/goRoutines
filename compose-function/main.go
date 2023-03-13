package main

import "fmt"

type TransformFunc func(string) string

type Server struct {
	filenameTransformFunc TransformFunc
}

func (s *Server) handleRequest(filename string) error {
	newFilename := s.filenameTransformFunc(filename)
	fmt.Println("new filename: ", newFilename)

	return nil
}

func PrefixFileName(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + "@" + filename
	}
}

func main() {

	s := &Server{
		// Set once to prefix.
		filenameTransformFunc: PrefixFileName("Prefix"),
	}

	s.handleRequest("cool_picture.jpg")

}
