package main

import "fmt"

func slicePointer(a []interface{}) (*[]interface{}) {
	return &a
}

func main () {
	// a byte is just an int
	b := []byte { 'a', 'b', 'c', 'd', 'e' }
	fmt.Println(b);

	s := []string {"a", "b", "c", "d", "e"}
	fmt.Printf("s is %q\n",s)

	// the slice operator : defaults to
	// 0:len(slice), which basically is the same
	// as the original slice
	s2 := s[:]
	fmt.Printf("s2=s[:], is %q\n", s2)

    s2[1] = "z"
    fmt.Printf("after changing s2[1]=z, s is %q\n", s)
    fmt.Printf("and s2 is %q\n", s2)

    // slicing s as s3 will not allocate a new array,
    // s3 points to the same array as s, but only
    // a short segment of s
    s3 := s[1:2]
    fmt.Printf("s3 = s[1:2]: %q %p\n" ,s3, s3)
    fmt.Printf("s3 len(%d) s3 cap(%d)\n", len(s3), cap(s3))
    s3[0] = "zz"
	fmt.Println("s3:" ,s3)
	fmt.Printf("s3 len(%d) s3 cap(%d)\n", len(s3), cap(s3))

    s4 := s3[0:cap(s3)]
    fmt.Printf("s4: len(%d) cap(%d) %q\n", len(s4), cap(s4), s4)

    m := []string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
    fmt.Printf("m is: %q\n", m)

    m1 := m[5:]
    fmt.Printf("m1 cap(%d), len(%d), m1 is: %q\n", cap(m1), len(m1), m1)

    m2 := m[:2]
    fmt.Printf("m2 cap(%d), len(%d), m2 is: %q\n", cap(m2), len(m2), m2)

    // this should give a runtime error as the length of m2 is 2
    m2[1] = "8"
    fmt.Printf("m2 is: %q\n", m2)

    f := make([]string, 10)
    copy(f, m)
	fmt.Printf("m is: %q\n", m)
	fmt.Printf("f is: %q\n", f)

	f = make([]string, 10)
	copy(f, m2)
	fmt.Printf("f is: %q\n", f)
	fmt.Printf("m2 is: %q\n", m2)

    //a := slicePointer(f)
    //fmt.Println("a:", a)
}
