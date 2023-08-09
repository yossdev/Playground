const std = @import("std");
const net = std.net;
const Server = std.http.Server;

// zig version 0.12.0-dev.25+36c57c3ba
pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const max_header_size = 8192;
    var server = Server.init(allocator, .{ .reuse_address = true });
    defer server.deinit();

    const address = try net.Address.parseIp("127.0.0.1", 8080);
    try server.listen(address);

    while (true) {
        const server_thread = try std.Thread.spawn(.{}, (struct {
                fn apply(s: *Server) !void {
                    var res = try s.accept(.{
                        .allocator = allocator,
                        .header_strategy = .{ .dynamic = max_header_size },
                    });
                    defer res.deinit();
                    defer _ = res.reset();
                    try res.wait();

                    const server_body: []const u8 = "message from server!\n";
                    res.transfer_encoding = .{ .content_length = server_body.len };
                    try res.headers.append("content-type", "text/plain");
                    try res.headers.append("connection", "close");
                    try res.do();

                    _ = try res.writer().writeAll(server_body);
                    try res.finish();
                }
            }).apply, .{&server});

        server_thread.join();
    }
}
