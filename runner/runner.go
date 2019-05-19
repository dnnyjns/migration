package runner

import (
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
)

type Runner struct {
	Db   *gorm.DB
	Migs Migraines
}

func (r *Runner) Run() {
	var (
		ids       = make([]string, len(r.Migs))
		persisted Migraines
	)

	for _, migraine := range r.Migs {
		ids = append(ids, migraine.Version)
	}
	r.Db.Model(&Migraine{}).Where(ids).Find(&persisted)

	persisted.sort()
	r.Migs.sort()

	for _, migraine := range r.Migs {
		length := len(persisted)
		i := sort.Search(length, func(i int) bool { return persisted[i].Version == migraine.Version })
		fmt.Println(i)
		if i == length {
			migraine.Run(r.Db)
		}
	}
}
