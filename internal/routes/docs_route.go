package routes

import "github.com/labstack/echo/v4"

func DocsRoute(e *echo.Echo) {
	e.GET("/docs", func(c echo.Context) error {
		return c.HTML(200, `
			<!doctype html>
			<html>
			<head>
			<title>API Docs</title>
			<meta charset="utf-8" />
			<meta
				name="viewport"
				content="width=device-width, initial-scale=1" />
			</head>

			<body>
			<div id="app"></div>

			<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>

			<script>
				Scalar.createApiReference('#app', {
				url: '/swagger/doc.json',
				theme: 'purple',
				})
			</script>
			</body>
			</html>
		`)
	})
}