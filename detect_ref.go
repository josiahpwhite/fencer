// +build !amd64,!386 gccgo

package fencer

func initCPU() {
        lfence = func() {}
        mfence = func() {}
        sfence = func() {}
}

