package gostagram

import (
	"testing"
)

var (
	media_id  = ""
	shortcode = ""
)

func LogMedia(media *Media, t *testing.T) {
	if (*media).MediaType().IsImage() {
		mediaImage := (*media).(MediaImage)

		t.Log("------------------ Start Image Media ------------------")
		t.Log("Id: ", DefaultStringValueIfEmpty(mediaImage.Id))
		t.Log("Type: ", DefaultStringValueIfEmpty(mediaImage.Type))
		t.Log("Link: ", DefaultStringValueIfEmpty(mediaImage.Link))
		t.Log("Filter: ", DefaultStringValueIfEmpty(mediaImage.Filter))
		t.Log("Created Time: ", DefaultStringValueIfEmpty(mediaImage.CreatedTime))

		t.Log("------------------ Media Start User ------------------")
		LogUser(&mediaImage.User, t)
		t.Log("------------------ Media  End User ------------------")

		t.Log("User Has Liked: ", mediaImage.UserHasLiked)
		t.Log("Attribution: ", mediaImage.Attribution)
		t.Log("Tags: ", mediaImage.Tags)

		t.Log("User in Photo: ")
		for _, tmp := range mediaImage.UserInPhoto {
			t.Log("User in Photo User: ")
			LogUser(&tmp.User, t)
			t.Log("User In Photo Position X: ", tmp.Position.X)
			t.Log("User In Photo Position Y: ", tmp.Position.Y)
		}

		t.Log("Comments Count: ", mediaImage.Comments.Count)
		t.Log("Caption Id: ", DefaultStringValueIfEmpty(mediaImage.Caption.Id))
		t.Log("Caption Created Time: ", DefaultStringValueIfEmpty(mediaImage.Caption.CreatedTime))
		t.Log("Caption Created Text: ", DefaultStringValueIfEmpty(mediaImage.Caption.Text))
		t.Log("Caption User: ")
		LogUser(&mediaImage.Caption.From, t)
		t.Log("Likes Count: : ", mediaImage.Likes.Count)

		t.Log("------------------ Start Thumbnail Videos ------------------")
		LogImage(&mediaImage.Images.Thumbnail, t)
		t.Log("------------------ End Thumbnail Videos ------------------")

		t.Log("------------------ Start Low Resolution Image ------------------")
		LogImage(&mediaImage.Images.LowResolution, t)
		t.Log("------------------ End Low Resolution Image ------------------")

		t.Log("------------------ Start Standard Resolution Image ------------------")
		LogImage(&mediaImage.Images.StandardResolution, t)
		t.Log("------------------ End Standard Resolution Image ------------------")

		t.Log("Location Id: ", DefaultStringValueIfEmpty(mediaImage.Location.Id))
		t.Log("Location Name: ", DefaultStringValueIfEmpty(mediaImage.Location.Name))
		t.Log("Location Latitude: ", mediaImage.Location.Latitude)
		t.Log("Location Longitude: ", mediaImage.Location.Longitude)
		t.Log("Location Street Address: ", DefaultStringValueIfEmpty(mediaImage.Location.StreetAddress))

		t.Log("------------------ End Image Media ------------------")

	} else if (*media).MediaType().IsVideo() {
		mediaVideo := (*media).(MediaVideo)

		t.Log("------------------ Start Video Media ------------------")
		t.Log("Id: ", DefaultStringValueIfEmpty(mediaVideo.Id))
		t.Log("Type: ", DefaultStringValueIfEmpty(mediaVideo.Type))
		t.Log("Link: ", DefaultStringValueIfEmpty(mediaVideo.Link))
		t.Log("Filter: ", DefaultStringValueIfEmpty(mediaVideo.Filter))
		t.Log("Created Time: ", DefaultStringValueIfEmpty(mediaVideo.CreatedTime))

		t.Log("------------------ Media Start User ------------------")
		LogUser(&mediaVideo.User, t)
		t.Log("------------------ Media  End User ------------------")

		t.Log("User Has Liked: ", mediaVideo.UserHasLiked)
		t.Log("Attribution: ", mediaVideo.Attribution)
		t.Log("Tags: ", mediaVideo.Tags)

		t.Log("User in Photo: ")
		for _, tmp := range mediaVideo.UserInPhoto {
			t.Log("User in Photo User:")
			LogUser(&tmp.User, t)
			t.Log("User In Photo Position X: ", tmp.Position.X)
			t.Log("User In Photo Position Y: ", tmp.Position.Y)
		}

		t.Log("Comments Count: ", mediaVideo.Comments.Count)
		t.Log("Caption Id: ", DefaultStringValueIfEmpty(mediaVideo.Caption.Id))
		t.Log("Caption Created Time: ", DefaultStringValueIfEmpty(mediaVideo.Caption.CreatedTime))
		t.Log("Caption Created Text: ", DefaultStringValueIfEmpty(mediaVideo.Caption.Text))
		t.Log("Caption User: ")
		LogUser(&mediaVideo.Caption.From, t)
		t.Log("Likes Count: : ", mediaVideo.Likes.Count)

		t.Log("------------------ Start Thumbnail Videos ------------------")
		LogImage(&mediaVideo.Images.Thumbnail, t)
		t.Log("------------------ End Thumbnail Videos ------------------")

		t.Log("------------------ Start Low Resolution Image ------------------")
		LogImage(&mediaVideo.Images.LowResolution, t)
		t.Log("------------------ End Low Resolution Image ------------------")

		t.Log("------------------ Start Standard Resolution Image ------------------")
		LogImage(&mediaVideo.Images.StandardResolution, t)
		t.Log("------------------ End Standard Resolution Image ------------------")

		t.Log("Location Id: ", DefaultStringValueIfEmpty(mediaVideo.Location.Id))
		t.Log("Location Name: ", DefaultStringValueIfEmpty(mediaVideo.Location.Name))
		t.Log("Location Latitude: ", mediaVideo.Location.Latitude)
		t.Log("Location Longitude: ", mediaVideo.Location.Longitude)
		t.Log("Location Street Address: ", DefaultStringValueIfEmpty(mediaVideo.Location.StreetAddress))

		t.Log("------------------ Start Standard Resolution Videos ------------------")
		LogVideo(&mediaVideo.Videos.StandardResolution, t)
		t.Log("------------------ End Standard Resolution Videos ------------------")

		t.Log("------------------ End Start Resolution Videos ------------------")
		LogVideo(&mediaVideo.Videos.LowResolution, t)
		t.Log("------------------ End Low Resolution Videos ------------------")

		t.Log("------------------ End Video Media ------------------")
	}
}

func LogImage(image *Image, t *testing.T) {
	t.Log("Image Url: ", DefaultStringValueIfEmpty(image.Url))
	t.Log("Image Height: ", image.Height)
	t.Log("Image Width: ", image.Width)
}

func LogVideo(video *VideoResolution, t *testing.T) {
	t.Log("Video Url: ", DefaultStringValueIfEmpty(video.Url))
	t.Log("Video Width: ", video.Width)
	t.Log("Video Height: ", video.Height)
}

func IterateMedia(media_arr []*Media, t *testing.T) {
	if len(media_arr) > 0 {
		for _, media := range media_arr {
			LogMedia(media, t)
		}
	} else {
		t.Log("No content.")
	}
}

func TestClient_GetCurrentUserRecentMedia(t *testing.T) {
	client := CreateClient(t)
	media_arr, err := client.GetCurrentUserRecentMedia("1", "1", 1)
	PanicIfError(err, t)
	IterateMedia(media_arr, t)
}

func TestClient_GetUserMedia(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	media_arr, err := client.GetUserMedia(other_user_id, 1, 1, 1)
	PanicIfError(err, t)
	IterateMedia(media_arr, t)
}

func TestClient_GetCurrentUserMediaLiked(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	media_arr, err := client.GetCurrentUserMediaLiked(media_id, 1)
	PanicIfError(err, t)
	IterateMedia(media_arr, t)
}

func TestClient_GetMediaById(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	media, err := client.GetMediaById(media_id)
	PanicIfError(err, t)
	LogMedia(media, t)
}

func TestClient_GetMediaByShortcode(t *testing.T) {
	FatalIfEmptyString(shortcode, "shortcode cannot be empty.", t)
	client := CreateClient(t)
	media, err := client.GetMediaByShortcode(shortcode)
	PanicIfError(err, t)
	LogMedia(media, t)
}

func TestClient_SearchMedia(t *testing.T) {
	client := CreateClient(t)
	media_arr, err := client.SearchMedia(0, 0, 300)
	PanicIfError(err, t)
	IterateMedia(media_arr, t)
}

func TestClient_GetRecentMediaTaggedByTagName(t *testing.T) {
	FatalIfEmptyString(tagname, "tagname cannot be empty.", t)
	client := CreateClient(t)
	media_arr, err := client.GetRecentMediaTaggedByTagName(tagname)
	PanicIfError(err, t)
	IterateMedia(media_arr, t)
}

func TestClient_GetRecentMediaLocation(t *testing.T) {
	FatalIfEmptyString(location_id, "Location id cannot be empty.", t)
	client := CreateClient(t)
	media_arr, err := client.GetRecentMediaLocation(location_id)
	PanicIfError(err, t)
	IterateMedia(media_arr, t)
}
