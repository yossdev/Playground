use std::time::SystemTime;
mod fib_big_uint;

fn main() {
    let start = SystemTime::now();
    // let mut collections: Vec<u64> = Vec::new();
    // const X: u32 = 500_000;
    const N: usize = 301;
    // let mut i = 0;
    // while i < X {
    //     let res = fib(N);
    //     collections.push(res as u64);
    //     i += 1;
    // }

    // with BigUint
    let res = fib_big_uint::fib_big_uint(N);
    print!("{}th fib: {}\n", N, res);

    let end = SystemTime::now();
    let elapsed = end.duration_since(start);
    // print!("Last collections: {}\n", collections[(X - 1) as usize]);
    print!(
        "Time took: {}ms, n: {}\n",
        elapsed.unwrap_or_default().as_millis(),
        N
    )
}

// fn fib(n: usize) -> usize {
//     if n <= 1 {
//         return n;
//     }
//     let mut f: Vec<usize> = vec![0; n];
//     f[0] = 0;
//     f[1] = 1;
//     for i in 2..n {
//         let prev = i - 2;
//         let current = i - 1;
//         f[i] = f[prev] + f[current];
//     }
//     return f[n - 1];
// }
