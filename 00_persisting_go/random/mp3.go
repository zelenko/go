package main

import (
	"fmt"
	"log"

	"github.com/bogem/id3v2"
)

func main() {
	// Open file and parse tag in it.
	tag, err := id3v2.Open("file.mp3", id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Read frames.
	// fmt.Println(tag.Artist())
	// fmt.Println(tag.Title())

	// Set simple text frames.
	// tag.SetArtist("New artist")
	// tag.SetTitle("New title")

	// tags := make(map[string]string)
	// for i, e := range id3v2.V24CommonIDs {
	// 	tags[e] = i
	// }

	tags23 := make(map[string]string)
	for i, e := range id3v2.V23CommonIDs {
		tags23[e] = i
	}

	all := tag.AllFrames()
	for i, e := range all {
		//fmt.Printf("%s %+v\n", i, e)
		for _, fr := range e {
			bpm, ok := fr.(id3v2.TextFrame)
			if !ok {
				fmt.Println("Couldn't assert bpm frame")
				continue
			}
			// tagName := tags[i]
			// tagName23 := tags23[i]

			// fmt.Print(tagName23, "\t==>\t", bpm.Text, "\t(", tagName, ")", "\n")
			//fmt.Print(tags23[i], "\t==>\t", bpm.Text, "\t(", tags[i], ")", "\n")
			fmt.Print(tags23[i], "\t==>\t", bpm.Text, "\n")
			//fmt.Print(tags23[i], "\t==>\t", bpm.Text, "\t", bpm.Encoding.Name, "\n")
		}
	}

	// // Set comment frame.
	// comment := id3v2.CommentFrame{
	// 	Encoding:    id3v2.EncodingUTF8,
	// 	Language:    "eng",
	// 	Description: "My opinion",
	// 	Text:        "Very good song",
	// }
	// tag.AddCommentFrame(comment)

	// // Write it to file.
	// if err = tag.Save(); err != nil {
	// 	log.Fatal("Error while saving a tag: ", err)
	// }
}
