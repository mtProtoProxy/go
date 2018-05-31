package mtproto

import "net"

// #define tls_push()      { struct tl_out_state *tlio_out = tl_out_state_alloc ();
// #define tls_pop()       tl_out_state_free (tlio_out); }
// #define TLS_START(C)            tls_push(); tls_init_tcp_raw_msg (tlio_out, C, 0);
// #define TLS_START_UNALIGN(C)    tls_push(); tls_init_tcp_raw_msg_unaligned (tlio_out, C, 0);
// #define TLS_END         tl_store_end_ext (0); tls_pop();

/* DH key exchange protocol data structures */
const (
	CodeReqPq             = 0x60469778
	CodeReqPqMulti        = 0xbe7e8ef1
	CodeReqDhParams       = 0xd712e4be
	CodeSetClientDhParams = 0xf5045f1f
)

/* RPC for front/proxy */
const (
	RpcProxyReq  = 0x36cef1ee
	RpcProxyAns  = 0x4403da0d
	RpcCloseConn = 0x1fcf425d
	RpcCloseExt  = 0x5eb634a2
	RpcSimpleAck = 0x3bac409b
)

/* not really a limit, for struct encrypted_message only */
// const MAX_MESSAGE_INTS     16384
const (
	MaxMessageInts      = 1048576
	MaxProtoMessageInts = 524288
)

// #pragma pack(push,4)
type EncryptedMessage struct {
	// unencrypted header
	AuthKeyID int64  // long long auth_key_id;
	MsgKey    net.IP // char msg_key[16];

	// encrypted part, starts with encrypted header
	ServerSalt int64 // long long server_salt;
	SessionID  int64 // long long session_id;

	// first message follows
	MsgID   int64 // long long msg_id;
	SeqNo   int
	MsgLen  int                     // int msg_len; // divisible by 4
	Message [MaxMessageInts + 8]int // int message[MAX_MESSAGE_INTS + 8];
}

const MaxProxyExtraBytes = 16384

type RpcProxyReq struct {
	Type       int    // int type; // RPC_PROXY_REQ
	Flags      int    // int flags;
	ExtConnID  int64  // long long ext_conn_id;
	RemoteIPv6 net.IP // unsigned char remote_ipv6[16];
	RemotePort int    // int remote_port;
	OurIPv6    net.IP // unsigned char our_ipv6[16];
	OurPort    int    // int our_port;
	// union {
	Data int // int data[0];
	// struct {
	ExtraBytes int                         // int extra_bytes;
	Extra      [MaxProxyExtraBytes / 4]int // int extra[MAX_PROXY_EXTRA_BYTES / 4];
	// }
	// };
}

type RpcProxyAns struct {
	Type      int   // int type; // RPC_PROXY_ANS
	Flags     int   // int flags; // +16 = small error packet, +8 = flush immediately
	ExtConnID int64 // long long ext_conn_id;
	Data      int   // int data[];
}

type RpcCloseConn struct {
	Type      int   // int type; // RPC_CLOSE_CONN
	ExtConnID int64 // long long ext_conn_id;
}

type RpcCloseExt struct {
	Type      int   // int type; // RPC_CLOSE_EXT
	ExtConnID int64 // long long ext_conn_id;
}

type RpcSimpleAck struct {
	Type       int   // int type; // RPC_SIMPLE_ACK
	ExtConnID  int64 // long long ext_conn_id;
	ConfirmKey int   // int confirm_key;
}

// #pragma pack(pop)
