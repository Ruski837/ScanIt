package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/amyruski/ScanIt/utils"
)

func main() {
	MusicExt := []string{".mp3", ".wav"}
	ImageExt := []string{".jpg", ".pdf", ".png"}
	MovieExt := []string{".mov", ".mp4"}
	var MusicSlice []string
	var ImageSlice []string
	var MovieSlice []string
	fmt.Println("Starting to walk through a filesystem...")
	root := flag.String("root-dir", "nil", "The root dir where you want your search to start.")
	ImagePath := flag.String("image-dir", "nil", "The root dir where you want your search to start.")
	MoviePath := flag.String("movie-dir", "nil", "The root dir where you want your search to start.")
	MusicPath := flag.String("music-dir", "nil", "The root dir where you want your search to start.")

	flag.Parse()
	fmt.Println("root:", *root)
	if *root == "nil" {
		log.Fatal("Trying to run cli with no root specified. Please set the root directory.")
	}
	err := filepath.WalkDir(*root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("error encountered while walking the path", err)
			return err
		}

		if utils.Contains(MusicExt, filepath.Ext(path)) && *MusicPath != "nil" {
			p, err := filepath.Abs(path)
			if err != nil {
				log.Fatal(err)
			}
			MusicSlice = append(MusicSlice, p)
		}
		if utils.Contains(ImageExt, filepath.Ext(path)) && *ImagePath != "nil" {
			p, err := filepath.Abs(path)
			if err != nil {
				log.Fatal(err)
			}
			ImageSlice = append(ImageSlice, p)
		}
		if utils.Contains(MovieExt, filepath.Ext(path)) && *MoviePath != "nil" {
			p, err := filepath.Abs(path)
			if err != nil {
				log.Fatal(err)
			}
			MovieSlice = append(MovieSlice, p)
		}

		return nil

	})
	if err != nil {
		fmt.Println("error walking the path", err)
	}
	if *MusicPath != "nil" {
		for _, ms := range MusicSlice {
			_, v := filepath.Split(ms)
			fmt.Println("moving", ms, "to ", filepath.Join(*MusicPath, v))
			err = os.Rename(ms, filepath.Join(*MusicPath, v))
			if err != nil {
				log.Fatal("Cannot move files:", err)
			}
		}
	} else {
		fmt.Println("Skiping Music files because the music-dir was not set.")
	}
	if *MoviePath != "nil" {
		for _, mvs := range MovieSlice {
			_, m := filepath.Split(mvs)
			fmt.Println("moving", mvs, "to ", filepath.Join(*MoviePath, m))
			err = os.Rename(mvs, filepath.Join(*MoviePath, m))
			if err != nil {
				log.Fatal("Cannot move files:", err)
			}
		}
	} else {
		fmt.Println("Skiping movie files because the movie-dir was not set.")
	}
	if *ImagePath != "nil" {
		for _, imgs := range ImageSlice {
			_, f := filepath.Split(imgs)
			fmt.Println("moving", imgs, "to ", filepath.Join(*ImagePath, f))

			fmt.Println(f)
			err = os.Rename(imgs, filepath.Join(*ImagePath, f))
			if err != nil {
				log.Fatal("Cannot move files:", err)
			}
		}
	} else {
		fmt.Println("Skiping image files because the image-dir was not set.")
	}
}

// env GOOS=windows GOARCH=amd64
