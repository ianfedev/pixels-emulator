package ephemeral

import (
	healthcheck "pixels-emulator/healthcheck/scheduler"
)

func Cron() {

	healthcheck.SchedulePing()

}
