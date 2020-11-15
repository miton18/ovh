package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Page struct {
	headers http.Header
}

func NewPage(headers http.Header) Page {
	return Page{headers}
}

// Get current page
func (p Page) Current() int64 {
	c, _ := strconv.ParseInt(p.headers.Get("X-Pagination-Number"), 10, 64)
	return c
}

// Current page size
func (p Page) Size() int64 {
	t, _ := strconv.ParseInt(p.headers.Get("X-Pagination-Size"), 10, 64)
	return t
}

// Total page count
func (p Page) Total() int64 {
	t, _ := strconv.ParseInt(p.headers.Get("X-Pagination-Total"), 10, 64)
	return t
}

func (p Page) ItemsCount() int64 {
	t, _ := strconv.ParseInt(p.headers.Get("X-Pagination-Elements"), 10, 64)
	return t
}

func (p Page) LastUpdate() time.Time {
	s := p.headers.Get("X-Pagination-Cache-Update")
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func (p Page) String() string {
	return fmt.Sprintf("Page %d/%d (%d items) %s", p.Current(), p.Total(), p.Size(), time.Since(p.LastUpdate()))
}

type Cursor struct {
	headers http.Header
}

// Get current cursor
func (c Cursor) Current() string {
	return c.headers.Get("X-Pagination-Cursor")
}

// Current next cursor
func (c Cursor) Next() string {
	return c.headers.Get("X-Pagination-Cursor-Next")
}

// Current items count
func (c Cursor) Size() int64 {
	t, _ := strconv.ParseInt(c.headers.Get("X-Pagination-Size"), 10, 64)
	return t
}

// Last cache update
func (c Cursor) LastUpdate() time.Time {
	s := c.headers.Get("X-Pagination-Cache-Update")
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func (c Cursor) String() string {
	return fmt.Sprintf("Cursor %s(%d items) -> %s %s", c.Current(), c.Size(), c.Next(), time.Since(c.LastUpdate()))
}

/* TODO: Find a better implementation
type Iterator struct {
	error error
	items int
	item  int
}

func NewIterator(client ovh.Client, path string) (*Iterator){
	limit := 100
	offset := 0
	hasNext := true
	stream := make(chan interface{})

	for hasNext {
		req, err := client.NewRequest("GET", path, nil, true)
		if err != nil {
			return nil, err
		}

		req.Header.Set("X-Pagination-Mode", "CachedObjectList-Pages")
		req.Header.Set("X-Pagination-Size", fmt.Sprintf("%d", limit))
		if offset != 0 {
			req.Header.Set("X-Pagination-Number", fmt.Sprintf("%d", offset))
		}

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		res.

		offset++
	}



	err = client.UnmarshalResponse(res, &loadbalancers)
}
*/
