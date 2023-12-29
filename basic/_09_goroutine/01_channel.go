package _09_goroutine

import (
	"fmt"
	"time"
)

/**
 * 反面案例(1): 死锁
 *     无缓冲的 channel 在写入数据时会阻塞程序后续执行,
 *     直到其他协程读取这个 channel 中的数据后, 程序会在阻塞处继续往后执行。
 *     但是读取 channel 数据的逻辑在阻塞逻辑后面, 得不到执行的机会, 所以程序会永远阻塞, 引发死锁。
 */
func channel01() {
	ch := make(chan int)
	ch <- 1 // fatal error: all goroutines are asleep - deadlock!
	v := <-ch
	fmt.Println(v)
	time.Sleep(1 * time.Second)
}

/**
 * 正确案例(1): 读写无缓冲的 channel 都会发生阻塞
 */
func channel02() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	v := <-ch // 此处读取 channel 时会发生阻塞, 直到上面的协程执行 channel 写入的那一刻, 此处会重新调度, 执行后续逻辑
	fmt.Println(v)
	time.Sleep(1 * time.Second)
}

/**
 * 反面案例(2): 读取已关闭的 channel 得到值是 channel 类型的零值
 *    注意: 读取已关闭的 channel 值会导致无法区分数据时候有意义
 *          因为 int 零值和正常写入 channel 的 0 都是 0, 后者有正常的业务意义
 */
func channel03() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch) // 关闭 channel
	}()
	v := <-ch      // 此处读取 channel 时会发生阻塞, 直到上面的协程执行 channel 写入的那一刻, 此处会重新调度, 执行后续逻辑
	fmt.Println(v) // 1
	v = <-ch       // 已关闭的 channel 依旧可以读取到值, 值为 channel 类型的零值, 如 int 的零值为 0
	fmt.Println(v) // 0
	time.Sleep(1 * time.Second)
}

/**
 * 正确案例(2): 读取已关闭的 channel 得到值是 channel 类型的零值
 */
func channel04() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch) // 关闭 channel
	}()
	v := <-ch      // 此处读取 channel 时会发生阻塞, 直到上面的协程执行 channel 写入的那一刻, 此处会重新调度, 执行后续逻辑
	fmt.Println(v) // 1
	v, ok := <-ch  // 读取到了其他协程写入的值, ok = true, 反之 ok = false
	if ok {
		fmt.Println(v) // 0
	} else {
		fmt.Println("channel 已关闭")
	}
	time.Sleep(1 * time.Second)
}

/**
 * 正确案例(3): 使用 for ... range 来读取 channel
 *    Ps: for ... range 不会读取关闭后的零值
 *        所以有明确数量数据的 channel 推荐使用 for ... range 来读取
 *        无明确数量数据(如无限 for 循环)的 channel 则不适用
 */
func channel05() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
