const std = @import("std");
const math = std.math;

fn isPrime(x: usize) bool {
    const root = math.sqrt(x);

    for (2..root) |i| {
        if (x % i == 0) return false;
    }

    return true;
}

pub fn main() !void {
    const target = 600_851_475_143;
    const sqrt_target = math.sqrt(target);

    var i: u64 = @intCast(@as(u64, sqrt_target));

    while (i >= 2) {
        if (isPrime(i) and target % i == 0) {
            std.debug.print("Result: {d}", .{i});
            break;
        }

        i -= 1;
    }
}
