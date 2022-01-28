package cmd

import (
	"context"
	"fmt"
	"github.com/redhatxl/kubectl-img/pkg/kube"
	"github.com/redhatxl/kubectl-img/pkg/mtable"
	"github.com/spf13/cobra"
	kv1 "k8s.io/api/apps/v1"
	jobv1 "k8s.io/api/batch/v1"
	cjobv1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "show resource image",
	Long:  `show k8s resource image`,
	RunE:  image,
}

func init() {
	rootCmd.AddCommand(imageCmd)
}

func image(cmd *cobra.Command, args []string) error {

	clientSet := kube.ClientSet(KubernetesConfigFlags)
	ns, _ := rootCmd.Flags().GetString("namespace")
	// 生命一个全局资源列表
	var rList []interface{}

	if flag, _ := cmd.Flags().GetBool("deployments"); flag {
		deployList, err := clientSet.AppsV1().Deployments(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list deployments error: %s", err.Error())
		}
		rList = append(rList, deployList)
	}

	if flag, _ := cmd.Flags().GetBool("daemonsets"); flag {
		deployList, err := clientSet.AppsV1().DaemonSets(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list daemonsets error: %s", err.Error())
		}
		rList = append(rList, deployList)
	}
	if flag, _ := cmd.Flags().GetBool("statefulsets"); flag {
		deployList, err := clientSet.AppsV1().StatefulSets(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list statefulsets error: %s", err.Error())
		}
		rList = append(rList, deployList)
	}
	if flag, _ := cmd.Flags().GetBool("jobs"); flag {
		deployList, err := clientSet.BatchV1().Jobs(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list jobs error: %s", err.Error())
		}
		rList = append(rList, deployList)
	}
	if flag, _ := cmd.Flags().GetBool("cronjobs"); flag {
		deployList, err := clientSet.BatchV1beta1().CronJobs(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Printf("list cronjobs error: %s", err.Error())
		}
		rList = append(rList, deployList)
	}
	deployMapList := make([]map[string]string, 0)
	for i := 0; i < len(rList); i++ {
		switch t := rList[i].(type) {
		case *kv1.DeploymentList:
			for k := 0; k < len(t.Items); k++ {
				for j := 0; j < len(t.Items[k].Spec.Template.Spec.Containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NAMESPACE"] = ns
					deployMap["TYPE"] = "deployment"
					deployMap["RESOURCE_NAME"] = t.Items[k].GetName()
					deployMap["CONTAINER_NAME"] = t.Items[k].Spec.Template.Spec.Containers[j].Name
					deployMap["IMAGE"] = t.Items[k].Spec.Template.Spec.Containers[j].Image
					deployMapList = append(deployMapList, deployMap)
				}
			}
		case *kv1.StatefulSetList:
			for k := 0; k < len(t.Items); k++ {
				for j := 0; j < len(t.Items[k].Spec.Template.Spec.Containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NAMESPACE"] = ns
					deployMap["TYPE"] = "statefulset"
					deployMap["RESOURCE_NAME"] = t.Items[k].GetName()
					deployMap["CONTAINER_NAME"] = t.Items[k].Spec.Template.Spec.Containers[j].Name
					deployMap["IMAGE"] = t.Items[k].Spec.Template.Spec.Containers[j].Image
					deployMapList = append(deployMapList, deployMap)
				}
			}
		case *kv1.DaemonSetList:
			for k := 0; k < len(t.Items); k++ {
				for j := 0; j < len(t.Items[k].Spec.Template.Spec.Containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NAMESPACE"] = ns
					deployMap["TYPE"] = "daemonset"
					deployMap["RESOURCE_NAME"] = t.Items[k].GetName()
					deployMap["CONTAINER_NAME"] = t.Items[k].Spec.Template.Spec.Containers[j].Name
					deployMap["IMAGE"] = t.Items[k].Spec.Template.Spec.Containers[j].Image
					deployMapList = append(deployMapList, deployMap)
				}
			}
		case *jobv1.JobList:
			for k := 0; k < len(t.Items); k++ {
				for j := 0; j < len(t.Items[k].Spec.Template.Spec.Containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NAMESPACE"] = ns
					deployMap["TYPE"] = "job"
					deployMap["RESOURCE_NAME"] = t.Items[k].GetName()
					deployMap["CONTAINER_NAME"] = t.Items[k].Spec.Template.Spec.Containers[j].Name
					deployMap["IMAGE"] = t.Items[k].Spec.Template.Spec.Containers[j].Image
					deployMapList = append(deployMapList, deployMap)
				}
			}
		case *cjobv1.CronJobList:
			for k := 0; k < len(t.Items); k++ {
				for j := 0; j < len(t.Items[k].Spec.JobTemplate.Spec.Template.Spec.Containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NAMESPACE"] = ns
					deployMap["TYPE"] = "cronjob"
					deployMap["RESOURCE_NAME"] = t.Items[k].GetName()
					deployMap["CONTAINER_NAME"] = t.Items[k].Spec.JobTemplate.Spec.Template.Spec.Containers[j].Name
					deployMap["IMAGE"] = t.Items[k].Spec.JobTemplate.Spec.Template.Spec.Containers[j].Image
					deployMapList = append(deployMapList, deployMap)
				}
			}
		}
	}
	json, _ := cmd.Flags().GetBool("json")

	if len(deployMapList) == 0 && json {
		fmt.Printf(`
no resource
Usage:
  kubectl-img image [flags]

Flags:
  -b, --cronjobs       show cronjobs image
  -e, --daemonsets     show daemonsets image
  -d, --deployments    show deployments image
  -h, --help           help for image
  -o, --jobs           show jobs image
  -f, --statefulsets   show statefulsets image

  -j, --json           show json format
`)
	}

	// gen table
	table := mtable.GenTable(deployMapList)
	// json format
	if json {
		jsonStr, _ := table.Json(2)
		fmt.Println(jsonStr)
		return nil
	}
	table.PrintTable()
	return nil
}
