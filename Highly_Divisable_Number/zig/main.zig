const std = @import("std");
const math = std.math;
const debug = std.debug;

fn isPrime(a: u64) bool {
    if (a < 2) return false;
    if (a == 2) return true;

    var i: usize = 2;

    while (i * i <= a) : (i += 1) {
        if (a % i == 0) return false;
    }

    return true;
}

fn countDivisors(a: u64) !u64 {
    switch (a) {
        1 => return 1,
        2 => return 2,
        else => {
            const half: u64 = try math.divFloor(u64, a, 2);
            if (half < 2) return 2;
            var n: u64 = 0;
            for (2..half + 1) |divisor| {
                if (a % divisor == 0) n += 1;
            }

            return n + 2;
        },
    }
}

pub fn main() !void {
    var i: u64 = 100;

    while (true) : (i += 1) {
        if (isPrime(i)) continue;

        const divisors = try countDivisors(i);

        if (divisors == 500) {
            debug.print("The first number to have 500 divisors: {d}", .{i});
            break;
        }
    }
}
