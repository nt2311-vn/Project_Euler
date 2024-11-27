const std = @import("std");
const debug = std.debug;

fn isPalindrome(x: usize) !bool {
    var buffer: [12]u8 = undefined;
    const x_to_str = try std.fmt.bufPrint(&buffer, "{d}", .{x});

    var i: usize = 0;
    var j: usize = x_to_str.len - 1;

    while (i < j) {
        if (x_to_str[i] != x_to_str[j]) return false;
        i += 1;
        j -= 1;
    }

    return true;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    var palin_lst = std.ArrayList(usize).init(allocator);
    defer palin_lst.deinit();

    for (100..1000) |x| {
        for (100..1000) |y| {
            const product = x * y;

            if (try isPalindrome(product)) {
                try palin_lst.append(product);
            }
        }
    }

    var max_num: usize = 0;

    for (palin_lst.items) |palin_num| {
        if (max_num < palin_num) {
            max_num = palin_num;
        }
    }

    debug.print("Result: {d}", .{max_num});
}
