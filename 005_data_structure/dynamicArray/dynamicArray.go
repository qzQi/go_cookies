/*
使用slice实现动态数组，练习使用slice
*/

package dynamicarray

import (
	"errors"
)

// DynamicArray struct
type DynamicArray struct {
	Size        int           //default 0
	Capacity    int           //default 0
	ElementData []interface{} //nil
}

var defaultCap = 10

func (d *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= d.Size {
		return errors.New("out of range")
	}
	return nil
}

func (d *DynamicArray) NewCapacity() {
	if d.Capacity == 0 {
		d.Capacity = defaultCap
	} else {
		d.Capacity = d.Capacity << 1
	}
	newArray := make([]interface{}, d.Capacity)

	copy(newArray, d.ElementData)
	d.ElementData = newArray
}

func (d *DynamicArray) Set(index int, elem interface{}) error {
	err := d.CheckRangeFromIndex(index)
	if err != nil {
		return err
	}
	d.ElementData[index] = elem
	return nil
}

func (d *DynamicArray) Insert(index int, elem interface{}) error {
	if d.Size == d.Capacity {
		d.NewCapacity()
	}
	err := d.CheckRangeFromIndex(index)
	if err != nil {
		return err
	}
	temp := append(d.ElementData[:index], elem)
	d.ElementData = append(temp, d.ElementData[index:]...)
	d.Size++
	return nil
}

func (d *DynamicArray) Get(index int) (interface{}, error) {
	err := d.CheckRangeFromIndex(index)
	if err != nil {
		return nil, err
	}
	return d.ElementData[index], nil
}

func (d *DynamicArray) Remove(index int) error {
	err := d.CheckRangeFromIndex(index)
	if err != nil {
		return nil
	}
	d.ElementData = append(d.ElementData[:index], d.ElementData[:index+1]...)
	d.Size--
	return nil
}

func (d *DynamicArray) Data() []interface{} {
	return d.ElementData
}

func (d *DynamicArray) Empty() bool {
	return d.Size == 0
}
