const std = @import("std");
const debug = std.debug;
const math = std.math;

fn sumOfSquare(x: usize) usize {
    var sum: usize = 0;
    var i: usize = 1;

    while (i <= x) : (i += 1) {
        sum += i * i;
    }

    return sum;
}
fn squareOfSum(x: usize) usize {
    var sum: usize = 0;
    var i: usize = 1;

    while (i <= x) : (i += 1) {
        sum += i;
    }

    return sum * sum;
}

pub fn main() !void {
    const a = sumOfSquare(100);
    const b = squareOfSum(100);

    debug.print("a: {d}\n", .{a});
    debug.print("b: {d}\n", .{b});
    debug.print("Result: {d}", .{b - a});
}
