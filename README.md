This is sample code for article : https://novrian.substack.com/p/understanding-mutex-in-go-what-why
In concurrent programming for example when you implement goroutine or multihtreading, managing access to shared resources is critical to avoid data races and inconsistent behavior. Go makes it easy to run concurrent code using goroutines, but without proper synchronization, it’s also easy to introduce bugs.

We will explore and understand what mutex is, why we need and when to use it, and how to use it. One tool Go provides for synchronizing access to shared data is using mutex. We will also see two versions of the same program, one using a mutex and one without,to clearly illustrate the difference.
