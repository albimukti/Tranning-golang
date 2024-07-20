package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/albimukti/assignment_3_albi/internal/cache"
	"github.com/albimukti/assignment_3_albi/internal/database"
	"github.com/albimukti/assignment_3_albi/internal/proto"
	"github.com/teris-io/shortid"
)

type URLShortenerService struct {
	proto.UnimplementedURLShortenerServer
}

func (s *URLShortenerService) ShortenURL(ctx context.Context, req *proto.ShortenURLRequest) (*proto.ShortenURLResponse, error) {
	longURL := req.LongUrl
	shortURL, err := shortid.Generate()
	if err != nil {
		return nil, err
	}

	_, err = database.DB.Exec(context.Background(), "INSERT INTO url_mappings (short_url, long_url) VALUES ($1, $2)", shortURL, longURL)
	if err != nil {
		return nil, err
	}

	cache.C.Set(shortURL, longURL, 5*time.Minute)

	return &proto.ShortenURLResponse{ShortUrl: shortURL}, nil
}

func (s *URLShortenerService) GetLongURL(ctx context.Context, req *proto.GetLongURLRequest) (*proto.GetLongURLResponse, error) {
	shortURL := req.ShortUrl

	if longURL, found := cache.C.Get(shortURL); found {
		return &proto.GetLongURLResponse{LongUrl: longURL.(string)}, nil
	}

	var longURL string
	err := database.DB.QueryRow(context.Background(), "SELECT long_url FROM url_mappings WHERE short_url = $1", shortURL).Scan(&longURL)
	if err == sql.ErrNoRows {
		return nil, errors.New("short URL not found")
	} else if err != nil {
		return nil, err
	}

	cache.C.Set(shortURL, longURL, 5*time.Minute)
	return &proto.GetLongURLResponse{LongUrl: longURL}, nil
}
