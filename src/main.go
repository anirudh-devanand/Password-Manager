package main

import(
	"fmt"
	"github.com/anirudh-devanand/PwdMngr-Go/src/models"
)



func main(){

// EntryDB is a map used to store website and password as key-value pairs, if EntryDB doesnt exist we must create
	var EntryDB = make(map[string]string)	
	
	fmt.Println("SYNTAX: <OPERATION> <WEBSITE>* <PASSWORD>** \nOPERATIONS:\nCREATE\tREAD\tUPDATE\tDELETE\texit\n*Using the READ operation with website as \"NULL\" and a password will return webiste\n** Using UPDATE operation with passowrd as \"\" will delete website entry")
	for {
		var (
			ops string
			entry models.Entry
			web string
			pwd string
		)

		fmt.Scanf("%s %s %s", &ops, &web, &pwd)

		if ops == "exit"{
			break;
		}

		entry.Website=web
		entry.Password=pwd
		entry.EntryOps(ops, EntryDB)

	}

}