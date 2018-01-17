package subprocess

import (
	"runtime"
	"testing"
)

// Run() function tests

func TestRunCliMockStdoutZero(t *testing.T) {
	response := Run("climock", "--stdout", "This is a test")
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Expected mock exit code to be zero and it was %d", response.ExitCode)
	}
	if response.StdOut != "This is a test" {
		t.Errorf("[FAIL] Expected mock std out to be 'This is a test' and it was actually '%s'", response.StdOut)
	}
	if len(response.StdErr) != 0 {
		t.Errorf("[FAIL] Expeceted no std err output but received '%s'", response.StdErr)
	}
}

func TestRunValidCommand(t *testing.T) {
	response := Run("git", "--help")
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
	}
	if len(response.StdErr) > 0 {
		t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
	}
	if len(response.StdOut) == 0 {
		t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
	}
}

func TestRunInValidCommandBadArgument(t *testing.T) {
	response := Run("git", "--bogus")
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Expected invalid argument to executable to return non-0 exit status code and instead it returned %d", response.ExitCode)
	}
	if len(response.StdErr) == 0 {
		t.Errorf("[FAIL] Expected invalid argument to executable to return standard error output and instead it returned an empty string")
	}
	if len(response.StdOut) > 0 {
		t.Errorf("[FAIL] Expected invalid argument to return no standard output but instead it returned %s.", response.StdOut)
	}
}

func TestRunInvalidCommandMissingExecutable(t *testing.T) {
	response := Run("bogus", "--help")
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Expected invalid command to return non-0 exit status code and instead it returned %d", response.ExitCode)
	}
	if len(response.StdErr) == 0 {
		t.Errorf("[FAIL] Expected invalid command to return standard error output and instead it returned an empty string")
	}
	if len(response.StdOut) > 0 {
		t.Errorf("[FAIL] Expected invalid command to return no standard output but instead it returned %s.", response.StdOut)
	}
}

// RunShell() function tests - *nix platform

func TestRunShellUnixValidDefaultShellCommandOneString(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("", "", "git --help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixValidDefaultShellCommandOneStringWithShellFlag(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("", "-c", "git --help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixValidDefaultShellCommandTwoStrings(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("", "", "git", "--help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixValidDefaultShellCommandTwoStringsWithShellFlag(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("", "-c", "git", "--help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixValidBashShellCommandOneString(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("/bin/bash", "", "git --help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixValidBashShellCommandOneStringWithShellFlag(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("/bin/bash", "-c", "git --help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixValidBashShellCommandTwoString(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("/bin/bash", "", "git", "--help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixValidBashShellCommandTwoStringWithShellFlag(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("/bin/bash", "-c", "git", "--help")
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellUnixInvalidShell(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("/usr/local/bin/bogusshell", "", "git", "--help")
		if response.ExitCode == 0 {
			t.Errorf("[FAIL] Expected command to return non-0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) == 0 {
			t.Errorf("[FAIL] Expected command to return standard error output and instead it returned empty string")
		}
		if len(response.StdOut) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard output but instead it returned %s.", response.StdOut)
		}
	}
}

func TestRunShellUnixInvalidExecutable(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("", "", "completelybogus", "--help")
		if response.ExitCode == 0 {
			t.Errorf("[FAIL] Expected command to return non-0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) == 0 {
			t.Errorf("[FAIL] Expected command to return standard error output and instead it returned empty string")
		}
		if len(response.StdOut) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard output but instead it returned %s.", response.StdOut)
		}
	}
}

func TestRunShellUnixInvalidExecutableArgument(t *testing.T) {
	if runtime.GOOS != "windows" {
		response := RunShell("", "", "git", "--bogus")
		if response.ExitCode == 0 {
			t.Errorf("[FAIL] Expected command to return non-0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) == 0 {
			t.Errorf("[FAIL] Expected command to return standard error output and instead it returned empty string")
		}
		if len(response.StdOut) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard output but instead it returned %s.", response.StdOut)
		}
	}
}

////////////////////////////////////////////////
//  RunShell() function tests - Windows platform
////////////////////////////////////////////////

func TestRunShellWindowsValidDefaultShellCommandOneString(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("", "", "dir /AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsValidDefaultShellCommandOneStringWithShellFlag(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("", "/C", "dir /AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsValidDefaultShellCommandTwoStrings(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("", "", "dir", "/AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsValidDefaultShellCommandTwoStringsWithShellFlag(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("", "/C", "dir", "/AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsValidMockShellCommandOneString(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("cmd.exe", "", "dir /AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsValidMockShellCommandOneStringWithShellFlag(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("cmd.exe", "/C", "dir /AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsValidMockShellCommandTwoString(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("cmd.exe", "", "dir", "/AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsValidMockShellCommandTwoStringWithShellFlag(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("cmd.exe", "/C", "dir", "/AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode != 0 {
			t.Errorf("[FAIL] Expected command to return 0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard error output and instead it returned %v", response.StdErr)
		}
		if len(response.StdOut) == 0 {
			t.Errorf("[FAIL] Expected command to return standard output but instead it was empty.")
		}
	}
}

func TestRunShellWindowsInvalidShell(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("totallybogusshell", "", "dir", "/AD")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode == 0 {
			t.Errorf("[FAIL] Expected command to return non-0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) == 0 {
			t.Errorf("[FAIL] Expected command to return standard error output and instead it returned empty string")
		}
		if len(response.StdOut) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard output but instead it returned %s.", response.StdOut)
		}
	}
}

func TestRunShellWindowsInvalidExecutable(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("", "", "completelybogus", "--help")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode == 0 {
			t.Errorf("[FAIL] Expected command to return non-0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) == 0 {
			t.Errorf("[FAIL] Expected command to return standard error output and instead it returned empty string")
		}
		if len(response.StdOut) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard output but instead it returned %s.", response.StdOut)
		}
	}
}

func TestRunShellWindowsInvalidExecutableArgument(t *testing.T) {
	if runtime.GOOS == "windows" {
		response := RunShell("", "", "dir", "/ZZZ")
		t.Logf("%d\n", response.ExitCode)
		t.Logf("%s\n", response.StdOut)
		t.Logf("%s\n", response.StdErr)
		if response.ExitCode == 0 {
			t.Errorf("[FAIL] Expected command to return non-0 exit status code and instead it returned %d", response.ExitCode)
		}
		if len(response.StdErr) == 0 {
			t.Errorf("[FAIL] Expected command to return standard error output and instead it returned empty string")
		}
		if len(response.StdOut) > 0 {
			t.Errorf("[FAIL] Expected command to return no standard output but instead it returned %s.", response.StdOut)
		}
	}
}
