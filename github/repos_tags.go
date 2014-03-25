package github

import "fmt"

// RepositoryTag represents a short description of tag
// for a commit, tree or blob in a repository.
type RepositoryTagShort struct {
	Name       *string    `json:"name,omitempty"`
	CommitTag  *CommitTag `json:"commit,omitempty"`
	ZipballUrl *string    `json:"zipball_url,omitempty"`
	TarballUrl *string    `json:"zipball_url,omitempty"`
}

type RepositoryTag struct {
	Tag       *string    `json:"tag,omitempty"`
	SHA       *string    `json:"sha,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Message   *string    `json:"message,omitempty"`
	Tagger    *Tagger    `json:"tagger,omitempty"`
	ObjectTag *ObjectTag `json:"object,omitempty"`
}
type Tagger struct {
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`
	Date  *Timestamp `json:"date,omitempty"`
}
type ObjectTag struct {
	Name *string `json:"type,omitempty"`
	SHA  *string `json:"sha,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// Commit represents a GitHub commit.
type CommitTag struct {
	SHA *string `json:"sha,omitempty"`
	URL *string `json:"url,omitempty"`
}

type DataTag struct {
	Tag     *string `json:"tag,omitempty"`
	Message *string `json:"message,omitempty"`
	Object  *string `json:"object,omitempty"`
	Type    *string `json:"type,omitempty"`
	Tagger  *Tagger `json:"tagger,omitempty"`
}

func (r RepositoryTag) String() string {
	return Stringify(r)
}

// ListTags lists all tags for the repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-tags
func (s *RepositoriesService) ListTags(owner, repo string) ([]RepositoryTagShort, *Response, error) {

	u := fmt.Sprintf("repos/%v/%v/tags", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tags := new([]RepositoryTagShort)
	resp, err := s.client.Do(req, tags)
	if err != nil {
		return nil, resp, err
	}
	return *tags, resp, err
}

// GetTag get a tag for the repository.
//
// GitHub API docs: http://developer.github.com/v3/git/tags/#get-a-tag
// Seems deprecated
func (s *RepositoriesService) GetTag(owner, repo string, SHA string) (*RepositoryTag, *Response, error) {

	u := fmt.Sprintf("repos/%v/%v/git/tags/%v", owner, repo, SHA)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tag := new(RepositoryTag)
	resp, err := s.client.Do(req, tag)
	if err != nil {
		return nil, resp, err
	}
	return tag, resp, err
}

// CreateHook creates a Hook for the specified repository.
// Name and Config are required fields.
//
// GitHub API docs: http://developer.github.com/v3/git/tags/#create-a-tag-object

func (s *RepositoriesService) CreateTag(owner, repo string, tag *DataTag) (*RepositoryTag, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/git/tags", owner, repo)
	req, err := s.client.NewRequest("POST", u, tag)
	if err != nil {
		return nil, nil, err
	}

	rt := new(RepositoryTag)
	resp, err := s.client.Do(req, rt)
	if err != nil {
		return nil, resp, err
	}

	return rt, resp, err
}
