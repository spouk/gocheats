package main

import (
	. "bitbucket.org/cyberspouk/learning/concurency/library"
	"fmt"
	"gopkg.in/fsnotify.v1"
	"os"
)

func main() {
	//---------------------------------------------------------------------------
	//  example fsnotify 
	//---------------------------------------------------------------------------
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Errorf("%v\n", err.Error())
	}
	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Printf("Event: %#v : %v\n", event, event.Op.String())

			case event_errors := <-watcher.Errors:
				fmt.Printf("Event: %v\n", event_errors)
			}
		}
	}()

	err = watcher.Add("/home/spouk/tmp/one")
	if err != nil {
		fmt.Print(err)
	}
	<- done

	os.Exit(1)

	defer func() {
		fmt.Printf("EXIT MAIN PROGRAMM DEFER FUNCTION\n")
	}()
	var listFiles = []string{
		"/home/spouk/tmp/one", "/home/spouk/tmp/two", "/home/spouk/tmp/three",
	}
	m := NewMasterControlFile(10, 1, listFiles)
	m.RunControl()

}
