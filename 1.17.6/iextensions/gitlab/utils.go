package gitlab

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	gitlab "github.com/xanzy/go-gitlab"
)

func NewGitLabClientByBasicAuth() (GitLabClient, error) {
	client, err := gitlab.NewBasicAuthClient(
		"root",
		"****",
		gitlab.WithBaseURL("****"),
	)

	if err != nil {
		log.Fatal(err)
	}

	r := GitLabClient{
		Client: client,
	}

	return r, nil
}

func PrintSqlsForSyncGitLabToMysql(
	userId int,
	email string,
	name string,
	username string,
	patId int,
	patName string,
	patValue string,
	keyId int,
	keyName string,
	publicKeyStr string,
	privateKeyStr string,
) {
	cur := time.Now().UnixNano() / 1e6

	sql := fmt.Sprintf(
		"insert into account values(%d, \"%s\", \"%s\", \"%s\", %d, \"%s\", \"%s\", %d, \"%s\", \"%s\", \"%s\", %d, %d);",
		userId,
		email,
		name,
		username,
		patId,
		patName,
		patValue,
		keyId,
		keyName,
		publicKeyStr,
		privateKeyStr,
		cur,
		cur,
	)

	fmt.Println(sql)
}

func Sync() {
	/*
		how to get all user info by xdp api NOT xdp local database operate manually

		sep-r
		分隔符 "\t" & "    "
		// ? arr := strings.Split(line, "    ") + "\t"

		HEX(id) & uuid from mysql -> uuid string
	*/

	cli, _ := NewGitLabClientByBasicAuth()

	var lambda func(id, email string) = func(id, email string) {
		usrs, _, _ := cli.GitLabUsersList(
			&gitlab.ListUsersOptions{
				Search: &email,
			},
		)

		if usrs.TotalItems > 1 {
			log.Println(email)
			return
		}

		// email -> hash & id -> password
		usr, _, _ := cli.GitLabUsersCreate(
			&gitlab.CreateUserOptions{
				Email: &email,
				// Username: &hash,
				// Name:     &hash,
				Password: &id,
			},
		)
		if usr == nil {
			log.Println(email)
			return
		}

		token, _, _ := cli.GitLabTokensCreate(usr.ID)
		if token == nil {
			log.Println(email)
			return
		}

		key, resp, _ := cli.GitLabSshKeysCreate(usr.ID)
		if resp.StatusCode/100 != 2 {
			log.Println(email)
			return
		}

		PrintSqlsForSyncGitLabToMysql(
			usr.ID,
			usr.Email,
			usr.Name,
			usr.Username,
			token.ID,
			token.Name,
			token.Token,
			key.ID,
			key.Title,
			key.Key,
			"// C-TODO",
		)
	}

	fp, _ := os.OpenFile("sync.20220422", os.O_RDWR|os.O_APPEND, 0666)
	defer fp.Close()
	reader := bufio.NewReader(fp)

	count := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		arr := strings.Split(line, "\t")
		email := arr[1]

		uuidBs, _ := uuid.Parse(arr[0])
		id := uuidBs.String()

		lambda(id, email)

		count++
	}

	fmt.Println("count: ", count)
}

func UpdateGitLabProjectsLimitForUsers() {
	/*
		update users max project size limit
	*/

	cli, _ := NewGitLabClientByBasicAuth()

	limit := 1000
	opt := gitlab.ModifyUserOptions{
		ProjectsLimit: &limit,
	}

	for i := 2; i < 150; i++ {
		cli.GitLabUsersUpdate(i, &opt)
	}
}

func UpdateGitLabSshKeyForUsers() {
	/*
		update users ssh key
		BUT only update mysql AND new key for gitlab

		AND sql maybe better WHERE email NOT user_id from gitlab
	*/

	cli, _ := NewGitLabClientByBasicAuth()

	for i := 2; i < 150; i++ {

		key, resp, _ := cli.GitLabSshKeysCreate(i)

		if resp.StatusCode/100 != 2 {
			continue
		}

		sql := fmt.Sprintf(
			"UPDATE account SET public_key_id=%d,public_key_title='%s',public_key_str='%s',private_key_str='%s' WHERE user_id=%d;",
			key.ID,
			key.Title,
			key.Key,
			"// C-TODO",
			i,
		)

		fmt.Println(sql)
	}
}
