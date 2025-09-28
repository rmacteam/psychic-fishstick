package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Link struct {
	ID        string
	Original  string
	Short     string
	CreatedAt time.Time
	Clicks    int
}

type LinkStore struct {
	Links map[string]*Link
}

func NewLinkStore() *LinkStore {
	return &LinkStore{Links: make(map[string]*Link)}
}

func generateShort(original string) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	short := make([]rune, 6)
	for i := range short {
		short[i] = chars[rand.Intn(len(chars))]
	}
	return string(short)
}

func (store *LinkStore) Add(original string) string {
	short := generateShort(original)
	link := &Link{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Original:  original,
		Short:     short,
		CreatedAt: time.Now(),
		Clicks:    0,
	}
	store.Links[short] = link
	return short
}

func (store *LinkStore) Click(short string) {
	if l, ok := store.Links[short]; ok {
		l.Clicks++
	}
}

func (store *LinkStore) Analytics(short string) *Link {
	return store.Links[short]
}

func (store *LinkStore) ExportJSON() string {
	b, _ := json.MarshalIndent(store.Links, "", "  ")
	return string(b)
}

func main() {
	store := NewLinkStore()
	short := store.Add("https://github.com/Zarokin/LinkShare-Shortener")
	fmt.Println("Short URL:", short)
	store.Click(short)
	store.Click(short)
	link := store.Analytics(short)
	fmt.Printf("Original: %s\nClicks: %d\n", link.Original, link.Clicks)
	fmt.Println("Exported JSON:")
	fmt.Println(store.ExportJSON())
}