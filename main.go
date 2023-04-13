package main

import (
	//"context"
	"fmt"
	//"log"
	"os"
	//"strconv"

	//"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

// uncomment this section for the simple age bot
// func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
// 	for event := range analyticsChannel {
// 		fmt.Println("command events")
// 		fmt.Println(event.Timestamp)
// 		fmt.Println(event.Command)
// 		fmt.Println(event.Parameters)
// 		fmt.Println(event.Event)
// 		fmt.Println()
// 	}
// }

// simple age bot
// func main() {
// 	//xoxb-5110025749844-5101150388934-v4bfo3JKpmru6VGa2ngLAaMG - file-bot token
// 	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5110025749844-5131271778944-1qf5JUzsqWAGicM0rZ72fiSy")
// 	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A053V73354G-5093012038919-cd0e1c40b2ce63d14244561ba9ffcbb74eb44fac9ce200d223c0051b64d66850")

// 	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

// 	go printCommandEvents(bot.CommandEvents())

// 	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
// 		Description: "yob calculator",
// 		// Examples:    "My yob is 2020",
// 		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
// 			year := request.Param("year")
// 			yob, err := strconv.Atoi(year)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			age := 2023 - yob
// 			r := fmt.Sprintf("age is %d", age)
// 			response.Reply(r)
// 		},
// 	})

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err := bot.Listen(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

//uncomment this function for file bot upload
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5110025749844-5101150388934-v4bfo3JKpmru6VGa2ngLAaMG")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A053V73354G-5093012038919-cd0e1c40b2ce63d14244561ba9ffcbb74eb44fac9ce200d223c0051b64d66850")
	os.Setenv("SLACK_CHANNEL_ID", "C0535G6LMB4")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}
	fileArr := []string{"Aquib_SE_UAE.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s \n", err)
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
