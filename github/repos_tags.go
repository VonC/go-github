package github

import "fmt"

// RepositoryTag represents a tag for a commit, tree or blob in a repository.
type RepositoryTag struct {
	Name       *string    `json:"name,omitempty"`
	CommitTag  *CommitTag `json:"commit,omitempty"`
	ZipballUrl *string    `json:"zipball_url,omitempty"`
	TarballUrl *string    `json:"zipball_url,omitempty"`
}

// Commit represents a GitHub commit.
type CommitTag struct {
	SHA *string `json:"sha,omitempty"`
	Url *string `json:"url,omitempty"`
}

func (r RepositoryTag) String() string {
	return Stringify(r)
}

// ListComments lists all the comments for the repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-tags
func (s *RepositoriesService) ListTags(owner, repo string) ([]RepositoryTag, *Response, error) {

	u := fmt.Sprintf("repos/%v/%v/tags", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tags := new([]RepositoryTag)
	resp, err := s.client.Do(req, tags)
	if err != nil {
		return nil, resp, err
	}
	return *tags, resp, err
}
