package gh

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ghClient = NewGithubClient()

func TestGetAccessToken(t *testing.T) {
	result := ghClient.GetAccessToken("access_token.txt")
	assert.NotNilf(t, result, "read access token successfully!")
}

func TestGetRepoStats(t *testing.T) {
	stats, response, err := ghClient.GetRepoStats("qdriven", "go-for-qa")
	if err != nil {
		return
	}
	fmt.Println(stats, response)
}

func TestGetStarredRepos(t *testing.T) {
	starred := ghClient.GetAllStarredRepos(1, 2)
	print(starred) //filter by topics and description
}

func TestGetFollowingUsers(t *testing.T) {
	users := ghClient.GetAllFollowing(0, 50)
	print(users)
}
