const std = @import("std");
const math = std.math;
const debug = std.debug;

fn isSpecialTriplet(a: usize, b: usize, c: usize) bool {
    return a > 0 and a < b and b < c and a + b + c == 1000 and a * a + b * b == c * c;
}

pub fn main() !void {
    var c: usize = 300;
    while (c < 450) : (c += 1) {
        var b: usize = 300;

        while (b < c and b < 450) : (b += 1) {
            const a: usize = if (1000 - b - c > 0) 1000 - b - c else break;

            if (isSpecialTriplet(a, b, c)) {
                debug.print("Found special tripler a: {d}, b: {d}, c: {d}\n", .{ a, b, c });
                break;
            }
        }
    }
}
