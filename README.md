# httpware
mini http framework in golang

### introduce
httpware is a mini non-invasive http framework written in go.

1. easy and safe. httpware uses same interface as offical http interface, you can always use httpware instead of offical http package or roll back with no any cost
2. strong and high degree of freedom. httpware supports custom router and additional process named as midware.

### usage

```
    server := httpware.NewServer(8083)
    router := http.NewServeMux()
    router.HandleFunc("/service/list", handlefunc)
    server.SetRouter(router)
    server.Run()
```

### todo
1. logger interface 
2. more midware implements
3. a router more easy to use
4. a context implement with higher performance
