type DynamicArrayElement struct {
	Value int
	IsSet bool
}

type DynamicArray struct {
	Data []DynamicArrayElement
}

func (da *DynamicArray) Capacity() int {
	return len(da.Data)
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		Data: make([]DynamicArrayElement, capacity),
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.Data[i].Value
}

func (da *DynamicArray) Set(i int, n int) {
	if i > da.Capacity() {
		return
	}
	da.Data[i].Value = n
	if !da.Data[i].IsSet {
		da.Data[i].IsSet = true
	}
}

func (da *DynamicArray) Pushback(n int) {
	lastIdx := da.Capacity() - 1
	if da.Data[lastIdx].IsSet {
		da.resize()
		lastIdx++ // potential cost savings shortcut: set to lastIdx = da.Capacity() - 1 if this doesn't work.
	}
	// find the index of the last set value (the "end" of the array)
	lastElement := 0
	for i := lastIdx; i >= 0; i-- {
		if da.Data[i].IsSet {
			lastElement = i + 1
			break
		}
	}
	da.Data[lastElement] = DynamicArrayElement{
		Value: n,
		IsSet: true,
	}
}

func (da *DynamicArray) Popback() int {
	for i := da.Capacity() - 1; i >= 0; i-- {
		if da.Data[i].IsSet {
			popbackValue := da.Data[i].Value
			da.Data[i] = DynamicArrayElement{
				Value: 0,
				IsSet: false,
			}
			return popbackValue
		}
	}
	return 0
}

func (da *DynamicArray) resize() {
	newCapacity := da.Capacity() * 2
	newArray := make([]DynamicArrayElement, newCapacity)
	for i, val := range da.Data {
		newArray[i] = val
	}
	da.Data = newArray
}

func (da *DynamicArray) GetSize() int {
	sizeCounter := 0
	for _, val := range da.Data {
		if val.IsSet {
			sizeCounter++
		}
	}
	return sizeCounter
}

func (da *DynamicArray) GetCapacity() int {
	return da.Capacity()
}
