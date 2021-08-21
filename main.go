package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

type File struct {
	FileName       string
	PhotoDNA       []byte
	BestMatchIndex int
	BestMatchScore int
	lock           sync.Mutex
}

func main() {
	log.Println("loading data")
	hashCsv, err := os.Open("hashes.csv")
	if err != nil {
		panic(err)
	}
	defer hashCsv.Close()

	csvLines, err := csv.NewReader(hashCsv).ReadAll()
	if err != nil {
		panic(err)
	}

	files := make([]File, 0, len(csvLines))

	for _, line := range csvLines {
		pdna, _ := base64.StdEncoding.DecodeString(line[1])
		if len(pdna) != 144 {
			panic("wrong hash length for file: " + line[0])
		}
		matchInfo := File{
			FileName:       line[0],
			PhotoDNA:       pdna,
			BestMatchIndex: -1,
			BestMatchScore: int(^uint(0) >> 1),
		}
		files = append(files, matchInfo)
	}

	// log.Println("loaded", len(files), "files")

	var wg sync.WaitGroup

	log.Println("starting work")

	for index, _ := range files {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			hashForFile(index, files)
		}(index)
	}

	wg.Wait()

	log.Println("workers finished")

	for i, file := range files {
		fmt.Println(i, file.FileName, file.BestMatchScore, "from", file.BestMatchIndex, files[file.BestMatchIndex].FileName)
	}
}

func hashForFile(index int, files []File) {
	thisPdna := files[index].PhotoDNA

	for i := index + 1; i < len(files); i++ {
		distance := calcDistance(thisPdna, files[i].PhotoDNA)

		if distance == 0 {
			// log.Println("skipping identical images", files[index].FileName, files[i].FileName)
			continue
		}

		if distance > files[index].BestMatchScore && distance > files[i].BestMatchScore {
			// bad match, don't bother locking for updates
			continue
		}

		files[index].lock.Lock()
		if distance < files[index].BestMatchScore || (distance == files[index].BestMatchScore && i < files[index].BestMatchIndex) {
			// log.Println("distance:", distance, "from", i)
			files[index].BestMatchIndex = i
			files[index].BestMatchScore = distance
		}
		files[index].lock.Unlock()

		files[i].lock.Lock()
		if distance < files[i].BestMatchScore || (distance == files[i].BestMatchScore && index < files[i].BestMatchIndex) {
			// log.Println("updating from other side:", i, "is", distance, "from", index)
			files[i].BestMatchIndex = index
			files[i].BestMatchScore = distance
		}
		files[i].lock.Unlock()
	}

	if index%1000 == 0 {
		log.Println("worker finished:", index)
	}
}

func calcDistance(h1, h2 []byte) (distance int) {
	if len(h1) != 144 {
		panic("h1 wrong length")
	}
	if len(h2) != 144 {
		panic("h2 wrong length")
	}
	for i := 0; i < 144; i++ {
		distance += (int(h1[i]) - int(h2[i])) * (int(h1[i]) - int(h2[i]))
	}
	return
}
