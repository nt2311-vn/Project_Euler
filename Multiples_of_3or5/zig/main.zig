const std = @import("std");

pub fn main() !void {
    var sum: usize = 0;

    for (1..1000) |num| {
        if (num % 3 == 0 or num % 5 == 0) {
            sum += num;
        }
    }

    std.debug.print("Result: {d}\n", .{sum});
}
