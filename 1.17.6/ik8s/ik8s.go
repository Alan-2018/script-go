package ik8s

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
)

type (
	jmap   = map[string]interface{}
	jarray = []interface{}
)

type K8sClient struct {
	K8sClientSet *kubernetes.Clientset
}

const (
	KubeConfigFilePath = ""
)

var (
	Kc *K8sClient = new(K8sClient)
)

func init() {
	cfg, err := clientcmd.BuildConfigFromFlags("", KubeConfigFilePath)
	if err != nil {
		panic(err)
	}

	cli, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	Kc.K8sClientSet = cli
}

func (c *K8sClient) GetPodsByLabel(namespace, label string) jarray {
	var (
		pods jarray
	)

	podList, err := c.K8sClientSet.CoreV1().Pods(namespace).List(
		context.TODO(),
		metav1.ListOptions{LabelSelector: label},
	)

	if err != nil {
		panic(err)
	}

	for idx, _ := range podList.Items {
		i := make(jmap)
		i["pod_name"] = podList.Items[idx].Name
		i["status"] = podList.Items[idx].Status.Phase
		i["create_time"] = podList.Items[idx].CreationTimestamp
		i["labels"] = podList.Items[idx].Labels

		pods = append(pods, i)
	}

	return pods
}

func (c *K8sClient) GetLogsByPodName(namespace, podName string) {
	/*
		TailLines 为 nil 则返回现存全部
		Follow 为 true
			HTTP连接关闭，io流是否关闭
			但是，实际上，浏览器 & Postman 都不支持，一直 pending 状态
			curl 命令行 可行

			rancher 也是通过 websocket 实现 流处理
		Container 指定容器名

		etc
			`io.Copy(ctx.Response().Writer, r)` work BUT buffer to be late than for loop

			HTTP Content-type application/octet-stream 是否支持“双向”流
			若不支持，如何处理

			k8s golang client 没有现成 Follow 多个 pods logs
			但是命令行可以指定labels
			如何处理
				同时读多个输入流，写一个输出流？
	*/

	var (
		isFollow bool

		cnt int64 = 10000
	)

	req := c.K8sClientSet.CoreV1().Pods(namespace).GetLogs(
		podName,
		&v1.PodLogOptions{
			TailLines: &cnt,
			Follow:    isFollow,
			// Previous:  true,
			// Container: "",
		},
	)

	r, err := req.Stream(context.TODO())
	if err != nil {
		panic(err)
	}

	if !isFollow {
		defer r.Close()
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, r)
		if err != nil {
			panic(err)
		}

		fmt.Println(
			buf.String(),
		)

	} else {
		reader := bufio.NewReader(r)

		for {
			// bytes, err := reader.ReadSlice('\n')
			bytes, err := reader.ReadBytes('\n')
			if err != nil {
				// if err is io.EOF then ...
				fmt.Printf("%+v", string(bytes))
				panic(err)
			}

			fmt.Printf("%+v", string(bytes))

			time.Sleep(200 * time.Nanosecond)
		}
	}
}
