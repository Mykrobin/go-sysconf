// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_cgotest

/*
#include <gnu/libc-version.h>
#include <stdlib.h>
#include <unistd.h>
*/
import "C"

import (
	"strconv"
	"strings"
	"testing"

	"github.com/tklauser/go-sysconf"
	"golang.org/x/sys/unix"
)

type testCase struct {
	goVar int
	cVar  C.int
	name  string
}

func testSysconfCgoMatch(t *testing.T) {
	testCases := []testCase{
		{sysconf.SC_AIO_LISTIO_MAX, C._SC_AIO_LISTIO_MAX, "AIO_LISTIO_MAX"},
		{sysconf.SC_AIO_MAX, C._SC_AIO_MAX, "AIO_MAX"},
		{sysconf.SC_AIO_PRIO_DELTA_MAX, C._SC_AIO_PRIO_DELTA_MAX, "AIO_PRIO_DELTA_MAX"},
		{sysconf.SC_ARG_MAX, C._SC_ARG_MAX, "ARG_MAX"},
		{sysconf.SC_ATEXIT_MAX, C._SC_ATEXIT_MAX, "ATEXIT_MAX"},
		{sysconf.SC_BC_BASE_MAX, C._SC_BC_BASE_MAX, "BC_BASE_MAX"},
		{sysconf.SC_BC_DIM_MAX, C._SC_BC_DIM_MAX, "BC_DIM_MAX"},
		{sysconf.SC_BC_SCALE_MAX, C._SC_BC_SCALE_MAX, "BC_SCALE_MAX"},
		{sysconf.SC_BC_STRING_MAX, C._SC_BC_STRING_MAX, "BC_STRING_MAX"},
		{sysconf.SC_CHILD_MAX, C._SC_CHILD_MAX, "CHILD_MAX"},
		{sysconf.SC_CLK_TCK, C._SC_CLK_TCK, "CLK_TCK"},
		{sysconf.SC_COLL_WEIGHTS_MAX, C._SC_COLL_WEIGHTS_MAX, "COLL_WEIGHTS_MAX"},
		{sysconf.SC_DELAYTIMER_MAX, C._SC_DELAYTIMER_MAX, "DELAYTIMER_MAX"},
		{sysconf.SC_EXPR_NEST_MAX, C._SC_EXPR_NEST_MAX, "EXPR_NEST_MAX"},
		{sysconf.SC_HOST_NAME_MAX, C._SC_HOST_NAME_MAX, "HOST_NAME_MAX"},
		{sysconf.SC_IOV_MAX, C._SC_IOV_MAX, "IOV_MAX"},
		{sysconf.SC_LINE_MAX, C._SC_LINE_MAX, "LINE_MAX"},
		{sysconf.SC_LOGIN_NAME_MAX, C._SC_LOGIN_NAME_MAX, "LOGIN_NAME_MAX"},
		{sysconf.SC_MQ_OPEN_MAX, C._SC_MQ_OPEN_MAX, "MQ_OPEN_MAX"},
		{sysconf.SC_MQ_PRIO_MAX, C._SC_MQ_PRIO_MAX, "MQ_PRIO_MAX"},
		{sysconf.SC_NGROUPS_MAX, C._SC_NGROUPS_MAX, "NGROUPS_MAX"},
		{sysconf.SC_OPEN_MAX, C._SC_OPEN_MAX, "OPEN_MAX"},
		{sysconf.SC_PAGE_SIZE, C._SC_PAGE_SIZE, "PAGE_SIZE"},
		{sysconf.SC_PAGESIZE, C._SC_PAGESIZE, "PAGESIZE"},
		{sysconf.SC_THREAD_DESTRUCTOR_ITERATIONS, C._SC_THREAD_DESTRUCTOR_ITERATIONS, "PTHREAD_DESTRUCTOR_ITERATIONS"},
		{sysconf.SC_THREAD_KEYS_MAX, C._SC_THREAD_KEYS_MAX, "PTHREAD_KEYS_MAX"},
		{sysconf.SC_THREAD_PRIO_INHERIT, C._SC_THREAD_PRIO_INHERIT, "POSIX_THREAD_PRIO_INHERIT"},
		{sysconf.SC_THREAD_PRIO_PROTECT, C._SC_THREAD_PRIO_PROTECT, "POSIX_THREAD_PRIO_PROTECT"},
		{sysconf.SC_THREAD_STACK_MIN, C._SC_THREAD_STACK_MIN, "PTHREAD_STACK_MIN"},
		{sysconf.SC_THREAD_THREADS_MAX, C._SC_THREAD_THREADS_MAX, "PTHREAD_THREADS_MAX"},
		{sysconf.SC_RE_DUP_MAX, C._SC_RE_DUP_MAX, "RE_DUP_MAX"},
		{sysconf.SC_RTSIG_MAX, C._SC_RTSIG_MAX, "RTSIG_MAX"},
		{sysconf.SC_SEM_NSEMS_MAX, C._SC_SEM_NSEMS_MAX, "SEM_NSEMS_MAX"},
		{sysconf.SC_SEM_VALUE_MAX, C._SC_SEM_VALUE_MAX, "SEM_VALUE_MAX"},
		{sysconf.SC_SIGQUEUE_MAX, C._SC_SIGQUEUE_MAX, "SIGQUEUE_MAX"},
		{sysconf.SC_STREAM_MAX, C._SC_STREAM_MAX, "STREAM_MAX"},
		{sysconf.SC_SYMLOOP_MAX, C._SC_SYMLOOP_MAX, "SYMLOOP_MAX"},
		{sysconf.SC_TIMER_MAX, C._SC_TIMER_MAX, "TIMER_MAX"},
		{sysconf.SC_TTY_NAME_MAX, C._SC_TTY_NAME_MAX, "TTY_NAME_MAX"},
		{sysconf.SC_TZNAME_MAX, C._SC_TZNAME_MAX, "TZNAME_MAX"},

		{sysconf.SC_ADVISORY_INFO, C._SC_ADVISORY_INFO, "_POSIX_ADVISORY_INFO"},
		{sysconf.SC_ASYNCHRONOUS_IO, C._SC_ASYNCHRONOUS_IO, "_POSIX_ASYNCHRONOUS_IO"},
		{sysconf.SC_BARRIERS, C._SC_BARRIERS, "_POSIX_BARRIERS"},
		{sysconf.SC_CLOCK_SELECTION, C._SC_CLOCK_SELECTION, "_POSIX_CLOCK_SELECTION"},
		{sysconf.SC_CPUTIME, C._SC_CPUTIME, "_POSIX_CPUTIME"},
		{sysconf.SC_FSYNC, C._SC_FSYNC, "_POSIX_FSYNC"},
		{sysconf.SC_IPV6, C._SC_IPV6, "_POSIX_IPV6"},
		{sysconf.SC_JOB_CONTROL, C._SC_JOB_CONTROL, "_POSIX_JOB_CONTROL"},
		{sysconf.SC_MAPPED_FILES, C._SC_MAPPED_FILES, "_POSIX_MAPPED_FILES"},
		{sysconf.SC_MEMLOCK, C._SC_MEMLOCK, "_POSIX_MEMLOCK"},
		{sysconf.SC_MEMLOCK_RANGE, C._SC_MEMLOCK_RANGE, "_POSIX_MEMLOCK_RANGE"},
		{sysconf.SC_MEMORY_PROTECTION, C._SC_MEMORY_PROTECTION, "_POSIX_MEMORY_PROTECTION"},
		{sysconf.SC_MESSAGE_PASSING, C._SC_MESSAGE_PASSING, "_POSIX_MESSAGE_PASSING"},
		{sysconf.SC_MONOTONIC_CLOCK, C._SC_MONOTONIC_CLOCK, "_POSIX_MONOTONIC_CLOCK"},
		{sysconf.SC_PRIORITIZED_IO, C._SC_PRIORITIZED_IO, "_POSIX_PRIORITIZED_IO"},
		{sysconf.SC_PRIORITY_SCHEDULING, C._SC_PRIORITY_SCHEDULING, "_POSIX_PRIORITY_SCHEDULING"},
		{sysconf.SC_RAW_SOCKETS, C._SC_RAW_SOCKETS, "_POSIX_RAW_SOCKETS"},
		{sysconf.SC_READER_WRITER_LOCKS, C._SC_READER_WRITER_LOCKS, "_POSIX_READER_WRITER_LOCKS"},
		{sysconf.SC_REALTIME_SIGNALS, C._SC_REALTIME_SIGNALS, "_POSIX_REALTIME_SIGNALS"},
		{sysconf.SC_REGEXP, C._SC_REGEXP, "_POSIX_REGEXP"},
		{sysconf.SC_SAVED_IDS, C._SC_SAVED_IDS, "_POSIX_SAVED_IDS"},
		{sysconf.SC_SEMAPHORES, C._SC_SEMAPHORES, "_POSIX_SEMAPHORES"},
		{sysconf.SC_SHELL, C._SC_SHELL, "_POSIX_SHELL"},
		{sysconf.SC_THREAD_CPUTIME, C._SC_THREAD_CPUTIME, "_POSIX_THREAD_CPUTIME"},
		{sysconf.SC_THREADS, C._SC_THREADS, "_POSIX_THREADS"},
		{sysconf.SC_TIMEOUTS, C._SC_TIMEOUTS, "_POSIX_TIMEOUTS"},
		{sysconf.SC_TIMERS, C._SC_TIMERS, "_POSIX_TIMERS"},
		{sysconf.SC_VERSION, C._SC_VERSION, "_POSIX_VERSION"},

		{sysconf.SC_2_C_DEV, C._SC_2_C_DEV, "_POSIX2_C_DEV"},
		// glibc version < 2.22 does not define _POSIX2_C_VERSION
		// despite supporting the standard, see glibc bug BZ 438. Test it
		// in testPosix2CVersion below.
		// {sysconf.SC_2_C_VERSION, C._SC_2_C_VERSION, "_POSIX2_C_VERSION"},
		{sysconf.SC_2_FORT_DEV, C._SC_2_FORT_DEV, "_POSIX2_FORT_DEV"},
		{sysconf.SC_2_FORT_RUN, C._SC_2_FORT_RUN, "_POSIX2_FORT_RUN"},
		{sysconf.SC_2_LOCALEDEF, C._SC_2_LOCALEDEF, "_POSIX2_LOCALEDEF"},
		{sysconf.SC_2_SW_DEV, C._SC_2_SW_DEV, "_POSIX2_SW_DEV"},
		{sysconf.SC_2_UPE, C._SC_2_UPE, "_POSIX2_UPE"},
		{sysconf.SC_2_VERSION, C._SC_2_VERSION, "_POSIX2_VERSION"},

		{sysconf.SC_XOPEN_CRYPT, C._SC_XOPEN_CRYPT, "_XOPEN_CRYPT"},
		{sysconf.SC_XOPEN_REALTIME, C._SC_XOPEN_REALTIME, "_XOPEN_REALTIME"},
		{sysconf.SC_XOPEN_UNIX, C._SC_XOPEN_UNIX, "_XOPEN_UNIX"},
		{sysconf.SC_XOPEN_VERSION, C._SC_XOPEN_VERSION, "_XOPEN_VERSION"},
		{sysconf.SC_XOPEN_XCU_VERSION, C._SC_XOPEN_XCU_VERSION, "_XOPEN_XCU_VERSION"},

		// non-standard
		{sysconf.SC_PHYS_PAGES, C._SC_PHYS_PAGES, "_PHYS_PAGES"},
		// AV_PHYS_PAGES might change between calling Go and C version
		// of sysconf. Don't test it for now.
		{sysconf.SC_NPROCESSORS_CONF, C._SC_NPROCESSORS_CONF, "_NPROCESSORS_CONF"},
		{sysconf.SC_NPROCESSORS_ONLN, C._SC_NPROCESSORS_ONLN, "_NPROCESSORS_ONLN"},
		{sysconf.SC_UIO_MAXIOV, C._SC_UIO_MAXIOV, "UIO_MAXIOV"},
	}

	for _, tc := range testCases {
		testSysconfGoCgo(t, tc)
	}

	testPosix2CVersion(t)
}

func testSysconfGoCgo(t *testing.T, tc testCase) {
	if tc.goVar != int(tc.cVar) {
		t.Errorf("sysconf variable %v values in Go and C don't match: %v <-> %v", tc.name, tc.goVar, tc.cVar)
	}
	goVal, goErr := sysconf.Sysconf(tc.goVar)
	if goErr != nil && goErr != unix.EINVAL {
		t.Errorf("Sysconf(%s/%d): %v", tc.name, tc.goVar, goErr)
	}
	if goErr != unix.EINVAL {
		t.Logf("%s = %v", tc.name, goVal)
	}

	cVal, cErr := C.sysconf(tc.cVar)
	if cErr != nil && cErr != unix.EINVAL {
		t.Fatalf("Sysconf(%s/%d): %v", tc.name, tc.goVar, cErr)
	}

	if goVal != int64(cVal) {
		t.Errorf("values in Go and C for %v don't match: %v <-> %v", tc.name, goVal, cVal)
	}
	if goErr != cErr {
		t.Errorf("error values in Go and C for %v don't match: %v <-> %v", tc.name, goErr, cErr)
	}
}

func testPosix2CVersion(t *testing.T) {
	glibcVer := C.GoString(C.gnu_get_libc_version())
	majMin := strings.Split(glibcVer, ".")
	if len(majMin) < 2 {
		t.Fatalf("unexpected glibc version: %v\n", glibcVer)
	}
	t.Logf("glibc version: %v\n", glibcVer)

	maj, _ := strconv.Atoi(majMin[0])
	min, _ := strconv.Atoi(majMin[1])
	// _POSIX2_C_VERSION was fixed in glibc 2.22, see glibc commit
	// 4e5f9259f352 ("Restore _POSIX2_C_VERSION definition (bug 438).")
	if maj <= 2 && min < 22 {
		t.Logf("skipping _POSIX2_C_VERSION test on glibc < 2.22")
	} else {
		tc := testCase{sysconf.SC_2_C_VERSION, C._SC_2_C_VERSION, "_POSIX2_C_VERSION"}
		testSysconfGoCgo(t, tc)
	}
}
