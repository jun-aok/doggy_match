package calc

// https://selfnote.work/20210211/programming/intersection-union-with-golang/

// 積集合を取得する
func Intersect(l1, l2 []int) []int {
	s := make(map[int]struct{}, len(l1))

	// list1をmap形式に変換
	for _, data := range l1 {
		// struct{}{}何もない空のデータ
		s[data] = struct{}{}
	}

	r := make([]int, 0, len(l1))

	for _, data := range l2 {
		// mapにデータがない場合は、スキップ
		// okにはデータの存在有無true/falseで入る
		if _, ok := s[data]; !ok {
			continue
		}

		// 積集合のデータを格納
		r = append(r, data)
	}
	return r
}
