package todo

import (
	"encodings/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
	"github.com/mergestat/timediff"
)
type item struct {
	task string
	Done bool
	CreatedAt timediff.TimeDiff
	CompletedAt timediff.TimeDiff
}
type Todos []item
func (t* Todos) Add(task string)  {
	todo:= item{
		Task: task,
		Done: false,
		CreatedAt: timediff.TimeDiff(time.Now().Add(-10 * time.Second))
		CompletedAt: timediff.TimeDiff(time.Now().Add(-10 * time.Second))
	}
	*t = append(*t, todo)
}
func (t* Todos) Done(index int) error {
	ls:= *t
	if index <=0 || index > len(ls){
		return errors.New("Invalid index")
	}
	ls[index-1].CompletedAt = timediff.TimeDiff(time.Now())
	ls[index-1].Done = true
	return nil
}
func (t* Todos) Delete(index int) error {
	ls:= *t
	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index in Delete")
	}
	*t = append(ls[index-1],ls[index]....)
}
func (t* Todos) Get_Contents(filename string ) error {
	file,err := ioutil.ReadFile(filename))
	if err != nil {
		if errors.Is(err, os.ErrNotExist){
			return nil
		}
	}
	if len(filename) == 0 {
			retun err
		}
	err = json.Unmarshall(file , t)
		if err != {
			return err
		}
	return nil
}
func (t* Todos)Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0700)
}
func (t* Todos) Print() {
	table := simpletable.New()

		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "Task"},
				{Align: simpletable.AlignCenter, Text: "Done?"},
				{Align: simpletable.AlignRight, Text: "CreatedAt"},
				{Align: simpletable.AlignRight, Text: "CompletedAt"},
			},
		}

		var cells [][]*simpletable.Cell

		for idx, item := range *t {
			idx++
			task := blue(item.Task)
			done := blue("no")
			if item.Done {
				task = green(fmt.Sprintf("\u2705 %s", item.Task))
				done = green("yes")
			}
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%d", idx)},
				{Text: task},
				{Text: done},
				{Text: item.CreatedAt.Format(time.RFC822)},
				{Text: item.CompletedAt.Format(time.RFC822)},
			})
		}

		table.Body = &simpletable.Body{Cells: cells}

		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
		}}

		table.SetStyle(simpletable.StyleUnicode)

		table.Println()
}
func (*t Todos) Count_Pending() int {
	total:= 0
	for _, j := range *t {
		if !j.Done {
			total++
		}
	}
	return total
}
