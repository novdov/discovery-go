package ch05

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Example_marshallJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2020, time.January, 4, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"Title":"Laundry","Status":2,"Deadline":"2020-01-04T15:43:00Z"}
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":2,"Deadline":"2020-01-04T15:43:00Z"}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Buy Milk
	// 2
	// 2020-01-04 15:43:00 +0000 UTC
}

func Example_gob() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N": "J"}
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	const width = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("% x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	fmt.Println(restored)
	// Output:
	// 0e ff 81 04 01 02 ff 82 00 01 0c 01 0c 00 00 08
	// ff 82 00 01 01 4e 01 4a
	// map[N:J]
}