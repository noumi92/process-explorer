package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.4" )

func ROXbUMXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itSMwHiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0STLeM9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rX4rTBcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTR2OfTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IajbAk0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glKQy2fdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNA0ux9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNjNjabEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYzavOygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6ML1rOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFRyqqhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGj1aIbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTSik4rAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dd4t7tJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkdCj0TVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cX25DeJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDljWubmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxYZwZxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhFOj8lfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLxdoWYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6eOvHISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsFVnh55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fQqmttRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLGwCtSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATUD1jSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6MIT3mLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDmCsZf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2Us14CRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1BcINekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dORyFYBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yr4IYpWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxcVpVGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtBcEPYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TgiU0rJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJkOwQo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cG6s7yG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZiiajTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJJiKDgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1naTvkSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZWbPFekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtkLCq3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUclu4IdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETCqT3JtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkHLjNWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rAiPV2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqhWCYcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlCzMrOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkKTAfU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2U8Fl2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opkitqThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0iTpjQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbJHVlqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HY9EkRX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAYz5Wu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9TOnj2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5ogAHBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fd788GmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvC6z8g9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lsqa1xtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5CK4JlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cyyim7q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRRVIgcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOj6TsKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9b3APGB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3paHHEUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MxfYfDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wK1MZetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFDCIOGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4F8tqrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19DeEHueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kz5kjvheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzpxzlNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xt6qWLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkmpqHXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSAp4jd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wodnlLE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBKYaZfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bf1I7eefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzsW4L91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gq2Y2e0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Oe43SAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QxDQcXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRQ0bmXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhcxR7xJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQ1pjcxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmTfjFkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9caYguvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybyknpExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCU1CXHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTzCFFDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGJJRz9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyRjrwA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2RxrRUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TRS3EtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5bLsKT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxXxd0bgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJjosI9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWxpYdA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doNY1dvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKR210jbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func av9OZbllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUtsG64qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLShiPJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixXgzZ7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iW1MdYhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXArk7dKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lipGmTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEywJB7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtsiVIgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kKifFbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGLkOD8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxfcTfJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZAXBWubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGnmMxXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyfCh3CUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIirKvzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6A9p6KyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5mTcuUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjoWkervWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIGU04PyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7N1wPcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLxWyDplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcZTCQZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoJoovmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85PnrUEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IL7aoTVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hJD5a7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5Ji483oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BagJFdDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJW6OIduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzrTWGz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTV2tHVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhmBa4xgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoMe5J9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7lnYOzpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMbEwDiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilvCrbCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6ZH4UARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZCrqQZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqP2N60UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hyqlyN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAD4Cfz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZYfOTP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqMFWXh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJR6n19qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EL1Y3A1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCkgT7YhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uC0jymoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxDZKDUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TS8kTOZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePKNfJN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySs3nZoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isIVTarDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRCsKJd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7f7omuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2KD4YgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3NyS1yDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqToC4b3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asFzVrEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zAaW6vuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3kyk1vZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZuaR3jsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 297c2sxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVvZ0jA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPiSW0uOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHUDxekpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9ddYSQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRtBmXhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jk22BTrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8K2SZO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZL4HSdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jY7CrZKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGutebfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 352ixpQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBJZPeX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sm7ybOUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFq3ZYvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TL4ySIOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5qCwUZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUlXvU3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QOo7uGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlMmfabWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQPssDEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeDyTl6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhyQLIOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v10BcjUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLbWqxaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3Gx6sR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trkXY5vgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jfHTCCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlXIZHI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFOoK8juWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31s7cxohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfYOgZG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2BMdMZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jsXBijF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zsRbcAReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6icFKH0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmrlulDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZOd6QynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWQyfDIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDPBqUoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kru1aOyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUNDJAvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3jLuYe0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDIvoKzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvDcEtlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dY77LCMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTqwatjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xQqrZw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOP8XyvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzrTYDiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyEPGPdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ag02RFTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiMEJsrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngpngQm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCl20H3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJx7yuKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIxcNfJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjSyfO3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEFFHYipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7hW36vMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOPgsN4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sc6oqsWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4PMRgDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7t5mN4KvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsMeSnDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmfLsgVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfODRIdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WX6QnjUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msADVNmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICWCMd6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sD5jQVmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AnOOglBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1irHsdyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SJg5SewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FZs7LZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHCRnhwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsDZtgY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMaQPLwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVB60WTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thqueUgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvMu3vWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ns7XHGuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XUCpArrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpaQDoyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5hH3pdOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTkv44qgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPI7IJuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MclaZYELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlliQ5muWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erECYkgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OzDTFq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcvqzZjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTxEYM08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHp7sxc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8vycot9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRD4t2d8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sTxOzGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REZb3qwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hs1GovBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLA6wHMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOKuYXm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWDIM2igWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwDiH1DJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olxh0uj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YX3nbgVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1fUmIHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UxNwEUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjDEWJXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRZdDPDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zgLVsn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWSvbyfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rlPs5j4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmWguf2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fxdHW3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9EzkjmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnD4k4QWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBF4A9jYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3gA0fmb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CHpzhKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Utb7wiGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04m8ah5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjiJQ67XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvDss1XTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eukrgr5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cI565yurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exdV0yQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bz3L1fIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrIWRY8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbPb2gW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlFJ3RIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgOVcf7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRobAsaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0hTv8VlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEgnBkxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMjkT4pOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmU6MU0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwB997eSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwLojSxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THSZO27pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XRLo3CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8jcvzJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60KdNozRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnigaMk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEKVufdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDl4eRQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDp0gAIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xopR4nwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZUg0vUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8SyE95CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LM59q9vFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LbRR79TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjSaDALxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNl79RSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sr14HPBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFvPz3LxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sXZVvEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVSdGh7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWnwsXnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LznFYGP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFlIvyhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A09KdGhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s46pzlZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFiGyYpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0j5PwjMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLA4rn1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZO9RqhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHxStW7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FF71QEAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sGZGIMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2Hrtp3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGryPEE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSxNGJ5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGX5189aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vlQH6uyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKehz5T5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAAHCRuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuhzbVtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAojA1evWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0I8DFq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TLbeT6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12fwyFAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nzdN2isWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3XC0FXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENS25RZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqMavMDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTgMwPhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xb3NicY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usMVoh3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnaSPtxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLwA7s3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmHzHKT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IWGXB77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func piVwaTLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAgttOaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1lPwLvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlokYxMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukrv6ucZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afR8TsZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAUJzOb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IH7XW5bJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5deLBzmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJQploXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOXWnr3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCjR1UaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEbEjfP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkO6HzwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGepGefZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rgfyv2SUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzNr5NiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHfuG3adWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inWV60BMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0k2mBgOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3ZKEMfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuwzXs3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4Uf4FXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjB6zRCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcmag0ESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bu5BeYc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lC9jUkSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ld26NgjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBIiWtJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u00ohfoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HPQNgR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQFvQuSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGS4K7pRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bjvuk71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sn8rUwKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXE34dozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DfoxhfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBdu8m1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G59J4ArvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAv4eu8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJ3ZGwlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ErSjtW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEFew2W8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prUdi4zsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwagJ1y4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJ7nQE5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VgbN9S6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaUYu0aaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKKtwcHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCHXnurKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6GZOXkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKvE6zITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzMBO85KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKj34CQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nr3yPPiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOZo12n6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfkWPY1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTnKO5EwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmW3ouR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upK3o5TCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4sdfDpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2pYYL16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6y1Y4gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cgce42rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDg6s3UvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7inP049Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EX84OzkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HhlfTmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxd5Tq8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vb36lvcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMgo30fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbpdEZqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBO4HN5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQvxXxnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcPNjroKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DGHw5BRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2wkJgvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZ7cnkcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BZS19C1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXmcSS0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEdP4mcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DsAdwEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWr1q7IxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAUXNMTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8EkuNN5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlgGe0DFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnlBrdfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBdLnRMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyP5JZLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULx80nyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPNMRvdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NO4Nsf5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTZgulwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbFi5SqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puuvbZWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Waj3tfETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rv4ZJPJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJMM274tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLKnBpaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4Ftj5BSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCicfjTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwH0LHOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qY2njvxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcaoNt3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYK4vIuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9QOi90GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAiqg7t7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JTLxjolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsKUs0bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LX22yhs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0khsrQ1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iB8fr0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIjqJPK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPkP4jHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9sM3kfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZ8GCOjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vou6kOIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kjdmwas7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHHy5MLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UD1O6WVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fL2hUBioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBdOir2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrxe4OGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77PwrL7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izI7elN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQU1ozxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGO965RFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNyCx2r4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahyw34fQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3RKiLZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Q54M5NBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EOo197UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrv66LXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYtBGpDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jwT8ERpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owfApyCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ztae0U5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dw7qXb23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5ximjBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjlDAVxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYUWVAODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTgo0nPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXs1GhqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzOfGBqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zyz603iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWAVhwWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVv9wOT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwDUR0g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQza2CMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAXzkrrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syX5mqp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNlcHC0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEtHJ1deWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHm4DcfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2U3ZFjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gp8dx2WEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xyHSlgL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HixOedBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTkP2MirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncyrAdYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IeB6IBIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIWOyWvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LndPCpA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39vzNyWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nsLDU6yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98eR5rwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbbUNnamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4nCWcFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZCBtlnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRhEWFFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6kmlf5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27QVyyuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAGJ4ENtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LU2Bxnc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oF0IHG3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4XiHoqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqbgfetVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHPC8bjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6s6doGSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSoH1K8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06OUsfDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHZrgBXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuv8j9oqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASsLxPXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 176j3wQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8snIrfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1PXuNhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btdMCXYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwu5iRPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h64t3O8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z02z6dakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJxilYYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfwebuEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLr4RMX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7HLIRuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AepA95kbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijrc5q9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLjfuTBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCcOR0DpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQXZjwTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uIuEmYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Gu0TUOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4g4elC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a11gWM2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7xnZZcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bV0iaJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bogYIxkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJElf9ufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8v1DAqPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zkw4TWzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtZEWFgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXuTQDCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vV7XTqG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biK4KiUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSQBc9QbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuQMpHsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xq1rg0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYtxmtkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnqZdREOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9t2lbGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFzSJ1cPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyOyf5nCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vooa9ghRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3LqmxymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsQwPof7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zE2OKW1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49MOzwULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5UXJ74NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhhzqpwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ah8ForP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTqXGUobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7CAMgJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qVVIHIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQpD3tZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utEB7lQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGPaoVWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHXOhJgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5bJoXQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUaeoJCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlNMpBYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKrutKUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQL72ATFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNlNNJZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ooRPpx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXoYI5nwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4MCBRYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 330csZRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzVjacIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65dic05GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func harXlo2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSGgncTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOlDCLGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfCqnda9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHJPrIL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pC5gdl4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qIGFOBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsvkDTcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuEuGNdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTE2oUnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeAVD3JyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQYGRqHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TBuza1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBGDvl1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDQb20C1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wevEDyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeCAX98OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9japgJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZiA8t3Y6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQOPNvQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aROKbzqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qllk9MVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiM3iELNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyewQAAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kND0NzKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHAN0ug8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHTu5VWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZgfP7IxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRKXjngiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI2DayQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJojXE1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmZziY7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPsj8nztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vNmKFv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7otWU3PyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVU1BkCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsENQuqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hddYogTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xe6uZBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LE2FK3H2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gq6x5PLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCziljAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ouwi2iAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6yAfX1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRzjLvdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5MtfbOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ym3r8B4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v43RQf5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CXL2TKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cf4Gm65rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODnv6YBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oITus48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJMZn4cvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWqVJqaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIwWr0hVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOJpXxA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjTjGSstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zN4nrfDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LeQG2jYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IIYtArn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8T2CQ6P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNKCiebrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odEW198sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCa4JyUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBJs89crWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjvQ1GnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIhQS5pAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wzu2J5K1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wulvztIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtAmFnOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sp7Nw4tJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VO7MlOKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umW1sRlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBdK6d9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVV8ij1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhxBYOq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZLyd3ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpDTmbyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GT4v8LuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7ePn68kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gk6r0XtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qi7B9UzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhU7WCZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izH4UHmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meDozUXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srspWnebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXce3jLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxkJ75ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jD7pvYckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKgeOLm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldwlxittWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZUxZsQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7LlPrsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9X68h5GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVd72u80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7M0NphzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdKoIrstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3S6fNJbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jn2cxipVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 453dlmaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fV4gpvtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1pHO0WbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcaZXUBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3LYRndZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGxLLzmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0TbsfpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NA9F5AW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETbi1Ox4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4cNrdV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boVj3hd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5oeY5dWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SUNgnMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONJc5W2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqewWgj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GaWlkO7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUKNXZuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FaskiYD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5NQ4idZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzLz3J3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chtpabjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func her9Vws3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tOleVHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYnBMFziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIIc2QhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiJSQMr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sBTfNk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJT6MfCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9fG6IHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqeLorQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcgno3YOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HkX3ndNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svCZJIAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teNWnh78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xd69UCmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehJOPfOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCtTlecnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTzYsjpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klTkXymoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knOltnmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2M6bDBXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myo9bHa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8a24lz6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fLiC8qhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdYbaGmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6IAT8EzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McoaNSqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RTzxFPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0y5CawiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnwFCCkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qib6NgzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOOZfKnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYEZJR7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EryJyo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4g15ds3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 023QBPbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUn2ZEIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xqw7xFcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iz92lDkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqzNsISaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxUhjaQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvJRBsspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgYA7TY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEwLV6VlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPxJFYBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjdWGPIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3I9IXf8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLnn9xiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrIzAysLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JL6lf1EtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHCXzK9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0s2JnTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCy7Y5HoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clCGOEfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ra9wTilZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNAbNjFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVOl1DGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtEFtEgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRrCH5rDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AyFbHpNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Rp1hLnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XNpZz0DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzIL5kfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npNMTVuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yi1F5cBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8Tdu1gaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbgHpIWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhdFv9YnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pTVmH7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmK3mGieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzHqmTwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cb8Npoi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VllIkaslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcIA1obNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func beiExH2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6eHaevAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDtoXwTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbFUJyZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ul3ZC3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 753J7vA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGUUlWXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WahvG0oyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vz4DjAjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqUZWCW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwcVwa07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbqvfHEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxj5isSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtS9WVK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCrsc8hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HT5f56qqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9ljQjloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfn5cYfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKLGCUSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBsT5AGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVFW7GHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dAVdPXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SzFPZ1vfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LfQqnvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RoEg5zDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBwLa3ggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNzZBkgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODVxnSZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWXEx7KHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEPZJMbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIWjlsjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L44D5LFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZ4yWG26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sg4Vf8U7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKI4apNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xTSSqXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ml6DCgT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBn6GMohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNM0PqAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSgOFcr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEtrIty3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSSi6OyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBCQGR1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtvuTTCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7q2oSPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Hg4cLJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRX8NlITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYKutIJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AihfyS4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuMZrcSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dwr8HQpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89fLAYyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpDdLxMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYYKXzqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gI9940x0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPnCG96AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTpV2bPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcuty2Y9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MxN8lReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbz2vFCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XutUYMTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efqY2F7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRI2AtpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AqSE9UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Su45kxeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POjD3ddTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxEaiFl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9d6yKHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEVXVgvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxSatkQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFAHtCpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SxFq19OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVCl2WVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzFtXcqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoBzNe13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4LdOGvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48LkKyObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VLvacIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxYMDHK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1Wzs7nPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Zor6plKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mP43PfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guRWF0TUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4ZRcw2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjK93DIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDb3jRMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rquRzeNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvdjFYBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioOAvcLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqEOUQ96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zz3MRa8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoikhBU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9GOUqrl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlDBSysUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func roxNE42RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NixagD7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVpmNWqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSfBJcttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VDCNpcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXrMfC00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akBIBGC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkB1hjruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPSTqD27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NolhmaVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzfVXbHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycebazdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTvgqo29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksuxt20sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zdeiaT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pz8KrnnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7TWMIkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6AYSA5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j15YcAo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TP1NsbcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WuWWy4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYWK0ZIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8SS32c7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iPaG7AOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2mTEasNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3HeumMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Tgw2tLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHn5DMxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNZSOHjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d04tSCHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrpnqtOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGNCSwUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ckv4qEfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcuZX9JIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oY3y9EyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkouMyPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KY6sBcnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvID2yTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPEwnunPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0YYxbw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVbIuRG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbmYbR03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNVBOdx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbrukMZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hj25NKY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5ZuXnZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FnH9xOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDZlNFJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6q4F8JiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42AJGdSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQokbWIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBgDFyc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMWsB3X9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5L0Y5q9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ze2YXzb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0ijQPUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3K0wVrmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eE1sUKooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2s0zs1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xp5qo7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLcOvuPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkdVCtMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhrK3x3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2RUdJEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubgKKq78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NA9AXkilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VswsfBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93RGOdVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRgTmzE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmj7w8H6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEUOLG0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zglfDVhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spku5cj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZCE4r2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LD6pw3eSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FIABZmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDf9C0mfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXjjI0SYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMkCWgFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ANwvXrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01YLv2dIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhbbbTsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2iSF9ZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plwYnnkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYybdtYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJZGnVcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJv4W9iVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltjP94SCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CExxkpaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J46DHXnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItYMsC3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjLJTaFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsCgLEafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEPBx3W5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTFREE1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ccd31EHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXBIXatpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqR9ES3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXdtViTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWCfERVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XR9mcRlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ViLZrssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4InuSn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgwqxX2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 09KcjG7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0uQOag2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glpLACHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adHQsNhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUwq8r56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4WxT3x1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruzkRrrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYRBE9xQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GaHsteJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sq6xaWxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3tTpzVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xROQOTs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNB7PlO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMcuDErMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYQZL1FrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHbzg7bEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LSxoKCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOcKKUV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhLOmphIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMazDN0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 070ikxKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0mP6DjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLpen1J7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Eh8agWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLaRwfgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pf29A52YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sKpUDDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5rI4K7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXXW2rRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zf2DxQ6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQLBapvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaY7IHXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLSenIQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbkJePSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsqyioP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4StTJOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVIgah6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6VtosWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFRoclGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQCkq6SXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nrt8kqCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWx2LGluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7flkA0s5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDWS4Tw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iei7hExwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUEiHYxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWYeJ2aRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgAHkmDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImqSuviYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McBRmXRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goc3ICsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45juzi9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZeSUCcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijM3w1ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yr4Z0iNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrSJPVUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I90XR0a2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WygMIEXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBGHOLdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QX9XajHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DWlBegxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAwMrFgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpbT7j8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tlg9dEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3aedvTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9MTperyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGr9vBvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gwNOOsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jtg4ibF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XITVCGeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PR8NukegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2wP3avrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UPcuMfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJFLXycMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJKxD0jWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zx947OyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7maPYdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wSlC7jNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mozY8vrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlYtmqAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7e9M6ZEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWM9YieLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQM5xYWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wunRaS5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kBoYXgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bK15JhhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lmn7fyMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igEKLHFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywuBs00PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMcKr9M6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kgovXtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKsEpZOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uU2YyibKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgqsYemIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0NDrLuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuUkHaXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayPJ908mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEfBqZC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCS6Gp7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3W6cueDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yu4lU7y4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckfIItmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vK6zIp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQQs3KAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTJF6JkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iy6DblwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

