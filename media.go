package gostagram

import (
	"fmt"
	"strconv"
	"errors"

	"github.com/mitchellh/mapstructure"
)

// Media is a generic interface that represents
// all valid instagram media resources
// (Images, videos and carousel images).
type Media interface {
	MediaType() MediaType
}


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
	return mt == carouselMediaType
}

const (
	imageMediaType    MediaType = 1
	videoMediaType    MediaType = 2
	carouselMediaType MediaType = 3

	maxDistance = 5000
)


var (
	maxDistanceError = errors.New("Maximun distance is 5km.")
)

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

// MediaImage struct represents
// an Image resource that instagram
// endpoint returns.
type MediaImage struct {
	BaseMediaResource `mapstructure:",squash"`
}

func (mi MediaImage) MediaType() MediaType {
	return imageMediaType
}

// MediaVideo struct represents
// an Image resource that instagram
// endpoint returns.
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

// MediaCarousel struct represents
// an Image resource that instagram
// endpoint returns.
type MediaCarousel struct {
	BaseMediaResource `mapstructure:",squash"`

	CarouselMedia []Media
}

func (mi MediaCarousel) MediaType() MediaType {
	return carouselMediaType
}

// Carousel resource can have both image and video
// resources with it, those resources base attributes
// are with are in BaseMediaCarousel struct.
type BaseMediaCarousel struct {
	Type string

	UserInPhoto []struct {
		User User

		Position struct {
			X int
			Y int
		}
	} `mapstructure:"user_in_photo"`
}

type MediaCarouselImage struct {
	BaseMediaCarousel `mapstructure:",squash"`

	Images struct {
		Thumbnail          Image
		LowResolution      Image `mapstructure:"low_resolution"`
		StandardResolution Image `mapstructure:"standard_resolution"`
	}
}

func (mi MediaCarouselImage) MediaType() MediaType {
	return imageMediaType
}

type MediaCarouselVideo struct {
	BaseMediaCarousel `mapstructure:",squash"`

	Videos struct {
		LowResolution      VideoResolution `mapstructure:"low_resolution"`
		LowBandwidth       VideoResolution `mapstructure:"low_bandwidth"`
		StandardResolution VideoResolution `mapstructure:"standard_resolution"`
	}
}

func (mi MediaCarouselVideo) MediaType() MediaType {
	return videoMediaType
}

func (c Client) getMedia(uri string) ([]*Media, error) {
	tmp, _, err := c.get(uri)

	if err != nil {
		return nil, err
	}

	var tmpMediaCollection []interface{}

	// checking if media response is an
	// interface array or a map of interfaces.
	switch (*tmp).(type) {
	case []interface{}:
		tmpMediaCollection = (*tmp).([]interface{})
		break
	case map[string]interface{}:
		tmpMediaCollection = append(tmpMediaCollection, (*tmp).(map[string]interface{}))
		break
	}

	var mediaCollection []*Media
	for _, tmpMediaInterface := range tmpMediaCollection {
		tmpMedia := tmpMediaInterface.(map[string]interface{})
		mediaType := tmpMedia["type"].(string)
		var media Media

		// check what kind of media resource,
		// was returned. (video, image or carousel image.)
		if mediaType == "image" {

			// carousel and image resources
			// are both an image type.

			if tmpMedia["carousel_media"] != nil {
				var mediaCarousel MediaCarousel

				if err := mapstructure.Decode(tmpMediaInterface, &mediaCarousel); err != nil {
					return nil, err
				}

				mediaCarousel.CarouselMedia = []Media{}
				carouselMediaType := tmpMedia["carousel_media"].([]map[string]interface{})

				for _, tmpcarouselMedia := range carouselMediaType {
					tmpcarouselType := tmpcarouselMedia["type"].(string)
					var media2 Media

					if tmpcarouselType == "image" {
						var mediaCarouselImage MediaCarouselImage

						if err := mapstructure.Decode(tmpcarouselMedia, &mediaCarouselImage); err != nil {
							return nil, err
						}

						media2 = Media(mediaCarouselImage)

					} else if tmpcarouselType == "video" {
						var mediaCarouselVideo MediaCarouselVideo

						if err := mapstructure.Decode(tmpcarouselMedia, &mediaCarouselVideo); err != nil {
							return nil, err
						}

						media2 = Media(mediaCarouselVideo)
					}

					// appending resources to carousel.
					mediaCarousel.CarouselMedia = append(mediaCarousel.CarouselMedia, media2)
				}

				media = Media(mediaCarousel)
			} else {
				var mediaImage MediaImage

				if err := mapstructure.Decode(tmpMedia, &mediaImage); err != nil {
					return nil, err
				}

				media = Media(mediaImage)
			}
		} else if mediaType == "video" {
			var mediaVideo MediaVideo

			if err := mapstructure.Decode(tmpMedia, &mediaVideo); err != nil {
				return nil, err
			}

			media = Media(mediaVideo)
		}

		mediaCollection = append(mediaCollection, &media)
	}

	return mediaCollection, nil
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
func (c Client) GetCurrentUserRecentMedia(params Parameters) ([]*Media, error) {
	tmp := "%susers/self/media/recent/?access_token=%s"

	if params != nil {
		if params["max_id"] != "" {
			tmp += fmt.Sprintf("&max_id=%s", params["max_id"])
		}

		if params["min_id"] != "" {
			tmp += fmt.Sprintf("&min_id=%s", params["min_id"])
		}

		if params["count"] != "" {
			tmp += fmt.Sprintf("&count=%s", params["count"])
		}
	}

	return c.getMedia(fmt.Sprintf(tmp, apiUrl, c.access_token))
}

// Get media resources from respective
// user_id, for more information about it, go to
// https://www.instagram.com/developer/endpoints/users/#get_users_media_recent
func (c Client) GetUserMedia(user_id string, params Parameters) ([]*Media, error) {
	tmp := "%susers/%s/media/recent/?access_token=%s"
	if params != nil {
		if params["max_id"] != "" {
			tmp += fmt.Sprintf("&max_id=%s", params["max_id"])
		}

		if params["min_id"] != "" {
			tmp += fmt.Sprintf("&min_id=%s", params["min_id"])
		}

		if params["count"] != "" {
			tmp += fmt.Sprintf("&count=%s", params["count"])
		}
	}

	return c.getMedia(fmt.Sprintf(tmp, apiUrl, user_id, c.access_token))
}

// Get the recent media liked by the current
// user, for more information aboit it, go to
// https://www.instagram.com/developer/endpoints/users/#get_users_feed_liked
func (c Client) GetCurrentUserMediaLiked(max_like_id string, parameters Parameters) ([]*Media, error) {
	tmp := "%susers/self/media/liked?max_like_id=%s&access_token=%s"
	if parameters != nil {
		if parameters["count"] != "" {
			tmp += fmt.Sprintf("&count=%s", parameters["count"])
		}
	}

	return c.getMedia(fmt.Sprintf(tmp,
		apiUrl, max_like_id, c.access_token,
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
func (c Client) SearchMedia(lat, long string, params Parameters) ([]*Media, error) {
	tmp := "%smedia/search?lat=%s&lng=%s&access_token=%s"
	if params != nil {
		if params["distance"] != "" {
			distance, err := strconv.Atoi(params["distance"])
			if err != nil {
				return nil, err
			}

			if distance > maxDistance {
				return nil, maxDistanceError
			}

			tmp += fmt.Sprintf("&distance=%d", distance)
		}
	}

	return c.getMedia(fmt.Sprintf(tmp, apiUrl, lat, long, c.access_token))
}

// Get media resources that has hashtags equal to 'tagname',
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/tags/#get_tags_media_recent
func (c Client) GetRecentMediaTaggedByTagName(tagname string, params Parameters) ([]*Media, error) {
	tmp := "%stags/%s/media/recent?access_token=%s"
	if params != nil {
		if params["max_tag_id"] != "" {
			tmp += fmt.Sprintf("&max_tag_id=%s", params["max_tag_id"])
		}

		if params["min_tag_id"] != "" {
			tmp += fmt.Sprintf("&min_tag_id=%s", params["min_tag_id"])
		}

		if params["count"] != "" {
			tmp += fmt.Sprintf("&count=%s", params["count"])
		}
	}

	return c.getMedia(fmt.Sprintf(tmp, apiUrl, tagname, c.access_token))
}

// Get media resources from a respective location id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/locations/#get_locations_media_recent
func (c Client) GetRecentMediaLocation(location_id string, params Parameters) ([]*Media, error) {
	tmp := "%slocations/%s/media/recent?access_token=%s"
	if params != nil {
		if params["max_id"] != "" {
			tmp += fmt.Sprintf("&max_id=%s", params["max_id"])
		}

		if params["min_id"] != "" {
			tmp += fmt.Sprintf("&min_id=%s", params["min_id"])
		}
	}

	return c.getMedia(fmt.Sprintf(tmp,
		apiUrl, location_id, c.access_token))
}
