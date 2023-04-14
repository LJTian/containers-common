//go:build linux && sw64
// +build linux,sw64

// Special signal handling for mips architecture
package signal

// Copyright 2013-2018 Docker, Inc.

// NOTE: this package has originally been copied from github.com/docker/docker.

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	sigrtmin = 34
	sigrtmax = 64

	SIGWINCH = syscall.SIGWINCH
)

// signalMap is a map of Linux signals.
var signalMap = map[string]syscall.Signal{
	//1) SIGHUP       2) SIGINT       3) SIGQUIT      4) SIGILL       5) SIGTRAP
	"HUP":  unix.SIGHUP,
	"INT":  unix.SIGINT,
	"QUIT": unix.SIGQUIT,
	"ILL":  unix.SIGILL,
	"TRAP": unix.SIGTRAP,
	//6) SIGABRT      7) SIGEMT       8) SIGFPE       9) SIGKILL     10) SIGBUS
	"ABRT": unix.SIGABRT,
	"EMT ": unix.SIGEMT,
	"FPE ": unix.SIGFPE,
	"KILL": unix.SIGKILL,
	"BUS":  unix.SIGBUS,
	//11) SIGSEGV     12) SIGSYS      13) SIGPIPE     14) SIGALRM     15) SIGTERM
	"SEGV": unix.SIGSEGV,
	"SYS":  unix.SIGSYS,
	"PIPE": unix.SIGPIPE,
	"ALRM": unix.SIGALRM,
	"TERM": unix.SIGTERM,
	//16) SIGURG      17) SIGSTOP     18) SIGTSTP     19) SIGCONT     20) SIGCHLD
	"URG":  unix.SIGURG,
	"STOP": unix.SIGSTOP,
	"TSTP": unix.SIGTSTP,
	"CONT": unix.SIGCONT,
	"CHLD": unix.SIGCHLD,
	//21) SIGTTIN     22) SIGTTOU     23) SIGIO       24) SIGXCPU     25) SIGXFSZ
	"TTIN":  unix.SIGTTIN,
	"TTOU ": unix.SIGTTOU,
	"IO":    unix.SIGIO,
	"XCPU":  unix.SIGXCPU,
	"XFSZ":  unix.SIGXFSZ,
	//26) SIGVTALRM   27) SIGPROF     28) SIGWINCH    29) SIGINFO     30) SIGUSR1
	"VTALRM": unix.SIGVTALRM,
	"PROF ":  unix.SIGPROF,
	"WINCH":  unix.SIGWINCH,
	"INFO ":  unix.SIGINFO,
	"USR1":   unix.SIGUSR1,
	//31) SIGUSR2     34) SIGRTMIN    35) SIGRTMIN+1  36) SIGRTMIN+2  37) SIGRTMIN+3
	"USR2 ":   unix.SIGUSR2,
	"RTMIN":   sigrtmin,
	"RTMIN+1": sigrtmin + 1,
	"RTMIN+2": sigrtmin + 2,
	"RTMIN+3": sigrtmin + 3,
	//38) SIGRTMIN+4  39) SIGRTMIN+5  40) SIGRTMIN+6  41) SIGRTMIN+7  42) SIGRTMIN+8
	"RTMIN+4": sigrtmin + 4,
	"RTMIN+5": sigrtmin + 5,
	"RTMIN+6": sigrtmin + 6,
	"RTMIN+7": sigrtmin + 7,
	"RTMIN+8": sigrtmin + 8,
	//43) SIGRTMIN+9  44) SIGRTMIN+10 45) SIGRTMIN+11 46) SIGRTMIN+12 47) SIGRTMIN+13
	"RTMIN+9":  sigrtmin + 9,
	"RTMIN+10": sigrtmin + 10,
	"RTMIN+11": sigrtmin + 11,
	"RTMIN+12": sigrtmin + 12,
	"RTMIN+13": sigrtmin + 13,
	//48) SIGRTMIN+14 49) SIGRTMIN+15 50) SIGRTMAX-14 51) SIGRTMAX-13 52) SIGRTMAX-12
	"RTMIN+14": sigrtmin + 14,
	"RTMIN+15": sigrtmin + 15,
	"RTMAX-14": sigrtmax - 14,
	"RTMAX-13": sigrtmax - 13,
	"RTMAX-12": sigrtmax - 12,
	//53) SIGRTMAX-11 54) SIGRTMAX-10 55) SIGRTMAX-9  56) SIGRTMAX-8  57) SIGRTMAX-7
	"RTMAX-11": sigrtmax - 11,
	"RTMAX-10": sigrtmax - 10,
	"RTMAX-9":  sigrtmax - 9,
	"RTMAX-8":  sigrtmax - 8,
	"RTMAX-7":  sigrtmax - 7,
	//58) SIGRTMAX-6  59) SIGRTMAX-5  60) SIGRTMAX-4  61) SIGRTMAX-3  62) SIGRTMAX-2
	"RTMAX-6": sigrtmax - 6,
	"RTMAX-5": sigrtmax - 5,
	"RTMAX-4": sigrtmax - 4,
	"RTMAX-3": sigrtmax - 3,
	"RTMAX-2": sigrtmax - 2,
	//63) SIGRTMAX-1  64) SIGRTMAX
	"RTMAX-1": sigrtmax - 1,
	"RTMAX":   sigrtmax,
}

// CatchAll catches all signals and relays them to the specified channel.
func CatchAll(sigc chan os.Signal) {
	handledSigs := make([]os.Signal, 0, len(signalMap))
	for _, s := range signalMap {
		handledSigs = append(handledSigs, s)
	}
	signal.Notify(sigc, handledSigs...)
}

// StopCatch stops catching the signals and closes the specified channel.
func StopCatch(sigc chan os.Signal) {
	signal.Stop(sigc)
	close(sigc)
}
