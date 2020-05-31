 
package notify

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"

	"fmt"

	"github.com/pkg/errors"
	"github.com/martinlindhe/notify"
)

const channelName = "ligastavok/notify"

type NotifyFlutterPlugin struct {
	channel *plugin.MethodChannel
}

var _ flutter.Plugin = &NotifyFlutterPlugin{}

func (p *NotifyFlutterPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	p.channel.HandleFunc("notify", handleNotify)

	return nil
}

func handleNotify(arguments interface{}) (reply interface{}, err error) {
	var ok bool
	var argsMap map[interface{}]interface{}
	if argsMap, ok = arguments.(map[interface{}]interface{}); !ok {
		return nil, errors.New("invalid arguments")
	}

	// params
	var appNameParam string
	var titleParam string
	var textParam string
	var iconPathParam string
	var modeParam string

	if appName, ok := argsMap["appName"]; ok {
		appNameParam = appName.(string)
	}

	if title, ok := argsMap["title"]; ok {
		titleParam = title.(string)
	}

	if text, ok := argsMap["text"]; ok {
		textParam = text.(string)
	}

	if iconPath, ok := argsMap["iconPath"]; ok {
		iconPathParam = iconPath.(string)
	}

	if mode, ok := argsMap["mode"]; ok {
		modeParam = mode.(string)
	}

	switch modeParam {
		case "notify":
			notify.Notify(appNameParam, titleParam, textParam, iconPathParam)
		case "alert":
			notify.Alert(appNameParam, titleParam, textParam, iconPathParam)
		default:
			fmt.Println("Unknown notify mode. Choose one of notify or alert")
			// return nil, errors.New("Unknown notify mode. Choose one of notify or alert")
	}

	return nil, nil
}
