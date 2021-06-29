package task14

type StringSet map[string]bool

func newStringSet(initStrings ...string) *StringSet {
	set := make(StringSet, 0)
	set.AddSlice(initStrings...)
	return &set
}

// Add new elem to set. bool indicates existance of elem before inserting
func (ss *StringSet) Add(new string) bool {
	_, ok := (*ss)[new]
	if !ok {
		(*ss)[new] = true
	}
	return ok
}

// AddSlice of strings
func (ss *StringSet) AddSlice(news ...string) {
	for i := 0; i < len(news); i++ {
		ss.Add(news[i])
	}
}

// ConvertToSlice is
func (ss StringSet) ConvertToSlice() []string {
	result := make([]string, len(ss))
	for k := range ss {
		result = append(result, k)
	}
	return result
}
