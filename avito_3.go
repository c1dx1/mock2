package main

import (
    "context"
	  "fmt"
	  "log"
	  "math/rand"
	  "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
    rnd := rand.Int63n(5000)
    time.Sleep(time.Duration(rnd) * time.Millisecond)

    return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc(ctx context.Context) (int64, error) {
   start := time.Now()
   defer func() {
       log.Printf("Время выполнения запроса: %v\n", time.Since(start))
   }()
   
   resCh := make(chan int64)
   go func() {
        defer close(resCh)
        val := unpredictableFunc()
        resCh <- val
   }()
    
    select {
        case <-ctx.Done():
            return 0, ctx.Err()
        case val := <-resCh:
            return val, nil
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
    defer cancel()
    
    val, err := predictableFunc(ctx)
    fmt.Printf("Value: %v Error: %v\n", val, err)
}
