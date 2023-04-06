package gitlab

import (
	"fmt"
	"log"

	"github.com/flower/script-go/iutils"
	gitlab "github.com/xanzy/go-gitlab"
)

/* **** gitlab tests **** */

/*
	go test -> *_test.go
*/

func TestGitLabTokens() {
	/*
		! some token unauth such as "KHFsQmyUNED44GbM383C"
		it seems that only admin token work
	*/

	// NewClient NOT NewBasicAuthClient
	client, _ := gitlab.NewClient(
		"C9QLYFX7CjB8pfv6mczy",
		gitlab.WithBaseURL("****"),
	)
	cli := GitLabClient{client}

	// 1 is admin
	project, _, _ := cli.GitLabProjectsCreate(1, "test-00000", "test-00000")

	log.Printf("%+v\n", project)
}

func TestGitLabProjects() {
	cli, _ := NewGitLabClientByBasicAuth()

	project, resp, _ := cli.GitLabProjectsCreate(1, "test-X0000", "test-X0000")
	log.Printf("%+v\n", project)

	/*
		结构体 嵌套
		resp.StatusCode -> resp.Response.StatusCode
	*/
	log.Println(resp.StatusCode, resp.Response.StatusCode, resp.StatusCode == resp.Response.StatusCode)

	orderBy := "updated_at"
	sort := "desc"
	search := "test-X0000"
	r, resp, _ := cli.GitLabProjectsList(
		1,
		&gitlab.ListProjectsOptions{
			OrderBy: &orderBy,
			Sort:    &sort,
			Search:  &search,
		},
	)

	for idx, _ := range r.Projects {

		iutils.Log(
			(*r.Projects[idx]).ID,
			(*r.Projects[idx]).Name,
			(*r.Projects[idx]).Description,
			(*r.Projects[idx]).Path,
			(*r.Projects[idx]).PathWithNamespace,
			(*r.Projects[idx]).CreatedAt.UnixNano()/1e6,
			(*r.Projects[idx]).LastActivityAt.UnixNano()/1e6,
		)

	}

	fmt.Println()
}
