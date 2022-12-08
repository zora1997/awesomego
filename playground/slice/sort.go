package main

import (
	"fmt"
	"sort"
)

// 音效数据库内字段
type DBSoundEffect struct {
	SoundEffectID int32
	Sequence      uint32
}

type SoundEffect []*DBSoundEffect

func (so SoundEffect) Len() int {
	return len(so)
}
func (so SoundEffect) Swap(i, j int) {
	so[i], so[j] = so[j], so[i]
}
func (so SoundEffect) Less(i, j int) bool {
	// 从小到大排序
	return so[j].Sequence > so[i].Sequence
}

func main() {
	soundEffects := make(SoundEffect, 0)

	soundEffects = append(soundEffects, &DBSoundEffect{
		SoundEffectID: 111,
		Sequence:      1,
	},
		&DBSoundEffect{
			SoundEffectID: 222,
			Sequence:      3,
		},
		&DBSoundEffect{
			SoundEffectID: 333,
			Sequence:      2,
		})

	for i := range soundEffects {
		fmt.Printf("sound before : %+v\n", soundEffects[i])
	}

	sort.Sort(soundEffects)

	for i := range soundEffects {
		fmt.Printf("sound after: %+v\n", soundEffects[i])
	}
}
