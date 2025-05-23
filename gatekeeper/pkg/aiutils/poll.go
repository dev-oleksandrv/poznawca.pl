package aiutils

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"time"
)

var (
	ErrPollRunStatusFailedToGetRunStatus = errors.New("failed_to_get_run_status")
	ErrPollRunStatusRunFailedOrCancelled = errors.New("run_failed_or_cancelled")
)

func PollRunStatus(ctx context.Context, client *openai.Client, threadID, runID string) (*openai.Run, error) {
	interval := 2

	for {
		run, err := client.RetrieveRun(ctx, threadID, runID)
		if err != nil {
			return nil, ErrPollRunStatusFailedToGetRunStatus
		}

		switch run.Status {
		case openai.RunStatusCompleted:
			return &run, nil
		case openai.RunStatusFailed, openai.RunStatusCancelled:
			return nil, ErrPollRunStatusRunFailedOrCancelled
		}

		time.Sleep(time.Second * time.Duration(interval))

		if interval < 5 {
			interval += 1
		}
	}
}
