package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestSimpleHTTPServer(t *testing.T) {
	srv := &http.Server{
		Addr: "127.0.0.1:8443",
		Handler: http.TimeoutHandler(DefaultHandler(),2 *time.Minute,""),
		IdleTimeout: 5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil{
		t.Fatal(err)
	}

	go func() {
		err := srv.Serve(l)
		if err != http.ErrServerClosed{
			t.Error(err)
		}
	}()

   testCases := []struct{
     method string
     body io.Reader
     code  int
     response string
   }{
   //GET tests
   {http.MethodGet,nil,http.StatusFound,"get/order"},
   {http.MethodGet,nil,http.StatusFound,"get/order/orderId"},
   {http.MethodGet,nil,http.StatusFound,"get/payment"},
   {http.MethodGet,nil,http.StatusFound,"get/user"},
   //POST tests
   {http.MethodPost,nil,http.StatusCreated,"post/order"},
   {http.MethodPost,nil,http.StatusCreated,"post/order/cancel"},
   {http.MethodPost,nil,http.StatusCreated,"post/payment"},
   {http.MethodPost,nil,http.StatusCreated,"post/user"},
   //PUT tests
   {http.MethodPut,nil,http.StatusFound,"put/order"},
   {http.MethodPut,nil,http.StatusFound,"put/payment"},
   {http.MethodPut,nil,http.StatusFound,"put/user"},
   //DELETE tests
   {http.MethodDelete,nil,http.StatusAccepted,"delete/order"},
   {http.MethodDelete,nil,http.StatusAccepted,"delete/user"},
   
   
   }
   

   client := new(http.Client)
   //GET url
   path1 := fmt.Sprintf("http://%s/orders", srv.Addr)
   path2 := fmt.Sprintf("http://%s/order/orderId",srv.Addr)
   path3 := fmt.Sprintf("http://%s/order/payment",srv.Addr)
   path4 := fmt.Sprintf("http://%s/get/user",srv.Addr)
   //POST url
   path5 := fmt.Sprintf("http://%s/order", srv.Addr)
   path6 := fmt.Sprintf("http://%s/order/cancel", srv.Addr)
   path7 := fmt.Sprintf("http://%s/order/pay", srv.Addr)
   path8 := fmt.Sprintf("http://%s/create/user", srv.Addr)
   //PUT url
   path9 := fmt.Sprintf("http://%s/order", srv.Addr)
   path10 := fmt.Sprintf("http://%s/order/payment", srv.Addr)
   path11 := fmt.Sprintf("http://%s/update/user", srv.Addr)
   //DELETE url
   path12 := fmt.Sprintf("http://%s/order/orderId", srv.Addr)
   path13 := fmt.Sprintf("http://%s/delete/user", srv.Addr)
   path := []string{path1,path2,path3,path4,path5,path6,path7,path8,path9,path10,path11,path12,path13}
  
   
   for i, c := range testCases{
     testPath := path[i]    
   r, err := http.NewRequest(c.method,testPath,c.body)
   if err != nil{
   t.Errorf("%d: %v",i, err)
   continue
   }

   resp, err := client.Do(r)
   if err != nil{
   t.Errorf("%d: %v",i, err)
   continue
   }

   if resp.StatusCode !=  c.code{
    t.Errorf("%d unexpected status code: %q", i, resp.Status)
   }
  
  data , err := io.ReadAll(resp.Body)
     if err != nil{
     t.Errorf("%d: %v",i, err)
     continue
     }
     _ = resp.Body.Close()

     if c.response != string(data){
     t.Errorf("%d: expected %q; actual %q",i, c.response,data)
     }
   }
   
	
}