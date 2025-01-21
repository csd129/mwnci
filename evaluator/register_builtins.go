package evaluator

import (
	"mwnci/object"
)

func init() {
	RegisterBuiltin("abs", func(env *object.Environment, args ...object.Object) object.Object { return (Abs(args...)) })
	RegisterBuiltin("accept", func(env *object.Environment, args ...object.Object) object.Object { return (Accept(args...)) })
	RegisterBuiltin("acos", func(env *object.Environment, args ...object.Object) object.Object { return (mathAcos(args...)) })
	RegisterBuiltin("args", func(env *object.Environment, args ...object.Object) object.Object { return (Args(args...)) })
	RegisterBuiltin("asctime", func(env *object.Environment, args ...object.Object) object.Object { return (AscTime(args...)) })
	RegisterBuiltin("asin", func(env *object.Environment, args ...object.Object) object.Object { return (mathAsin(args...)) })
	RegisterBuiltin("assert", func(env *object.Environment, args ...object.Object) object.Object { return (Assert(args...)) })
	RegisterBuiltin("atan", func(env *object.Environment, args ...object.Object) object.Object { return (mathAtan(args...)) })
	RegisterBuiltin("basename", func(env *object.Environment, args ...object.Object) object.Object { return (basenameFun(args...)) })
	RegisterBuiltin("base64dec", func(env *object.Environment, args ...object.Object) object.Object { return (Base64Dec(args...)) })
	RegisterBuiltin("base64enc", func(env *object.Environment, args ...object.Object) object.Object { return (Base64Enc(args...)) })
	RegisterBuiltin("batch", func(env *object.Environment, args ...object.Object) object.Object { return (Batch(args...)) })
	RegisterBuiltin("bin", func(env *object.Environment, args ...object.Object) object.Object { return (Bin(args...)) })
	RegisterBuiltin("bind", func(env *object.Environment, args ...object.Object) object.Object { return (Bind(args...)) })
	RegisterBuiltin("bool", func(env *object.Environment, args ...object.Object) object.Object { return (Bool(args...)) })
	RegisterBuiltin("cargs", func(env *object.Environment, args ...object.Object) object.Object { return (Cargs(args...)) })
	RegisterBuiltin("cat", func(env *object.Environment, args ...object.Object) object.Object { return (ConCat(args...)) })
	RegisterBuiltin("cbrt", func(env *object.Environment, args ...object.Object) object.Object { return (mathCbrt(args...)) })
	RegisterBuiltin("cd", func(env *object.Environment, args ...object.Object) object.Object { return (Cd(args...)) })
	RegisterBuiltin("checkip", func(env *object.Environment, args ...object.Object) object.Object { return (CheckIP(args...)) })
	RegisterBuiltin("chmod", func(env *object.Environment, args ...object.Object) object.Object { return (chmodFun(args...)) })
	RegisterBuiltin("chown", func(env *object.Environment, args ...object.Object) object.Object { return (Chown(args...)) })
	RegisterBuiltin("chr", func(env *object.Environment, args ...object.Object) object.Object { return (Chr(args...)) })
	RegisterBuiltin("close", func(env *object.Environment, args ...object.Object) object.Object { return (Close(args...)) })
	RegisterBuiltin("connect", func(env *object.Environment, args ...object.Object) object.Object { return (Connect(args...)) })
	RegisterBuiltin("copy", func(env *object.Environment, args ...object.Object) object.Object { return (acopyFun(args...)) })
	RegisterBuiltin("cos", func(env *object.Environment, args ...object.Object) object.Object { return (mathCos(args...)) })
	RegisterBuiltin("cosh", func(env *object.Environment, args ...object.Object) object.Object { return (mathCosh(args...)) })
	RegisterBuiltin("count", func(env *object.Environment, args ...object.Object) object.Object { return (Count(args...)) })
	RegisterBuiltin("cp", func(env *object.Environment, args ...object.Object) object.Object { return (FCp(args...)) })
	RegisterBuiltin("cumulative", func(env *object.Environment, args ...object.Object) object.Object { return (Cumulate(args...)) })
	RegisterBuiltin("cut", func(env *object.Environment, args ...object.Object) object.Object { return (FuncCut(args...)) })
	RegisterBuiltin("date", func(env *object.Environment, args ...object.Object) object.Object { return (Date(args...)) })
	RegisterBuiltin("delete", func(env *object.Environment, args ...object.Object) object.Object { return (harrayDelete(args...)) })
	RegisterBuiltin("difference", func(env *object.Environment, args ...object.Object) object.Object { return (ArrDiff(args...)) })
	RegisterBuiltin("dirpath", func(env *object.Environment, args ...object.Object) object.Object { return (Dirpath(args...)) })
	RegisterBuiltin("divmod", func(env *object.Environment, args ...object.Object) object.Object { return (Divmod(args...)) })
	RegisterBuiltin("eval", func(env *object.Environment, args ...object.Object) object.Object { return (evalFun(env, args...)) })
	RegisterBuiltin("exit", func(env *object.Environment, args ...object.Object) object.Object { return (exitFun(args...)) })
	RegisterBuiltin("extend", func(env *object.Environment, args ...object.Object) object.Object { return (ArrExtend(args...)) })
	RegisterBuiltin("fields", func(env *object.Environment, args ...object.Object) object.Object { return (Fields(args...)) })
	RegisterBuiltin("file", func(env *object.Environment, args ...object.Object) object.Object { return (File(args...)) })
	RegisterBuiltin("findfile", func(env *object.Environment, args ...object.Object) object.Object { return (FindFile(args...)) })
	RegisterBuiltin("first", func(env *object.Environment, args ...object.Object) object.Object { return (First(args...)) })
	RegisterBuiltin("float", func(env *object.Environment, args ...object.Object) object.Object { return (floatFun(args...)) })
	RegisterBuiltin("fopen", func(env *object.Environment, args ...object.Object) object.Object { return (FOpen(args...)) })
	RegisterBuiltin("getegid", func(env *object.Environment, args ...object.Object) object.Object { return (Getegid(args...)) })
	RegisterBuiltin("getenv", func(env *object.Environment, args ...object.Object) object.Object { return (getEnvFun(args...)) })
	RegisterBuiltin("geteuid", func(env *object.Environment, args ...object.Object) object.Object { return (Geteuid(args...)) })
	RegisterBuiltin("getgid", func(env *object.Environment, args ...object.Object) object.Object { return (Getgid(args...)) })
	RegisterBuiltin("getos", func(env *object.Environment, args ...object.Object) object.Object { return (GetOs(args...)) })
	RegisterBuiltin("getpid", func(env *object.Environment, args ...object.Object) object.Object { return (getPidFun(args...)) })
	RegisterBuiltin("getppid", func(env *object.Environment, args ...object.Object) object.Object { return (getPpidFun(args...)) })
	RegisterBuiltin("getuid", func(env *object.Environment, args ...object.Object) object.Object { return (Getuid(args...)) })
	RegisterBuiltin("getwd", func(env *object.Environment, args ...object.Object) object.Object { return (Getwd(args...)) })
	RegisterBuiltin("glob", func(env *object.Environment, args ...object.Object) object.Object { return (Glob(args...)) })
	RegisterBuiltin("hasprefix", func(env *object.Environment, args ...object.Object) object.Object { return (HasPrefix(args...)) })
	RegisterBuiltin("hassuffix", func(env *object.Environment, args ...object.Object) object.Object { return (HasSuff(args...)) })
	RegisterBuiltin("hex", func(env *object.Environment, args ...object.Object) object.Object { return (Hex(args...)) })
	RegisterBuiltin("hostname", func(env *object.Environment, args ...object.Object) object.Object { return (Hostname(args...)) })
	RegisterBuiltin("httpget", func(env *object.Environment, args ...object.Object) object.Object { return (HttpGet(args...)) })
	RegisterBuiltin("id", func(env *object.Environment, args ...object.Object) object.Object { return (IdOf(args...)) })
	RegisterBuiltin("include", func(env *object.Environment, args ...object.Object) object.Object { return (Include(env, args...)) })
	RegisterBuiltin("index", func(env *object.Environment, args ...object.Object) object.Object { return (Index(args...)) })
	RegisterBuiltin("input", func(env *object.Environment, args ...object.Object) object.Object { return (Input(args...)) })
	RegisterBuiltin("insert", func(env *object.Environment, args ...object.Object) object.Object { return (ArrayInsert(args...)) })
	RegisterBuiltin("int", func(env *object.Environment, args ...object.Object) object.Object { return (intFun(args...)) })
	RegisterBuiltin("inttoip", func(env *object.Environment, args ...object.Object) object.Object { return (I2IP(args...)) })
	RegisterBuiltin("ipincidr", func(env *object.Environment, args ...object.Object) object.Object { return (Ipincidr(args...)) })
	RegisterBuiltin("iptoint", func(env *object.Environment, args ...object.Object) object.Object { return (IP2I(args...)) })
	RegisterBuiltin("isreadable", func(env *object.Environment, args ...object.Object) object.Object { return (IsReadable(args...)) })
	RegisterBuiltin("iswriteable", func(env *object.Environment, args ...object.Object) object.Object { return (IsWriteable(args...)) })
	RegisterBuiltin("issorted", func(env *object.Environment, args ...object.Object) object.Object { return (Issorted(args...)) })
	RegisterBuiltin("join", func(env *object.Environment, args ...object.Object) object.Object { return (Joiner(args...)) })
	RegisterBuiltin("items", func(env *object.Environment, args ...object.Object) object.Object { return (hashItems(args...)) })
	RegisterBuiltin("keys", func(env *object.Environment, args ...object.Object) object.Object { return (hashKeys(args...)) })
	RegisterBuiltin("kill", func(env *object.Environment, args ...object.Object) object.Object { return (Kill(args...)) })
	RegisterBuiltin("last", func(env *object.Environment, args ...object.Object) object.Object { return (Last(args...)) })
	RegisterBuiltin("len", func(env *object.Environment, args ...object.Object) object.Object { return (FunLen(args...)) })
	RegisterBuiltin("listcidrips", func(env *object.Environment, args ...object.Object) object.Object { return (Listcidrips(args...)) })
	RegisterBuiltin("listen", func(env *object.Environment, args ...object.Object) object.Object { return (Listen(args...)) })
	RegisterBuiltin("log10", func(env *object.Environment, args ...object.Object) object.Object { return (mathLog10(args...)) })
	RegisterBuiltin("loge", func(env *object.Environment, args ...object.Object) object.Object { return (mathLogE(args...)) })
	RegisterBuiltin("ltrim", func(env *object.Environment, args ...object.Object) object.Object { return (LTrim(args...)) })
	RegisterBuiltin("match", func(env *object.Environment, args ...object.Object) object.Object { return (matchFun(args...)) })
	RegisterBuiltin("md5sum", func(env *object.Environment, args ...object.Object) object.Object { return (Md5sum(args...)) })
	RegisterBuiltin("mkdir", func(env *object.Environment, args ...object.Object) object.Object { return (mkdirFun(args...)) })
	RegisterBuiltin("mkdirhier", func(env *object.Environment, args ...object.Object) object.Object { return (Mkdirhier(args...)) })
	RegisterBuiltin("mkfifo", func(env *object.Environment, args ...object.Object) object.Object { return (Mkfifo(args...)) })
	RegisterBuiltin("nslookup", func(env *object.Environment, args ...object.Object) object.Object { return (NSLookup(args...)) })
	RegisterBuiltin("oct", func(env *object.Environment, args ...object.Object) object.Object { return (Oct(args...)) })
	RegisterBuiltin("open", func(env *object.Environment, args ...object.Object) object.Object { return (openFun(args...)) })
	RegisterBuiltin("ord", func(env *object.Environment, args ...object.Object) object.Object { return (Ord(args...)) })
	RegisterBuiltin("pop", func(env *object.Environment, args ...object.Object) object.Object { return (Pop(args...)) })
	RegisterBuiltin("print", func(env *object.Environment, args ...object.Object) object.Object { return (putsFun(args...)) })
	RegisterBuiltin("printf", func(env *object.Environment, args ...object.Object) object.Object { return (printfFun(args...)) })
	RegisterBuiltin("println", func(env *object.Environment, args ...object.Object) object.Object { return (Println(args...)) })
	RegisterBuiltin("product", func(env *object.Environment, args ...object.Object) object.Object { return (Product(args...)) })
	RegisterBuiltin("push", func(env *object.Environment, args ...object.Object) object.Object { return (fpushFun(args...)) })
	RegisterBuiltin("random", func(env *object.Environment, args ...object.Object) object.Object { return (Random(args...)) })
	RegisterBuiltin("range", func(env *object.Environment, args ...object.Object) object.Object { return (Range(args...)) })
	RegisterBuiltin("read", func(env *object.Environment, args ...object.Object) object.Object { return (Read(args...)) })
	RegisterBuiltin("readpw", func(env *object.Environment, args ...object.Object) object.Object { return (Readpw(args...)) })
	RegisterBuiltin("regexp", func(env *object.Environment, args ...object.Object) object.Object { return (Regexp(args...)) })
	RegisterBuiltin("repeat", func(env *object.Environment, args ...object.Object) object.Object { return (Repeat(args...)) })
	RegisterBuiltin("replace", func(env *object.Environment, args ...object.Object) object.Object { return (Replace(args...)) })
	RegisterBuiltin("replaceall", func(env *object.Environment, args ...object.Object) object.Object { return (Replaceall(args...)) })
	RegisterBuiltin("reverse", func(env *object.Environment, args ...object.Object) object.Object { return (Rev(args...)) })
	RegisterBuiltin("rtrim", func(env *object.Environment, args ...object.Object) object.Object { return (RTrim(args...)) })
	RegisterBuiltin("seek", func(env *object.Environment, args ...object.Object) object.Object { return (Seek(args...)) })
	RegisterBuiltin("set", func(env *object.Environment, args ...object.Object) object.Object { return (setFun(args...)) })
	RegisterBuiltin("setenv", func(env *object.Environment, args ...object.Object) object.Object { return (setEnvFun(args...)) })
	RegisterBuiltin("shift", func(env *object.Environment, args ...object.Object) object.Object { return (Shift(args...)) })
	RegisterBuiltin("shuffle", func(env *object.Environment, args ...object.Object) object.Object { return (ShuffleFun(args...)) })
	RegisterBuiltin("sin", func(env *object.Environment, args ...object.Object) object.Object { return (mathSin(args...)) })
	RegisterBuiltin("sinh", func(env *object.Environment, args ...object.Object) object.Object { return (mathSinh(args...)) })
	RegisterBuiltin("sleep", func(env *object.Environment, args ...object.Object) object.Object { return (Sleep(args...)) })
	RegisterBuiltin("slice", func(env *object.Environment, args ...object.Object) object.Object { return (Slice(args...)) })
	RegisterBuiltin("socket", func(env *object.Environment, args ...object.Object) object.Object { return (Socket(args...)) })
	RegisterBuiltin("sort", func(env *object.Environment, args ...object.Object) object.Object { return (Sorted(args...)) })
	RegisterBuiltin("split", func(env *object.Environment, args ...object.Object) object.Object { return (Split(args...)) })
	RegisterBuiltin("splithostport", func(env *object.Environment, args ...object.Object) object.Object { return (SplitHP(args...)) })
	RegisterBuiltin("sprintf", func(env *object.Environment, args ...object.Object) object.Object { return (sprintfFun(args...)) })
	RegisterBuiltin("sqrt", func(env *object.Environment, args ...object.Object) object.Object { return (mathSqrt(args...)) })
	RegisterBuiltin("string", func(env *object.Environment, args ...object.Object) object.Object { return (strFun(args...)) })
	RegisterBuiltin("substr", func(env *object.Environment, args ...object.Object) object.Object { return (subStrFun(args...)) })
	RegisterBuiltin("sum", func(env *object.Environment, args ...object.Object) object.Object { return (Sum(args...)) })
	RegisterBuiltin("swap", func(env *object.Environment, args ...object.Object) object.Object { return (Swapper(args...)) })
	RegisterBuiltin("syslog", func(env *object.Environment, args ...object.Object) object.Object { return (Syslog(args...)) })
	RegisterBuiltin("system", func(env *object.Environment, args ...object.Object) object.Object { return (System(args...)) })
	RegisterBuiltin("tan", func(env *object.Environment, args ...object.Object) object.Object { return (mathTan(args...)) })
	RegisterBuiltin("tanh", func(env *object.Environment, args ...object.Object) object.Object { return (mathTanh(args...)) })
	RegisterBuiltin("time", func(env *object.Environment, args ...object.Object) object.Object { return (Time(args...)) })
	RegisterBuiltin("tmpnam", func(env *object.Environment, args ...object.Object) object.Object { return (Tmpnam(args...)) })
	RegisterBuiltin("tolower", func(env *object.Environment, args ...object.Object) object.Object { return (Lower(args...)) })
	RegisterBuiltin("toupper", func(env *object.Environment, args ...object.Object) object.Object { return (Upper(args...)) })
	RegisterBuiltin("touch", func(env *object.Environment, args ...object.Object) object.Object { return (Touch(args...)) })
	RegisterBuiltin("trim", func(env *object.Environment, args ...object.Object) object.Object { return (Trim(args...)) })
	RegisterBuiltin("trimprefix", func(env *object.Environment, args ...object.Object) object.Object { return (Trimprefix(args...)) })
	RegisterBuiltin("trimsuffix", func(env *object.Environment, args ...object.Object) object.Object { return (TrimSuff(args...)) })
	RegisterBuiltin("type", func(env *object.Environment, args ...object.Object) object.Object { return (typeFun(args...)) })
	RegisterBuiltin("ucount", func(env *object.Environment, args ...object.Object) object.Object { return (Ucount(env, args...)) })
	RegisterBuiltin("union", func(env *object.Environment, args ...object.Object) object.Object { return (ArrUnion(args...)) })
	RegisterBuiltin("uniq", func(env *object.Environment, args ...object.Object) object.Object { return (Uniq(args...)) })
	RegisterBuiltin("unlink", func(env *object.Environment, args ...object.Object) object.Object { return (unlinkFun(args...)) })
	RegisterBuiltin("unsetenv", func(env *object.Environment, args ...object.Object) object.Object { return (unsetEnvFun(args...)) })
	RegisterBuiltin("values", func(env *object.Environment, args ...object.Object) object.Object { return (hashValues(args...)) })
	RegisterBuiltin("whoami", func(env *object.Environment, args ...object.Object) object.Object { return (Whoami(args...)) })
	RegisterBuiltin("write", func(env *object.Environment, args ...object.Object) object.Object { return (Write(args...)) })
	RegisterBuiltin("writefile", func(env *object.Environment, args ...object.Object) object.Object { return (WriteFile(args...)) })
	RegisterBuiltin("zip", func(env *object.Environment, args ...object.Object) object.Object { return (zipFun(args...)) })
}
