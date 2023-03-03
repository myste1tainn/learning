package main

type MyClass struct {
	cache map[string]string
}

func (this MyClass) greet(name string) string {
	if this.cache[name] != "" {
		return this.cache[name]
	} else {
		this.cache[name] = "Hello, " + name
		return this.cache[name]
	}
}

func greet(name string) string {
	return "Hello, " + name
}

var cache map[string]string

func cacheable[P1 any, R any](key string, ttl int, x func(parameter1 P1) R) func(name P1) R {
	return func(parameter1 P1) R1 {
		if cache[parameter1.(string)] != "" {
			return cache[parameter1.(string)]
		} else {
			cache[parameter1] = x(parameter1)
			return cache[parameter1.(string)]
		}
	}
}

type MyClass2 struct {
	greet    func(x string) string
	farewell func(x Person) string
}

type Person struct {
	first string
	last  string
}

func asdf() {
	var c = MyClass2{}
	c.greet = cacheable(func(parameter1 string) string {
		return "Hello, " + parameter1
	})
	c.farewell = cacheable(func(parameter1 Person) string {
		return "Farewell, " + parameter1.first
	})
	c.greet("Ice")
}
