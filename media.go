package gostagram

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Represents every media file's type,
// normaly used to differentiate video, image and
// carousel resources from a
// Media interface array ([]Media).
type MediaType uint8

func (mt MediaType) IsImage() bool {
	return mt == imageMediaType
}

func (mt MediaType) IsVideo() bool {
	return mt == videoMediaType
}

func (mt MediaType) IsCarousel() bool {
	return mt == videoMediaType
}

const (
	imageMediaType    MediaType = 1
	videoMediaType    MediaType = 2
	carouselMediaType MediaType = 3
)

// Media is a generic interface that represents
// all valid instagram media resources
// (Images, videos and carousel images).
type Media interface {
	MediaType() MediaType
}

type Image struct {
	Url    string
	Width  int
	Height int
}

type VideoResolution struct {
	Url    string
	Width  int
	Height int
}

// BaseMediaResource represent all
// attributes that all media resources
// may has.
type BaseMediaResource struct {
	Id          string
	Type        string
	Link        string
	Filter      string
	CreatedTime string `mapstructure:"created_time"`

	User         User
	UserHasLiked bool `mapstructure:"user_has_liked"`
	Attribution  interface{}
	Tags         []string

	UserInPhoto []struct {
		User User

		Position struct {
			X int
			Y int
		}
	} `mapstructure:"user_in_photo"`

	Comments struct {
		Count int
	}

	Caption struct {
		From        User
		Id          string
		Text        string
		CreatedTime string
	}

	Likes struct {
		Count int
	}

	Images struct {
		Thumbnail          Image
		LowResolution      Image `mapstructure:"low_resolution"`
		StandardResolution Image `mapstructure:"standard_resolution"`
	}

	Location struct {
		Id            string
		Name          string
		Latitude      float64
		Longitude     float64
		StreetAddress string `mapstructure:"street_address"`
	}
}

type MediaImage struct {
	BaseMediaResource `mapstructure:",squash"`
}

func (mi MediaImage) MediaType() MediaType {
	return imageMediaType
}

type MediaVideo struct {
	BaseMediaResource `mapstructure:",squash"`

	Videos struct {
		LowResolution      VideoResolution `mapstructure:"low_resolution"`
		StandardResolution VideoResolution `mapstructure:"standard_resolution"`
	}
}

func (mi MediaVideo) MediaType() MediaType {
	return videoMediaType
}

type MediaCarousel struct {
	BaseMediaResource `mapstructure:",squash"`

	CarouselMedia []struct{}
}

func (mi MediaCarousel) MediaType() MediaType {
	return carouselMediaType
}

func (c Client) getMedia(uri string) ([]*Media, error) {
	tmp, _, err := c.get(uri)

	if err != nil {
		return nil, err
	}

	var tmpMediaArray []interface{}

	// checking if media response is an
	// interface array or a map of interfaces.
	switch (*tmp).(type) {
	case []interface{}:
		tmpMediaArray = (*tmp).([]interface{})
		break
	case map[string]interface{}:
		tmpMediaArray = append(tmpMediaArray, (*tmp).(map[string]interface{}))
		break
	}

	var media_array []*Media
	for _, tmpMedia := range tmpMediaArray {
		tmp := tmpMedia.(map[string]interface{})
		mediaType := tmp["type"].(string)

		// check what kind of media resource,
		// was returned. (video, image or carousel image.)
		if mediaType == "image" {
			if tmp["carousel_media"] != nil {

			} else {
				var media MediaImage

				if err := mapstructure.Decode(tmpMedia, &media); err != nil {
					return nil, err
				}

				tt := Media(media)
				media_array = append(media_array, &tt)
			}
		} else if mediaType == "video" {
			var media MediaVideo

			if err := mapstructure.Decode(tmpMedia, &media); err != nil {
				return nil, err
			}

			tt := Media(media)
			media_array = append(media_array, &tt)
		}
	}

	return media_array, nil
}

func (c Client) getOnlyOneMediaContent(uri string) (*Media, error) {
	media, err := c.getMedia(uri)
	if err != nil {
		return nil, err
	}

	return media[0], nil
}

// Get current user media resources
// and how many resources want to return.
func (c Client) GetCurrentUserRecentMedia(max_id, min_id string, count int) ([]*Media, error) {
	return c.getMedia(fmt.Sprintf("%susers/self/media/recent/?max_id=%s&min_id=%s&count=%d&access_token=%s",
		apiUrl, max_id, min_id, count, c.access_token,
	))
}

// Get media resources from respective
// user_id, for more information about it, go to
// https://www.instagram.com/developer/endpoints/users/#get_users_media_recent
func (c Client) GetUserMedia(user_id string, max, min, count int) ([]*Media, error) {
	return c.getMedia(fmt.Sprintf("%susers/%s/media/recent/?max_id=%d&min_id=%d&count=%d&access_token=%s",
		apiUrl, user_id, max, min, count, c.access_token,
	))
}

// Get the recent media liked by the current
// user, for more information aboit it, go to
// https://www.instagram.com/developer/endpoints/users/#get_users_feed_liked
func (c Client) GetCurrentUserMediaLiked(max_like_id string, count int) ([]*Media, error) {
	return c.getMedia(fmt.Sprintf("%susers/self/media/liked?max_like_id=%s&count=%d&access_token=%s",
		apiUrl, max_like_id, count, c.access_token,
	))
}

// Get media resource by id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/media/#get_media
func (c Client) GetMediaById(media_id string) (*Media, error) {
	return c.getOnlyOneMediaContent(fmt.Sprintf("%smedia/%s?access_token=%s",
		apiUrl, media_id, c.access_token))
}

// Get media resouce by its shortcode,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/media/#get_media_by_shortcode
func (c Client) GetMediaByShortcode(short_code string) (*Media, error) {
	return c.getOnlyOneMediaContent(fmt.Sprintf("%smedia/shortcode/%s?access_token=%s",
		apiUrl, short_code, c.access_token))
}

// Get media resouces by latitude, longitude and distance,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/media/#get_media_search
func (c Client) SearchMedia(lat, long, distance float64) ([]*Media, error) {
	return c.getMedia(fmt.Sprintf("%smedia/search?lat=%f&lng=%f&distance=%f&access_token=%s",
		apiUrl, lat, long, distance, c.access_token))
}

// Get media resources that has hashtags equal to 'tagname',
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/tags/#get_tags_media_recent
func (c Client) GetRecentMediaTaggedByTagName(tagname string) ([]*Media, error) {
	return c.getMedia(fmt.Sprintf("%stags/%s/media/recent?access_token=%s", apiUrl, tagname, c.access_token))
}

// Get media resources from a respective location id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/locations/#get_locations_media_recent
func (c Client) GetRecentMediaLocation(location_id string) ([]*Media, error) {
	return c.getMedia(fmt.Sprintf("%slocations/%s/media/recent?access_token=%s",
		apiUrl, location_id, c.access_token))
}
