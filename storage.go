package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

//var storageFileDir = "/var/lib/circuitbreaker/"

type Storage struct {
	storageFileDir string
	fileName string
	storageFilePath string
	mutex *sync.Mutex
}

func NewStorage(storageFileDir, fileName string) *Storage {

	storageFilePath := storageFileDir + fileName

	storage := &Storage{
		storageFileDir: storageFileDir,
		fileName: fileName,
		storageFilePath : storageFileDir + "/" + fileName,
		mutex: &sync.Mutex{},
	}

	if _, err := os.Stat(storageFilePath); err == nil {
		return storage
	}

	storage.initStorage()

	return storage
}

func (s *Storage) initStorage() {

	if _, err := os.Stat(s.storageFileDir); os.IsNotExist(err) {
		err := os.MkdirAll(s.storageFileDir, 0644)
		if err != nil {
			log.Fatalf("데이터 파일 디렉터리 생성에 실패했습니다 %s", err.Error())
			panic(err)
		}
	}

	err := ioutil.WriteFile(s.storageFilePath, []byte{}, 0644)
	if err != nil {
		log.Errorf("데이터 파일을 생성할 수 없습니다. %s", err.Error())
		panic(err)
	}
}

func (s *Storage) getBrakedHosts() []string {
	s.mutex.Lock()
	file, err := os.Open(s.storageFilePath)
	defer file.Close()

	if err != nil {
		log.Errorf("데이터 파일을 읽을 수 없습니다. %s", err.Error())
		s.initStorage()
		s.mutex.Unlock()
		return []string{}
	}

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)
	var lines []string

	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	s.mutex.Unlock()
	return lines

}

func (s *Storage) recoverBrakedHost(host string) {

}


func (s *Storage) isMetricHostAlreadyBraked(reqHost string) bool {
	hosts := s.getBrakedHosts()
	for _, host := range hosts {
		if strings.Contains(host, reqHost) {
			return true
		}
	}
	return false
}
