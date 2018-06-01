package main

import "flag"

var (
	flagIPv6                   = flag.Bool("6", false, "enables ipv6 TCP/UDP support")
	flagMaxSpecialConnectionss = flag.Int("C", 0, "sets maximal number of accepted client connections per worker")
	flagHttpPorts              = flag.String("H", "", "comma-separated list of client (HTTP) ports to listen")
	flagSlaves                 = flag.String("M", "", "spawn several slave workers")
	flagProxyTag               = flag.String("P", "", "16-byte proxy tag in hex mode to be passed along with all forwarded queries")
	flagMTProtoSecret          = flag.String("S", "", "16-byte secret in hex mode")
	flagPingInterval           = flag.Float64("T", 5.000, "sets ping interval in second for local TCP connections")
	flagWindowClamp            = flag.String("W", "", "sets window clamp for client TCP connections")
	flagBacklog                = flag.Int("b", 0, "sets backlog size")
	flagConnections            = flag.Int("c", 0, "sets maximal connections number")
	flagDaemonize              = flag.Bool("d", false, "changes between daemonize/not daemonize mode")
	flagHelp                   = flag.Bool("h", false, "prints help and exits")
	flagLog                    = flag.String("l", "", "sets log file name")
	flagPort                   = flag.String("p", "", "<port> or <sport>:<eport> sets listening port number or port range")
	flagUser                   = flag.String("u", "", "sets user name to make setuid")
	flagVerbosity              = flag.Int("v", 0, "sets or increases verbosity level")
	flagAesPwd                 = flag.String("aes-pwd", "", "sets custom secret.conf file")
	flagNice                   = flag.Bool("nice", false, "sets niceness")
	flagMsgBuffersSize         = flag.Int64("msg-buffers-size", 268435456, "sets maximal buffers size")
	flagDisableTcp             = flag.Bool("disable-tcp", false, "do not open listening tcp socket")
	flagCrc32c                 = flag.Bool("crc32c", false, "try to use crc32c instead of crc32 in tcp rpc")
	flagCpuThreads             = flag.Int("cpu-threads", 8, "number of CPU threads (1-64)")
	flagIoThreads              = flag.Int("io-threads", 16, "number of I/O threads (1-64)")
	flagAllowSkipDh            = flag.Bool("allow-skip-dh", false, "Allow skipping DH during RPC handshake")
	flagForceDh                = flag.Bool("force-dh", false, "Force using DH for all outbound RPC connections")
	flagMaxAcceptRate          = flag.Int("max-accept-rate", 0, "max number of connections per second that is allowed to accept")
	flagMaxDhAcceptRate        = flag.Int("max-dh-accept-rate", 0, "max number of DH connections per second that is allowed to accept")
	flagMultithread            = flag.Bool("multithread", false, "run in multithread mode")
	flagTcpCpuThreads          = flag.Int("tcp-cpu-threads", 0, "number of tcp-cpu threads")
	flagTcpIothreads           = flag.Int("tcp-iothreads", 0, "number of tcp-io threads")
	flagNatInfo                = flag.String("nat-info", "", "<local-addr>:<global-addr> sets network address translation for RPC protocol handshake")
	flagAddress                = flag.String("address", "", "tries to bind socket only to specified address")
)

func main() {
	flag.Parse()
}
