package sort

type InterFace interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntSlice []int

func (p IntSlice) Len() int              { return len(p) }
func (p IntSlice) Less(i, j int) bool    { return p[i] > p[j] }
func (p IntSlice) Swap(i, j int)         { p[i], p[j] = p[j], p[i] }
func (p IntSlice) Get(i int) interface{} { return p[i] }
