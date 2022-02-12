package main

import (
	"log"
	"strings"
)

func main() {
	path := "m/44H/60H/0H/0/0"
	paths := strings.Split(path, "/")
	pathsWithMethod := make([][]string, 0)
	for _, v := range paths {
		if v == "m" {
			pathsWithMethod = append(pathsWithMethod, []string{v})
			continue
		}
		if strings.Contains(v, "H") {
			hSplit := strings.TrimSuffix(v, "H")
			pathsWithMethod = append(pathsWithMethod, []string{hSplit, "H"})
		} else {
			pathsWithMethod = append(pathsWithMethod, []string{v})
		}
	}
	log.Fatal(pathsWithMethod)
	// app.Init(".env", "logs")
	// logwrapper.Log.Info("Starting app")
	// port := envutil.MustGetEnv("PORT")
	// err := app.GinApp.Run(":" + port)
	// if err != nil {
	// 	logwrapper.Log.Fatal(err)
	// }
}
