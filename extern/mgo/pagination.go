package main

import (
	"flag"
	"fmt"
	"labix.org/v2/mgo"
)

type Number struct {
	N int
}

func main() {

	// Setup some command line flags.
	var page = flag.Int("p", 1, "Specify a page number to display.")
	var ppage = flag.Int("pp", 15, "Specify the number of items to display per page.")
	flag.Parse()

	// Start a database session.
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Get the numbers collection.
	c := session.DB("test").C("numbers")

	// Get the total number of items.
	count, err := c.Count()
	if err != nil {
		panic(err)
	}

	// Calculate the total number of pages
	var tpages int
	if count%*ppage != 0 {
		tpages = count / *ppage + 1
	} else {
		tpages = count / *ppage
	}

	if tpages < *page {
		*page = tpages
	}

	// The number of documents to skip.
	skip := *ppage * (*page - 1)

	result := []Number{}
	err = c.Find(nil).Limit(*ppage).Skip(skip).All(&result)
	if err != nil {
		panic(err)
	}

	total, err := c.Count()
	if err != nil {
		panic(err)
	}

	// Calculate the highest value displayed.
	var hval int
	if skip+*ppage > total {
		hval = total
	} else {
		hval = skip + *ppage
	}

	fmt.Printf("Showing results %d to %d of %d\n", skip+1, hval, total)
	fmt.Println(result)
	fmt.Printf("Showing page %d of %d\n", *page, tpages)
}
