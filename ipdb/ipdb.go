package ipdb

import (
	"bufio"
	"os"
	"strings"
)

type DB map[string]string

func Load(filename string) (DB, error) {
	var macids = make(map[string]string)
	if len(filename) == 0 {
		return macids, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		if len(elements) >= 2 {
			mac := strings.TrimSpace(elements[0])
			mac = strings.ToUpper(strings.ReplaceAll(mac, ":", "-"))
			ip := strings.TrimSpace(elements[1])
			macids[mac] = ip
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// for mac, ip := range macids {
	// 	log.Printf("%s\t%s", mac, ip)
	// }
	return macids, nil
}

func (db DB) Lookup(mac string) string {
	if result, ok := db[mac]; ok {
		return result
	}
	return ""
}
