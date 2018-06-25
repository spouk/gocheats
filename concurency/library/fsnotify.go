package library

import (
	"sync"
	"os"
	"log"
	"os/signal"
	"syscall"
	"time"
	"fmt"
)

type MasterControlFile struct {
	sync.WaitGroup
	ListFiles []*os.File
	Command   chan string
	Log       *log.Logger
	sigs      chan os.Signal
	timer     *time.Timer
	ticker    *time.Ticker
}

func NewMasterControlFile(tickPeriod, timePeriod time.Duration, listFiles []string, ) *MasterControlFile {
	//make new instance
	var mm = &MasterControlFile{
		Command: make(chan string, len(listFiles)),
		sigs:    make(chan os.Signal),
		Log:     log.New(os.Stdout, "[mastercontrolfiles]", log.Ltime|log.Ldate|log.Lshortfile),
		ticker:  time.NewTicker(time.Duration(time.Second * tickPeriod)),
		timer:   time.NewTimer(time.Duration(time.Minute * timePeriod)),
	}
	//check + open files
	files, err := mm.checkFiles(listFiles)
	if err != nil {
		mm.Log.Fatal(err)
	}
	mm.ListFiles = files

	return mm

}
func (m *MasterControlFile) checkFiles(listFiles []string) ([]*os.File, error) {
	//check files exists + open file
	var result []*os.File
	for _, x := range listFiles {
		if _, err := os.Stat(x); err != nil {
			if err != nil {
				return nil, err
			}
		} else {
			f, err := os.Open(x)
			if err != nil {
				return nil, err
			}
			result = append(result, f)
		}
	}
	return result, nil
}
func (m *MasterControlFile) RunControl() {

	//signal notify catch exit program
	signal.Notify(m.sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-m.sigs
		m.Log.Printf("[CATCHER]: CATCH Os.SIGNAL `%v`\n", sig)
		for _, x := range m.ListFiles {
			if x != nil {
				x.Close()
			}
		}
		m.Log.Printf("[CATCHER]: all open handlers success closed\n")
		os.Exit(0)
	}()

	//start timer worker
	go m.timerChecker()
	m.Add(1)

	//start files workers
	for i, x := range m.ListFiles {
		go m.worker(fmt.Sprintf("W#%d", i), x)
		m.Add(1)
	}

	//wait end workers
	m.Wait()
}
func (m *MasterControlFile) checkError(err error) {
	if err != nil {
		m.Log.Printf(err.Error())
		os.Exit(-1)
	}
	return
}
func (m *MasterControlFile) timerChecker() {
	fmt.Printf("[master.timerChecker] start...\n")
	defer func() {
		fmt.Printf("[master.timerChecker] end...\n")
		m.Done()
	}()

	for {
		select {
		case <-m.timer.C:
			fmt.Printf("[master.timerChecker] get timer signal\n")
			for x:=0; x < len(m.ListFiles); x ++ {
				m.Command <- "exit"
			}
			fmt.Printf("[master.timerChecker] send command `exit` all workers in channel \n")
			return
		}
	}
}
func (m *MasterControlFile) worker(name string, obj *os.File) {
	var ff  = func(msg string) string {
		return 	fmt.Sprintf("[%s]  %s\n", name, msg)
	}
	m.Log.Printf("%s", ff("starting"))
	defer func() {
		m.Done()
		m.Log.Printf("%s", ff("success exit"))
	}()
	for {
		select {
		case com := <-m.Command:
			if com == "exit" {
				return
			}
		case tick := <-m.ticker.C:
			m.Log.Printf("%s%v", ff(""), tick)
		default:
			m.Log.Printf("[%s] Obj: %v\n", name, obj.Name())
			time.Sleep(time.Second * 10)
		}
	}
}
