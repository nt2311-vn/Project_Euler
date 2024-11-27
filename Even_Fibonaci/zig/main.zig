const std = @import("std");

pub fn main() !void {
    var a: usize = 1;
    var b: usize = 1;
    var c: usize = undefined;
    var sum: usize = 0;

    while (b < 4_000_000) {
        c = a + b;
        a = b;
        b = c;

        if (c % 2 == 0) {
            sum += c;
        }
    }

    std.debug.print("Result: {d}", .{sum});
}
