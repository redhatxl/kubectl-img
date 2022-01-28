package kube

type Resource struct {
	short        string
	resourceType string
	funcName     string
}

func NewResource(short, resourceType, funcName string) Resource {
	return Resource{
		short:        short,
		resourceType: resourceType,
		funcName:     funcName,
	}
}

//func GenResources() map[string]Resource {
//	r := make(map[string]Resource)
//	r["deploy"] = NewResource("deploy", "deployments", "Deployments")
//	r["sts"] = NewResource("sts", "statefulsets", "")
//	r["ds"] = NewResource("ds", "daemonsets", "")
//	r["jobs"] = NewResource("jobs", "jobs", "job")
//	r["cj"] = NewResource("cj", "cronjobs", "job")
//	return r
//}
