package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	"github.com/jackpal/bencode-go"
	"github.com/labstack/gommon/color"
)

// File available as part of the torrent
type File struct {
	Length int    `bencode:"length"`
	Md5sum string `bencode:"md5sum"`
	Path   string `bencode:"path"`
}

// MetaInfoData hold data about the download itself
type MetaInfoData struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Private     int    `bencode:"private"`
	Length      int    `bencode:"length"`
	Md5sum      string `bencode:"md5sum"`
	Files       []File `bencode:"files"`
}

// MetaInfo contain .torrent file description.
type MetaInfo struct {
	Announce     string       `bencode:"announce"`
	AnnounceList [][]string   `bencode:"announce-list"`
	Info         MetaInfoData `bencode:"info"`
	Encoding     string       `bencode:"encoding"`
	CreationDate int          `bencode:"creation date"`
	CreatedBy    string       `bencode:"created by"`
}

// Torrent struct
type Torrent struct {
	Path string
	Data MetaInfo
	Hash []byte
}

func newTorrent(path string) (Torrent, error) {
	file, err := os.Open(path)
	if err != nil {
		return Torrent{}, errors.New("Failed to open torrent file: " + err.Error())
	}
	defer file.Close()

	info := MetaInfo{}
	err = bencode.Unmarshal(file, &info)
	if err != nil {
		return Torrent{}, errors.New("Failed to decode torrent file: " + err.Error())
	}

	return Torrent{Path: path, Data: info}, nil
}

func readTorrentFile(path string, worker *sync.WaitGroup) {
	defer worker.Done()
	torrent, err := newTorrent(path)
	if err != nil {
		fmt.Println(crossPre +
			color.Yellow(" [") +
			color.Red(stats.Index) +
			color.Yellow("/") +
			color.Red(stats.NbFiles) +
			color.Yellow("] ") +
			color.Red("Error: ") +
			color.Yellow(err.Error()))
	}
	atomic.AddUint64(&stats.FileCount, uint64(len(torrent.Data.Info.Files)))

	for _, file := range torrent.Data.Info.Files {
		atomic.AddUint64(&stats.Size, uint64(file.Length))
	}
	fmt.Println(checkPre +
		color.Yellow(" [") +
		color.Green(stats.Index) +
		color.Yellow("/") +
		color.Green(stats.NbFiles) +
		color.Yellow("] ") +
		color.Green(" Extracted infos from ") +
		color.Yellow(path))
	stats.Index++
}
