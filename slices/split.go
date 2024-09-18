package slices

func Split[T interface{}](s []T, chunk int) [][]T {
	length := len(s)
	if length == 0 || length <= chunk {
		return [][]T{s[:]}
	}

	var data [][]T
	for i := 0; i < length; i += chunk {
		next := i + chunk
		if next > length {
			data = append(data, s[i:])
			break
		}
		data = append(data, s[i:i+chunk])
	}

	return data
}
