const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    var timer = try std.time.Timer.start();
    const n = 184; // max u128 fib
    var i: u16 = 0;
    while (i < n) : (i += 1) {
        print("{d}\n", .{fib(i)});
    }

    const elapsed = timer.lap() / std.time.ns_per_ms;
    print("Time lap: {d}ms\n", .{elapsed});
}

fn fib(n: u16) u128 {
    if (n <= 1) {
        return n;
    }

    var a: u128 = 0;
    var b: u128 = 1;

    var i: u16 = 2;
    while (i <= n) : (i += 1) {
        var c = a + b;
        a = b;
        b = c;
    }

    return b;
}
