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

var (
wg = &sync.WaitGroup{}
pool1, pool2 int
inPool = make(chan int,10000)
outPool = make(chan int,10000)
)
type DataValue struct{
	in,
	out int
	
}
type counter struct{
    reqId,
	respId int
	 sync.Mutex
}
func main(){

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := &counter{}
	//send out 4-goroutines at once in each iteration
	// so that 4-hits for GET,4-hits for POST,3-hits for PUT,2-hits for DELETE 
	// so that's 13 requests per iteration

	for i:=1;i<=100;i++{
		wg.Add(4)
		go c.getSimulated(ctx,wg)
		go c.postSimulated(ctx,wg)
		go c.putSimulated(ctx,wg)
		go c.deleteSimulated(ctx,wg)
		
	}

   wg.Wait()
  
  
 	
  
  
   avg := (c.respId/c.reqId)*100
   fmt.Println(c.reqId)
   fmt.Println(c.respId)
   fmt.Printf("success rate :%v",avg)

	

	
}


//simulate GET across all paths
func (c *counter)getSimulated(ctx context.Context,wg *sync.WaitGroup){
	defer func(){
		 wg.Done()
	}()
	
	urlPaths := []string{
		"http://localhost:8443/orders",
		"http://localhost:8443/order/orderId",
		"http://localhost:8443/order/payment",
		"http://localhost:8443/get/user",
		
	}
  
   for _,k := range urlPaths{
      req, err := http.NewRequestWithContext(ctx,http.MethodGet,k,nil)
      if err != nil{
      log.Println(err)
      continue
      }
      c.Lock()
      c.reqId++
      c.Unlock()
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
      
       fmt.Println(resp.StatusCode, string(data))
       c.Lock()
       c.respId++
       c.Unlock()
    
   }
     
  
}

//simulate POST across all paths
func (c *counter)postSimulated(ctx context.Context,wg *sync.WaitGroup){
	
	defer func(){
			wg.Done()
	}()
	
	urlPaths := []string{
		"http://localhost:8443/order",
		"http://localhost:8443/order/cancel",
		"http://localhost:8443/order/pay",
		"http://localhost:8443/create/user",
	}
	
	for _, k := range urlPaths{
		req, err := http.NewRequestWithContext(ctx,http.MethodPost,k,nil)
		if err != nil{
		//	log.Println(color.Style{color.BgBlack,color.FgHiRed,color.OpConcealed}.Sprint(err))
		    log.Println(err)
			continue
		}
		c.Lock()
		c.reqId++
		c.Unlock()
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
   
   fmt.Println(resp.StatusCode, string(data))
   c.Lock()
   c.respId++
   c.Unlock()
	}
}


  


//simulate PUT across all paths
func (c *counter)putSimulated(ctx context.Context,wg *sync.WaitGroup){
	
	defer func(){
			wg.Done()
	}()
	
	urlPaths := []string{
		"http://localhost:8443/order",
		"http://localhost:8443/order/payment",
		"http://localhost:8443/update/user",
	}
	
	for _, k := range urlPaths{
		req, err := http.NewRequestWithContext(ctx,http.MethodPut,k,nil)
		if err != nil{
			log.Println(err)
			continue
		}
		c.Lock()
		c.reqId++
		c.Unlock()
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
     
   fmt.Println(resp.StatusCode, string(data))
   c.Lock()
   c.respId++
   c.Unlock()
	}



  
}

//simulate DELETE across all the paths
func (c *counter)deleteSimulated(ctx context.Context,wg *sync.WaitGroup){

	defer func(){
	
		wg.Done()		
		}()
	
	urlPaths := []string{
		"http://localhost:8443/order/orderId",
		"http://localhost:8443/delete/user",
		}
	
	for _, k := range urlPaths{
		req, err := http.NewRequestWithContext(ctx,http.MethodDelete,k,nil)
		if err != nil{
			log.Println(err)
			continue
		}
	    c.Lock()
	    c.reqId++
		c.Unlock()
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
      
    fmt.Println(resp.StatusCode, string(data))
   c.Lock()
   c.respId++
   c.Unlock()
    
	}



  
}