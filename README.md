
# Explore Golang features [![Go Report Card](https://goreportcard.com/badge/zelenko/go)](https://goreportcard.com/report/zelenko/go)
This repository is viewable on [sourcegraph.com](https://sourcegraph.com/github.com/zelenko/go).

![Go](https://raw.githubusercontent.com/zelenko/go/master/37_html_template/pub/img/toby.jpg)

## Contents

| Name | Description |
| :---- | :---- |
| 00_data_type                | type, error, recursion, reference, slice, switch, pointer    |
| 01_concurrency_channel      | channel, waiting, concurrency, sleep, close, count, queue |
| 02_mongodb_aggregate_cli    | aggregate data from mongodb displayed in CLI |
| 03_mongodb_find_sort_cli    | MongoDB Find All, Sort commands, results in CLI |
| 04_get_url_variable         | Get "FormValue" variable from URL |
| 05_mongodb_crud_cli         | Insert, Update, Drop, Find, Index commands, results in CLI |
| 06_ajax_send_receive        | Two Ajax examples: read and write |
| 07_one_page_template        | GO code and HTML template in one file |
| 08_gowiki                   | Simple wiki example |
| 09_todo_list_html           | Todo list with struct and HTML template. |
| 10_mongodb                  | Mongodb query results on HTML page in one go file. |
| 11_socket_send_receive      | Send/receive text between client/server via socket. |
| 12_mgo_pipeline             | Mongodb pipeline query saved in one go file |
| 13_qr_barcode               | QR code generator displays PNG image in browser. |
| 14_read_txt_file            | compare ways to process TXT file |
| 15_png_or_svg_barchart      | generate bar chart, PNG or SVG |
| 16_drop_down_menu_form      | Form /w drop down menu, template with sub-templates |
| 17_mysql_user_login         | MySQL for user registration, login, and user list. |
| 18_cookie_authentication    | Authentication with cookies. |
| 19_importing_packages       | Importing functions from other packages. |
| 20_mongodb_crud             | Mongodb CRUD with REST using httprouter & HTML templates |
| 21_httprouter_template      | html template with httprouter and ServeFiles |
| 22_mongodb_crud_rest_html   | Mongodb CRUD with REST using httprouter & HTML templates |
| 23_file_uploader            | Upload file and save on server |
| 24_https_and_redirect       | Serve HTTP and HTTPS with redirect |
| 25_https_static_files       | Serve HTTP and HTTPS w/ NotFound for static files |
| 26_url_not_found_handler    | Custom Not Found handler. |
| 27_mongodb_bulk_upsert      | Mongodb bulk insert from TXT file |
| 28_markdown                 | Generate Markdown using blackfriday |
| 29_mongodb_crud_json_api    | REST API using JSON with httprouter |
| 30_mongodb_crud_json_api    | REST API using JSON, httprouter, and toml |
| 31_send_email               | Send email with attachment |
| 32_colorful_cli             | Create colorful CLI |
| 33_goroutine_waitgroup      | wait for goroutines to finish, run code once |
| 34_html_template_function   | Pass function to HTML template |
| 35_mongodb_pipeline_page    | One page MGO aggregation with pipeline |
| 36_json_html_javascript     | Display JSON on HTML page |
| 37_html_template            | Simple html template with Go |
|                             | |

## The proper way to copy a slice
```GO
package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	e := make([]string, len(a))
	copy(e, a)
	fmt.Println(e)
}
```
## How to find out the data type?
```GO
// Figure out what type it is: maps, slices, or arrays!

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Declaring local variables
	map1 := map[string]string{"name": "John", "desc": "Golang"}
	map2 := map[string]int{"apple": 23, "tomato": 13}
	slice1 := []int{1, 2, 3}
	array1 := [3]int{1, 2, 3}
	// var m map[string]int
	// m = make(map[string]int)
	// More info Here: https://blog.golang.org/go-maps-in-action

	// Type, such as map[string]string, []int, [3]int
	fmt.Println("map1:", reflect.TypeOf(map1))
	fmt.Println("map2:", reflect.TypeOf(map2))
	fmt.Println("slice1:", reflect.TypeOf(slice1))
	fmt.Println("array1:", reflect.TypeOf(array1))

	// Value, such as map, slice, array.
	fmt.Println("map1:", reflect.ValueOf(map1).Kind())
	fmt.Println("map2:", reflect.ValueOf(map2).Kind())
	fmt.Println("slice1:", reflect.ValueOf(slice1).Kind())
	fmt.Println("array1:", reflect.ValueOf(array1).Kind())

	// True/False statement inside Printf
	fmt.Printf("%v is a map? %v\n", map1, reflect.ValueOf(map1).Kind() == reflect.Map)
	fmt.Printf("%v is a map? %v\n", map2, reflect.ValueOf(map2).Kind() == reflect.Map)
	fmt.Printf("%v is a map? %v\n", slice1, reflect.ValueOf(slice1).Kind() == reflect.Map)

	/*  More about reflect package: https://golang.org/pkg/reflect/
		 	Invalid Kind = iota
	        Bool
	        Int
	        Int8
	        Int16
	        Int32
	        Int64
	        Uint
	        Uint8
	        Uint16
	        Uint32
	        Uint64
	        Uintptr
	        Float32
	        Float64
	        Complex64
	        Complex128
	        Array
	        Chan
	        Func
	        Interface
	        Map
	        Ptr
	        Slice
	        String
	        Struct
	        UnsafePointer
	*/
}
```
![Go](http://farm4.staticflickr.com/3774/11740822616_e435d02a54_o.gif)

### Go Programs and Apps
* **Go Programs and Apps"** [http://go-lang.cat-v.org/go-code](http://go-lang.cat-v.org/go-code)
* **Who is using Go?** https://github.com/golang/go/wiki/GoUsers
* **Stream music using Go:** [nf/goplayer](https://github.com/nf/goplayer/blob/master/player.go)
* **Reading Excel files:** [https://github.com/tealeg/xlsx](https://github.com/tealeg/xlsx)
*  **Generate PDF:** [https://github.com/signintech/gopdf](https://github.com/signintech/gopdf)
* **RSS:** [https://github.com/mmcdole/gofeed](https://github.com/mmcdole/gofeed)
* **File Upload:** [https://github.com/astaxie/build-web-application-with-golang/blob/master/en/04.5.md](https://github.com/astaxie/build-web-application-with-golang/blob/master/en/04.5.md)
* **LDAP:** [https://github.com/mmitton/ldap](https://github.com/mmitton/ldap)
* **Web Frameworks:** [https://golanglibs.com/category/web-framework?sort=top](https://golanglibs.com/category/web-framework?sort=top)
* **Go syntax and features:** [https://github.com/a8m/go-lang-cheat-sheet](https://github.com/a8m/go-lang-cheat-sheet)

# Idiomatic Go
Use _established conventions_ for programming in Go, such as **naming**, **formatting**, **program construction**, and so on, so that programs you write will be _easy_ for other Go programmers to understand.
* https://about.sourcegraph.com/go/idiomatic-go/
* http://idiomaticgo.com/post/testing/idiomatic-go-tests/
* https://pocketgophers.com/idiomatic-go/
* https://golang.org/doc/effective_go.html

## Libraries
Save yourself some time.  Do not reinvent the wheel.  Use existing open source libraries in your projects.
* **Go Libraries:** [https://golanglibs.com/top?q=cms](https://golanglibs.com/top?q=cms)
* **Go Libraries:** [http://go-search.org/](http://go-search.org/search?q=content)
* **Awesome Go:** [https://awesome-go.com/](https://awesome-go.com/)
