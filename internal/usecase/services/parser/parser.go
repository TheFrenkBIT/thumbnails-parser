package parser

import (
	"context"
	"fmt"
	"github.com/evrone/go-clean-template/internal/infrastructure/cache"
	"github.com/evrone/go-clean-template/internal/infrastructure/youtube"
	"sync"
)

const imgUrl = "https://img.youtube.com/vi/%s/hqdefault.jpg"

type Parser struct {
	client youtube.Interface
	cache  cache.Interface
	wg     sync.WaitGroup
	mutex  sync.Mutex
}

func New(client youtube.Interface, cache cache.Interface) *Parser {
	return &Parser{
		client: client,
		cache:  cache,
	}
}

func (p *Parser) Parse(ctx context.Context, urls []string) ([][]byte, error) {
	images := make([][]byte, 0)

	p.wg.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			defer p.wg.Done()
			id := p.pullId(u)
			image, err := p.cache.GetValue(ctx, id)
			if err != nil {
				fmt.Println("http")
				image, err = p.client.GetPreview(fmt.Sprintf(imgUrl, id))
				p.cache.SetValue(ctx, id, image)
			}

			if err == nil {
				p.mutex.Lock()
				images = append(images, image.([]byte))
				p.mutex.Unlock()
			}
		}(url)
	}

	p.wg.Wait()

	return images, nil
}

func (p *Parser) pullId(url string) string {
	urlRune := []rune(url)
	result := make([]rune, 0)
	isId := false
	for i := 0; i < len(urlRune); i++ {
		if isId && string(urlRune[i]) == "&" {
			return string(result)
		} else if isId {
			result = append(result, urlRune[i])
		}

		if string(urlRune[i]) == "=" {
			isId = true
		}
	}

	return string(urlRune)
}
