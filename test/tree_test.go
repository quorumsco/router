package iogo

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Trie", func() {
	It("should insert and retrieve values", func() {
		node := newNode(0, staticType)
		tests := [][3]string{
			{"/", "0", "/"},
			{"/me", "1", "/me"},
			{"/admin/test", "2", "/admin/test"},
			{"/admin/tester", "3", "/admin/tester"},
			{"/admin/:name", "4", "/admin/name"},
			{"/admin/:name/:id", "5", "/admin/name/id"},
		}

		f := func(w http.ResponseWriter, r *http.Request) {

		}
		for _, e := range tests {
			node.insert(e[0], f)
			//value, found, _ := node.find(e[2])
			//Expect(value).To(Equal(e[1]))
			//Expect(found).To(BeTrue())
		}
	})

	It("should not find a value not inserted", func() {
		//node := newNode(0, staticType)
		//node.insert("will u find me?", 5)
		//value, found, _ := node.find("will u find me")

		//Expect(value).To(BeNil())
		//Expect(found).To(BeFalse())
	})
})
