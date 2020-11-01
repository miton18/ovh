package utils

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
