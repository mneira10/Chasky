package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	commands "main/commands"

	"github.com/infracloudio/msbotbuilder-go/core"
	"github.com/infracloudio/msbotbuilder-go/core/activity"
	"github.com/infracloudio/msbotbuilder-go/schema"
)

var customHandler = activity.HandlerFuncs{
	OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
		log.Println("in custom handler")
		responseTxt, isJSONText := commands.HandleCommand(turn.Activity.Text)

		if isJSONText {

			log.Println("Response is JSON.")
			var obj map[string]interface{}
			byteTxt := []byte(responseTxt)

			err := json.Unmarshal(byteTxt, &obj)

			if err != nil {
				return schema.Activity{}, err
			}

			attachments := []schema.Attachment{
				{
					ContentType: "application/vnd.microsoft.card.adaptive",
					Content:     obj,
				},
			}
			return turn.SendActivity(activity.MsgOptionAttachments(attachments))
		}
		log.Println("Response is regular text")
		return turn.SendActivity(activity.MsgOptionText(responseTxt))

	},
}

// HTTPHandler handles the HTTP requests from then connector service
type HTTPHandler struct {
	core.Adapter
}

func (ht *HTTPHandler) processMessage(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	activity, err := ht.Adapter.ParseRequest(ctx, req)
	if err != nil {
		log.Println("Failed to parse request.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ht.Adapter.ProcessActivity(ctx, activity, customHandler)
	if err != nil {
		log.Println("Failed to process request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Request processed successfully.")
}

func main() {

	setting := core.AdapterSetting{
		AppID:       os.Getenv("APP_ID"),
		AppPassword: os.Getenv("APP_PASSWORD"),
	}

	port := os.Getenv("APP_PORT")

	log.SetPrefix("CHASKY_BOT:")

	adapter, err := core.NewBotAdapter(setting)
	if err != nil {
		log.Fatal("Error creating adapter: ", err)
	}

	httpHandler := &HTTPHandler{adapter}

	http.HandleFunc("/", httpHandler.processMessage)
	log.Printf("Starting server on port:%s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
