package queue

import (
	"fmt"
	"testing"
	"github.com/franela/goblin"
)

func TestQueue(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Queue", func() {
		var array = []int{9, 5, 3, 2, 8, 1, 2, 3}
		var obj = NewQueue()
		for _, v := range array {
			obj.Append(v)
		}
		obj.AppendLeft(0)
		//print
		fmt.Println(obj)

		g.It("should test items in the queue", func() {
			defer
				g.Assert(obj.IsEmpty()).IsFalse()
			g.Assert(obj.Size()).Equal(len(array) + 1)
			g.Assert(obj.First().(int)).Equal(0)
			g.Assert(obj.PopLeft().(int)).Equal(0)

			g.Assert(obj.First().(int)).Equal(9)
			g.Assert(obj.PopLeft().(int)).Equal(9)
			g.Assert(obj.PopLeft().(int)).Equal(5)
			g.Assert(obj.PopLeft().(int)).Equal(3)
			g.Assert(obj.PopLeft().(int)).Equal(2)
			g.Assert(obj.Last().(int)).Equal(3)
			fmt.Println(obj)

			obj.Empty()
			g.Assert(obj.PopLeft()).Equal(nil)
			g.Assert(obj.IsEmpty()).IsTrue()
			//print
			fmt.Println(obj)

		})
	})
}
