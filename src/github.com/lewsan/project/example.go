package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"time"
	"golang.org/x/sync/errgroup"
)

func actByUser(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func queryUser(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func postForm(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func postFormByIdAndPage(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")
	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}

func postWithMapParams(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
}

func uploadSingle(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func uploadBatch(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func loginEndpoint(c *gin.Context) {
	fmt.Println("loginEndpoint")
}

func submitEndpoint(c *gin.Context) {
	fmt.Println("submitEndpoint")
}

func readEndpoint(c *gin.Context) {
	fmt.Println("readEndpoint")
}

//func main() {
//	//log file
//	//gin.DisableConsoleColor()
//	//f, _ := os.Create("gin.log")
//	//gin.DefaultWriter = io.MultiWriter(f)
//
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.GET("/user/:name", func(c *gin.Context) {
//		name := c.Param("name")
//		c.String(http.StatusOK, "Hello %s", name)
//	})
//	r.GET("/user/:name/*action", actByUser)
//	r.GET("/welcome", queryUser)
//	r.POST("/form_post", postForm)
//	r.POST("/post", postFormByIdAndPage)
//	r.POST("/post_map", postWithMapParams)
//	r.POST("/single_upload", uploadSingle)
//	r.POST("batch_upload", uploadBatch)
//	v1 := r.Group("v1")
//	{
//		v1.POST("/login", loginEndpoint)
//		v1.POST("/submit", submitEndpoint)
//		v1.POST("/read", readEndpoint)
//	}
//	v2 := r.Group("v2")
//	{
//		v2.POST("/login", loginEndpoint)
//		v2.POST("/submit", submitEndpoint)
//		v2.POST("/read", readEndpoint)
//	}
//	r.Run()
//}

//Using Middleware

//func Logger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//		c.Set("example", "12345")
//		c.Next()
//		latency := time.Since(t)
//		log.Print(latency)
//		status := c.Writer.Status()
//		log.Println(status)
//	}
//}
//
//func MiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("before middleware")
//		c.Set("request", "client_request")
//		c.Next()
//		fmt.Println("after middleware")
//	}
//}
//
//func AuthMiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("before Authorization")
//		c.Set("authorize", "true")
//		c.Next()
//		fmt.Println("after authorization")
//	}
//}
//
//var secrets = gin.H{
//	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
//	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
//	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
//}
//
//func main() {
//	r := gin.New()
//	//r.Use(gin.Logger())
//	r.Use(Logger())
//	r.Use(gin.Recovery())
//	r.Use(MiddleWare())
//	{
//		r.GET("/middleware", func(c *gin.Context) {
//			request := c.MustGet("request").(string)
//			req, _ := c.Get("request")
//			c.JSON(http.StatusOK, gin.H{
//				"middle_request": request,
//				"request":        req})
//
//		})
//	}
//	r.GET("/test", func(c *gin.Context) {
//		example := c.MustGet("example").(string)
//		log.Println(example)
//	})
//	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"data": "home"})
//	})
//	//authorized := r.Group("/")
//	//authorized.Use(AuthMiddleWare())
//	//{
//	//	authorized.POST("/login", loginEndpoint)
//	//	authorized.POST("/submit", submitEndpoint)
//	//	authorized.POST("/read", readEndpoint)
//	//
//	//	testing := authorized.Group("testing")
//	//	testing.GET("/analytics", func(c *gin.Context) {
//	//		fmt.Println("analytics calling!")
//	//	})
//	//}
//
//	// Group using gin.BasicAuth() middleware
//	// gin.Accounts is a shortcut for map[string]string
//	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
//		"foo":    "bar",
//		"austin": "1234",
//		"lena":   "hello2",
//		"manu":   "4321",
//	}))
//	authorized.GET("/secrets", func(c *gin.Context) {
//		user := c.MustGet(gin.AuthUserKey).(string)
//		if secret, ok := secrets[user]; ok {
//			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET!"})
//		}
//	})
//
//	r.Run(":8080")
//}

//func main() {
//	router := gin.New()
//	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
//		return fmt.Sprintf("%s - [%s] \" %s %s %s %d %s \"%s\" %s\"\n",
//			param.ClientIP,
//			param.TimeStamp.Format(time.RFC1123),
//			param.Method,
//			param.Path,
//			param.Request.Proto,
//			param.StatusCode,
//			param.Latency,
//			param.Request.UserAgent(),
//			param.ErrorMessage,
//		)
//	}))
//	router.Use(gin.Recovery())
//	router.GET("/ping", func(c *gin.Context) {
//		c.String(http.StatusOK, "pong")
//	})
//	router.Run(":8080")
//}

//type Login struct {
//	User     string `form:"user" json:"user" xml:"user" binding:"required"`
//	Password string `form:"password" json:"password" xml:"password" binding:"required"`
//}
//
//func main() {
//	router := gin.Default()
//	router.POST("/loginJson", func(c *gin.Context) {
//		var json Login
//		if err := c.ShouldBindJSON(&json); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		if json.User != "manu" || json.Password != "123" {
//			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"status": "You are logged in!"})
//	})
//
//	router.POST("/loginXML", func(c *gin.Context) {
//		var xml Login
//		if err := c.ShouldBindXML(&xml); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		if xml.User != "manu" || xml.Password != "123" {
//			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"status": "You are logged in!"})
//	})
//
//	router.POST("/loginForm", func(c *gin.Context) {
//		var form Login
//		if err := c.ShouldBind(&form); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		if form.User != "manu" || form.Password != "123" {
//			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"status": "You are logged in!"})
//	})
//	router.Run(":8080")
//}

//type Booking struct {
//	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
//	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
//}
//
//func bookableDate(
//	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
//	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
//) bool {
//	if date, ok := field.Interface().(time.Time); ok {
//		today := time.Now()
//		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
//			return false
//		}
//	}
//	return true
//}
//
//func getBookable(c *gin.Context) {
//	var b Booking
//	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
//		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//	}
//}
//
//type Person struct {
//	ID       string    `uri:"id" binding:"required,uuid"`
//	Name     string    `form:"name"`
//	Address  string    `form:"address"`
//	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
//}
//
//func startPage(c *gin.Context) {
//	var person Person
//	//if c.ShouldBindQuery(&person) == nil {
//	if c.ShouldBind(&person) == nil {
//		log.Println("====== Only Bind By Query String ======")
//		log.Println(person.Name)
//		log.Println(person.Address)
//		log.Println(person.Birthday)
//	}
//	c.String(http.StatusOK, "Success")
//}
//
//type LoginForm struct {
//	User     string `form:"user" binding:"required"`
//	Password string `form:"password" binding:"required"`
//}
//
//func main() {
//	r := gin.Default()
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		v.RegisterValidation("bookabledate", bookableDate)
//	}
//	r.GET("/bookable", getBookable)
//	r.Any("/testing", startPage)
//	r.GET("/users/:name/:id", func(c *gin.Context) {
//		var person Person
//		if err := c.ShouldBindUri(&person); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
//	})
//	r.POST("/login", func(c *gin.Context) {
//		var form LoginForm
//		if c.ShouldBind(&form) == nil {
//			if form.User == "user" && form.Password == "password" {
//				c.JSON(http.StatusOK, gin.H{"status": "You're logged in."})
//			} else {
//				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//			}
//		}
//	})
//	r.Run(":8080")
//}

//func main() {
//	r := gin.Default()
//	r.GET("/someJSON", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
//	})
//	r.GET("/moreJSON", func(c *gin.Context) {
//		var msg struct {
//			Name    string `json:"user"`
//			Message string
//			Number  int
//		}
//		msg.Name = "Lena"
//		msg.Message = "hey"
//		msg.Number = 123
//		c.JSON(http.StatusOK, msg)
//	})
//	r.GET("/someXML", func(c *gin.Context) {
//		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
//	})
//	r.GET("/someYAML", func(c *gin.Context) {
//		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
//	})
//	r.GET("/someProtoBuf", func(c *gin.Context) {
//		reps := []int64{int64(1), int64(2)}
//		label := "test"
//		data := &protoexample.Test{
//			Label: &label,
//			Reps:  reps,
//		}
//		c.ProtoBuf(http.StatusOK, data)
//	})
//	r.Run(":8080")
//}

//func main() {
//	router := gin.Default()
//	router.Delims("{[{", "}]}")
//	router.SetFuncMap(template.FuncMap{
//		"formatAsDate": func(t time.Time) string {
//			year, month, day := t.Date()
//			return fmt.Sprintf("%d/%02d/%02d", year, month, day)
//		},
//	})
//	router.LoadHTMLFiles("./testdata/template/raw.tmpl")
//	router.GET("/raw", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
//			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
//		})
//	})
//	//router.LoadHTMLGlob("templates/**/*")
//	//router.GET("/posts/index", func(c *gin.Context) {
//	//	c.HTML(http.StatusOK, "posts/index.html", gin.H{
//	//		"title": "Posts",
//	//	})
//	//})
//	//router.GET("/users/index", func(c *gin.Context) {
//	//	c.HTML(http.StatusOK, "users/index.html", gin.H{
//	//		"title": "Users",
//	//	})
//	//})
//	//router.GET("/index", func(c *gin.Context) {
//	//	c.HTML(http.StatusOK, "index.html", gin.H{
//	//		"title": "Main website",
//	//	})
//	//})
//	//router.GET("/someDataFromReader", func(c *gin.Context) {
//	//	url := "https://raw.githubusercontent.com/gin-gonic/logo/master/color.png"
//	//	response, err := http.Get(url)
//	//	if err != nil || response.StatusCode != http.StatusOK {
//	//		c.Status(http.StatusServiceUnavailable)
//	//		return
//	//	}
//	//	reader := response.Body
//	//	contentLength := response.ContentLength
//	//	contentTyp := response.Header.Get("Content-Type")
//	//	extraHeaders := map[string]string{
//	//		"Content-Disposition": `attachment;filename="gopher.png"`,
//	//	}
//	//	c.DataFromReader(http.StatusOK, contentLength, contentTyp, reader, extraHeaders)
//	//})
//	router.Run(":8080")
//}

// Goroutines inside a middleware

//func main() {
//	r := gin.Default()
//	r.GET("/long_async", func(c *gin.Context) {
//		cCp := c.Copy()
//		go func() {
//			time.Sleep(5 * time.Second)
//			log.Println("Done! in path " + cCp.Request.URL.Path)
//		}()
//	})
//	r.GET("/long_sync", func(c *gin.Context) {
//		time.Sleep(5 * time.Second)
//		log.Println("Done! in path " + c.Request.URL.Path)
//	})
//	r.Run(":8080")
//}

// Run multiple service using Gin
var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01!",
			},
		)
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02!",
			},
		)
	})
	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
