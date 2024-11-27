const std = @import("std");
const math = std.math;
const debug = std.debug;

fn isPrime(x: u64) bool {
    if (x <= 2) return true;

    var i: u64 = 2;

    while (i < x) {
        if (x % i == 0) {
            return false;
        }
        i += 1;
    }
    return true;
}

fn findPrimeFactor(x: u64, alloc: std.mem.Allocator) !std.ArrayList(u64) {
    var prime_factor = std.ArrayList(u64).init(alloc);
    defer prime_factor.deinit();

    if (x <= 2) {
        return prime_factor;
    }

    for (2..x) |i| {
        if (isPrime(i) and x % i == 0) {
            try prime_factor.append(i);
        }
    }

    return prime_factor;
}

fn toPrimeFactor(x: u64, alloc: std.mem.Allocator) !std.AutoHashMap(u64, u64) {
    var factor_map = std.AutoHashMap(u64, u64).init(alloc);
    defer factor_map.deinit();

    var num = x;
    const prime_factor = try findPrimeFactor(x, alloc);
    defer prime_factor.deinit();

    for (prime_factor.items) |factor| {
        while (num % factor == 0) {
            try factor_map.put(factor, factor_map.get(factor).? + 1);
            num /= factor;
        }
    }

    return factor_map;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    var prime_map = std.AutoHashMap(u64, u64).init(allocator);
    defer prime_map.deinit();

    var non_prime = std.ArrayList(u64).init(allocator);
    defer non_prime.deinit();

    for (2..20) |i| {
        if (isPrime(i)) {
            try prime_map.put(i, 1);
        } else {
            try non_prime.append(i);
        }
    }

    for (non_prime.items) |num| {
        var factor_nonprime = try toPrimeFactor(num, allocator);
        defer factor_nonprime.deinit();

        var it = factor_nonprime.iterator();

        while (it.next()) |entry| {
            const key = entry.key_ptr.*;
            const value = entry.value_ptr.*;

            const current_val = prime_map.get(key) orelse 0;

            if (current_val < value) {
                try prime_map.put(key, value);
            }
        }
    }

    var smallest: u64 = 1;
    var it = prime_map.iterator();

    while (it.next()) |entry| {
        const prime = entry.key_ptr.*;
        const exponent = entry.value_ptr.*;

        smallest *= math.pow(u64, prime, exponent);
    }

    debug.print("Smallest: {d}", .{smallest});
}
