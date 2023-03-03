package com.example.learn;

class MyClass {
    @Cacheable(key = "#name", unless = "null")
    String greet(String name) {
        return "Hello, " + name;
    }
}
