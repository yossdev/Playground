use num_bigint::BigUint;

pub fn fib_big_uint(n: usize) -> BigUint {
    if n <= 1 {
        return BigUint::from(n);
    }
    let mut f: Vec<BigUint> = vec![BigUint::from(0u16); 2];
    f[0] = BigUint::from(0u16);
    f[1] = BigUint::from(1u16);
    let mut i = 2;
    while i < n {
        let prev = 0;
        let recent = 1;
        let sum = &f[prev] + &f[recent];
        f.append(&mut vec![sum]);
        f.remove(0);
        i += 1;
    }
    return f[f.len() - 1].clone();
}
