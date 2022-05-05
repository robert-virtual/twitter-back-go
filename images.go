package main

import "fmt"

func findImages(postId *string) ([]image, error) {
	var images []image
	rows, error := db.Query("SELECT * FROM images where postId = ?", postId)
	if error != nil {
		return nil, fmt.Errorf("images %q:%v", *postId, error)
	}
	defer rows.Close()
	for rows.Next() {
		var image image
		if err := rows.Scan(&image.Id, &image.Url, &image.Url); err != nil {
			return nil, fmt.Errorf("images %q:%v", *postId, err)
		}
		images = append(images, image)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("images %q:%v", *postId, err)
	}
	return images, nil

}
