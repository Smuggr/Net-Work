package bridger

import (
	"network/common/pluginer"

	"github.com/charmbracelet/log"
)

func Tester() {
	r := pluginer.NewRouter()

	gettersGroup := r.Group("getters", nil)
	{
		barGettersGroup := gettersGroup.Group("bar", map[string][]pluginer.HandlerFunction{
			"GET": { func(ctx *pluginer.Context) { ctx.JSON(200, "fab!") }, },
		})
		{
			group := barGettersGroup.GET("fab", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fab!")
			})

			group.GET("/fem", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fem!")
			})

			barGettersGroup.GET("fam", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fam!")
			})
		}

		farGettersGroup := gettersGroup.Group(":far", nil)
		{
			group1 := farGettersGroup.Group(":fob", nil)
			group1.GET("feb", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fob!")
            })

			group2 := farGettersGroup.Group(":for", nil)
			group2.GET("fer", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fob!")
            })
		}
	}

	group := r.GetGroup("/getters/bar/fab")
	log.Warn("found", "group", group)

	group = r.GetGroup("/getters/asd/sex/sssssssssssssssssssssss/sex/sex/feb")
	log.Warn("found", "group", group)
}