//backend
package main 
 import(
"fmt"
"os"
"log"
"net/https"
"database/sql"


 
func main ()
//reading file
 entries, err := os.Open("./home/dried/Documents/drieds-web-repos/repo-folder/hippos/")
 
 if err != nil {
	 log.Fatal(err)
 }

 for _, e := range entries {
		 fmt.Println(e.Name())
 }



) 