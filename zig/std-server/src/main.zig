const std = @import("std");
const net = std.net;
const Server = std.http.Server;

fn handler(res: *Server.Response) !void {
    while (true) {
        defer res.reset();

        const message = "Hello, World! AAAAAAAAAAAAA\n";
        try res.wait();
        // res.headers.transfer_encoding = .{ .content_length = 14 };
        // res.headers.connection = res.request.headers.connection;
        res.transfer_encoding = .{ .content_length = message.len };
        try res.do();
        _ = try res.write(message);

        if (res.connection.conn.closing) break;
    }
}

// zig version 0.11.0-dev.2834+13101295b
pub fn main() !void {
    var server = Server.init(std.heap.page_allocator, .{ .reuse_address = true });
    defer server.deinit();

    try server.listen(try net.Address.parseIp("127.0.0.1", 8080));

    while (true) {
        const res = try server.accept(.{ .dynamic = 8192 });

        const thread = try std.Thread.spawn(.{}, handler, .{res});
        thread.detach();
    }
}
