package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"


	"github.com/pedramN/HackBrowserData/utils"
	"github.com/jszwec/csvutil"
)

var (
	utf8Bom = []byte{239, 187, 191}
)

func (b *Bookmarks) outPutJson(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "bookmark", "json")
	sort.Slice(b.Bookmarks, func(i, j int) bool {
		return b.Bookmarks[i].ID < b.Bookmarks[j].ID
	})
	err := writeToJson(filename, b.Bookmarks)
	if err != nil {
		return err
	}
	fmt.Printf("%s Get %d Bookmarks, filename is %s \n", utils.Prefix, len(b.Bookmarks), filename)
	return nil
}

func (h *HistoryData) outPutJson(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "history", "json")
	sort.Slice(h.History, func(i, j int) bool {
		return h.History[i].VisitCount > h.History[j].VisitCount
	})
	err := writeToJson(filename, h.History)
	if err != nil {
		return err
	}
	fmt.Printf("%s Get %d history, filename is %s \n", utils.Prefix, len(h.History), filename)
	return nil
}

func (d *downloads) outPutJson(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "download", "json")
	err := writeToJson(filename, d.downloads)
	if err != nil {
		return err
	}
	fmt.Printf("%s Get %d history, filename is %s \n", utils.Prefix, len(d.downloads), filename)
	return nil
}

func (p *passwords) outPutJson(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "password", "json")
	err := writeToJson(filename, p.logins)
	if err != nil {
		return err
	}
	fmt.Printf("%s Get %d passwords, filename is %s \n", utils.Prefix, len(p.logins), filename)
	return nil
}

func (c *Cookies) outPutJson(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "cookie", "json")
	err := writeToJson(filename, c.Cookies)
	if err != nil {
		return err
	}
	fmt.Printf("%s Get %d cookies, filename is %s \n", utils.Prefix, len(c.Cookies), filename)
	return nil
}

func (c *creditCards) outPutJson(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "credit", "json")
	err := writeToJson(filename, c.cards)
	if err != nil {
		return err
	}
	fmt.Printf("%s Get %d credit cards, filename is %s \n", utils.Prefix, len(c.cards), filename)
	return nil
}

func writeToJson(filename string, data interface{}) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	w := new(bytes.Buffer)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	err = enc.Encode(data)
	if err != nil {
		return err
	}
	_, err = f.Write(w.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (b *Bookmarks) outPutCsv(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "bookmark", "csv")
	if err := writeToCsv(filename, b.Bookmarks); err != nil {
		return err
	}
	fmt.Printf("%s Get %d Bookmarks, filename is %s \n", utils.Prefix, len(b.Bookmarks), filename)
	return nil
}

func (h *HistoryData) outPutCsv(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "history", "csv")
	if err := writeToCsv(filename, h.History); err != nil {
		return err
	}
	fmt.Printf("%s Get %d history, filename is %s \n", utils.Prefix, len(h.History), filename)
	return nil
}

func (d *downloads) outPutCsv(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "download", "csv")
	if err := writeToCsv(filename, d.downloads); err != nil {
		return err
	}
	fmt.Printf("%s Get %d download history, filename is %s \n", utils.Prefix, len(d.downloads), filename)
	return nil
}

func (p *passwords) outPutCsv(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "password", "csv")
	if err := writeToCsv(filename, p.logins); err != nil {
		return err
	}
	fmt.Printf("%s Get %d passwords, filename is %s \n", utils.Prefix, len(p.logins), filename)
	return nil
}

func (c *Cookies) outPutCsv(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "cookie", "csv")
	var tempSlice []cookie
	for _, v := range c.Cookies {
		tempSlice = append(tempSlice, v...)
	}
	if err := writeToCsv(filename, tempSlice); err != nil {
		return err
	}
	fmt.Printf("%s Get %d cookies, filename is %s \n", utils.Prefix, len(c.Cookies), filename)
	return nil
}

func (c *creditCards) outPutCsv(browser, dir string) error {
	filename := utils.FormatFileName(dir, browser, "credit", "csv")
	var tempSlice []card
	for _, v := range c.cards {
		tempSlice = append(tempSlice, v...)
	}
	if err := writeToCsv(filename, tempSlice); err != nil {
		return err
	}
	fmt.Printf("%s Get %d credit cards, filename is %s \n", utils.Prefix, len(c.cards), filename)
	return nil
}

func writeToCsv(filename string, data interface{}) error {
	var d []byte
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(utf8Bom)
	if err != nil {
		return err
	}
	d, err = csvutil.Marshal(data)
	if err != nil {
		return err
	}
	_, err = f.Write(d)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bookmarks) outPutConsole() {
	for _, v := range b.Bookmarks {
		fmt.Printf("%+v\n", v)
	}
}

func (c *Cookies) outPutConsole() {
	for host, value := range c.Cookies {
		fmt.Printf("%s\n%+v\n", host, value)
	}
}

func (h *HistoryData) outPutConsole() {
	for _, v := range h.History {
		fmt.Printf("%+v\n", v)
	}
}

func (d *downloads) outPutConsole() {
	for _, v := range d.downloads {
		fmt.Printf("%+v\n", v)
	}
}

func (p *passwords) outPutConsole() {
	for _, v := range p.logins {
		fmt.Printf("%+v\n", v)
	}
}

func (c *creditCards) outPutConsole() {
	for _, v := range c.cards {
		fmt.Printf("%+v\n", v)
	}
}
