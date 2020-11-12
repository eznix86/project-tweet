package repository

import (
	"errors"

	"github.com/theArtechnology/project-tweet/src/entity"
)

// TweetRepository is a Repository (meaning Storage) for all my tweets
type TweetRepository struct {
	TweetList []entity.Tweet // in memory TweetList
}

// Find is
func (r *TweetRepository) Find(id int) (*entity.Tweet, error) {
	for _, tweet := range r.TweetList {
		if tweet.ID == id {
			return &tweet, nil
		}
	}
	return nil, errors.New("tweet not found")
}

// Add is
func (r *TweetRepository) Add(entity entity.Tweet) error {
	r.TweetList = append(r.TweetList, entity)
	return nil
}

// Remove is
func (r *TweetRepository) Remove(id int) error {
	for index, tweet := range r.TweetList {
		if tweet.ID == id {
			r.TweetList = append(r.TweetList[:index], r.TweetList[index:]...)
			return nil
		}
	}

	return errors.New("tweet not found")
}

func (r *TweetRepository) getAll(entity entity.Tweet) ([]entity.Tweet, error) {
	return r.TweetList, nil
}
