package main

import (
	"strconv"

	"github.com/dinever/golf"
)

func mainHandler(ctx *golf.Context) {
	name, err := ctx.Session.Get("name")
	ctx.SetHeader("Content-Type", "text/html;charset=UTF-8")
	if err != nil {
		//ctx.Send("Hello World! Please <a href=\"/login\">log in</a>. Current sessions: " + strconv.Itoa(ctx.App.SessionManager.Count()))
		ctx.Redirect("/login")
	} else {
		ctx.Send("Hello " + name.(string) + ". Current sessions: " + strconv.Itoa(ctx.App.SessionManager.Count()) +
			`<br><a href="/logout">logout</a>
			
			
		`)
	}
}

func loginHandler(ctx *golf.Context) {
	ctx.Loader("default").Render("login.gohtml", make(map[string]interface{}))
}

func logoutHandler(ctx *golf.Context) {
	err := ctx.Session.Delete("name")
	if err != nil {
		ctx.SetHeader("Content-Type", "text/html;charset=UTF-8")
		ctx.Send("Error: " + err.Error())
	}
	ctx.Redirect("/")
}

func loginHandlerPost(ctx *golf.Context) {
	ctx.Session.Set("name", ctx.Request.FormValue("name"))
	//ctx.Send("Hi, " + ctx.Request.FormValue("name"))
	ctx.Redirect("/")
}

func main() {
	app := golf.New()
	app.View.SetTemplateLoader("default", ".")
	app.SessionManager = golf.NewMemorySessionManager()
	app.Use(golf.SessionMiddleware)

	app.Get("/", mainHandler)
	app.Post("/login", loginHandlerPost)
	app.Get("/login", loginHandler)
	app.Get("/logout", logoutHandler)
	app.Run(":80")
}

/*

	// nice piece of code
	// Config control for the application.
	type Config struct {
		mapping map[string]interface{}
	}

	// NewConfig creates a new configuration instance.
	func NewConfig() *Config {
		mapping := make(map[string]interface{})
		return &Config{mapping}
	}

	// ConfigFromJSON creates a Config instance from a JSON io.reader.
	func ConfigFromJSON(reader io.Reader) (*Config, error) {
		jsonBytes, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		var obj map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &obj); err != nil {
			return nil, err
		}
		return &Config{obj}, nil
	}


	// supports multiple data types
	func (ctx *Context) Send(body interface{}) {

		switch body.(type) {
		case []byte:
			ctx.Response.Write(body.([]byte))
		case string:
			ctx.Response.Write([]byte(body.(string)))
		case *bytes.Buffer:
			ctx.Response.Write(body.(*bytes.Buffer).Bytes())
		default:
			panic(fmt.Errorf("Body type not supported"))
		}
	}


*/
