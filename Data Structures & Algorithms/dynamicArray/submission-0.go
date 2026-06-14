type DynamicArray struct {
	arr 	 []int
	capacity int
	size 	 int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		arr: make([]int, capacity),
		capacity: capacity,
		size: 0,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if(da.size == da.capacity){
		da.resize()
	}
	da.arr[da.size] = n
	da.size++
}

 func (da *DynamicArray) Popback() int {
        da.size--
        return da.arr[da.size]
  }

  func (da *DynamicArray) resize() {
        da.capacity *= 2
        newArr := make([]int, da.capacity)
        copy(newArr, da.arr)
        da.arr = newArr
  }

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
