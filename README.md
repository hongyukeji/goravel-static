# Goravel Static

> An Static For A Goravel Extend Package

> 将public目录下的文件挂载至应用根路径

## Directory Structure

This is a directory standard, but you can change it if you like.

| Directory        | Action           |
| -----------      | --------------   |
| config           | Store the config files   |
| middleware        | Store the middleware files   |

## Install

1. Add package

```
go get -u github.com/hongyukeji/goravel-static
```

2. Use middleware

```
// app/http/kernel.go
import staticmiddleware "github.com/hongyukeji/goravel-static/middleware"

func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{
		staticmiddleware.Static(), // add static middleware
	}
}
```

3. Publish Configuration

```
go run . artisan vendor:publish --package=github.com/hongyukeji/goravel-static
```

4. Testing

```
访问格式：[你的应用网址]/[public目录下的文件名]
例如：http://127.0.0.1:3000/logo.png
```

The console will print `file`.