package print

import "fmt"

func P(v interface{},header ... string){

	if len(header) > 0 {
		fmt.Printf("%s: %+v\n",header[0],v)
	}else{
		fmt.Printf("%+v\n",v)
	}
}
func PA(v interface{}){
	fmt.Printf("%#v\n",v)
}