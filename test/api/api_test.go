package main

import (
	"github.com/reality95/cf-predictor/src/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComments(t *testing.T) {
	comments, err := api.GetComments(69)
	if err != nil {
		t.Errorf("Expected no error while extracting comments for blogId 69, got %s\n", err.Error())
	}
	if len(comments) > 0 {
		t.Errorf("Blog with blogId 69 has no comments, received comments\n")
	}

	comments, err = api.GetComments(666)
	if err != nil {
		t.Errorf("Expected no error while extracting comments for blogId 666, got %s\n", err.Error())
	}
	if len(comments) != 5 {
		t.Errorf("Blog with blogId 666 has 5 comments, received %d comments\n", len(comments))
	}

	comments, err = api.GetComments(666666)
	if err == nil {
		t.Errorf("Expected an error while extracting comments for blogId 666666, received none\n")
	}
	if err.Error() != "blogEntryId: Blog entry with id 666666 not found" {
		t.Errorf("Expected a different error while extracting comments for blogId 666666\n")
	}
}

func TestGetBlog(t *testing.T) {
	blog, err := api.GetBlog(69)
	if err != nil {
		t.Errorf("Expected no error while extracting BlogEntry for blogId 69, got %s\n", err.Error())
	}
	assert := assert.New(t)
	assert.Equal(blog.OriginalLocale, "ru")
	assert.Equal(blog.AllowViewHistory, false)
	assert.Equal(blog.CreationTimeSeconds, 1265647999)
	assert.Equal(blog.Rating, 1)
	assert.Equal(blog.AuthorHandle, "MikeMirzayanov")
	assert.Equal(blog.ModificationTimeSeconds, 1265653678)
	assert.Equal(blog.Id, 69)
	assert.Equal(blog.Title, "Back home!")
	assert.Equal(blog.Locale, "en")
	assert.Equal(blog.Tags, []string{"2010", "acm", "acm-icpc", "back home", "saratov"})
}
