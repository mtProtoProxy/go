package common

const tlStat = 0x9d56e6b2

const (
	rpcInvokeReq        = 0x2374df3d
	rpcInvokeKphpReq    = 0x99a37fda
	rpcReqRunning       = 0x346d5efa
	rpcReqError         = 0x7ae432f5
	rpcReqResult        = 0x63aeda4e
	rpcReady            = 0x6a34cac7
	rpcStopReady        = 0x59d86654
	rpcSendSessionMsg   = 0x1ed5a3cc
	rpcResponseIndirect = 0x2194f56e
	rpcPing             = 0x5730a2df
	rpcPong             = 0x8430eaa7
)

const (
	rpcDestActor      = 0x7568aabd
	rpcDestActorFlags = 0xf0a5acf7
	rpcDestFlags      = 0xe352035e
	rpcReqResultFlags = 0x8cc84ce1
)

const maxTlStringLength = 0xffffff

const tlErrorRetry = 503

const (
	tlBoolTrue  = 0x997275b5
	tlBoolFalse = 0xbc799737
)

const tlBoolStat = 0x92cbcbfa

const tlTrue = 0x3fedd339

const (
	tlInt    = 0xa8509bda
	tlLong   = 0x22076cba
	tlDouble = 0x2210c154
	tlString = 0xb5286e24
)

const (
	tlMaybeTrue  = 0x3f9c8ef8
	tlMaybeFalse = 0x27930a7b
)

const (
	tlVector      = 0x1cb5c415
	tlVectorTotal = 0x10133f47
	tlTuple       = 0x9770768a
)

const tlDictionary = 0x1f4c618f

// Error codes
// Query syntax errors -1000...-1999
const (
	tErrorSyntax         = -1000
	tlErrorExtraData     = -1001
	tlErrorHeader        = -1002
	tlErrorWrongQueryID  = -1003
	tlErrorNotEnoughData = -1004
)

// Syntax ok, bad can not start query. -2000...-2999
const (
	tlErrorUnknownFunctionID   = -2000
	tlErrorProxyNoTarget       = -2001
	tlErrorWrongActorID        = -2002
	tlErrorTooLongString       = -2003
	tlErrorValueNotInRange     = -2004
	tlErrorQueryIncorrect      = -2005
	tlErrorBadValue            = -2006
	tlErrorBinlogDisabled      = -2007
	tlErrorFeatureDisabled     = -2008
	tlErrorQueryIsEmpty        = -2009
	tlErrorInvalidConnectionID = -2010
	tlErrorWrongSplit          = -2011
	tlErrorTooBigOffset        = -2012
)

// Error processing query -3000...-3999
const (
	tlErrorQueryTimeout         = -3000
	tlErrorProxyInvalidResponse = -3001
	tlErrorNoConnections        = -3002
	tlErrorInternal             = -3003
	tlErrorAioFail              = -3004
	tlErrorAioTimeout           = -3005
	tlErrorBinlogWaitTimeout    = -3006
	tlErrorAioMaxRetryExceeded  = -3007
	tlErrorTTL                  = -3008
	tlErrorBadMetafile          = -3009
	tlErrorNotReady             = -3010
)

const (
	tlErrorStorageCacheMiss          = -3500
	tlErrorStorageCacheNoMtprotoConn = -3501
)

// Different errors -4000...-4999
const tlErrorUnknown = -4000

// TL_IS_USER_ERROR = (x) ((x) <= -1000 && (x) > -3000)

const tlNamespace = "TL_"

// CONCAT = (a,b) a ## b
// TLN = (nspc,name) CONCAT (nspc, name)
// TLG = (name) TL_ ## name
// TL = (x) TLN (TL_NAMESPACE, x)
