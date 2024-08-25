package pprof

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	spsio "github.com/SPSZerone/sps-go-zerone/io"
)

func ListenAndServe(pprofIP string, handler http.Handler) {
	err := http.ListenAndServe(pprofIP, handler)
	if err != nil {
		fmt.Printf("start pprof failed on %+v\n", pprofIP)
	}
}

func NewPProf(ip string, port int, savePath string, tickSaveInterval int) (*PProf, error) {
	param := &Param{
		ip:               ip,
		port:             port,
		savePath:         savePath,
		tickSaveInterval: tickSaveInterval}
	pprof := &PProf{param: param}
	err := pprof.Init()
	return pprof, err
}

type Param struct {
	ip               string
	port             int
	savePath         string
	tickSaveInterval int // (in minute)
}

func (p *Param) Check() bool {
	if "" == p.ip || 0 == p.port || "" == p.savePath {
		return false
	}
	return true
}

func (p *Param) GetUrlForProfile() string {
	return fmt.Sprintf("http://%s:%d/debug/pprof/profile", p.ip, p.port)
}

func (p *Param) GetUrlForHeap() string {
	return fmt.Sprintf("http://%s:%d/debug/pprof/heap?gc=1", p.ip, p.port)
}

func (p *Param) CreateSavePath() error {
	savePath := p.savePath
	if _, err := os.Stat(savePath); err != nil {
		err = os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			fmt.Printf("PProfParam CreateSavePath %s error: %+v\n", savePath, err)
			return err
		}
	}
	return nil
}

type Type int

const (
	TypeProfile = 1 << iota
	TypeHeap
)

type PProf struct {
	param *Param
}

func (p *PProf) Init() error {
	if p.CheckParam() {
		return p.param.CreateSavePath()
	}
	return errors.New("PProf CheckParam failed")
}

func (p *PProf) Start() {
	p.Save()

	if p.param.tickSaveInterval > 0 {
		ticker := time.NewTicker(time.Duration(p.param.tickSaveInterval) * time.Minute)

		go func() {
			for {
				select {
				case <-ticker.C:
					p.Save()
				}
			}
		}()

		select {}
	}
}

func (p *PProf) Save() {
	fmt.Println("========== PProf Save [START] ==========")

	var wg sync.WaitGroup

	pprofTypes := []Type{TypeProfile, TypeHeap}
	for _, pprofType := range pprofTypes {
		wg.Add(1)
		go func(ppType Type) {
			defer wg.Done()

			err := p.save(ppType)
			if err != nil {
				fmt.Printf("save failed %d %+v\n", ppType, err)
			}
		}(pprofType)
	}

	wg.Wait()

	fmt.Println("========== PProf Save [DONE] ==========")
}

func (p *PProf) save(pprofType Type) error {
	typeName := p.getTypeName(pprofType)
	fmt.Printf("PProf save %s | START\n", typeName)

	dataBytes, err := p.getDataByType(pprofType)
	if err != nil {
		fmt.Printf("PProf save %s | getDataByType failed: %+v\n", typeName, err)
		return err
	}

	nowTimeFormat := time.Now().Format("2006.01.02_15.04.05")
	fileName := fmt.Sprintf("%s.%s.pprof", typeName, nowTimeFormat)
	fullFileName := filepath.Join(p.param.savePath, fileName)
	err = spsio.WriteBytes(dataBytes, fullFileName, false)
	if err != nil {
		fmt.Printf("PProf save %s | save file failed: %+v\n", typeName, err)
		return err
	}

	fmt.Printf("PProf save %s | SUCCESSED to %s\n", typeName, fullFileName)
	return nil
}

func (p *PProf) CheckParam() bool {
	if p.param == nil {
		return false
	}
	return p.param.Check()
}

func (p *PProf) getTypeName(pprofType Type) string {
	switch pprofType {
	case TypeProfile:
		return "profile"
	case TypeHeap:
		return "heap"
	}
	return "unknown"
}

func (p *PProf) getDataByType(pprofType Type) ([]byte, error) {
	var url string
	switch pprofType {
	case TypeProfile:
		url = p.param.GetUrlForProfile()
	case TypeHeap:
		url = p.param.GetUrlForHeap()
	default:
		return nil, errors.New(fmt.Sprintf("Unknown PProfType %v\n", pprofType))
	}
	return p.getData(url)
}

func (p *PProf) getData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("PProf getData | http.Get error: %+v\n", err)
		return nil, err
	}
	defer func(body io.ReadCloser) {
		err = body.Close()
		if err != nil {
			fmt.Printf("PProf getData | resp.Body.Close error: %+v\n", err)
		}
	}(resp.Body)

	bodyBytes, err := spsio.ReadBytes(resp.Body, 10240)
	if err != nil {
		fmt.Printf("PProf getData | ReadBytes error: %+v\n", err)
		return nil, err
	}

	return bodyBytes, nil
}
