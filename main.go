package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	controller "github.com/scraper_v2/controllers"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.Use(func(c *fiber.Ctx) {
		c.Fasthttp.Response.Header.Add("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	grp := app.Group("api")
	grp.Get("/1337x", controller.Controller1337x)
	grp.Get("/eztv", controller.EztvController)
	grp.Get("/horriblesubs", controller.HorriblesubsController)
	grp.Get("/nyaa", controller.NyaaController)
	grp.Get("/kickass", controller.KickassController)
	grp.Get("/limetorrents", controller.LimetorrentsController)
	grp.Get("/thepiratebay", controller.PirateBayController)
	grp.Get("/torrentdownloads", controller.TorrentdownloadsController)
	grp.Get("/tgx", controller.TorrentGalaxyController)
	grp.Get("/yts", controller.YtsController)
	grp.Get("/tgxmov", controller.RecentMoviesController)
	grp.Get("/tgxseries", controller.RecentShowsController)
	grp.Get("/imdb", controller.ImdbController)
	grp.Get("/rarbg", func(c *fiber.Ctx) { c.Status(204) })
	grp.Get("/appversion", controller.AppUpdateController)
	grp.Get("/zooqle", controller.ZooqleController)

	port := "8080"
	app.Settings.CaseSensitive = true
	app.Settings.StrictRouting = true
	app.Listen(":" + port)

}
