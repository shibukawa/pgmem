import net from "node:net";

/**
 * A minimal PostgreSQL client for the simple query protocol, so the package
 * tests need no driver. query() resolves to rows of strings.
 */
export async function connectWire(url) {
  const u = new URL(url);
  const socket = net.connect({ host: u.hostname, port: Number(u.port) });
  await new Promise((resolve, reject) => {
    socket.once("connect", resolve);
    socket.once("error", reject);
  });
  let buf = Buffer.alloc(0);
  let rows = [];
  let error = null;
  let waiter = null;
  const settle = (err) => {
    const w = waiter;
    const r = rows;
    waiter = null;
    rows = [];
    error = null;
    w?.(err, r);
  };
  socket.on("data", (chunk) => {
    buf = Buffer.concat([buf, chunk]);
    while (buf.length >= 5 && buf.length >= 1 + buf.readInt32BE(1)) {
      const type = String.fromCharCode(buf[0]);
      const body = buf.subarray(5, 1 + buf.readInt32BE(1));
      buf = buf.subarray(1 + buf.readInt32BE(1));
      if (type === "D") {
        const row = [];
        let off = 2;
        for (let i = body.readInt16BE(0); i > 0; i--) {
          const len = body.readInt32BE(off);
          off += 4;
          row.push(len < 0 ? null : body.toString("utf8", off, off + len));
          off += Math.max(len, 0);
        }
        rows.push(row);
      } else if (type === "E") {
        const fields = Object.fromEntries(
          body
            .toString()
            .split("\0")
            .filter(Boolean)
            .map((f) => [f[0], f.slice(1)]),
        );
        error = Object.assign(new Error(fields.M), { code: fields.C });
      } else if (type === "Z") {
        settle(error);
      }
    }
  });
  socket.on("error", () => {});
  socket.on("close", () => settle(error ?? new Error("connection closed")));
  const send = (sql) =>
    new Promise((resolve, reject) => {
      waiter = (err, r) => (err ? reject(err) : resolve(r));
      if (sql === undefined) return;
      const text = Buffer.from(sql + "\0");
      const msg = Buffer.alloc(5 + text.length);
      msg.write("Q");
      msg.writeInt32BE(4 + text.length, 1);
      text.copy(msg, 5);
      socket.write(msg);
    });
  const params = Buffer.from(`user\0${decodeURIComponent(u.username)}\0database\0${u.pathname.slice(1)}\0\0`);
  const startup = Buffer.alloc(8 + params.length);
  startup.writeInt32BE(startup.length, 0);
  startup.writeInt32BE(196608, 4);
  params.copy(startup, 8);
  const ready = send(undefined);
  socket.write(startup);
  await ready;
  return {
    query: send,
    close: () =>
      new Promise((resolve) => {
        if (socket.destroyed) return resolve();
        socket.once("close", resolve);
        socket.end(Buffer.from([0x58, 0, 0, 0, 4]));
      }),
  };
}
