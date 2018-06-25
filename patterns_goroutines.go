package main

import (
	"fmt"
	"sync"
	"time"
)

type CorePatterns struct {
	sync.WaitGroup
	start  chan bool
	end    chan bool
	timer  *time.Timer
	ticker *time.Ticker
}

func NewCorePatterns(timer, tickerTime time.Duration) *CorePatterns {
	return &CorePatterns{
		start:  make(chan bool),
		end:    make(chan bool),
		timer:  time.NewTimer(timer * time.Second),
		ticker: time.NewTicker(tickerTime * time.Second),
	}
}

func (p *CorePatterns) worker(name string) {
	fmt.Printf("[%v] start...\n", name)
	defer func() {
		p.Done()
		fmt.Printf("[%v] end...\n", name)
	}()
	//wait starting work
	<-p.start

	for {
		select {
		case <-p.end:
			return
		default:
			fmt.Printf("[%v] WORKER RUN SELECT...\n", name)
			time.Sleep(time.Duration(time.Second * 1))
		}
	}
}
func (p *CorePatterns) runWorkers(count int) {
	go p.runTicker()
	p.Add(1)

	go p.runManager()
	p.Add(1)

	for i := 0; i < count; i ++ {
		go p.worker(fmt.Sprintf("%d", i))
		p.Add(1)
	}
	fmt.Print("Стартует цикл начала старта работы горутин...\n")
	time.Sleep(time.Duration(time.Second * 5))
	close(p.start)
	fmt.Print("Старт пошел...\n")

	p.Wait()
	fmt.Print("Работа завершена...\n")
}
func (p *CorePatterns) runTicker() {
	defer func() {
		p.Done()
	}()
	for {
		select {
		case t := <-p.ticker.C:
			fmt.Printf("Ticker: %v\n", t)
		case <-p.end:
			fmt.Printf("ticker the END\n")
			return
		default:
			time.Sleep(time.Second * 3)
		}
	}
}
func (p *CorePatterns) runManager() {

	defer func() {
		p.Done()
	}()
	<-p.timer.C
	close(p.end)

	//for {
	//	select {
	//	case  tick := <- p.timer.C:
	//		fmt.Print("TICK: %v\n",tick)
	//		//fmt.Print("---manager close timer channel \n")
	//		//close(p.end)
	//	default:
	//		fmt.Printf("++++MANAGER\n")
	//		time.Sleep(time.Duration(time.Second * 2))
	//
	//	}
	//}
}

func main() {
	p := NewCorePatterns(time.Duration(10), 1)
	p.runWorkers(10)
}
