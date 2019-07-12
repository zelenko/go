
# Exploring Golang features [![Go Report Card](https://goreportcard.com/badge/zelenko/go)](https://goreportcard.com/report/zelenko/go)
You can incorporate these code snippets into your larger programming modules. This repository is viewable on [sourcegraph.com](https://sourcegraph.com/github.com/zelenko/go).

![Go](https://raw.githubusercontent.com/zelenko/go/master/37_html_template/pub/img/toby.jpg)

## Contents

| Name | Description |
| :---- | :---- |
| 00_persisting_go            | Installing Go, persisting on systemd, Upstart |
| 01_data_type                | type, error, recursion, reference, sort, switch, type assertion |
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
| 14_read_txt_file            | Compare ways to process TXT file |
| 15_png_or_svg_barchart      | generate bar chart, PNG or SVG |
| 16_drop_down_menu_form      | Form /w drop down menu, template with sub-templates |
| 17_mysql_user_login         | MySQL for user registration, login, and user list. |
| 18_cookie_authentication    | Authentication with cookies. |
| 19_os_global_variables      | Displaying all global system variables. |
| 20_mongodb_crud             | Mongodb CRUD with REST using httprouter & HTML templates |
| 21_httprouter_template      | html template with httprouter and ServeFiles |
| 22_mongodb_crud_rest_html   | Mongodb CRUD with REST using httprouter & HTML templates |
| 23_file_uploader            | Upload file and save on server |
| 24_calculate_time           | Find out how much time it takes for function to start and finish |
| 25_https_static_files       | Serve HTTP and HTTPS w/ NotFound for static files |
| 26_url_not_found_handler    | Custom Not Found handler. |
| 27_mongodb_bulk_upsert      | Mongodb bulk insert from TXT file |
| 28_markdown                 | Generate Markdown using blackfriday |
| 29_go_crud_json_api         | REST API using JSON with httprouter.  JavaScript is used to view, create, edit, and delete records. |
| 30_mongodb_crud_json_api    | REST API using JSON, httprouter, and toml, i/o to MongoDB |
| 31_send_email               | Send email with attachment |
| 32_colorful_cli             | Create colorful CLI |
| 33_testing                  | Testing package example |
| 34_channels                 | Buffered/unbuffered channels, forking channel, ranging over closed channel. |
| 35_mongodb_pipeline_page    | One page MGO aggregation with pipeline |
| 36_concurrency_channel      | Channel, waiting, concurrency, sleep, close, count, queue |
| 37_html_template            | Simple html template with Go |
| 38_url_request_JSON         | Convert data to/from JSON, get and parse file from URL |
| 39_read_directory_content   | List files adn sub folders in a given folder |
| 40_cron_scheduler           | Schedule processes |
| 41_cli_arguments            | Run cli utility with options using os.Args |
| 42_upload_many_files        | Upload multiple files from browser form to folder on a server. |
| 43_resize_jpg_png_image     | Resize images |
| 44_csv_file                 | Read and write to CSV file.  Parse CSV file to slice of objects. |
| 45_image_exif_data          | Get image attributes for each image in a folder. |
| 46_video_capture            | Capture video from web camera and display live. |
| 47_download_slice_as_csv    | Download link generates CSV or Tab Delimited file that can be saved localy on your computer. |
| 48_keyboard_driver          | Testing IOT devices |
| 49_constructor              | Example creating new package with allocation/constructor that accepts multiple types using interface. Experimenting with Readers and Writers. |
| 50_golf_framework           | A fast, simple and lightweight micro-web framework for Go |
| 51_blur_image               | Blur, Rotate, and Generate Thumbnails. |
| 52_jpg_image_watermark      | Add watermark to image |
| 53_regular_expression       | Validation, Find and replace, security |
| 54_rotate_image             | Image rotation in degrees from 1 to 360. |
| 55_html_template_std_lib    | HTML template using standard library packages |
| 56_html_formatter           | Work in similar fashion as `go fmt`, but on HTML files. |
| 57_valid_interface		  | Interface as parameter. |
| 58_GO_HTML_template         | Go HTML template examples. |
| 59_zip_and_unzip            | Example for archive/zip package. |
| 60_http_response_as_file    | HTTP handler responds with a copied file. |
| 61_logging_middleware       | Save logs and error logs to file or database. |
| 62_download_progress        | Progress shown on CLI. |
| 63_graphql_todo_example     | One file GraphQL example. |
| 64_date_and_time            | Using time package. |
| 65_books_example            | MongoDB CRUD example. |
| 66_server_sent_events       | Live logs (events) from server to browser using the [EventSource](https://developer.mozilla.org/en-US/docs/Web/API/EventSource) HTML interface. |
| 67_stringutil				  | Reverse a string, test included |
| 68_iota                     | Iota identifier is used in const declarations to simplify definitions of incrementing numbers. |

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
}
```
![Go](http://farm4.staticflickr.com/3774/11740822616_e435d02a54_o.gif)

> In Go, the code does exactly what it says on the page.

> It’s the simplicity that makes Go awesome.

> Go strives to keep things small and beautiful.

> What I would have done in Python, Java, Ruby, PHP, C, C# or C++, I’m now doing in Go.

> The code must be like a piece of music.

# TODO :
- [X] MongoDB connection
- [X] MySQL connection
- [X] Resize Images
- [X] Set/Get Image tags
- [ ] PostgreSQL connection
- [ ] SQLite connection
- [ ] MS SQL Server connection
- [ ] React frontend
