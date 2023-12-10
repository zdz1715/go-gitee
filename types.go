package gitee

const (
	CloudEndpoint = "https://gitee.com"
)

type MergeMethod string

const (
	MergeMethodMerge  MergeMethod = "merge" // default
	MergeMethodSquash MergeMethod = "squash"
	MergeMethodRebase MergeMethod = "rebase"
)
