/*

   @author #Kkk
   @File code_server
   @Description:
   @version 0.1
   @date 2022/8/27 10:27

*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	// 实现interface alt+insert implement
	Name string
}

// 遇事不决 用指针.

func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	//TODO implement me
	http.HandleFunc(pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	//TODO implement me
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

func main() {
	server := NewHttpServer("test-server")
	server.Route("/", Home)

	server.Start(":8088")

}

func Home(writer http.ResponseWriter, request *http.Request) {
	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintf(writer, "read request err : %v ", err)
		return
	}
	fmt.Fprintf(writer, "read request body : %s \n ", string(bytes))
}
