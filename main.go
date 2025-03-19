package main

import "wliafdew/go-2025/fakeData"

func main() {
	// init db connection
	fakeData.ImportFakeUserPg()
	// just import the fake data importer want to use

}
