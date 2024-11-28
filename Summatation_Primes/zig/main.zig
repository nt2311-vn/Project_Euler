const std = @import("std");
const debug = std.debug;

fn isPrime(a: usize) bool {
    switch (a) {
        2 => return true,
        else => {
            var i: usize = 2;
            while (i * i <= a) : (i += 1) {
                if (a % i == 0) return false;
            }

            return true;
        },
    }
}
pub fn main() !void {
    var sum: usize = 0;
    var i: usize = 2;

    while (i < 2_000_000) : (i += 1) {
        if (isPrime(i)) {
            sum += i;
        }
    }

    debug.print("Sumatation of primes under two million: {d}\n", .{sum});
}
