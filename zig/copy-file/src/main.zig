const std = @import("std");
const print = std.debug.print;
const dbg_assert = std.debug.assert;

pub fn main() !void {
    var timer = try std.time.Timer.start();

    var heap_arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer heap_arena.deinit();
    const arena_alloc = heap_arena.allocator();

    var args = try std.process.argsAlloc(arena_alloc);
    defer std.process.argsFree(arena_alloc, args);

    const ROOT = "D:/Code/Playground";
    const CWD = "zig/copy-file";
    const DIR = "dst";

    var cleanup = false;
    if (args.len > 1) {
        cleanup = std.mem.eql(u8, args[1], "clean");
    }

    if (cleanup) {
        const path_dir = try std.fmt.allocPrint(arena_alloc, "{s}/{s}/{s}", .{ ROOT, CWD, DIR });
        // defer arena_alloc.free(path_dir);

        try cleanUp(path_dir);
    } else {
        const count = 10_000;

        const src = try std.fmt.allocPrint(arena_alloc, "{s}/common/src/test.md", .{ROOT});
        // defer arena_alloc.free(src);

        try copyFile(count, src, ROOT, CWD, DIR);
    }

    const elapsed = timer.lap() / std.time.ns_per_ms;
    print("Time took: {d}ms\n", .{elapsed});
}

fn copyFile(n: u16, src: []const u8, root: *const [18:0]u8, cwd: *const [13:0]u8, dir: *const [3:0]u8) !void {
    print("----- COPYING -----\n", .{});

    var alloc = std.heap.page_allocator;

    var sum_n_bytes: u64 = 0;

    var i: u16 = 0;
    while (i < n) : (i += 1) {
        const dest = try std.fmt.allocPrint(alloc, "{s}/{s}/{s}/test{d}.md", .{ root, cwd, dir, i });
        defer alloc.free(dest);

        try std.fs.copyFileAbsolute(src, dest, .{});

        const file = try std.fs.openFileAbsolute(dest, .{});
        defer std.fs.File.close(file);

        const fstat = try std.fs.File.stat(file);
        sum_n_bytes += fstat.size;
    }

    print("Total writen: {d} bytes\n", .{sum_n_bytes});
}

fn cleanUp(path_dir: []u8) !void {
    print("----- CLEANING -----\n", .{});
    try std.fs.deleteTreeAbsolute(path_dir);
    try std.fs.makeDirAbsolute(path_dir);
}
