" Vim Syntax File
" Language:     mwnci
" Creator:      James Mills, prologic at shortcircuit dot net dot au
" Last Change:  31st January 2019

if version < 600
    syntax clear
elseif exists("b:current_syntax")
    finish
endif

syntax case match

syntax keyword xType OS ARCH ARGV ARGC PID NULL PROGNAME STDIN STDOUT STDERR AF_BLACK AF_RED AF_GREEN AF_YELLOW AF_BLUE AF_MAGENTA AF_CYAN AF_WHITE AF_RESET AB_BLACK AB_RED AB_GREEN AB_YELLOW AB_BLUE AB_MAGENTA AB_CYAN AB_WHITE AB_RESET AE_BOLD_ON AE_BOLD_OFF AE_DIM_ON AE_DIM_OFF AE_UNDER_ON AE_UNDER_OFF AE_BLINK_ON AE_BLINK_OFF AE_REV_ON AE_REV_OFF AE_HIDE_ON AE_HIDE_OFF AE_STRIKE_ON AE_STRIKE_OFF PI PI_2 PI_4 E TAU PHI SQRT2 SQRTE SQRTPI SQRTPHI LN2 LOG2E LN10 LOGH10E true   false   null   PI   NULL   ARGC   ARGV   PROGNAME   AF_BLACK   AF_RED   AF_GREEN   AF_YELLOW   AF_BLUE   AF_MAGENTA   AF_CYAN   AF_WHITE   AF_RESET   AB_BLACK   AB_RED   AB_GREEN   AB_YELLOW   AB_BLUE   AB_MAGENTA   AB_CYAN   AB_WHITE   AB_RESET   AE_BOLD_ON   AE_BOLD_OFF   AE_DIM_ON   AE_DIM_OFF   AE_UNDER_ON   AE_UNDER_OFF   AE_BLINK_ON   AE_BLINK_OFF   AE_REV_ON   AE_REV_OFF   AE_HIDE_ON   AE_HIDE_OFF   AE_STRIKE_ON   AE_STRIKE_OFF   OS   ARCH   PID   STDIN   STDOUT   STDERR   PI_2   PI_4   E   TAU   PHI   SQRT2   SQRTE   SQRTPI   SQRTPHI   LN2   LOG2E   LN10   LOGH10E

syntax keyword xKeyword case const default else fn foreach function if in let noop return switch while

syntax keyword xFunction myip openports cidrrange ping addresscount lookupport conncount port2pid extractcert readsitecert readsitechain certsans sitecertsans sitecertexp certexpiredays certtojson iostat getgrent groups getgrgid getgrnam hashgrp getgroups termsize uname df osrelease loadavg tojson loadjson savejson nsip4 nsip6 route netstat nettrans httpcontent httpsave getweb httpstat checkfunctype getenv hex2dec rest bin2dec getblock getfileblock head interpol intersect isrsorted lines map max min rotl rotr rsort tail usort lookupcname lookuphost lookupip lookuptxt lookupptr lookupns lookupmx permsym oct2sym rsortips oct2bin oct2dec enumerate rindex sortips cclass tr issubset ucount bsearch randwords tty first last divmod sqlget sqlcmd getpname hashspw getspnam getbin isdoable filetype istype isfile isblock isdir ischr isfifo islink issock isexist isempty issetuid issetgid issticky isnewer isolder isowner isgroup isexec isequal writetofile uptime free hashstats getcpu datetoepoch vmstat loadyaml saveyaml uptime mpstat iostat pidstat sndev procexists process pnametopid procstat covariance variance stdev mode average mean median harmonicmean geomean getpname hashspw getspnam ljust rjust center title isalpha isalnum isnum isspace islower istitle isupper swapcase capitalize partition zfill pidmem getpwent getpwnam getpwuid hashpw fma ceil degrees dim exp exp2 expm1 factorial floor fmod gcd highest hypot isprime lcm lowest modf pow pow10 radians round trunc free getcpu datetoepoch vmstat netstat route nettrans
syntax keyword xFunction abs accept acos arch args asctime asin assert atan basename base64dec base64enc batch bin bind bool bzcat cargs cat cbrt cd checkip chmod chown chr close cls connect cont[ains] copy cos cosh count cp cumulative cut date delete difference dirpath env eval exit extend fields file findfile float fopen fromkeys getegid geteuid getgid getos getpid getppid getuid getwd glob graph gzcat hasprefix hassuffix hex hostname httpget httppost id include index input insert int inttoip ipincidr iptoint isreadable isjson iswriteable issorted jsontoyaml join items keys kill len listcidrips listen log10 loge ltrim match md5sum memstat mindex mkdir mkdirhier mkfifo nslookup oct open ord pjson pop print printat printf println product push random randsample range read readpw regexp repeat replace replaceall reverse rtrim sec2time seek set setenv shift shuffle sin sinh sleep slice socket sorted sort split splithostport sprintf sqrt stac stat string substr sum swap syslog system tac tan tanh time tmpnam tolower toupper touch trim trimprefix trimsuffix type union uniq unlink unsetenv unzip values whoami write writefile yamltojson zip

syntax match xOperator "=="
syntax match xOperator "\v!\="
syntax match xOperator "\v<"
syntax match xOperator "\v>"
syntax match xOperator "\v!"
syntax match xOperator "\v\+"
syntax match xOperator "\v-"
syntax match xOperator "\v\*"
syntax match xOperator "\v/"
syntax match xOperator "\v:\="
syntax match xOperator "\v\="
syntax match xOperator "\v&"
syntax match xOperator "\v\|"
syntax match xOperator "\v^"
syntax match xOperator "\v\~"
syntax match xOperator "\v&&"
syntax match xOperator "\v\|\|"

syntax region xString start=/"/ skip=/\\./ end=/"/

syntax region xComment start='#' end='$' keepend
syntax region xComment start='//' end='$' keepend

highlight link xOperator Operator
highlight link xType Type
highlight link xKeyword Keyword
highlight link xFunction Function
highlight link xString String
highlight link xComment Comment


let b:current_syntax = "mwnci"
