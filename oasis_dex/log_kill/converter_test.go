package log_kill_test

import (
	"github.com/8thlight/oasis_watcher/oasis_dex/log_kill"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vulcanize/vulcanizedb/pkg/core"
)

var _ = Describe("WatchedEvent to LogKill Model", func() {
	It("Converts a WatchedEvent to LogKillModel", func() {
		watchedEvent := core.WatchedEvent{
			Topic1: "0x0000000000000000000000000000000000000000000000000000000000009f06",
		}
		cnvtr := log_kill.LogKillConverter{}
		lkm, err := cnvtr.ToModel(watchedEvent)
		Expect(err).ToNot(HaveOccurred())
		var topic1AsInt uint64 = 40710
		Expect(lkm.ID).To(Equal(topic1AsInt))
	})

	It("Converts another WatchedEvent to LogKillModel", func() {
		watchedEvent := core.WatchedEvent{
			Topic1: "0x0000000000000000000000000000000000000000000000000000000000009eda",
		}
		cnvtr := log_kill.LogKillConverter{}
		lkm, err := cnvtr.ToModel(watchedEvent)
		Expect(err).ToNot(HaveOccurred())
		var topic1AsInt uint64 = 40666
		Expect(lkm.ID).To(Equal(topic1AsInt))
	})
})
