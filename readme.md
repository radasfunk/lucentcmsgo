# Go Driver for LucentCMS

### Usage

Required keys, `channel_id`, `access_key`, `secret_key` and `locale` are required to communicate with the lucent api.

### Initializing 

```go
import "github.com/radasfunk/lucentcmsgo"
```

### Creating a lucent client

```go

channel := env.Get("LUCENTV3_CHANNEL")
secret := env.Get("LUCENTV3_SECRET")
accessKey := env.Get("LUCENTV3_ACCESS_KEY")
locale := env.Get("LUCENTV3_LOCALE")

dur := time.Duration(2 * time.Second) // timeout 

lc := lucentcmsgo.NewLucentClient(channel, secret, accessKey, locale, dur)
```

### Creating a request 

**lucentcmsgo** package follows the approach of creating a request and then making the request. 

Retriving all the documents,

```go
lc := lucentcmsgo.NewLucentClient(channel, secret, user, locale, dur)

request, err := lc.NewRequest("documents", nil)

res, err := request.Get()

if err != nil {
    // handle your error
    panic(err)
}

fmt.Println(res) 
```

Lucent requests will return a `Response` or `LucentListResponse`. `Response` is when you create or update a resource and get a single value returned. 

`LucentListResponse` is for every other response. 

**Note** If your api request results in an error, it will still return a `Response` or `LucentListResponse` depending on the request type, and the `error` value will be `nil`. 
But if it has other errors like having problem encoding or something like that, or maybe the request had send malformatted data and go is having problem to decode it, it will return an error with an empty `Response` or `LucentListResponse`.

### Adding request data with get request

For a get request, data is added as params

```go
lc := lucentcmsgo.NewLucentClient(channel, secret, user, locale, dur)

requestData := make(map[string]string)
requestData["key"] = "some value"

request, err := lc.NewRequest("documents", nil)
request.AddParams(requestData)

res, err := request.Get()
```


### Adding request data with post request

**Note** For post request the data structure is a map of string to interfaces{}, depending on your need.
But, at the top level, it needs to have a key with data

```go
requestWith := make(map[string]interface{})
requestWith["key"] = "some value"

request, err := lc.NewRequest("documents", nil)

requestData := make(map[string]interface{})
requestData["data"] = requestWith

lc := lucentcmsgo.NewLucentClient(channel, secret, user, locale, dur)
	request.AddData(requestData)

res, err := request.Post()
```

You can also add data in `lc.NewRequest("documents",$here)`. Only for non-get methods.


### Skip
To skip values for paginations or any other reason,

```go
request, err := lc.NewRequest("documents", nil)

request.setSkip(20)
```

### Limit
To limit the number of values,

```go
request, err := lc.NewRequest("documents", nil)

request.SetLimit(20)
```

### Meta
To set meta value for the request,

```go
request, err := lc.NewRequest("documents", nil)

request.SetMeta("articles")
```

### Include
To include related documents,

```go
request, err := lc.NewRequest("documents", nil)

request.SetInclude("articles")
// or
request.SetInclude("articles,categories")
```

### IncludeAll
To include all related documents,

```go
request, err := lc.NewRequest("documents", nil)

request.SetIncludeAll()
```
### Adding headers

```go
additionalHeaders := make(map[string]string)

additionalHeaders["Custom-Key"] = "Custom-Value"
additionalHeaders["Custom-Another-Key"] = "custom value 123"

request, err := lc.NewRequest("documents", nil)
request.AddHeaders(additionalHeaders)
```

**Note** Some headers are protected and can not be overridden, them being 
`Lucent-Channel` and `Lucent-User`.


### Adding where filters
```go
request, err := lc.NewRequest("documents", nil)

request.Where("title","hello world")
```

### Adding orWhere filters
```go
request, err := lc.NewRequest("documents", nil)

request.OrWhere("user_id","123-456-789-101")
```

### Available filter methods
```go
Where(key string, value interface{})

OrWhere(key string, value interface{})

In(key string, value string)

Regex(key string, value string)

Exists(key string)

NExists(key string)

Eq(key string, value interface{})

Ne(key string, value interface{})

Nin(key string, value interface{})

Lt(key string, value interface{})

Lte(key string, value interface{})

Gt(key string, value interface{})

Gte(key string, value interface{})

True(key string)

False(key string)

Null(key string)

Nil(key string)

Empty(key string)
```

### Uploading files from disk 

```go

path := []string{
	"absolute/path/to/file.extension",
	"/home/dev/pikachu.png", // like so
}

res, err := request.UploadFromPath(path) // Response will be an Upload Response
```

### Response structure
```go
type LucentListResponse struct {
	Data Document
	Errors, Links  []string
	Meta, Included map[string]interface{} 
}
```
### LucentListResponse structure

```go
type LucentListResponse struct {
	Data []Document
	Errors, Links  []string
	Meta, Included map[string]interface{} 
}
```
### UploadResponse structure

```go
type UploadResponse struct {
	Data []File `json:"data"`
	Errors, Links  []string
	Meta, Included map[string]interface{} 
}
```

### Document structure

```go
type Document struct {
	ID            string      `json:"id"`
	RequestLocale string      `json:"requestLocale"`
	Locale        string      `json:"locale"`
	Schema        string      `json:"schema"`
	Creator       string      `json:"creator"`
	Editor        string      `json:"editor"`
	Status        string      `json:"status"`
	Version       int         `json:"version"`
	PublishedAt   time.Time   `json:"publishedAt"`
	Behind        bool        `json:"behind"`
	Content       Content     `json:"content"` // map[string]interface{}
	Subdocs       interface{} `json:"subdocs"`
	Relationships interface{} `json:"relationships"`
	Channel       string      `json:"channel"`
	Resource      string      `json:"resource"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	CreatedAt     time.Time   `json:"createdAt"`
}
```

### File structure 

```go
type File struct {
	ID           string      `json:"id"`
	OriginalName string      `json:"originalName"`
	Filename     string      `json:"filename"`
	Path         string      `json:"path"`
	Mime         string      `json:"mime"`
	URL          string      `json:"url"`
	Image        string      `json:"image"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Size         int         `json:"size"`
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	Alt          string      `json:"alt"`
	Credits      interface{} `json:"credits"`
	Checksum     string      `json:"checksum"`
	Copyright    string      `json:"copyright"`
	Tags         []string    `json:"tags"`
	UploaderID   string      `json:"uploaderId"`
	Channel      string      `json:"channel"`
	UpdatedAt    time.Time   `json:"updatedAt"`
	CreatedAt    time.Time   `json:"createdAt"`
	Resource     string      `json:"resource"`
}
```

### Methods with response (Response and LucentListResponse)

```go
GetIncluded() map[string]interface{}
GetErrors() []string 
Error() string // returns the first error if exists
HasErrors() bool
```

### GetData()
`GetData` method is available with both `Response` and `LucentListResponse` but they return different sturctures, `LucentListResponse` will return you an `array of Document` while `Response` will return you a `single Document`. 

### Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

### Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details.

### Security

If you discover any security related issues, please email hey@lucentcms.com instead of using the issue tracker.