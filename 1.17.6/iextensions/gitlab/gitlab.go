package gitlab

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/flower/script-go/iutils"
	gitlab "github.com/xanzy/go-gitlab"
)

/*
	go-gitlab "github.com/xanzy/go-gitlab"
	? how to retry
*/

type GitLabClient struct {
	Client *gitlab.Client
}

/*
	users
*/

func (c GitLabClient) GitLabUsersGet(userId int) (usr *gitlab.User, resp *gitlab.Response, err error) {
	/*
		user_id start from 1 not 0
	*/

	usr, resp, err = c.Client.Users.GetUser(
		userId,
		gitlab.GetUsersOptions{},
	)

	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabUsersGet err: %w", err))
		return nil, resp, err
	}

	iutils.Log(
		usr,
		resp,
	)

	return usr, resp, err
}

type GitLabUsersListResult struct {
	Users []*gitlab.User

	TotalItems   int
	TotalPages   int
	ItemsPerPage int
	CurrentPage  int
	NextPage     int
	PreviousPage int
}

func (c GitLabClient) GitLabUsersList(opt *gitlab.ListUsersOptions) (users *GitLabUsersListResult, resp *gitlab.Response, err error) {
	/*
		Active: true
		PerPage: 0 & Page: 0 -> return nil
		Search:
				/users?search=John
				/users?search=fake@fake.com

		opt := gitlab.ListUsersOptions{
			Username: username,
			Search:   search,
			ListOptions: gitlab.ListOptions{
				PerPage: perpage,
				Page:    page,
			},
		}
	*/

	usrs, resp, err := c.Client.Users.ListUsers(opt)

	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabUsersList err: %w", err))
		return nil, resp, err
	}

	iutils.Log(
		users,
		resp,
	)

	r := GitLabUsersListResult{
		usrs,
		resp.TotalItems,
		resp.TotalPages,
		resp.ItemsPerPage,
		resp.CurrentPage,
		resp.NextPage,
		resp.PreviousPage,
	}

	return &r, resp, err
}

func (c GitLabClient) GitLabUsersCreate(opt *gitlab.CreateUserOptions) (usr *gitlab.User, resp *gitlab.Response, err error) {
	/*
		skipConfirmation := true
		admin := false
		projectsLimit := 1000
		forceRandomPassword := true

		opt := gitlab.CreateUserOptions{
			Email:               &email,
			Name:                &name,
			Username:            &username,
			Password:            &password,
			SkipConfirmation:    &skipConfirmation,
			Admin:               &admin,
			ProjectsLimit:       &projectsLimit,
			ForceRandomPassword: &forceRandomPassword,
		}

		POST http://xxxx/api/v4/users: 409 {message: Email has already been taken}
	*/

	usr, resp, err = c.Client.Users.CreateUser(opt)
	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabUsersCreate err: %w", err))

		if resp == nil || resp.Body == nil {
			// pass
		} else {
			// http: read on closed response body
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			log.Printf("E GitLabUsersCreate body: %+v\n", string(body))
		}

		return nil, resp, err
	}

	iutils.Log(
		usr,
		usr.CanCreateGroup,
		usr.CanCreateProject,
		usr.State,
		usr.Organization,
		usr.Identities,

		resp,
	)

	return
}

func (c GitLabClient) GitLabUsersUpdate(userId int, opt *gitlab.ModifyUserOptions) (usr *gitlab.User, resp *gitlab.Response, err error) {
	/*
		PUT http://xxxx/api/v4/users/151: 404 {message: 404 User Not Found}
	*/

	usr, resp, err = c.Client.Users.ModifyUser(userId, opt)
	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabUsersUpdate err: %w", err))
		return nil, resp, err
	}

	iutils.Log(
		usr,
		resp,
	)

	return
}

func (c GitLabClient) GitLabUsersDelete(userId int) (resp *gitlab.Response, err error) {
	resp, err = c.Client.Users.DeleteUser(userId)
	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabUsersDelete err: %w", err))
		return nil, err
	}

	iutils.Log(
		resp,
	)

	return
}

func (c GitLabClient) GitLabUsersUpdateForBatch() {
	/*
		确定 输入、输出，以方便使用
	*/
}

/*
	tokens
*/

func (c GitLabClient) GitLabTokensCreate(userId int) (token *gitlab.PersonalAccessToken, resp *gitlab.Response, err error) {
	name := strconv.Itoa(userId)
	scopes := []string{"write_repository"}

	opt := gitlab.CreatePersonalAccessTokenOptions{
		Name:   &name,
		Scopes: &scopes,
	}

	token, resp, err = c.Client.Users.CreatePersonalAccessToken(userId, &opt)
	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabTokensCreate err: %w", err))
		return nil, resp, err
	}

	/*
		gitlab.PersonalAccessToken{
			ID:15,
			Name:"14",
			Revoked:false,
			CreatedAt:time.Time{wall:, ext:},
			Scopes:["write_repository"],
			UserID:14,
			Active:true,
			Token:"-7GBzVpqDsDcqCv4Z8NR"
		}
	*/
	return
}

/*
	ssh
*/

func (c GitLabClient) GitLabSshKeysCreate(userId int) (key *gitlab.SSHKey, resp *gitlab.Response, err error) {
	/*
		title: GitLab UID | workspace id | workspace name | email | ...

		key := ""
		400 {message: {fingerprint: [cannot be generated]}, {key: [can't be blank, is invalid, type is forbidden. Must be RSA, DSA, ECDSA, or ED25519]}}

		生成公钥 没有 comment i.e. email 等等
		尝试，但是失败
		key = fmt.Sprintf("%s {}", strings.TrimRight(publicKeyStr, "\n"))
	*/

	var (
		cnt int = 3
	)

	title := strconv.Itoa(userId)

	for i := 0; i < cnt; i++ {
		publicKeyStr, _, _ := iutils.GenerateKeysByRsa(2048)

		key, resp, err = c.Client.Users.AddSSHKeyForUser(
			userId,
			&gitlab.AddSSHKeyOptions{
				Title: &title,
				Key:   &publicKeyStr,
			},
		)

		if err != nil && strings.Contains(err.Error(), "400 {message: {fingerprint: [has already been taken]}}") {
			continue
		} else if err != nil {
			err = iutils.LogError(fmt.Errorf("E GitLabSshKeysCreate err: %w", err))
			break
		}

		return
	}

	iutils.Log(
		key,

		key.ID,
		key.Title,
		key.Key,
		key.CreatedAt,
		key.ExpiresAt,
	)

	return
}

/*
	project
*/

func (c GitLabClient) GitLabProjectsCreate(userId int, name, description string) (project *gitlab.Project, resp *gitlab.Response, err error) {
	/*
		// C-TODO project option template
		// C-TODO private-visibility-CAUSE-OF-default-user-default-group-AND-internal-visibility-eq-public-visibility

		POST http://xxxx/api/v4/projects/user/150: 400 {message: {name: [has already been taken]}, {path: [has already been taken]}}
	*/

	opt := gitlab.CreateProjectForUserOptions{
		Name:                     gitlab.String(name),
		Description:              gitlab.String(description),
		Visibility:               gitlab.Visibility(gitlab.PrivateVisibility),
		MergeRequestsAccessLevel: gitlab.AccessControl(gitlab.EnabledAccessControl),
		SnippetsAccessLevel:      gitlab.AccessControl(gitlab.EnabledAccessControl),
	}
	project, resp, err = c.Client.Projects.CreateProjectForUser(userId, &opt)
	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabProjectsCreate err: %w", err))
		return nil, resp, err
	}

	iutils.Log(
		project,
		resp,
	)

	return project, resp, nil
}

type GitLabProjectsListResult struct {
	Projects []*gitlab.Project

	TotalItems   int
	TotalPages   int
	ItemsPerPage int
	CurrentPage  int
	NextPage     int
	PreviousPage int
}

func (c GitLabClient) GitLabProjectsList(userId int, opt *gitlab.ListProjectsOptions) (projects *GitLabProjectsListResult, resp *gitlab.Response, err error) {
	/*
		sort-r		default sort desc by create_time of project
						&& sort-search-args-r maybe more good for query-r OR search-r
						&& maybe query-r is better
					etc:
						owned
						search
						simple
						visibility

		number-r 	default page 1 & per_page 20 as same as usersList
						per_page only work for `lte 100`


		opt = &gitlab.ListProjectsOptions{
			Simple: gitlab.Bool(simple),

			OrderBy:          gitlab.String(orderBy),
			Sort:             gitlab.String(sort),
			Search:           gitlab.String(search),
			SearchNamespaces: gitlab.Bool(true),

			ListOptions: gitlab.ListOptions{
				PerPage: perPage,
				Page:    page,
			},
		}
	*/

	r, resp, err := c.Client.Projects.ListUserProjects(userId, opt)
	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabProjectsList err: %w", err))
		return nil, resp, err
	}

	iutils.Log(
		r,
		resp,
	)

	projects = &GitLabProjectsListResult{
		r,
		resp.TotalItems,
		resp.TotalPages,
		resp.ItemsPerPage,
		resp.CurrentPage,
		resp.NextPage,
		resp.PreviousPage,
	}

	return projects, resp, nil
}

func (c GitLabClient) GitLabProjectsDelete(identifier interface{}) (resp *gitlab.Response, err error) {
	/*
		identifier is project_id OR project_path_with_namespace

		DELETE http://xxxx/api/v4/projects/118: 404 {message: 404 Project Not Found}
	*/

	resp, err = c.Client.Projects.DeleteProject(identifier)
	if err != nil {
		err = iutils.LogError(fmt.Errorf("E GitLabProjectsDelete err: %w", err))
		return resp, err
	}

	return resp, nil
}
