package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)


func main(){
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
    
	for i:=1; i <= 10000;i++{
        wg.Add(1)		
		go getSimulated(ctx,wg,i)
	}
	for k := 1; k <=10000;k++{
		wg.Add(1)
		go postSimulated(ctx,wg,k)
	}
	for j := 1; j <= 10000; j++{
		wg.Add(1)
		go putSimulated(ctx,wg,j)
	}
	for n := 1; n <= 10000; n++{
		wg.Add(1)
		go deleteSimulated(ctx,wg,n)
	}

	wg.Wait()

	
}

//simulate GET across all paths
func getSimulated(ctx context.Context,wg *sync.WaitGroup,id int){
	defer wg.Done()
	
	urlPaths := []string{
		"http://localhost:8443/order",
		"http://localhost:8443/order/orderId",
		"http://localhost:8443/payment",
		"http://localhost:8443/user",
		
	}
  
   for _,k := range urlPaths{
      req, err := http.NewRequestWithContext(ctx,http.MethodGet,k,nil)
      if err != nil{
      log.Println(err)
      continue
      }
      resp, err := http.DefaultClient.Do(req)
      if err != nil{
      log.Println(err)
      continue
      }
      time.Sleep(2 * time.Second)
      data, err := io.ReadAll(resp.Body)
      if err != nil{
      log.Println(err)
      }
     
      fmt.Println(id,resp.StatusCode,string(data))
   }
  
}

//simulate POST across all paths
func postSimulated(ctx context.Context,wg *sync.WaitGroup,id int){
	defer wg.Done()
	
	urlPaths := []string{
		"http://localhost:8443/order",
		"http://localhost:8443/order/cancel",
		"http://localhost:8443/payment",
		"http://localhost:8443/user",
	}
	
	for _, k := range urlPaths{
		req, err := http.NewRequestWithContext(ctx,http.MethodPost,k,nil)
		if err != nil{
			log.Println(err)
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil{
			log.Println(err)
			continue
		}
		time.Sleep(1*time.Second)
	 data, err := io.ReadAll(resp.Body)
   if err != nil{
   log.Println(err)
   }

   fmt.Println(id,resp.StatusCode, string(data))
	}

  
}

//simulate PUT across all paths
func putSimulated(ctx context.Context,wg *sync.WaitGroup,id int){
	defer wg.Done()
	
	urlPaths := []string{
		"http://localhost:8443/order",
		"http://localhost:8443/payment",
		"http://localhost:8443/user",
	}
	
	for _, k := range urlPaths{
		req, err := http.NewRequestWithContext(ctx,http.MethodPut,k,nil)
		if err != nil{
			log.Println(err)
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil{
			log.Println(err)
			continue
		}
		time.Sleep(1*time.Second)
   data, err := io.ReadAll(resp.Body)
   if err != nil{
   log.Println(err)
   }
   
   fmt.Println(id,resp.StatusCode, string(data))
	}

  
}

//simulate DELETE across all the paths
func deleteSimulated(ctx context.Context,wg *sync.WaitGroup,id int){
	defer wg.Done()
	
	urlPaths := []string{
		"http://localhost:8443/order",
		"http://localhost:8443/user",
		}
	
	for _, k := range urlPaths{
		req, err := http.NewRequestWithContext(ctx,http.MethodDelete,k,nil)
		if err != nil{
			log.Println(err)
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil{
			log.Println(err)
			continue
		}
		time.Sleep(1*time.Second)
   data, err := io.ReadAll(resp.Body)
   if err != nil{
   log.Println(err)
   }
  
   fmt.Println(id,resp.StatusCode, string(data))
	}

  
}