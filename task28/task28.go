package task28

import "fmt"

type apiMIT struct {
}

func (api *apiMIT) getStudentsIDsByRequset(token string, userID int) []int {
	return []int{1, 20, 100}
}

type apiMGU struct {
}

func (api *apiMGU) getStudentsIDs(url string) []int {
	return []int{1, 11, 111}
}

type univerApier interface {
	getStudentsIDs() []int
}

type adapterAPIMIT struct {
	api    *apiMIT
	token  string
	userID int
}

func (adapter *adapterAPIMIT) getStudentsIDs() []int {
	return adapter.api.getStudentsIDsByRequset(adapter.token, adapter.userID)
}

type adapterAPIMGU struct {
	api *apiMGU
	url string
}

func (adapter *adapterAPIMGU) getStudentsIDs() []int {
	return adapter.api.getStudentsIDs(adapter.url)
}

// Task28 demo
func Task28() {
	apiMGU := apiMGU{}
	apiMIT := apiMIT{}

	apiers := []univerApier{
		&adapterAPIMGU{&apiMGU, "ftp://source.get"},
		&adapterAPIMIT{&apiMIT, "nf3fh0f0ue1-ue1", 123},
	}

	ids := make([]int, 0)
	for _, api := range apiers {
		ids = append(ids, api.getStudentsIDs()...)
	}
	fmt.Println(ids)
}
