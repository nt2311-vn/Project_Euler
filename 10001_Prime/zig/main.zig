const std = @import("std");
const debug = std.debug;
const math = std.math;

fn isPrime(x: usize) bool {
    if (x < 2) return false;
    if (x == 2) return true;

    var i: usize = 2;
    while (i * i <= x) : (i += 1) {
        if (x % i == 0) return false;
    }

    return true;
}

pub fn main() !void {
    var n: usize = 0;
    var i: usize = 2;

    while (true) : (i += 1) {
        if (isPrime(i)) {
            n += 1;

            if (n == 10001) {
                debug.print("Result: {d}", .{i});
                break;
            }
        }
    }
}
