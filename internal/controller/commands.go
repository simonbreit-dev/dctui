package controller

import "dctui/internal/appinstance"

func executeCommand(cmd string) {

	switch cmd {
	case "q":
		fallthrough
	case "q!":
		appinstance.App.Stop()
		break
	}

}
