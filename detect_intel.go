// +build 386,!gccgo amd64,!gccgo

package fencer

func asmLFence()
func asmMFence()
func asmSFence()

func initCPU() {
        lfence = asmLFence
        mfence = asmMFence
        sfence = asmSFence
}
